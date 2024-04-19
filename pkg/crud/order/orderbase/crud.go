package order

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"

	"github.com/google/uuid"
)

type Req struct {
	EntID         *uuid.UUID
	AppID         *uuid.UUID
	UserID        *uuid.UUID
	GoodID        *uuid.UUID
	AppGoodID     *uuid.UUID
	GoodType      *goodtypes.GoodType
	ParentOrderID *uuid.UUID
	OrderType     *types.OrderType
	PaymentType   *types.PaymentType
	CreateMethod  *types.OrderCreateMethod
	Simulate      *bool
	DeletedAt     *uint32
}

//nolint:funlen,gocyclo
func CreateSet(c *ent.OrderBaseCreate, req *Req) *ent.OrderBaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
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
	if req.AppGoodID != nil {
		c.SetAppGoodID(*req.AppGoodID)
	}
	if req.GoodType != nil {
		c.SetGoodType(req.GoodType.String())
	}
	if req.ParentOrderID != nil {
		c.SetParentOrderID(*req.ParentOrderID)
	}
	if req.OrderType != nil {
		c.SetOrderType(req.OrderType.String())
	}
	if req.PaymentType != nil {
		c.SetPaymentType(req.PaymentType.String())
	}
	if req.Simulate != nil {
		c.SetSimulate(*req.Simulate)
	}
	if req.CreateMethod != nil {
		c.SetCreateMethod(req.CreateMethod.String())
	}

	return c
}

func UpdateSet(u *ent.OrderBaseUpdateOne, req *Req) *ent.OrderBaseUpdateOne {
	if req.PaymentType != nil {
		u.SetPaymentType(req.PaymentType.String())
	}
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
	AppID          *cruder.Cond
	UserID         *cruder.Cond
	GoodID         *cruder.Cond
	GoodIDs        *cruder.Cond
	AppGoodID      *cruder.Cond
	AppGoodIDs     *cruder.Cond
	ParentOrderID  *cruder.Cond
	ParentOrderIDs *cruder.Cond
	OrderType      *cruder.Cond
	PaymentType    *cruder.Cond
	PaymentTypes   *cruder.Cond
	Simulate       *cruder.Cond
	CreatedAt      *cruder.Cond
	UpdatedAt      *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderBaseQuery, conds *Conds) (*ent.OrderBaseQuery, error) {
	q.Where(entorderbase.DeletedAt(0))
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
			q.Where(entorderbase.ID(id))
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
				q.Where(entorderbase.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
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
			q.Where(entorderbase.EntID(id))
		case cruder.NEQ:
			q.Where(entorderbase.EntIDNEQ(id))
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
				q.Where(entorderbase.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
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
			q.Where(entorderbase.AppID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorderbase.UserID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entorderbase.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		if len(ids) > 0 {
			switch conds.GoodIDs.Op {
			case cruder.IN:
				q.Where(entorderbase.GoodIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.AppGoodID != nil {
		id, ok := conds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodid")
		}
		switch conds.AppGoodID.Op {
		case cruder.EQ:
			q.Where(entorderbase.AppGoodID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.AppGoodIDs != nil {
		ids, ok := conds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appgoodids")
		}
		if len(ids) > 0 {
			switch conds.AppGoodIDs.Op {
			case cruder.IN:
				q.Where(entorderbase.AppGoodIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.ParentOrderID != nil {
		id, ok := conds.ParentOrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid parentorderid")
		}
		switch conds.ParentOrderID.Op {
		case cruder.EQ:
			q.Where(entorderbase.ParentOrderID(id))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.ParentOrderIDs != nil {
		ids, ok := conds.ParentOrderIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid parentorderids")
		}
		if len(ids) > 0 {
			switch conds.ParentOrderIDs.Op {
			case cruder.IN:
				q.Where(entorderbase.ParentOrderIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid order field")
			}
		}
	}
	if conds.OrderType != nil {
		ordertype, ok := conds.OrderType.Val.(types.OrderType)
		if !ok {
			return nil, fmt.Errorf("invalid ordertype")
		}
		switch conds.OrderType.Op {
		case cruder.EQ:
			q.Where(entorderbase.OrderType(ordertype.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentType != nil {
		paymenttype, ok := conds.PaymentType.Val.(types.PaymentType)
		if !ok {
			return nil, fmt.Errorf("invalid paymenttype")
		}
		switch conds.PaymentType.Op {
		case cruder.EQ:
			q.Where(entorderbase.PaymentType(paymenttype.String()))
		case cruder.NEQ:
			q.Where(entorderbase.PaymentTypeNEQ(paymenttype.String()))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.PaymentTypes != nil {
		paymenttypes, ok := conds.PaymentTypes.Val.([]types.PaymentType)
		if !ok {
			return nil, fmt.Errorf("invalid paymenttypes")
		}
		_types := []string{}
		for _, _type := range paymenttypes {
			_types = append(_types, _type.String())
		}
		switch conds.PaymentTypes.Op {
		case cruder.IN:
			q.Where(entorderbase.PaymentTypeIn(_types...))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.Simulate != nil {
		val, ok := conds.Simulate.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid simulate")
		}
		switch conds.Simulate.Op {
		case cruder.EQ:
			q.Where(entorderbase.Simulate(val))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.CreatedAt != nil {
		at, ok := conds.CreatedAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid createdat")
		}
		switch conds.CreatedAt.Op {
		case cruder.LT:
			q.Where(entorderbase.CreatedAtLT(at))
		case cruder.LTE:
			q.Where(entorderbase.CreatedAtLTE(at))
		case cruder.GT:
			q.Where(entorderbase.CreatedAtGT(at))
		case cruder.GTE:
			q.Where(entorderbase.CreatedAtGTE(at))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	if conds.UpdatedAt != nil {
		at, ok := conds.UpdatedAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid updatedat")
		}
		switch conds.UpdatedAt.Op {
		case cruder.LT:
			q.Where(entorderbase.UpdatedAtLT(at))
		case cruder.LTE:
			q.Where(entorderbase.UpdatedAtLTE(at))
		case cruder.GT:
			q.Where(entorderbase.UpdatedAtGT(at))
		case cruder.GTE:
			q.Where(entorderbase.UpdatedAtGTE(at))
		default:
			return nil, fmt.Errorf("invalid order field")
		}
	}
	return q, nil
}
