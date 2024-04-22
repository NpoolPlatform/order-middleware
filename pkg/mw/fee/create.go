package feeorder

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	feeorderstate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee/state"
	ordercoupon1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/coupon"
	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/lock"
	orderbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderbase"
	orderstatebase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/statebase"
	paymentbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
	paymentbalance1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/balance"
	paymentbalancelock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/balance/lock"
	paymenttransfer1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/transfer"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	payWithBalance        bool
	sql                   string
	sqlOrderBase          string
	sqlOrderStateBase     string
	sqlFeeOrderState      string
	sqlLedgerLock         string
	sqlPaymentBalanceLock string
	sqlOrderCoupons       []string
	sqlPaymentBase        string
	sqlPaymentBalances    []string
	sqlPaymentTransfers   []string
}

func (h *createHandler) constructSQL() {
	h.sql = h.ConstructCreateSQL()
}

func (h *createHandler) constructOrderBaseSQL(ctx context.Context) {
	handler, _ := orderbase1.NewHandler(ctx)
	handler.Req = *h.OrderBaseReq
	h.sqlOrderBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructOrderStateBaseSQL(ctx context.Context) {
	handler, _ := orderstatebase1.NewHandler(ctx)
	handler.Req = *h.OrderStateBaseReq
	handler.Req.StartMode = func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }()
	h.sqlOrderStateBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructFeeOrderStateSQL(ctx context.Context) {
	handler, _ := feeorderstate1.NewHandler(ctx)
	handler.Req = *h.FeeOrderStateReq
	h.sqlFeeOrderState = handler.ConstructCreateSQL()
}

func (h *createHandler) constructLedgerLockSQL(ctx context.Context) {
	if h.LedgerLockReq.EntID == nil {
		return
	}
	handler, _ := orderlock1.NewHandler(ctx)
	handler.Req = *h.LedgerLockReq
	h.sqlLedgerLock = handler.ConstructCreateSQL()
}

func (h *createHandler) constructPaymentBalanceLockSQL(ctx context.Context) {
	if h.LedgerLockReq.EntID == nil {
		return
	}
	handler, _ := paymentbalancelock1.NewHandler(ctx)
	handler.Req = *h.PaymentBalanceLockReq
	h.sqlPaymentBalanceLock = handler.ConstructCreateSQL()
	h.payWithBalance = true
}

func (h *createHandler) constructOrderCouponSQLs(ctx context.Context) {
	for _, req := range h.OrderCouponReqs {
		handler, _ := ordercoupon1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderCoupons = append(h.sqlOrderCoupons, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) constructPaymentBaseSQL(ctx context.Context) {
	// TODO: some case do not need to create payment info
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.PaymentBaseReq
	h.sqlPaymentBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructPaymentBalanceSQLs(ctx context.Context) {
	for _, req := range h.PaymentBalanceReqs {
		handler, _ := paymentbalance1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentBalances = append(h.sqlPaymentBalances, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) constructPaymentTransferSQLs(ctx context.Context) {
	for _, req := range h.PaymentTransferReqs {
		handler, _ := paymenttransfer1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, handler.ConstructCreateSQL())
	}
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create powerrental: %v", err)
	}
	return nil
}

func (h *createHandler) createOrderBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderBase)
}

func (h *createHandler) createOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderStateBase)
}

func (h *createHandler) createFeeOrderState(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlFeeOrderState)
}

func (h *createHandler) createLedgerLock(ctx context.Context, tx *ent.Tx) error {
	if !h.payWithBalance {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlLedgerLock)
}

func (h *createHandler) createPaymentBalanceLock(ctx context.Context, tx *ent.Tx) error {
	if !h.payWithBalance {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlPaymentBalanceLock)
}

func (h *createHandler) createOrderCoupons(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlOrderCoupons {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) createPaymentBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlPaymentBase)
}

func (h *createHandler) createPaymentBalances(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentBalances {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) createPaymentTransfers(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentTransfers {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) createFeeOrder(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *createHandler) formalizeOrderID() {
	if h.OrderID != nil {
		return
	}
	h.OrderID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	h.OrderBaseReq.EntID = h.OrderID
	h.OrderStateBaseReq.OrderID = h.OrderID
	h.FeeOrderStateReq.OrderID = h.OrderID
	h.LedgerLockReq.OrderID = h.OrderID
	h.PaymentBaseReq.OrderID = h.OrderID
}

func (h *createHandler) formalizeEntIDs() {
	if h.OrderStateBaseReq.EntID == nil {
		h.OrderStateBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.FeeOrderStateReq.EntID == nil {
		h.FeeOrderStateReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.PaymentBalanceLockReq.EntID == nil {
		h.PaymentBalanceLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *createHandler) formalizeOrderCoupons() {
	for _, req := range h.OrderCouponReqs {
		req.OrderID = h.OrderID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentBalances() {
	for _, req := range h.PaymentBalanceReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentTransfers() {
	for _, req := range h.PaymentTransferReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *createHandler) formalizePaymentID() {
	if h.PaymentBaseReq.EntID != nil {
		return
	}
	h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	h.FeeOrderStateReq.PaymentID = h.PaymentBaseReq.EntID
	h.PaymentBalanceLockReq.PaymentID = h.PaymentBaseReq.EntID
}

func (h *Handler) CreateFeeOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &createHandler{
		Handler: h,
	}

	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}

	handler.formalizeOrderID()
	handler.formalizeEntIDs()
	handler.formalizeOrderCoupons()
	handler.formalizePaymentID()
	handler.formalizePaymentBalances()
	handler.formalizePaymentTransfers()

	handler.constructOrderBaseSQL(ctx)
	handler.constructOrderStateBaseSQL(ctx)
	handler.constructFeeOrderStateSQL(ctx)
	handler.constructLedgerLockSQL(ctx)
	handler.constructPaymentBalanceLockSQL(ctx)
	handler.constructOrderCouponSQLs(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	handler.constructPaymentTransferSQLs(ctx)
	handler.constructSQL()

	if err := handler.createOrderBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.createOrderStateBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.createFeeOrderState(ctx, tx); err != nil {
		return err
	}
	if err := handler.createLedgerLock(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentBalanceLock(ctx, tx); err != nil {
		return err
	}
	if err := handler.createOrderCoupons(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentBalances(ctx, tx); err != nil {
		return err
	}
	if err := handler.createPaymentTransfers(ctx, tx); err != nil {
		return err
	}
	return handler.createFeeOrder(ctx, tx)
}

func (h *Handler) CreateFeeOrder(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateFeeOrderWithTx(_ctx, tx)
	})
}
