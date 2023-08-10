package order

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                        *uuid.UUID
	AppID                     *uuid.UUID
	GoodID                    *uuid.UUID
	UserID                    *uuid.UUID
	ParentOrderID             *uuid.UUID
	PayWithParent             *bool
	Units                     *decimal.Decimal
	PromotionID               *uuid.UUID
	DiscountCouponID          *uuid.UUID
	UserSpecialReductionID    *uuid.UUID
	StartAt                   *uint32
	EndAt                     *uint32
	FixAmountCouponID         *uuid.UUID
	Type                      *basetypes.OrderType
	State                     *basetypes.OrderState
	CouponIDs                 *[]uuid.UUID
	LastBenefitAt             *uint32
	CreatedAt                 *uint32
	DeletedAt                 *uint32
	PaymentID                 *uuid.UUID
	PaymentAccountID          *uuid.UUID
	PaymentAccountStartAmount *decimal.Decimal
	PaymentAmount             *decimal.Decimal
	PayWithBalanceAmount      *decimal.Decimal
	PaymentCoinUSDCurrency    *decimal.Decimal
	PaymentLocalUSDCurrency   *decimal.Decimal
	PaymentLiveUSDCurrency    *decimal.Decimal
	PaymentCoinID             *uuid.UUID
	PaymentFinishAmount       *decimal.Decimal
	PaymentUserSetCanceled    *bool
	PaymentFakePayment        *bool
}

//nolint:gocyclo
func CreateSet(c *ent.OrderCreate, req *Req) *ent.OrderCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.ParentOrderID != nil {
		c.SetParentOrderID(*req.ParentOrderID)
	}
	if req.PayWithParent != nil {
		c.SetPayWithParent(*req.PayWithParent)
	}
	if req.Units != nil {
		c.SetUnitsV1(*req.Units)
	}
	if req.PromotionID != nil {
		c.SetPromotionID(*req.PromotionID)
	}
	if req.DiscountCouponID != nil {
		c.SetDiscountCouponID(*req.DiscountCouponID)
	}
	if req.UserSpecialReductionID != nil {
		c.SetUserSpecialReductionID(*req.UserSpecialReductionID)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.FixAmountCouponID != nil {
		c.SetFixAmountCouponID(*req.FixAmountCouponID)
	}
	if req.Type != nil {
		c.SetType(req.Type.String())
	}
	if req.State != nil {
		c.SetState(req.State.String())
	}
	if req.CouponIDs != nil {
		c.SetCouponIds(*req.CouponIDs)
	}
	if req.LastBenefitAt != nil {
		c.SetLastBenefitAt(*req.LastBenefitAt)
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.OrderUpdateOne, req *Req) *ent.OrderUpdateOne {
	if req.State != nil {
		u.SetState(req.State.String())
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
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID                     *cruder.Cond
	IDs                    *cruder.Cond
	AppID                  *cruder.Cond
	UserID                 *cruder.Cond
	GoodID                 *cruder.Cond
	Type                   *cruder.Cond
	State                  *cruder.Cond
	States                 *cruder.Cond
	FixAmountCouponID      *cruder.Cond
	DiscountCouponID       *cruder.Cond
	UserSpecialReductionID *cruder.Cond
	LastBenefitAt          *cruder.Cond
	CouponID               *cruder.Cond
	CouponIDs              *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderQuery, conds *Conds) (*ent.OrderQuery, error) {
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
			q.Where(entpayment.ID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
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
				q.Where(entpayment.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entpayment.AppID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entpayment.UserID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entpayment.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.FixAmountCouponID != nil {
		id, ok := conds.FixAmountCouponID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid fixamountcouponid")
		}
		switch conds.FixAmountCouponID.Op {
		case cruder.EQ:
			q.Where(entpayment.FixAmountCouponID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.DiscountCouponID != nil {
		id, ok := conds.DiscountCouponID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid discountcouponid")
		}
		switch conds.DiscountCouponID.Op {
		case cruder.EQ:
			q.Where(entpayment.DiscountCouponID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.UserSpecialReductionID != nil {
		id, ok := conds.UserSpecialReductionID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userspecialreductionid")
		}
		switch conds.UserSpecialReductionID.Op {
		case cruder.EQ:
			q.Where(entpayment.UserSpecialReductionID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.State != nil {
		state, ok := conds.State.Val.(basetypes.PaymentState)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(entpayment.State(state.String()))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.States != nil {
		states, ok := conds.States.Val.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid states")
		}
		if len(states) > 0 {
			switch conds.States.Op {
			case cruder.IN:
				q.Where(entpayment.StateIn(states...))
			default:
				return nil, fmt.Errorf("invalid payment field")
			}
		}
	}
	if conds.LastBenefitAt != nil {
		lastBenefitAt, ok := conds.LastBenefitAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid lastbenefitat")
		}
		switch conds.LastBenefitAt.Op {
		case cruder.EQ:
			q.Where(entpayment.LastBenefitAt(lastBenefitAt))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	q.Where(entpayment.DeletedAt(0))
	return q, nil
}