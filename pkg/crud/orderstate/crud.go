package orderstate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                   *uuid.UUID
	OrderID              *uuid.UUID
	OrderState           *basetypes.OrderState
	StartMode            *basetypes.OrderStartMode
	StartAt              *uint32
	EndAt                *uint32
	LastBenefitAt        *uint32
	BenefitState         *basetypes.BenefitState
	UserSetPaid          *bool
	UserSetCanceled      *bool
	AdminSetCanceled     *bool
	PaymentTransactionID *string
	PaymentFinishAmount  *decimal.Decimal
	PaymentState         *basetypes.PaymentState
	OutOfGasHours        *uint32
	CompensateHours      *uint32
	CreatedAt            *uint32
	DeletedAt            *uint32
}

//nolint:gocyclo
func CreateSet(c *ent.OrderStateCreate, req *Req) *ent.OrderStateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.OrderState != nil {
		c.SetOrderState(req.OrderState.String())
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
	if req.StartMode != nil {
		u.SetStartMode(req.StartMode.String())
	}
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
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
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
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
}

//nolint
func SetQueryConds(q *ent.OrderStateQuery, conds *Conds) (*ent.OrderStateQuery, error) {
	q.Where(entorderstate.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
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
		ids, ok := conds.IDs.Val.([]uuid.UUID)
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
		state, ok := conds.OrderState.Val.(basetypes.OrderState)
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
		startmode, ok := conds.StartMode.Val.(basetypes.OrderStartMode)
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
		state, ok := conds.BenefitState.Val.(basetypes.BenefitState)
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
		state, ok := conds.PaymentState.Val.(basetypes.PaymentState)
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
	return q, nil
}
