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
	paymenttransfer1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/transfer"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	sql                 string
	sqlOrderBase        string
	sqlOrderStateBase   string
	sqlFeeOrderState    string
	sqlLedgerLock       string
	sqlOrderCoupons     []string
	sqlPaymentBase      string
	sqlPaymentBalances  []string
	sqlPaymentTransfers []string
}

func (h *updateHandler) constructSQL() {
	h.sql = h.ConstructUpdateSQL()
}

func (h *updateHandler) constructOrderBaseSQL(ctx context.Context) {
	handler, _ := orderbase1.NewHandler(ctx)
	handler.Req = *h.OrderBaseReq
	h.sqlOrderBase = handler.ConstructUpdateSQL()
}

func (h *updateHandler) constructOrderStateBaseSQL(ctx context.Context) {
	handler, _ := orderstatebase1.NewHandler(ctx)
	handler.Req = *h.OrderStateBaseReq
	handler.Req.StartMode = func() *types.OrderStartMode { e := types.OrderStartMode_OrderStartInstantly; return &e }()
	h.sqlOrderStateBase = handler.ConstructUpdateSQL()
}

func (h *updateHandler) constructFeeOrderStateSQL(ctx context.Context) {
	handler, _ := feeorderstate1.NewHandler(ctx)
	handler.Req = *h.FeeOrderStateReq
	h.sqlFeeOrderState = handler.ConstructUpdateSQL()
}

func (h *updateHandler) constructLedgerLockSQL(ctx context.Context) {
	handler, _ := orderlock1.NewHandler(ctx)
	handler.Req = *h.LedgerLockReq
	h.sqlLedgerLock = handler.ConstructUpdateSQL()
}

func (h *updateHandler) constructOrderCouponSQLs(ctx context.Context) {
	for _, req := range h.OrderCouponReqs {
		handler, _ := ordercoupon1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderCoupons = append(h.sqlOrderCoupons, handler.ConstructUpdateSQL())
	}
}

func (h *updateHandler) constructPaymentBaseSQL(ctx context.Context) {
	handler, _ := paymentbase1.NewHandler(ctx)
	handler.Req = *h.PaymentBaseReq
	h.sqlPaymentBase = handler.ConstructUpdateSQL()
}

func (h *updateHandler) constructPaymentBalanceSQLs(ctx context.Context) {
	for _, req := range h.PaymentBalanceReqs {
		handler, _ := paymentbalance1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentBalances = append(h.sqlPaymentBalances, handler.ConstructUpdateSQL())
	}
}

func (h *updateHandler) constructPaymentTransferSQLs(ctx context.Context) {
	for _, req := range h.PaymentTransferReqs {
		handler, _ := paymenttransfer1.NewHandler(ctx)
		handler.Req = *req
		h.sqlPaymentTransfers = append(h.sqlPaymentTransfers, handler.ConstructUpdateSQL())
	}
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail update powerrental: %v", err)
	}
	return nil
}

func (h *updateHandler) updateOrderBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderBase)
}

func (h *updateHandler) updateOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOrderStateBase)
}

func (h *updateHandler) updateFeeOrderState(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlFeeOrderState)
}

func (h *updateHandler) updateLedgerLock(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlLedgerLock)
}

func (h *updateHandler) updateOrderCoupons(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlOrderCoupons {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) updatePaymentBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlPaymentBase)
}

func (h *updateHandler) updatePaymentBalances(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentBalances {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) updatePaymentTransfers(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPaymentTransfers {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) updateFeeOrder(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sql)
}

func (h *updateHandler) formalizeOrderID() {
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

func (h *updateHandler) formalizeEntIDs() {
	if h.OrderStateBaseReq.EntID == nil {
		h.OrderStateBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.FeeOrderStateReq.EntID == nil {
		h.FeeOrderStateReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.LedgerLockReq.EntID == nil {
		h.LedgerLockReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	if h.PaymentBaseReq.EntID == nil {
		h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
}

func (h *updateHandler) formalizeOrderCoupons() {
	for _, req := range h.OrderCouponReqs {
		req.OrderID = h.OrderID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentBalances() {
	for _, req := range h.PaymentBalanceReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentTransfers() {
	for _, req := range h.PaymentTransferReqs {
		req.PaymentID = h.PaymentBaseReq.EntID
		if req.EntID == nil {
			req.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
		}
	}
}

func (h *updateHandler) formalizePaymentID() {
	if h.PaymentBaseReq.EntID != nil {
		return
	}
	h.PaymentBaseReq.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	h.FeeOrderStateReq.PaymentID = h.PaymentBaseReq.EntID
}

func (h *Handler) UpdateFeeOrder(ctx context.Context) error {
	handler := &updateHandler{
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
	handler.constructOrderCouponSQLs(ctx)
	handler.constructPaymentBaseSQL(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	handler.constructPaymentTransferSQLs(ctx)
	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateOrderBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateOrderStateBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateFeeOrderState(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateLedgerLock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateOrderCoupons(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updatePaymentBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updatePaymentBalances(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updatePaymentTransfers(_ctx, tx); err != nil {
			return err
		}
		return handler.updateFeeOrder(_ctx, tx)
	})
}
