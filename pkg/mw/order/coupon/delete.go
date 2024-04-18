package ordercoupon

import (
	"context"
	"time"

	ordercouponcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/coupon"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteOrderCoupon(ctx context.Context, tx *ent.Tx) error {
	if _, err := ordercouponcrud.UpdateSet(
		tx.OrderCoupon.UpdateOneID(*h.ID),
		&ordercouponcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOrderCoupon(ctx context.Context) error {
	info, err := h.GetOrderCoupon(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	h.ID = &info.ID
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteOrderCoupon(_ctx, tx)
	})
}
