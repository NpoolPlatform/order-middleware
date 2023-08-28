package payment

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID          *uuid.UUID
	AppID       *uuid.UUID
	UserID      *uuid.UUID
	GoodID      *uuid.UUID
	OrderID     *uuid.UUID
	AccountID   *uuid.UUID
	StartAmount *decimal.Decimal
	CreatedAt   *uint32
	DeletedAt   *uint32
}

func CreateSet(c *ent.PaymentCreate, req *Req) *ent.PaymentCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.AccountID != nil {
		c.SetAccountID(*req.AccountID)
	}
	if req.StartAmount != nil {
		c.SetStartAmount(*req.StartAmount)
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.PaymentUpdateOne, req *Req) *ent.PaymentUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	IDs       *cruder.Cond
	AppID     *cruder.Cond
	UserID    *cruder.Cond
	GoodID    *cruder.Cond
	OrderID   *cruder.Cond
	AccountID *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.PaymentQuery, conds *Conds) (*ent.PaymentQuery, error) {
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
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entpayment.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	if conds.AccountID != nil {
		id, ok := conds.AccountID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid accountid")
		}
		switch conds.AccountID.Op {
		case cruder.EQ:
			q.Where(entpayment.AccountID(id))
		default:
			return nil, fmt.Errorf("invalid payment field")
		}
	}
	q.Where(entpayment.DeletedAt(0))
	return q, nil
}
