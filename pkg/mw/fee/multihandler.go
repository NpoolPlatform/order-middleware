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

func (h *MultiHandler) CreateFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, handler := range h.Handlers {
			if err := handler.CreateFeeOrderWithTx(_ctx, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (h *MultiHandler) UpdateFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, handler := range h.Handlers {
			if err := handler.UpdateFeeOrderWithTx(_ctx, tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (h *MultiHandler) DeleteFeeOrders(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, handler := range h.Handlers {
			if err := handler.DeleteFeeOrderWithTx(_ctx, tx); err != nil {
				return err
			}
		}
		return nil
	})
}
