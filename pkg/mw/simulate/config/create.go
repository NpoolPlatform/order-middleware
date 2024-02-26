package config

import (
	"context"
	"fmt"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"
	configcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/simulate/config"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkSimulateConfig(ctx context.Context) error {
	if h.Enabled == nil || !*h.Enabled {
		return nil
	}
	h.Conds = &configcrud.Conds{
		AppID:   &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		Enabled: &cruder.Cond{Op: cruder.EQ, Val: *h.Enabled},
	}
	exist, err := h.ExistSimulateConfigConds(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("repeated config")
	}
	return nil
}

func (h *createHandler) createSimulateConfig(ctx context.Context, tx *ent.Tx) error {
	if _, err := configcrud.CreateSet(
		tx.SimulateConfig.Create(),
		&configcrud.Req{
			EntID:                     h.EntID,
			AppID:                     h.AppID,
			SendCouponMode:            h.SendCouponMode,
			SendCouponProbability:     h.SendCouponProbability,
			EnabledCashableProfit:     h.EnabledCashableProfit,
			CashableProfitProbability: h.CashableProfitProbability,
			Enabled:                   h.Enabled,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateSimulateConfig(ctx context.Context) (*npool.SimulateConfig, error) {
	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	if err := handler.checkSimulateConfig(ctx); err != nil {
		return nil, err
	}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.createSimulateConfig(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetSimulateConfig(ctx)
}
