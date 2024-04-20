package powerrental

import (
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
