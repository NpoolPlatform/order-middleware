package paymentbase

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	paymentbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	paymentbasecrud.Req
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
			return err
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
			return err
		}
		h.OrderID = &_id
		return nil
	}
}

func WithObseleteState(e *types.PaymentObseleteState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid obseletestate")
			}
			return nil
		}
		switch *e {
		case types.PaymentObseleteState_PaymentObseleteNone:
		case types.PaymentObseleteState_PaymentObseleteWait:
		case types.PaymentObseleteState_PaymentObseleteUnlock:
		case types.PaymentObseleteState_PaymentObseleted:
		case types.PaymentObseleteState_PaymentObseleteFail:
		default:
			return wlog.Errorf("invalid obseletestate")
		}
		h.ObseleteState = e
		return nil
	}
}
