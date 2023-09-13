package orderlock

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID       *uuid.UUID
	AppID    *uuid.UUID
	UserID   *uuid.UUID
	OrderID  *uuid.UUID
	LockType *basetypes.OrderLockType
	Reqs     []*orderlockcrud.Req
	Conds    *orderlockcrud.Conds
	Offset   int32
	Limit    int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUserID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid userid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderID = &_id
		return nil
	}
}

func WithLockType(lockType *basetypes.OrderLockType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lockType == nil {
			if must {
				return fmt.Errorf("invalid locktype")
			}
			return nil
		}
		switch *lockType {
		case basetypes.OrderLockType_LockStock:
		case basetypes.OrderLockType_LockBalance:
		case basetypes.OrderLockType_LockCommission:
		default:
			return fmt.Errorf("invalid locktype")
		}
		h.LockType = lockType
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &orderlockcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
		}
		if conds.LockType != nil {
			switch conds.GetLockType().GetValue() {
			case uint32(basetypes.OrderLockType_LockBalance):
			case uint32(basetypes.OrderLockType_LockStock):
			case uint32(basetypes.OrderLockType_LockCommission):
			default:
				return fmt.Errorf("invalid locktype")
			}
			_type := conds.GetLockType().GetValue()
			h.Conds.LockType = &cruder.Cond{Op: conds.GetLockType().GetOp(), Val: basetypes.OrderLockType(_type)}
		}
		if conds.IDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: ids}
		}
		if conds.OrderIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetOrderIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.OrderIDs = &cruder.Cond{Op: conds.GetOrderIDs().GetOp(), Val: ids}
		}

		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}

//nolint:gocyclo
func WithReqs(reqs []*npool.OrderLockReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_reqs := []*orderlockcrud.Req{}
		for _, req := range reqs {
			if must {
				if req.AppID == nil {
					return fmt.Errorf("invalid appid")
				}
				if req.UserID == nil {
					return fmt.Errorf("invalid userid")
				}
				if req.OrderID == nil {
					return fmt.Errorf("invalid orderid")
				}
				if req.LockType == nil {
					return fmt.Errorf("invalid locktype")
				}
			}
			_req := &orderlockcrud.Req{}
			if req.ID != nil {
				id, err := uuid.Parse(*req.ID)
				if err != nil {
					return err
				}
				_req.ID = &id
			}
			if req.AppID != nil {
				id, err := uuid.Parse(*req.AppID)
				if err != nil {
					return err
				}
				_req.AppID = &id
			}
			if req.UserID != nil {
				id, err := uuid.Parse(*req.UserID)
				if err != nil {
					return err
				}
				_req.UserID = &id
			}
			if req.OrderID != nil {
				id, err := uuid.Parse(*req.OrderID)
				if err != nil {
					return err
				}
				_req.OrderID = &id
			}
			if req.LockType != nil {
				switch *req.LockType {
				case basetypes.OrderLockType_LockBalance:
				case basetypes.OrderLockType_LockStock:
				case basetypes.OrderLockType_LockCommission:
				default:
					return fmt.Errorf("invalid locktype")
				}
				_req.LockType = req.LockType
			}
			_reqs = append(_reqs, _req)
		}
		h.Reqs = _reqs
		return nil
	}
}
