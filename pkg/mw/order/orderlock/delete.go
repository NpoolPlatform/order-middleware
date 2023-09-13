package orderlock

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
	deletedAt uint32
}

func (h *deleteHandler) deleteOrderLock(ctx context.Context, tx *ent.Tx, id uuid.UUID) error {
	orderlock, err := tx.OrderLock.
		Query().
		Where(
			entorderlock.ID(id),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err := orderlockcrud.UpdateSet(
		orderlock.Update(),
		&orderlockcrud.Req{
			DeletedAt: &h.deletedAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOrderLocks(ctx context.Context) ([]*npool.OrderLock, error) {
	handler := &deleteHandler{
		Handler:   h,
		deletedAt: uint32(time.Now().Unix()),
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				return fmt.Errorf("invalid id")
			}
			if err := handler.deleteOrderLock(ctx, tx, *req.ID); err != nil {
				return err
			}
			ids = append(ids, *req.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &orderlockcrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrderLocks(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
