package order

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
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

var stateAllowMap = map[types.OrderState][]types.OrderState{
	types.OrderState_OrderStateCreated:     {types.OrderState_OrderStateWaitPayment},
	types.OrderState_OrderStateWaitPayment: {types.OrderState_OrderStateCheckPayment},
	types.OrderState_OrderStateCheckPayment: {
		types.OrderState_OrderStatePaymentTransferReceived,
		types.OrderState_OrderStatePreCancel,
		types.OrderState_OrderStatePaymentTimeout,
	},
	types.OrderState_OrderStatePaymentTransferReceived:      {types.OrderState_OrderStatePaymentTransferReceivedCheck},
	types.OrderState_OrderStatePaymentTransferReceivedCheck: {types.OrderState_OrderStatePaymentTransferBookKept},
	types.OrderState_OrderStatePaymentTransferBookKept:      {types.OrderState_OrderStatePaymentTransferBookKeptCheck},
	types.OrderState_OrderStatePaymentTransferBookKeptCheck: {types.OrderState_OrderStatePaymentBalanceSpent},
	types.OrderState_OrderStatePaymentBalanceSpent:          {types.OrderState_OrderStatePaymentBalanceSpentCheck},
	types.OrderState_OrderStatePaymentBalanceSpentCheck:     {types.OrderState_OrderStateGoodStockTransferred},
	types.OrderState_OrderStateGoodStockTransferred:         {types.OrderState_OrderStateGoodStockTransferredCheck},
	types.OrderState_OrderStateGoodStockTransferredCheck:    {types.OrderState_OrderStateCommissionAdded},
	types.OrderState_OrderStateCommissionAdded:              {types.OrderState_OrderStateCommissionAddedCheck},
	types.OrderState_OrderStateCommissionAddedCheck:         {types.OrderState_OrderStateAchievementBookKept},
	types.OrderState_OrderStateAchievementBookKept:          {types.OrderState_OrderStateAchievementBookKeptCheck},
	types.OrderState_OrderStateAchievementBookKeptCheck:     {types.OrderState_OrderStatePaid},
	types.OrderState_OrderStatePaid:                         {types.OrderState_OrderStateInService, types.OrderState_OrderStatePreCancel},
	types.OrderState_OrderStateInService:                    {types.OrderState_OrderStatePreExpired, types.OrderState_OrderStatePreCancel},
	types.OrderState_OrderStatePreExpired:                   {types.OrderState_OrderStatePreExpiredCheck},
	types.OrderState_OrderStatePreExpiredCheck:              {types.OrderState_OrderStateRestoreExpiredStock},
	types.OrderState_OrderStateRestoreExpiredStock:          {types.OrderState_OrderStateRestoreExpiredStockCheck},
	types.OrderState_OrderStateRestoreExpiredStockCheck:     {types.OrderState_OrderStateExpired},
	types.OrderState_OrderStateExpired:                      {},
	types.OrderState_OrderStatePreCancel:                    {types.OrderState_OrderStatePreCancelCheck},
	types.OrderState_OrderStatePreCancelCheck:               {types.OrderState_OrderStateRestoreCanceledStock},
	types.OrderState_OrderStateRestoreCanceledStock:         {types.OrderState_OrderStateRestoreCanceledStockCheck},
	types.OrderState_OrderStateRestoreCanceledStockCheck:    {types.OrderState_OrderStateCancelAchievement},
	types.OrderState_OrderStateCancelAchievement:            {types.OrderState_OrderStateCancelAchievementCheck},
	types.OrderState_OrderStateCancelAchievementCheck:       {types.OrderState_OrderStateReturnCanceledBalance},
	types.OrderState_OrderStateReturnCanceledBalance:        {types.OrderState_OrderStateReturnCanceledBalanceCheck},
	types.OrderState_OrderStateReturnCanceledBalanceCheck:   {types.OrderState_OrderStateCanceled},
	types.OrderState_OrderStateCanceled:                     {},
	types.OrderState_OrderStatePaymentTimeout:               {types.OrderState_OrderStatePreCancel},
}

var stateRollbackMap = map[types.OrderState]*types.OrderState{
	types.OrderState_OrderStatePaymentTransferBookKept: types.OrderState_OrderStatePaymentTransferReceivedCheck.Enum(),
	types.OrderState_OrderStatePaymentBalanceSpent:     types.OrderState_OrderStatePaymentTransferBookKeptCheck.Enum(),
	types.OrderState_OrderStateGoodStockTransferred:    types.OrderState_OrderStatePaymentBalanceSpentCheck.Enum(),
	types.OrderState_OrderStateCommissionAdded:         types.OrderState_OrderStateGoodStockTransferredCheck.Enum(),
	types.OrderState_OrderStateAchievementBookKept:     types.OrderState_OrderStateCommissionAddedCheck.Enum(),
	types.OrderState_OrderStatePaid:                    types.OrderState_OrderStateAchievementBookKeptCheck.Enum(),
	types.OrderState_OrderStateRestoreExpiredStock:     types.OrderState_OrderStatePreExpiredCheck.Enum(),
	types.OrderState_OrderStateExpired:                 types.OrderState_OrderStateRestoreExpiredStockCheck.Enum(),
	types.OrderState_OrderStateRestoreCanceledStock:    types.OrderState_OrderStatePreCancelCheck.Enum(),
	types.OrderState_OrderStateCancelAchievement:       types.OrderState_OrderStateRestoreCanceledStockCheck.Enum(),
	types.OrderState_OrderStateReturnCanceledBalance:   types.OrderState_OrderStateCancelAchievementCheck.Enum(),
	types.OrderState_OrderStateCanceled:                types.OrderState_OrderStateReturnCanceledBalanceCheck.Enum(),
}

func (h *updateHandler) checkOrderState(oldState types.OrderState, req *orderstatecrud.Req) error {
	allowedStates, ok := stateAllowMap[oldState]
	if !ok {
		return fmt.Errorf("invalid orderstate")
	}
	for _, state := range allowedStates {
		if state == *req.OrderState {
			return nil
		}
	}
	return fmt.Errorf("invalid orderstate")
}

func (h *updateHandler) checkOrderStateRollback(req *orderstatecrud.Req) (*types.OrderState, error) {
	rollbackState, ok := stateRollbackMap[*req.OrderState]
	if !ok {
		return nil, fmt.Errorf("invalid orderstate")
	}
	if rollbackState == nil {
		return nil, fmt.Errorf("invalid orderstate")
	}
	return rollbackState, nil
}

//nolint:gocyclo
func (h *updateHandler) updateOrderState(ctx context.Context, tx *ent.Tx, req *orderstatecrud.Req, rollback bool) error {
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
	startMode := types.OrderStartMode(types.OrderStartMode_value[orderstate.StartMode])
	if req.StartAt != nil && startMode == types.OrderStartMode_OrderStartTBD {
		startAt = *req.StartAt
	}
	endAt := startAt + duration
	if req.EndAt != nil && startMode == types.OrderStartMode_OrderStartTBD {
		endAt = *req.EndAt
	}

	if (req.UserSetCanceled != nil && !*req.UserSetCanceled) ||
		(req.AdminSetCanceled != nil && !*req.AdminSetCanceled) {
		return fmt.Errorf("permission denied")
	}

	_orderType := types.OrderType(types.OrderType_value[order.OrderType])
	_orderState := types.OrderState(types.OrderState_value[orderstate.OrderState])

	if req.UserSetCanceled != nil && *req.UserSetCanceled {
		switch _orderType {
		case types.OrderType_Normal:
		default:
			return fmt.Errorf("permission denied")
		}
		switch _orderState {
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStateCheckPayment:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateInService:
		default:
			return fmt.Errorf("invalid cancel state")
		}
	}

	if req.AdminSetCanceled != nil && *req.AdminSetCanceled {
		if req.OrderState != nil {
			return fmt.Errorf("permission denied")
		}
		switch _orderType {
		case types.OrderType_Offline:
		case types.OrderType_Airdrop:
		default:
			return fmt.Errorf("permission denied")
		}
		switch _orderState {
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateInService:
		default:
			return fmt.Errorf("invalid cancel state")
		}
	}

	if req.OrderState != nil {
		if rollback && *req.OrderState == _orderState {
			rollbackOrderState, err := h.checkOrderStateRollback(req)
			if err != nil {
				return err
			}
			req.OrderState = rollbackOrderState
		} else {
			err := h.checkOrderState(_orderState, req)
			if err != nil {
				return err
			}
		}
	}

	if req.PaymentState != nil {
		switch orderstate.PaymentState {
		case types.PaymentState_PaymentStateWait.String():
			switch *req.PaymentState {
			case types.PaymentState_PaymentStateDone:
			case types.PaymentState_PaymentStateCanceled:
			case types.PaymentState_PaymentStateTimeout:
			default:
				return fmt.Errorf("permission denied")
			}
		case types.PaymentState_PaymentStateDone.String():
			fallthrough //nolint
		case types.PaymentState_PaymentStateCanceled.String():
			fallthrough //nolint
		case types.PaymentState_PaymentStateTimeout.String():
			fallthrough //nolint
		case types.PaymentState_PaymentStateNoPayment.String():
			return fmt.Errorf("permission denied")
		}
	}

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			OrderState:           req.OrderState,
			CancelState:          req.CancelState,
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
		if err := handler.updateOrderState(_ctx, tx, req.OrderStateReq, h.Rollback != nil && *h.Rollback); err != nil {
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
			if err := handler.updateOrderState(ctx, tx, req.OrderStateReq, false); err != nil {
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
