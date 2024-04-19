package powerrental

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpowerrental "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrental"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                *uint32
	EntID             *uuid.UUID
	OrderID           *uuid.UUID
	AppGoodStockID    *uuid.UUID
	Units             *decimal.Decimal
	GoodValueUSD      *decimal.Decimal
	PaymentAmountUSD  *decimal.Decimal
	DiscountAmountUSD *decimal.Decimal
	PromotionID       *uuid.UUID
	Duration          *uint32
	InvestmentType    *types.InvestmentType
	DeletedAt         *uint32
}

func CreateSet(c *ent.PowerRentalCreate, req *Req) *ent.PowerRentalCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.AppGoodStockID != nil {
		c.SetAppGoodStockID(*req.AppGoodStockID)
	}
	if req.Units != nil {
		c.SetUnits(*req.Units)
	}
	if req.GoodValueUSD != nil {
		c.SetGoodValueUsd(*req.GoodValueUSD)
	}
	if req.PaymentAmountUSD != nil {
		c.SetPaymentAmountUsd(*req.PaymentAmountUSD)
	}
	if req.DiscountAmountUSD != nil {
		c.SetDiscountAmountUsd(*req.DiscountAmountUSD)
	}
	if req.PromotionID != nil {
		c.SetPromotionID(*req.PromotionID)
	}
	if req.Duration != nil {
		c.SetDuration(*req.Duration)
	}
	if req.InvestmentType != nil {
		c.SetInvestmentType(req.InvestmentType.String())
	}
	return c
}

func UpdateSet(u *ent.PowerRentalUpdateOne, req *Req) *ent.PowerRentalUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID             *cruder.Cond
	IDs            *cruder.Cond
	EntID          *cruder.Cond
	EntIDs         *cruder.Cond
	OrderID        *cruder.Cond
	OrderIDs       *cruder.Cond
	InvestmentType *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.PowerRentalQuery, conds *Conds) (*ent.PowerRentalQuery, error) {
	q.Where(entpowerrental.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.ID(id))
		default:
			return nil, fmt.Errorf("invalid powerrental field")
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
				q.Where(entpowerrental.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid powerrental field")
			}
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entpowerrental.EntID(id))
		case cruder.NEQ:
			q.Where(entpowerrental.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid powerrental field")
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
				q.Where(entpowerrental.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid powerrental field")
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
			q.Where(entpowerrental.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid powerrental field")
		}
	}
	if conds.OrderIDs != nil {
		ids, ok := conds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderids")
		}
		if len(ids) > 0 {
			switch conds.OrderIDs.Op {
			case cruder.IN:
				q.Where(entpowerrental.OrderIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid powerrental field")
			}
		}
	}
	if conds.InvestmentType != nil {
		_type, ok := conds.InvestmentType.Val.(types.InvestmentType)
		if !ok {
			return nil, fmt.Errorf("invalid investmenttype")
		}
		switch conds.InvestmentType.Op {
		case cruder.EQ:
			q.Where(entpowerrental.InvestmentType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid powerrental field")
		}
	}
	return q, nil
}
