package orderstm

import (
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
)

var rollbacks = map[types.OrderState]types.OrderState{
	types.OrderState_OrderStatePaymentSpendBalance:         types.OrderState_OrderStatePaymentTransferBookKeeping,
	types.OrderState_OrderStateTransferGoodStockLocked:     types.OrderState_OrderStatePaymentSpendBalance,
	types.OrderState_OrderStateAchievementBookKeeping:      types.OrderState_OrderStateTransferGoodStockLocked,
	types.OrderState_OrderStateAddCommission:               types.OrderState_OrderStateAchievementBookKeeping,
	types.OrderState_OrderStatePaymentUnlockAccount:        types.OrderState_OrderStateAddCommission,
	types.OrderState_OrderStatePaid:                        types.OrderState_OrderStatePaymentUnlockAccount,
	types.OrderState_OrderStateInService:                   types.OrderState_OrderStateTransferGoodStockWaitStart,
	types.OrderState_OrderStateRestoreExpiredStock:         types.OrderState_OrderStatePreExpired,
	types.OrderState_OrderStateExpired:                     types.OrderState_OrderStateRestoreExpiredStock,
	types.OrderState_OrderStateDeductLockedCommission:      types.OrderState_OrderStateRestoreCanceledStock,
	types.OrderState_OrderStateCancelAchievement:           types.OrderState_OrderStateDeductLockedCommission,
	types.OrderState_OrderStateReturnCanceledBalance:       types.OrderState_OrderStateCancelAchievement,
	types.OrderState_OrderStateCanceledTransferBookKeeping: types.OrderState_OrderStateReturnCanceledBalance,
	types.OrderState_OrderStateCancelUnlockPaymentAccount:  types.OrderState_OrderStateCanceledTransferBookKeeping,
	types.OrderState_OrderStateCanceled:                    types.OrderState_OrderStateCancelUnlockPaymentAccount,
}

type rollbackHandler struct {
	*orderQueryHandler
}

func (h *rollbackHandler) rollback() (*types.OrderState, error) {
	state, ok := rollbacks[h._ent.OrderState()]
	if !ok {
		return nil, wlog.Errorf("invalid orderstate")
	}
	return &state, nil
}
