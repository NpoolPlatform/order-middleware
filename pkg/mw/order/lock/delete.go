package orderlock

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/lock"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteOrderLock(ctx context.Context, tx *ent.Tx) error {
	_, err := orderlockcrud.UpdateSet(
		tx.OrderLock.UpdateOneID(*h.ID),
		&orderlockcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return wlog.WrapError(err)
}

func (h *Handler) DeleteOrderLockWithTx(ctx context.Context, tx *ent.Tx) error {
	info, err := h.GetOrderLock(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return nil
	}
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	h.ID = &info.ID
	return handler.deleteOrderLock(ctx, tx)
}

func (h *Handler) DeleteOrderLock(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteOrderLockWithTx(_ctx, tx)
	})
}
