package statebase

import (
	"context"
	"fmt"
	"time"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	statebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	statebasecrud.Req
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
				return fmt.Errorf("invalid id")
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
				return fmt.Errorf("invalid entid")
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

//nolint:gocyclo
func WithOrderState(state *types.OrderState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid orderstate")
			}
			return nil
		}
		switch *state {
		case types.OrderState_OrderStateCreated:
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaymentTransferReceived:
		case types.OrderState_OrderStatePaymentTransferBookKeeping:
		case types.OrderState_OrderStatePaymentSpendBalance:
		case types.OrderState_OrderStateTransferGoodStockLocked:
		case types.OrderState_OrderStateAddCommission:
		case types.OrderState_OrderStateAchievementBookKeeping:
		case types.OrderState_OrderStateUpdatePaidChilds:
		case types.OrderState_OrderStateChildPaidByParent:
		case types.OrderState_OrderStatePaymentUnlockAccount:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateTransferGoodStockWaitStart:
		case types.OrderState_OrderStateUpdateInServiceChilds:
		case types.OrderState_OrderStateChildInServiceByParent:
		case types.OrderState_OrderStateInService:
		case types.OrderState_OrderStatePaymentTimeout:
		case types.OrderState_OrderStatePreCancel:
		case types.OrderState_OrderStatePreExpired:
		case types.OrderState_OrderStateRestoreExpiredStock:
		case types.OrderState_OrderStateUpdateExpiredChilds:
		case types.OrderState_OrderStateChildExpiredByParent:
		case types.OrderState_OrderStateRestoreCanceledStock:
		case types.OrderState_OrderStateCancelAchievement:
		case types.OrderState_OrderStateDeductLockedCommission:
		case types.OrderState_OrderStateReturnCanceledBalance:
		case types.OrderState_OrderStateUpdateCanceledChilds:
		case types.OrderState_OrderStateChildCanceledByParent:
		case types.OrderState_OrderStateCanceledTransferBookKeeping:
		case types.OrderState_OrderStateCancelUnlockPaymentAccount:
		case types.OrderState_OrderStateCanceled:
		case types.OrderState_OrderStateExpired:
		default:
			return fmt.Errorf("invalid orderstate")
		}
		h.OrderState = state
		return nil
	}
}

func WithStartMode(startMode *types.OrderStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startMode == nil {
			if must {
				return fmt.Errorf("invalid startmode")
			}
			return nil
		}
		switch *startMode {
		case types.OrderStartMode_OrderStartConfirmed:
		case types.OrderStartMode_OrderStartTBD:
		case types.OrderStartMode_OrderStartInstantly:
		case types.OrderStartMode_OrderStartNextDay:
		case types.OrderStartMode_OrderStartPreset:
		default:
			return fmt.Errorf("invalid startmode")
		}
		h.StartMode = startMode
		return nil
	}
}

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startAt == nil {
			if must {
				return fmt.Errorf("invalid startat")
			}
			return nil
		}
		if *startAt < uint32(time.Now().Unix()) {
			return fmt.Errorf("invalid startat")
		}
		h.StartAt = startAt
		return nil
	}
}

func WithLastBenefitAt(lastBenefitAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if lastBenefitAt == nil {
			if must {
				return fmt.Errorf("invalid lastbenefitat")
			}
			return nil
		}
		h.LastBenefitAt = lastBenefitAt
		return nil
	}
}

func WithBenefitState(benefitState *types.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if benefitState == nil {
			if must {
				return fmt.Errorf("invalid benefitstate")
			}
			return nil
		}
		switch *benefitState {
		case types.BenefitState_BenefitWait:
		case types.BenefitState_BenefitCalculated:
		case types.BenefitState_BenefitBookKept:
		default:
			return fmt.Errorf("invalid benefitstate")
		}
		h.BenefitState = benefitState
		return nil
	}
}
