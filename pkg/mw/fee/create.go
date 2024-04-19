package feeorder

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	feeorderstate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee/state"
	ordercoupon1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/coupon"
	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/lock"
	orderbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderbase"
	orderstatebase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/statebase"
	paymentbalance1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/balance"
	paymenttransfer1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment/transfer"
)

type createHandler struct {
	*Handler
	sql                 string
	sqlOrderBase        string
	sqlOrderStateBase   string
	sqlFeeOrderState    string
	sqlLedgerLock       string
	sqlOrderCoupons     []string
	sqlPaymentBalances  []string
	sqlPaymentTransfers []string
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
	h.sqlOrderStateBase = handler.ConstructCreateSQL()
}

func (h *createHandler) constructFeeOrderStateSQL(ctx context.Context) {
	handler, _ := feeorderstate1.NewHandler(ctx)
	handler.Req = *h.FeeOrderStateReq
	h.sqlFeeOrderState = handler.ConstructCreateSQL()
}

func (h *createHandler) constructLedgerLockSQL(ctx context.Context) {
	handler, _ := orderlock1.NewHandler(ctx)
	handler.Req = *h.LedgerLockReq
	h.sqlLedgerLock = handler.ConstructCreateSQL()
}

func (h *createHandler) constructOrderCouponSQLs(ctx context.Context) {
	for _, req := range h.OrderCouponReqs {
		handler, _ := ordercoupon1.NewHandler(ctx)
		handler.Req = *req
		h.sqlOrderCoupons = append(h.sqlOrderCoupons, handler.ConstructCreateSQL())
	}
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
	return h.execSQL(ctx, tx, h.sqlLedgerLock)
}

func (h *createHandler) createOrderCoupons(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlOrderCoupons {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
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

func (h *Handler) CreateFeeOrder(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}

	handler.constructOrderBaseSQL(ctx)
	handler.constructOrderStateBaseSQL(ctx)
	handler.constructFeeOrderStateSQL(ctx)
	handler.constructLedgerLockSQL(ctx)
	handler.constructOrderCouponSQLs(ctx)
	handler.constructPaymentBalanceSQLs(ctx)
	handler.constructPaymentTransferSQLs(ctx)
	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOrderBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createOrderStateBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createFeeOrderState(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createLedgerLock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createOrderCoupons(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createPaymentBalances(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createPaymentTransfers(_ctx, tx); err != nil {
			return err
		}
		return handler.createFeeOrder(_ctx, tx)
	})
}
