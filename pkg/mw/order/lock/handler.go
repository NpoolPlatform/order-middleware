package orderlock

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/lock"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	orderlockcrud.Req
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderID = &_id
		return nil
	}
}

func WithLockType(lockType *types.OrderLockType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lockType == nil {
			if must {
				return wlog.Errorf("invalid locktype")
			}
			return nil
		}
		switch *lockType {
		case types.OrderLockType_LockStock:
		case types.OrderLockType_LockBalance:
		case types.OrderLockType_LockCommission:
		default:
			return wlog.Errorf("invalid locktype")
		}
		h.LockType = lockType
		return nil
	}
}
