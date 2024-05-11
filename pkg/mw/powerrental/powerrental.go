package powerrental

import (
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type PowerRental interface {
	OrderID() uuid.UUID
	PaymentID() uuid.UUID
	LedgerLockID() uuid.UUID
}

type powerRental struct {
	entOrderBase        *ent.OrderBase
	entOrderStateBase   *ent.OrderStateBase
	entPowerRental      *ent.PowerRental
	entPowerRentalState *ent.PowerRentalState
	entPaymentBase      *ent.PaymentBase
	entPaymentBalances  []*ent.PaymentBalance
	entPaymentTransfers []*ent.PaymentTransfer
	entLedgerLock       *ent.OrderLock
	entStockLock        *ent.OrderLock
	entOrderCoupons     []*ent.OrderCoupon
	payWithMeOrderIDs   []uuid.UUID
}

func (f *powerRental) OrderID() uuid.UUID {
	return f.entPowerRental.OrderID
}

func (f *powerRental) PaymentID() uuid.UUID {
	return f.entPaymentBase.EntID
}

func (f *powerRental) Exist() bool {
	return f.entOrderBase != nil
}

func (f *powerRental) LedgerLockID() uuid.UUID {
	return f.entLedgerLock.EntID
}

func (f *powerRental) OrderBaseID() uint32 {
	return f.entOrderBase.ID
}

func (f *powerRental) OrderStateBaseID() uint32 {
	return f.entOrderStateBase.ID
}

func (f *powerRental) PowerRentalID() uint32 {
	return f.entPowerRental.ID
}

func (f *powerRental) PowerRentalStateID() uint32 {
	return f.entPowerRentalState.ID
}

func (f *powerRental) PaymentState() types.PaymentState {
	return types.PaymentState(types.PaymentState_value[f.entPowerRentalState.PaymentState])
}

func (f *powerRental) OrderState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entOrderStateBase.OrderState])
}

func (f *powerRental) CancelState() types.OrderState {
	return types.OrderState(types.OrderState_value[f.entPowerRentalState.CancelState])
}

func (f *powerRental) UserSetCanceled() bool {
	return f.entPowerRentalState.UserSetCanceled
}

func (f *powerRental) AdminSetCanceled() bool {
	return f.entPowerRentalState.AdminSetCanceled
}

func (f *powerRental) PayWithMeOrderIDs() []uuid.UUID {
	return f.payWithMeOrderIDs
}
