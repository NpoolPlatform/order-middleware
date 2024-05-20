package payment

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	paymentbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	obseleteState types.PaymentObseleteState
}

func (h *updateHandler) validateObseleteState() error {
	if h.ObseleteState == nil {
		return nil
	}
	switch h.obseleteState {
	case types.PaymentObseleteState_PaymentObseleteNone:
		return wlog.Errorf("permission denied")
	case types.PaymentObseleteState_PaymentObseleteWait:
		if *h.ObseleteState == types.PaymentObseleteState_PaymentObseleteUnlock {
			return wlog.Errorf("permission denied")
		}
	case types.PaymentObseleteState_PaymentObseleteUnlock:
		switch *h.ObseleteState {
		case types.PaymentObseleteState_PaymentObseleted:
			fallthrough //nolint
		case types.PaymentObseleteState_PaymentObseleteFail:
		default:
			return wlog.Errorf("permission denied")
		}
	case types.PaymentObseleteState_PaymentObseleted:
		return wlog.Errorf("permission denied")
	case types.PaymentObseleteState_PaymentObseleteFail:
		return wlog.Errorf("permission denied")
	}
	return nil
}

func (h *Handler) UpdatePayment(ctx context.Context) error {
	info, err := h.GetPayment(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid payment")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.validateObseleteState(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_, err := paymentbasecrud.UpdateSet(
			cli.PaymentBase.UpdateOneID(*h.ID),
			&h.Req,
		).Save(_ctx)
		return wlog.WrapError(err)
	})
}
