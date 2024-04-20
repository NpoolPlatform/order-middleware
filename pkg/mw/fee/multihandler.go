package feeorder

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type MultiHandler struct {
	Handlers []*Handler
}

func (h *MultiHandler) AppendHandler(handler *Handler) {
	h.Handlers = append(h.Handlers, handler)
}

func (h *MultiHandler) CreateFeeOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.CreateFeeOrderWithTx(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (h *MultiHandler) CreateFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.CreateFeeOrdersWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) UpdateFeeOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.UpdateFeeOrderWithTx(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (h *MultiHandler) UpdateFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.UpdateFeeOrdersWithTx(_ctx, tx)
	})
}

func (h *MultiHandler) DeleteFeeOrdersWithTx(ctx context.Context, tx *ent.Tx) error {
	for _, handler := range h.Handlers {
		if err := handler.DeleteFeeOrderWithTx(ctx, tx); err != nil {
			return err
		}
	}
	return nil
}

func (h *MultiHandler) DeleteFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteFeeOrdersWithTx(_ctx, tx)
	})
}
