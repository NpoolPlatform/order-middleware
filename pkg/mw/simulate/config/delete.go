package config

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"
	configcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/simulate/config"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteSimulateConfig(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := configcrud.UpdateSet(
		tx.SimulateConfig.UpdateOneID(*h.ID),
		&configcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteSimulateConfig(ctx context.Context) (*npool.SimulateConfig, error) {
	info, err := h.GetSimulateConfig(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &deleteHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteSimulateConfig(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
