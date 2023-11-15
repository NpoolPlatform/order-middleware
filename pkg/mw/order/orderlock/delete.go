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
)

type deleteHandler struct {
	*Handler
	deletedAt uint32
}

func (h *deleteHandler) deleteOrderLock(ctx context.Context, tx *ent.Tx, id uint32) error {
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
	ids := []uint32{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return nil, fmt.Errorf("invalid id")
		}
		ids = append(ids, *req.ID)
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
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) > 0 && len(infos) < len(ids) {
		return nil, fmt.Errorf("atomic denied")
	}

	handler := &deleteHandler{
		Handler:   h,
		deletedAt: uint32(time.Now().Unix()),
	}
	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.deleteOrderLock(ctx, tx, *req.ID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
