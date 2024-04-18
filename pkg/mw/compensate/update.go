package compensate

import (
	"context"
	"fmt"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateCompensate(ctx context.Context, tx *ent.Tx) error {
	if _, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
			CompensateSeconds: h.CompensateSeconds,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCompensate(ctx context.Context) error {
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid compensate")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCompensate(ctx, tx)
	})
}
