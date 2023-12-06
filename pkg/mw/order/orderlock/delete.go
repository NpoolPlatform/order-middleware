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
	entIDs := []uuid.UUID{}
	for _, req := range h.Reqs {
		if req.EntID == nil && req.ID == nil {
			return nil, fmt.Errorf("invalid id")
		}
		if req.ID != nil {
			ids = append(ids, *req.ID)
			continue
		}
		if req.EntID != nil {
			entIDs = append(entIDs, *req.EntID)
		}
	}
	infos := []*npool.OrderLock{}
	if len(ids) > 0 {
		h.Conds = &orderlockcrud.Conds{IDs: &cruder.Cond{Op: cruder.IN, Val: ids}}
		h.Limit = int32(len(ids))
		orderlocks, _, err := h.GetOrderLocks(ctx)
		if err != nil {
			return nil, err
		}
		infos = append(infos, orderlocks...)
	}
	if len(entIDs) > 0 {
		h.Conds = &orderlockcrud.Conds{EntIDs: &cruder.Cond{Op: cruder.IN, Val: entIDs}}
		h.Limit = int32(len(entIDs))
		orderlocks, _, err := h.GetOrderLocks(ctx)
		if err != nil {
			return nil, err
		}
		infos = append(infos, orderlocks...)
	}
	if len(infos) == 0 {
		return nil, nil
	}
	if len(infos) != len(h.Reqs) {
		return nil, fmt.Errorf("atomic denied")
	}

	orderlockMap := map[string]*npool.OrderLock{}
	for _, val := range infos {
		orderlockMap[val.EntID] = val
	}

	handler := &deleteHandler{
		Handler:   h,
		deletedAt: uint32(time.Now().Unix()),
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.ID == nil {
				orderlock, ok := orderlockMap[req.EntID.String()]
				if !ok {
					return fmt.Errorf("orderlock not found")
				}
				req.ID = &orderlock.ID
			}
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
