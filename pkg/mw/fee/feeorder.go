package feeorder

import (
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type FeeOrder interface {
	OrderID() uuid.UUID
	PaymentID() uuid.UUID
}

type feeOrder struct {
	entOrderBase        *ent.OrderBase
	entOrderStateBase   *ent.OrderStateBase
	entFeeOrder         *ent.FeeOrder
	entFeeOrderState    *ent.FeeOrderState
	entPaymentBase      *ent.PaymentBase
	entPaymentBalances  []*ent.PaymentBalance
	entPaymentTransfers []*ent.PaymentTransfer
	entLedgerLock       *ent.OrderLock
	entOrderCoupons     []*ent.OrderCoupon
}

func (f *feeOrder) OrderID() uuid.UUID {
	return f.entFeeOrder.OrderID
}

func (f *feeOrder) PaymentID() uuid.UUID {
	return f.entPaymentBase.EntID
}
