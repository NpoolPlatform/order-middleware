package config

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"
	configcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/simulate/config"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateSimulateConfig(ctx context.Context, tx *ent.Tx) error {
	if _, err := configcrud.UpdateSet(
		tx.SimulateConfig.UpdateOneID(*h.ID),
		&configcrud.Req{
			Units:                 h.Units,
			SendCouponMode:        h.SendCouponMode,
			SendCouponProbability: h.SendCouponProbability,
			Enabled:               h.Enabled,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateSimulateConfig(ctx context.Context) (*npool.SimulateConfig, error) {
	handler := &updateHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateSimulateConfig(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSimulateConfig(ctx)
}
