package orderstate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                *uuid.UUID
	OrderID              *uuid.UUID
	OrderState           *types.OrderState
	CancelState          *types.OrderState
	StartMode            *types.OrderStartMode
	StartAt              *uint32
	EndAt                *uint32
	PaidAt               *uint32
	LastBenefitAt        *uint32
	BenefitState         *types.BenefitState
	UserSetPaid          *bool
	UserSetCanceled      *bool
	AdminSetCanceled     *bool
	PaymentTransactionID *string
	PaymentFinishAmount  *decimal.Decimal
	PaymentState         *types.PaymentState
	OutOfGasHours        *uint32
	CompensateHours      *uint32
	RenewState           *types.OrderRenewState
	CreatedAt            *uint32
	DeletedAt            *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.OrderStateCreate, req *Req) *ent.OrderStateCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.OrderState != nil {
		c.SetOrderState(req.OrderState.String())
	}
	if req.CancelState != nil {
		c.SetCancelState(req.CancelState.String())
	}
	if req.StartMode != nil {
		c.SetStartMode(req.StartMode.String())
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.LastBenefitAt != nil {
		c.SetLastBenefitAt(*req.LastBenefitAt)
	}
	if req.BenefitState != nil {
		c.SetBenefitState(req.BenefitState.String())
	}
	if req.UserSetPaid != nil {
		c.SetUserSetPaid(*req.UserSetPaid)
	}
	if req.UserSetCanceled != nil {
		c.SetUserSetCanceled(*req.UserSetCanceled)
	}
	if req.AdminSetCanceled != nil {
		c.SetAdminSetCanceled(*req.AdminSetCanceled)
	}
	if req.PaymentTransactionID != nil {
		c.SetPaymentTransactionID(*req.PaymentTransactionID)
	}
	if req.PaymentFinishAmount != nil {
		c.SetPaymentFinishAmount(*req.PaymentFinishAmount)
	}
	if req.PaymentState != nil {
		c.SetPaymentState(req.PaymentState.String())
	}
	if req.OutOfGasHours != nil {
		c.SetOutofgasHours(*req.OutOfGasHours)
	}
	if req.CompensateHours != nil {
		c.SetCompensateHours(*req.CompensateHours)
	}
	if req.RenewState != nil {
		c.SetRenewState(req.RenewState.String())
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}

	return c
}

//nolint:gocyclo
func UpdateSet(u *ent.OrderStateUpdateOne, req *Req) *ent.OrderStateUpdateOne {
	if req.OrderState != nil {
		u.SetOrderState(req.OrderState.String())
	}
	if req.CancelState != nil {
		u.SetCancelState(req.CancelState.String())
	}
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.PaidAt != nil {
		u.SetPaidAt(*req.PaidAt)
	}
	if req.LastBenefitAt != nil {
		u.SetLastBenefitAt(*req.LastBenefitAt)
	}
	if req.BenefitState != nil {
		u.SetBenefitState(req.BenefitState.String())
	}
	if req.UserSetPaid != nil {
		u.SetUserSetPaid(*req.UserSetPaid)
	}
	if req.UserSetCanceled != nil {
		u.SetUserSetCanceled(*req.UserSetCanceled)
	}
	if req.AdminSetCanceled != nil {
		u.SetAdminSetCanceled(*req.AdminSetCanceled)
	}
	if req.PaymentTransactionID != nil {
		u.SetPaymentTransactionID(*req.PaymentTransactionID)
	}
	if req.PaymentFinishAmount != nil {
		u.SetPaymentFinishAmount(*req.PaymentFinishAmount)
	}
	if req.PaymentState != nil {
		u.SetPaymentState(req.PaymentState.String())
	}
	if req.OutOfGasHours != nil {
		u.SetOutofgasHours(*req.OutOfGasHours)
	}
	if req.CompensateHours != nil {
		u.SetCompensateHours(*req.CompensateHours)
	}
	if req.RenewState != nil {
		u.SetRenewState(req.RenewState.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID                *cruder.Cond
	EntIDs               *cruder.Cond
	ID                   *cruder.Cond
	IDs                  *cruder.Cond
	OrderID              *cruder.Cond
	OrderState           *cruder.Cond
	StartMode            *cruder.Cond
	LastBenefitAt        *cruder.Cond
	BenefitState         *cruder.Cond
	PaymentTransactionID *cruder.Cond
	PaymentState         *cruder.Cond
	OrderStates          *cruder.Cond
	RenewState           *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderStateQuery, conds *Conds) (*ent.OrderStateQuery, error) {
	q.Where(entorderstate.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entorderstate.EntID(id))
		case cruder.NEQ:
			q.Where(entorderstate.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		if len(ids) > 0 {
			switch conds.EntIDs.Op {
			case cruder.IN:
				q.Where(entorderstate.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entorderstate.ID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		if len(ids) > 0 {
			switch conds.IDs.Op {
			case cruder.IN:
				q.Where(entorderstate.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entorderstate.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.OrderState != nil {
		state, ok := conds.OrderState.Val.(types.OrderState)
		if !ok {
			return nil, fmt.Errorf("invalid orderstate")
		}
		switch conds.OrderState.Op {
		case cruder.EQ:
			q.Where(entorderstate.OrderState(state.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.StartMode != nil {
		startmode, ok := conds.StartMode.Val.(types.OrderStartMode)
		if !ok {
			return nil, fmt.Errorf("invalid startmode")
		}
		switch conds.StartMode.Op {
		case cruder.EQ:
			q.Where(entorderstate.StartMode(startmode.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.BenefitState != nil {
		state, ok := conds.BenefitState.Val.(types.BenefitState)
		if !ok {
			return nil, fmt.Errorf("invalid benefitstate")
		}
		switch conds.BenefitState.Op {
		case cruder.EQ:
			q.Where(entorderstate.BenefitState(state.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentState != nil {
		state, ok := conds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return nil, fmt.Errorf("invalid paymentstate")
		}
		switch conds.PaymentState.Op {
		case cruder.EQ:
			q.Where(entorderstate.PaymentState(state.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentTransactionID != nil {
		paymentTransactionID, ok := conds.PaymentTransactionID.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid paymenttransactionid")
		}
		switch conds.PaymentTransactionID.Op {
		case cruder.EQ:
			q.Where(entorderstate.PaymentTransactionID(paymentTransactionID))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.LastBenefitAt != nil {
		lastBenefitAt, ok := conds.LastBenefitAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid lastbenefitat")
		}
		switch conds.LastBenefitAt.Op {
		case cruder.EQ:
			q.Where(entorderstate.LastBenefitAt(lastBenefitAt))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.OrderStates != nil {
		states, ok := conds.OrderStates.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid orderstates")
		}
		if len(states) > 0 {
			switch conds.OrderStates.Op {
			case cruder.IN:
				q.Where(entorderstate.OrderStateIn(states...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.RenewState != nil {
		state, ok := conds.RenewState.Val.(types.OrderRenewState)
		if !ok {
			return nil, fmt.Errorf("invalid renewstate")
		}
		switch conds.RenewState.Op {
		case cruder.EQ:
			q.Where(entorderstate.RenewState(state.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	return q, nil
}
