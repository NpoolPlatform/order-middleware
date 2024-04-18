package appconfig

import (
	"context"
	"time"

	appconfigcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/app/config"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteAppConfig(ctx context.Context, tx *ent.Tx) error {
	if _, err := appconfigcrud.UpdateSet(
		tx.AppConfig.UpdateOneID(*h.ID),
		&appconfigcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteAppConfig(ctx context.Context) error {
	info, err := h.GetAppConfig(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.deleteAppConfig(ctx, tx)
	})
}
