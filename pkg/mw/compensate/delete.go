package compensate

import (
	"context"
	"time"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteCompensate(ctx context.Context, tx *ent.Tx) error {
	if _, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteCompensate(ctx context.Context) error {
	info, err := h.GetCompensate(ctx)
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
		return handler.deleteCompensate(_ctx, tx)
	})
}
