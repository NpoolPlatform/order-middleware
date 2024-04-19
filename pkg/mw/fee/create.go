package feeorder

import (
	"context"

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
		return nil
	})
}
