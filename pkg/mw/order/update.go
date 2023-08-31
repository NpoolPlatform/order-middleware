package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *updateHandler) checkOrderState(oldState basetypes.OrderState) error {
	switch oldState {
	case basetypes.OrderState_OrderStateWaitPayment:
		if *h.OrderState != basetypes.OrderState_OrderStateCheckPayment {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateCheckPayment:
		if *h.OrderState != basetypes.OrderState_OrderStatePaid && *h.OrderState != basetypes.OrderState_OrderStatePaymentTimeout {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStatePaid:
		if *h.OrderState != basetypes.OrderState_OrderStateInService && *h.OrderState != basetypes.OrderState_OrderStatePreCancel {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateInService:
		if *h.OrderState != basetypes.OrderState_OrderStatePreExpired && *h.OrderState != basetypes.OrderState_OrderStatePreCancel {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStatePreExpired:
		if *h.OrderState != basetypes.OrderState_OrderStateRestoreExpiredStock {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateRestoreExpiredStock:
		if *h.OrderState != basetypes.OrderState_OrderStateExpired && *h.OrderState != basetypes.OrderState_OrderStateRestoreExpiredStock {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStatePaymentTimeout:
		if *h.OrderState != basetypes.OrderState_OrderStatePreCancel {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStatePreCancel:
		if *h.OrderState != basetypes.OrderState_OrderStateRestoreCanceledStock {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateRestoreCanceledStock:
		if *h.OrderState != basetypes.OrderState_OrderStateReturnCalceledBalance {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateReturnCalceledBalance:
		if *h.OrderState != basetypes.OrderState_OrderStateCanceled && *h.OrderState != basetypes.OrderState_OrderStateReturnCalceledBalance {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateCanceled:
		if *h.OrderState != basetypes.OrderState_OrderStateCanceled {
			return fmt.Errorf("invalid orderstate")
		}
	case basetypes.OrderState_OrderStateExpired:
		if *h.OrderState != basetypes.OrderState_OrderStateExpired {
			return fmt.Errorf("invalid orderstate")
		}
	default:
		return fmt.Errorf("invalid orderstate")
	}
	return nil
}

//nolint:gocyclo
func (h *updateHandler) updateOrderState(ctx context.Context, tx *ent.Tx, req *orderstatecrud.Req) error {
	orderstate, err := tx.OrderState.
		Query().
		Where(
			entorderstate.OrderID(*req.OrderID),
			entorderstate.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	order, err := tx.Order.
		Query().
		Where(
			entorder.ID(*req.OrderID),
			entorder.DeletedAt(0),
		).
		Only(ctx)
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("invalid order")
	}

	duration := orderstate.EndAt - orderstate.StartAt
	startAt := orderstate.StartAt
	startMode := basetypes.OrderStartMode(basetypes.OrderStartMode_value[orderstate.StartMode])
	if req.StartAt != nil && startMode == basetypes.OrderStartMode_OrderStartTBD {
		startAt = *req.StartAt
	}
	endAt := startAt + duration
	if req.EndAt != nil && startMode == basetypes.OrderStartMode_OrderStartTBD {
		endAt = *req.EndAt
	}

	if orderstate.PaymentState != basetypes.PaymentState_PaymentStateWait.String() &&
		order.OrderType == basetypes.OrderType_Normal.String() {
		if req.UserSetCanceled != nil && *req.UserSetCanceled {
			return fmt.Errorf("not wait payment")
		}
	}

	if h.OrderState != nil {
		_orderState := basetypes.OrderState(basetypes.OrderState_value[orderstate.OrderState])
		if err := h.checkOrderState(_orderState); err != nil {
			return err
		}

		if h.Rollback != nil && *h.Rollback && *h.OrderState == _orderState {
			switch _orderState {
			case basetypes.OrderState_OrderStateExpired:
				req.OrderState = basetypes.OrderState_OrderStateRestoreExpiredStock.Enum()
			case basetypes.OrderState_OrderStateRestoreExpiredStock:
				req.OrderState = basetypes.OrderState_OrderStatePreExpired.Enum()
			case basetypes.OrderState_OrderStateCanceled:
				req.OrderState = basetypes.OrderState_OrderStateReturnCalceledBalance.Enum()
			case basetypes.OrderState_OrderStateReturnCalceledBalance:
				req.OrderState = basetypes.OrderState_OrderStateRestoreCanceledStock.Enum()
			default:
				return fmt.Errorf("invalid orderstate")
			}
		}
	}

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			OrderState:           req.OrderState,
			StartMode:            req.StartMode,
			StartAt:              &startAt,
			EndAt:                &endAt,
			LastBenefitAt:        req.LastBenefitAt,
			BenefitState:         req.BenefitState,
			UserSetPaid:          req.UserSetPaid,
			UserSetCanceled:      req.UserSetCanceled,
			AdminSetCanceled:     req.AdminSetCanceled,
			PaymentTransactionID: req.PaymentTransactionID,
			PaymentFinishAmount:  req.PaymentFinishAmount,
			PaymentState:         req.PaymentState,
			OutOfGasHours:        req.OutOfGasHours,
			CompensateHours:      req.CompensateHours,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdateOrder(ctx context.Context) (*npool.Order, error) {
	req, err := h.ToOrderReq(ctx, false)
	if err != nil {
		return nil, err
	}
	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateOrderState(_ctx, tx, req.OrderStateReq); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOrder(ctx)
}

func (h *Handler) UpdateOrders(ctx context.Context) ([]*npool.Order, error) {
	handler := &updateHandler{
		Handler: h,
	}

	ids := []uuid.UUID{}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if req.OrderStateReq.OrderID == nil {
				return fmt.Errorf("invalid id")
			}
			if err := handler.updateOrderState(ctx, tx, req.OrderStateReq); err != nil {
				return err
			}
			ids = append(ids, *req.ID)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	h.Offset = 0
	h.Limit = int32(len(ids))

	infos, _, err := h.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	return infos, err
}
