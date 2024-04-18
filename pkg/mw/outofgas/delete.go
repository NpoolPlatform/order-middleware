package outofgas

import (
	"context"
	"time"

	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteOutOfGas(ctx context.Context, tx *ent.Tx) error {
	if _, err := outofgascrud.UpdateSet(
		tx.OutOfGas.UpdateOneID(*h.ID),
		&outofgascrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOutOfGas(ctx context.Context) error {
	info, err := h.GetOutOfGas(ctx)
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
		return handler.deleteOutOfGas(_ctx, tx)
	})
}
