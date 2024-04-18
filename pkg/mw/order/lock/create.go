package orderlock

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/lock"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) checkOrderExist(ctx context.Context, tx *ent.Tx, req *orderlockcrud.Req) error {
	exist, err := tx.Order.
		Query().
		Where(
			entorder.EntID(*req.OrderID),
			entorder.DeletedAt(0),
		).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("invalid order")
	}
	return nil
}

func (h *createHandler) checkLockExist(ctx context.Context, tx *ent.Tx, req *orderlockcrud.Req) error {
	stm := tx.OrderLock.
		Query().
		Where(
			entorderlock.OrderID(*req.OrderID),
			entorderlock.LockType(req.LockType.String()),
		)
	if *req.LockType == basetypes.OrderLockType_LockCommission {
		stm.Where(
			entorderlock.AppID(*req.AppID),
			entorderlock.UserID(*req.UserID),
		)
	}
	exist, err := stm.Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("orderlockid is exist")
	}
	return nil
}

func (h *createHandler) createOrderLock(ctx context.Context, tx *ent.Tx, req *orderlockcrud.Req) error {
	if _, err := orderlockcrud.CreateSet(
		tx.OrderLock.Create(),
		&orderlockcrud.Req{
			EntID:    req.EntID,
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
			if req.EntID == nil {
				req.EntID = &id
			}
			if err := handler.checkOrderExist(ctx, tx, req); err != nil {
				return err
			}
			if err := handler.checkLockExist(ctx, tx, req); err != nil {
				return err
			}
			if err := handler.createOrderLock(ctx, tx, req); err != nil {
				return err
			}
			ids = append(ids, *req.EntID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &orderlockcrud.Conds{
		EntIDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrderLocks(ctx)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
