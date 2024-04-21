package orderstm

import (
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
)

var forwards = map[types.OrderState][]types.OrderState{
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

type forwardHandler struct {
	*orderQueryHandler
}

func (h *forwardHandler) forward() (*types.OrderState, error) {
	states, ok := forwards[h._ent.OrderState()]
	if !ok {
		return nil, fmt.Errorf("invalid orderstate")
	}
	for _, state := range states {
		if state == *h.OrderState {
			return &state, nil
		}
	}
	return nil, fmt.Errorf("invalid orderstate")
}
