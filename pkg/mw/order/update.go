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

var stateMap = map[basetypes.OrderState][]basetypes.OrderState{
	basetypes.OrderState_OrderStateWaitPayment:                {basetypes.OrderState_OrderStateCheckPayment},
	basetypes.OrderState_OrderStateCheckPayment:               {basetypes.OrderState_OrderStatePaid, basetypes.OrderState_OrderStatePaymentTimeout},
	basetypes.OrderState_OrderStatePaid:                       {basetypes.OrderState_OrderStateInService, basetypes.OrderState_OrderStatePreCancel},
	basetypes.OrderState_OrderStateInService:                  {basetypes.OrderState_OrderStatePreExpired, basetypes.OrderState_OrderStatePreCancel},
	basetypes.OrderState_OrderStatePreExpired:                 {basetypes.OrderState_OrderStatePreExpiredCheck},
	basetypes.OrderState_OrderStatePreExpiredCheck:            {basetypes.OrderState_OrderStateRestoreExpiredStock},
	basetypes.OrderState_OrderStateRestoreExpiredStock:        {basetypes.OrderState_OrderStateRestoreExpiredStockCheck, basetypes.OrderState_OrderStateRestoreExpiredStock},
	basetypes.OrderState_OrderStateRestoreExpiredStockCheck:   {basetypes.OrderState_OrderStateExpired},
	basetypes.OrderState_OrderStateExpired:                    {basetypes.OrderState_OrderStateExpired},
	basetypes.OrderState_OrderStatePreCancel:                  {basetypes.OrderState_OrderStatePreCancelCheck},
	basetypes.OrderState_OrderStatePreCancelCheck:             {basetypes.OrderState_OrderStateRestoreCanceledStock},
	basetypes.OrderState_OrderStateRestoreCanceledStock:       {basetypes.OrderState_OrderStateRestoreCanceledStockCheck, basetypes.OrderState_OrderStateRestoreCanceledStock},
	basetypes.OrderState_OrderStateRestoreCanceledStockCheck:  {basetypes.OrderState_OrderStateReturnCanceledBalance},
	basetypes.OrderState_OrderStateReturnCanceledBalance:      {basetypes.OrderState_OrderStateReturnCanceledBalanceCheck, basetypes.OrderState_OrderStateReturnCanceledBalance},
	basetypes.OrderState_OrderStateReturnCanceledBalanceCheck: {basetypes.OrderState_OrderStateCanceled},
	basetypes.OrderState_OrderStateCanceled:                   {basetypes.OrderState_OrderStateCanceled},
	basetypes.OrderState_OrderStatePaymentTimeout:             {basetypes.OrderState_OrderStatePreCancel},
}

func (h *updateHandler) checkOrderState(oldState basetypes.OrderState) error {
	allowedStates := stateMap[oldState]
	for _, state := range allowedStates {
		if state == *h.OrderState {
			return nil
		}
	}
	return fmt.Errorf("invalid orderstate")
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
				req.OrderState = basetypes.OrderState_OrderStateReturnCanceledBalance.Enum()
			case basetypes.OrderState_OrderStateReturnCanceledBalance:
				req.OrderState = basetypes.OrderState_OrderStateRestoreCanceledStock.Enum()
			case basetypes.OrderState_OrderStateRestoreCanceledStock:
				req.OrderState = basetypes.OrderState_OrderStatePreCancel.Enum()
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
