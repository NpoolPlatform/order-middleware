package order

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
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
	parentCanceled    bool
	parentCancelState types.OrderState
}

var stateAllowMap = map[types.OrderState][]types.OrderState{
	types.OrderState_OrderStateCreated: {types.OrderState_OrderStateWaitPayment},
	types.OrderState_OrderStateWaitPayment: {
		types.OrderState_OrderStatePaymentTransferReceived,
		types.OrderState_OrderStatePreCancel,
		types.OrderState_OrderStatePaymentTimeout,
	},
	types.OrderState_OrderStatePaymentTransferReceived:     {types.OrderState_OrderStatePaymentTransferBookKeeping},
	types.OrderState_OrderStatePaymentTransferBookKeeping:  {types.OrderState_OrderStatePaymentSpendBalance},
	types.OrderState_OrderStatePaymentSpendBalance:         {types.OrderState_OrderStateTransferGoodStockLocked},
	types.OrderState_OrderStateTransferGoodStockLocked:     {types.OrderState_OrderStateAddCommission},
	types.OrderState_OrderStateAddCommission:               {types.OrderState_OrderStateAchievementBookKeeping},
	types.OrderState_OrderStateAchievementBookKeeping:      {types.OrderState_OrderStatePaymentUnlockAccount},
	types.OrderState_OrderStatePaymentUnlockAccount:        {types.OrderState_OrderStateUpdatePaidChilds},
	types.OrderState_OrderStateUpdatePaidChilds:            {types.OrderState_OrderStateChildPaidByParent, types.OrderState_OrderStatePaid},
	types.OrderState_OrderStateChildPaidByParent:           {types.OrderState_OrderStatePaid},
	types.OrderState_OrderStatePaid:                        {types.OrderState_OrderStateTransferGoodStockWaitStart, types.OrderState_OrderStatePreCancel},
	types.OrderState_OrderStateTransferGoodStockWaitStart:  {types.OrderState_OrderStateUpdateInServiceChilds},
	types.OrderState_OrderStateUpdateInServiceChilds:       {types.OrderState_OrderStateChildInServiceByParent, types.OrderState_OrderStateInService},
	types.OrderState_OrderStateChildInServiceByParent:      {types.OrderState_OrderStateInService},
	types.OrderState_OrderStateInService:                   {types.OrderState_OrderStatePreExpired, types.OrderState_OrderStatePreCancel},
	types.OrderState_OrderStatePreExpired:                  {types.OrderState_OrderStateRestoreExpiredStock},
	types.OrderState_OrderStateRestoreExpiredStock:         {types.OrderState_OrderStateUpdateExpiredChilds},
	types.OrderState_OrderStateUpdateExpiredChilds:         {types.OrderState_OrderStateChildExpiredByParent, types.OrderState_OrderStateExpired},
	types.OrderState_OrderStateChildExpiredByParent:        {types.OrderState_OrderStateExpired},
	types.OrderState_OrderStateExpired:                     {},
	types.OrderState_OrderStatePreCancel:                   {types.OrderState_OrderStateRestoreCanceledStock},
	types.OrderState_OrderStateRestoreCanceledStock:        {types.OrderState_OrderStateDeductLockedCommission},
	types.OrderState_OrderStateDeductLockedCommission:      {types.OrderState_OrderStateCancelAchievement},
	types.OrderState_OrderStateCancelAchievement:           {types.OrderState_OrderStateReturnCanceledBalance},
	types.OrderState_OrderStateReturnCanceledBalance:       {types.OrderState_OrderStateCanceledTransferBookKeeping},
	types.OrderState_OrderStateCanceledTransferBookKeeping: {types.OrderState_OrderStateCancelUnlockPaymentAccount},
	types.OrderState_OrderStateCancelUnlockPaymentAccount:  {types.OrderState_OrderStateUpdateCanceledChilds},
	types.OrderState_OrderStateUpdateCanceledChilds:        {types.OrderState_OrderStateChildCanceledByParent, types.OrderState_OrderStateCanceled},
	types.OrderState_OrderStateChildCanceledByParent:       {types.OrderState_OrderStateCanceled},
	types.OrderState_OrderStateCanceled:                    {},
	types.OrderState_OrderStatePaymentTimeout:              {types.OrderState_OrderStatePreCancel},
}

var stateRollbackMap = map[types.OrderState]*types.OrderState{
	types.OrderState_OrderStatePaymentSpendBalance:         types.OrderState_OrderStatePaymentTransferBookKeeping.Enum(),
	types.OrderState_OrderStateTransferGoodStockLocked:     types.OrderState_OrderStatePaymentSpendBalance.Enum(),
	types.OrderState_OrderStateAddCommission:               types.OrderState_OrderStateTransferGoodStockLocked.Enum(),
	types.OrderState_OrderStateAchievementBookKeeping:      types.OrderState_OrderStateAddCommission.Enum(),
	types.OrderState_OrderStatePaymentUnlockAccount:        types.OrderState_OrderStateAchievementBookKeeping.Enum(),
	types.OrderState_OrderStateUpdatePaidChilds:            types.OrderState_OrderStatePaymentUnlockAccount.Enum(),
	types.OrderState_OrderStateChildPaidByParent:           types.OrderState_OrderStateUpdatePaidChilds.Enum(),
	types.OrderState_OrderStatePaid:                        types.OrderState_OrderStateChildPaidByParent.Enum(),
	types.OrderState_OrderStateUpdateInServiceChilds:       types.OrderState_OrderStateTransferGoodStockWaitStart.Enum(),
	types.OrderState_OrderStateChildInServiceByParent:      types.OrderState_OrderStateUpdateInServiceChilds.Enum(),
	types.OrderState_OrderStateInService:                   types.OrderState_OrderStateChildInServiceByParent.Enum(),
	types.OrderState_OrderStateRestoreExpiredStock:         types.OrderState_OrderStatePreExpired.Enum(),
	types.OrderState_OrderStateUpdateExpiredChilds:         types.OrderState_OrderStateRestoreExpiredStock.Enum(),
	types.OrderState_OrderStateChildExpiredByParent:        types.OrderState_OrderStateUpdateExpiredChilds.Enum(),
	types.OrderState_OrderStateExpired:                     types.OrderState_OrderStateChildExpiredByParent.Enum(),
	types.OrderState_OrderStateDeductLockedCommission:      types.OrderState_OrderStateRestoreCanceledStock.Enum(),
	types.OrderState_OrderStateCancelAchievement:           types.OrderState_OrderStateDeductLockedCommission.Enum(),
	types.OrderState_OrderStateReturnCanceledBalance:       types.OrderState_OrderStateCancelAchievement.Enum(),
	types.OrderState_OrderStateCanceledTransferBookKeeping: types.OrderState_OrderStateReturnCanceledBalance.Enum(),
	types.OrderState_OrderStateCancelUnlockPaymentAccount:  types.OrderState_OrderStateCanceledTransferBookKeeping.Enum(),
	types.OrderState_OrderStateUpdateCanceledChilds:        types.OrderState_OrderStateCancelUnlockPaymentAccount.Enum(),
	types.OrderState_OrderStateChildCanceledByParent:       types.OrderState_OrderStateUpdateCanceledChilds.Enum(),
	types.OrderState_OrderStateCanceled:                    types.OrderState_OrderStateChildCanceledByParent.Enum(),
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

//nolint:funlen,gocyclo
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

	if req.PaymentTransactionID != nil {
		exist, err := tx.OrderState.
			Query().
			Where(
				entorderstate.PaymentTransactionID(*req.PaymentTransactionID),
				entorderstate.IDNEQ(orderstate.ID),
				entorderstate.DeletedAt(0),
			).
			Exist(ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("invalid paymenttransactionid")
		}
	}

	order, err := tx.Order.
		Query().
		Where(
			entorder.EntID(*req.OrderID),
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
		(req.AdminSetCanceled != nil && !*req.AdminSetCanceled) ||
		(req.UserSetPaid != nil && !*req.UserSetPaid) {
		return fmt.Errorf("permission denied")
	}

	if (orderstate.AdminSetCanceled || orderstate.UserSetCanceled) &&
		((req.UserSetCanceled != nil && *req.UserSetCanceled) ||
			(req.AdminSetCanceled != nil && *req.AdminSetCanceled)) {
		return fmt.Errorf("permission denied")
	}

	_orderType := types.OrderType(types.OrderType_value[order.OrderType])
	_orderState := types.OrderState(types.OrderState_value[orderstate.OrderState])
	_cancelState := types.OrderState(types.OrderState_value[orderstate.CancelState])

	switch _orderState {
	case types.OrderState_OrderStateExpired:
		fallthrough //nolint
	case types.OrderState_OrderStateCanceled:
		if !rollback {
			return fmt.Errorf("permission denied")
		}
	}

	if req.UserSetCanceled != nil && *req.UserSetCanceled {
		if order.PaymentType == types.PaymentType_PayWithParentOrder.String() {
			return fmt.Errorf("permission denied")
		}
		switch _orderType {
		case types.OrderType_Normal:
		default:
			return fmt.Errorf("permission denied")
		}
		switch _orderState {
		case types.OrderState_OrderStateWaitPayment:
		case types.OrderState_OrderStatePaid:
		case types.OrderState_OrderStateInService:
		default:
			return fmt.Errorf("invalid cancelstate")
		}
	}

	if req.AdminSetCanceled != nil && *req.AdminSetCanceled {
		if !h.parentCanceled && h.parentCancelState != types.OrderState_OrderStatePaymentTimeout {
			if order.PaymentType == types.PaymentType_PayWithParentOrder.String() {
				return fmt.Errorf("permission denied")
			}
			if req.OrderState != nil {
				return fmt.Errorf("permission denied")
			}
		} else {
			if order.PaymentType == types.PaymentType_PayWithParentOrder.String() {
				req.CancelState = &h.parentCancelState
			} else {
				req.CancelState = &_orderState
			}
		}
		switch _orderType {
		case types.OrderType_Offline:
		case types.OrderType_Airdrop:
		case types.OrderType_Normal:
		default:
			return fmt.Errorf("permission denied")
		}
		if !h.parentCanceled && h.parentCancelState != types.OrderState_OrderStatePaymentTimeout {
			switch _orderState {
			case types.OrderState_OrderStatePaid:
			case types.OrderState_OrderStateInService:
			default:
				return fmt.Errorf("invalid cancelstate")
			}
		}
	}

	if req.UserSetPaid != nil && *req.UserSetPaid {
		switch _orderType {
		case types.OrderType_Normal:
		default:
			return fmt.Errorf("permissioned denied")
		}
	}

	if req.OrderState != nil {
		if rollback && *req.OrderState == _orderState {
			if order.PaymentType != types.PaymentType_PayWithParentOrder.String() {
				rollbackOrderState, err := h.checkOrderStateRollback(req)
				if err != nil {
					return err
				}
				req.OrderState = rollbackOrderState
			}
		} else {
			if *req.OrderState == types.OrderState_OrderStatePreCancel {
				if _cancelState != types.OrderState_DefaultOrderState {
					return fmt.Errorf("permission denied")
				}
				switch _orderState {
				case types.OrderState_OrderStateWaitPayment:
				case types.OrderState_OrderStatePaymentTimeout:
				case types.OrderState_OrderStatePaid:
				case types.OrderState_OrderStateInService:
				default:
					return fmt.Errorf("permission denied")
				}
			}
			if order.PaymentType != types.PaymentType_PayWithParentOrder.String() {
				err := h.checkOrderState(_orderState, req)
				if err != nil {
					return err
				}
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

	if req.PaymentState != nil && *req.PaymentState == types.PaymentState_PaymentStateDone {
		now := uint32(time.Now().Unix())
		req.PaidAt = &now
	}

	req.StartAt = &startAt
	req.EndAt = &endAt
	if req.OrderState != nil && *req.OrderState == types.OrderState_OrderStatePreCancel {
		req.CancelState = &_orderState
	}
	if _, err := orderstatecrud.UpdateSet(orderstate.Update(), req).Save(ctx); err != nil {
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

	info, err := h.GetOrder(ctx)
	if err != nil {
		logger.Sugar().Warnw(
			"CreateOrders",
			"ID", *h.ID,
			"Error", err,
		)
	}
	return info, nil
}

func (h *updateHandler) checkOrders(ctx context.Context) error {
	reqs := []*OrderReq{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return fmt.Errorf("invalid id")
		}

		handler, err := NewHandler(ctx)
		if err != nil {
			return err
		}
		handler.ID = req.ID
		info, err := handler.GetOrder(ctx)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("invalid order")
		}
		// reset req
		orderID := uuid.MustParse(info.EntID)
		req.EntID = &orderID
		req.OrderStateReq.OrderID = &orderID
		reqs = append(reqs, req)
	}
	h.Reqs = reqs
	return nil
}

//nolint
func (h *updateHandler) checkChildOrderStates(ctx context.Context) error {
	orderIDs := []uint32{}
	updateState := false
	orderMap := map[uint32]*npool.Order{}

	for _, req := range h.Reqs {
		orderIDs = append(orderIDs, *req.ID)
		if req.OrderStateReq.OrderState != nil {
			updateState = true
		}
	}
	if !updateState {
		return nil
	}
	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: orderIDs},
	}
	h.Offset = 0
	h.Limit = int32(len(orderIDs))
	orders, _, err := h.GetOrders(ctx)
	if err != nil {
		return err
	}
	if len(orders) == 1 && orders[0].PaymentType != types.PaymentType_PayWithParentOrder {
		return nil
	}

	parentOrderID1 := uuid.Nil.String() // Child's parent ID
	parentOrderID2 := uuid.Nil.String() // Parent ID
	var parentOrder *npool.Order

	for _, order := range orders {
		orderMap[order.ID] = order
		if order.ParentOrderID != uuid.Nil.String() {
			continue
		}
		if order.EntID != uuid.Nil.String() && parentOrderID2 != uuid.Nil.String() {
			return fmt.Errorf("invalid parentorderid")
		}
		parentOrderID2 = order.EntID
		parentOrder = order
	}
	for _, order := range orders {
		if order.ParentOrderID == uuid.Nil.String() {
			continue
		}
		if parentOrderID1 != uuid.Nil.String() && order.ParentOrderID != parentOrderID1 {
			return fmt.Errorf("invalid parentorderid")
		}
		parentOrderID1 = order.ParentOrderID
	}
	if parentOrderID1 == uuid.Nil.String() {
		return fmt.Errorf("invalid parentorderid")
	}
	if parentOrderID2 != uuid.Nil.String() && parentOrderID2 != parentOrderID1 {
		return fmt.Errorf("invalid parentorderid")
	}
	if parentOrder == nil {
		id := uuid.MustParse(parentOrderID1)
		h.EntID = &id
		parentOrder, err = h.GetOrder(ctx)
		if err != nil {
			return err
		}
		if parentOrder == nil {
			return fmt.Errorf("invalid order")
		}
	}
	h.parentCanceled = parentOrder.AdminSetCanceled || parentOrder.UserSetCanceled
	h.parentCancelState = parentOrder.CancelState
	for _, req := range h.Reqs {
		if req.EntID.String() == parentOrderID1 {
			continue
		}
		order, ok := orderMap[*req.ID]
		if !ok {
			return fmt.Errorf("invalid order")
		}
		if order.PaymentType != types.PaymentType_PayWithParentOrder {
			return fmt.Errorf("invalid paymenttype")
		}
		if req.OrderStateReq.OrderState == nil {
			continue
		}
	}

	return nil
}

func (h *Handler) UpdateOrders(ctx context.Context) ([]*npool.Order, error) {
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.checkOrders(ctx); err != nil {
		return nil, err
	}

	if err := handler.checkChildOrderStates(ctx); err != nil {
		return nil, err
	}

	ids := []uint32{}
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
		logger.Sugar().Warnw(
			"GetOrders",
			"IDs", ids,
			"Error", err,
		)
	}

	return infos, nil
}
