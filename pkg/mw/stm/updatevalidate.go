package orderstm

import (
	"context"
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
)

type validateHandler struct {
	*Handler
	*orderQueryHandler
	*rollbackHandler
	*forwardHandler
}

func (h *validateHandler) validateFinaled() error {
	switch h._ent.OrderState() {
	case types.OrderState_OrderStateExpired:
		fallthrough //nolint
	case types.OrderState_OrderStateCanceled:
		if h.Rollback == nil || !*h.Rollback {
			return fmt.Errorf("permission denied")
		}
	}
	return nil
}

func (h *validateHandler) validateUserCancelable() error {
	if h.UserSetCanceled == nil {
		return nil
	}
	if !*h.UserSetCanceled {
		return fmt.Errorf("permission denied")
	}
	if *h.UserCanceled {
		return fmt.Errorf("permission denied")
	}
	if h._ent.PaymentType() == types.PaymentType_PayWithParentOrder {
		return fmt.Errorf("permission denied")
	}
	switch h._ent.OrderType() {
	case types.OrderType_Normal:
	default:
		return fmt.Errorf("permission denied")
	}
	switch h._ent.OrderState() {
	case types.OrderState_OrderStateWaitPayment:
	case types.OrderState_OrderStatePaid:
	case types.OrderState_OrderStateInService:
	default:
		return fmt.Errorf("invalid cancelstate")
	}
	return nil
}

func (h *validateHandler) validateAdminCancelable() error {
	if h.AdminSetCanceled == nil {
		return nil
	}
	if !*h.AdminSetCanceled {
		return fmt.Errorf("invalid cancelstate")
	}
	if !*h.AdminCanceled {
		return fmt.Errorf("invalid cancelstate")
	}
	switch h._ent.OrderType() {
	case types.OrderType_Offline:
	case types.OrderType_Airdrop:
	case types.OrderType_Normal:
	default:
		return fmt.Errorf("permission denied")
	}
	return nil
}

func (h *validateHandler) validateUserPayable() error {
	if h.UserSetPaid == nil {
		return nil
	}
	if !*h.UserSetPaid {
		return fmt.Errorf("permission denied")
	}
	switch h._ent.OrderType() {
	case types.OrderType_Normal:
	default:
		return fmt.Errorf("permissioned denied")
	}
	return nil
}

func (h *validateHandler) validatePaymentState() error {
	if h.NewPaymentState == nil {
		return nil
	}
	switch *h.CurrentPaymentState {
	case types.PaymentState_PaymentStateWait:
		switch *h.NewPaymentState {
		case types.PaymentState_PaymentStateDone:
		case types.PaymentState_PaymentStateCanceled:
		case types.PaymentState_PaymentStateTimeout:
		default:
			return fmt.Errorf("permission denied")
		}
	case types.PaymentState_PaymentStateDone:
		fallthrough //nolint
	case types.PaymentState_PaymentStateCanceled:
		fallthrough //nolint
	case types.PaymentState_PaymentStateTimeout:
		fallthrough //nolint
	case types.PaymentState_PaymentStateNoPayment:
		return fmt.Errorf("permission denied")
	}
	return nil
}

func (h *validateHandler) validateOrderState() (*types.OrderState, error) {
	if h.OrderState == nil {
		return nil, nil
	}
	if h.Rollback != nil && *h.Rollback && *h.OrderState == h._ent.OrderState() {
		return h.rollback()
	}
	return h.forward()
}

func (h *Handler) ValidateUpdateForNewState(ctx context.Context) (*types.OrderState, error) {
	handler := &validateHandler{
		Handler: h,
		orderQueryHandler: &orderQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireOrder(ctx); err != nil {
		return nil, err
	}
	handler.rollbackHandler = &rollbackHandler{
		orderQueryHandler: handler.orderQueryHandler,
	}
	handler.forwardHandler = &forwardHandler{
		orderQueryHandler: handler.orderQueryHandler,
	}

	if err := handler.validateFinaled(); err != nil {
		return nil, err
	}
	if err := handler.validateUserCancelable(); err != nil {
		return nil, err
	}
	if err := handler.validateAdminCancelable(); err != nil {
		return nil, err
	}
	if err := handler.validateUserPayable(); err != nil {
		return nil, err
	}
	if err := handler.validatePaymentState(); err != nil {
		return nil, err
	}
	return handler.validateOrderState()
}
