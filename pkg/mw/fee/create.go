package feeorder

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	orderbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderbase"
)

type createHandler struct {
	*Handler
	sql                  string
	sqlOrderBase         string
	sqlOrderStateBase    string
	sqlFeeOrderStateBase string
	sqlLedgerLock        string
	sqlOrderCoupons      []string
	sqlPaymentBalances   []string
	sqlPaymentTransfers  []string
}

func (h *createHandler) constructSQL() {
	h.sql = h.ConstructCreateSQL()
}

func (h *createHandler) constructOrderBaseSQL(ctx context.Context) {
	handler, _ := orderbase1.NewHandler(
		ctx,
	)
	handler.Req = *h.OrderBaseReq
	h.sqlOrderBase = handler.ConstructCreateSQL()
}

func (h *Handler) CreateFeeOrder(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}

	handler.constructOrderBaseSQL(ctx)
	handler.constructSQL()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return nil
	})
}
