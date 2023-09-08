package orderlock

import (
	"context"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createOrderLock(ctx context.Context, tx *ent.Tx, req *orderlockcrud.Req) error {
	if _, err := orderlockcrud.CreateSet(
		tx.OrderLock.Create(),
		&orderlockcrud.Req{
			ID:       req.ID,
			AppID:    req.AppID,
			UserID:   req.UserID,
			OrderID:  req.OrderID,
			LockType: req.LockType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateOrderLocks(ctx context.Context) ([]*npool.OrderLock, error) {
	handler := &createHandler{
		Handler: h,
	}
	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			id := uuid.New()
			if req.ID == nil {
				req.ID = &id
			}
			if err := handler.createOrderLock(ctx, tx, req); err != nil {
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
