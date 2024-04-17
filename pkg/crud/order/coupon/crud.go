package ordercoupon

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	"github.com/google/uuid"
)

type Req struct {
	ID         *uint32
	EntID      *uuid.UUID
	AppID      *uuid.UUID
	UserID     *uuid.UUID
	OrderID    *uuid.UUID
	CouponType *basetypes.OrderCouponType
	CreatedAt  *uint32
	DeletedAt  *uint32
}

func CreateSet(c *ent.OrderCouponCreate, req *Req) *ent.OrderCouponCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.CouponType != nil {
		c.SetCouponType(req.CouponType.String())
	}

	return c
}

func UpdateSet(u *ent.OrderCouponUpdateOne, req *Req) *ent.OrderCouponUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID      *cruder.Cond
	EntIDs     *cruder.Cond
	ID         *cruder.Cond
	AppID      *cruder.Cond
	UserID     *cruder.Cond
	OrderID    *cruder.Cond
	CouponType *cruder.Cond
	IDs        *cruder.Cond
	OrderIDs   *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderCouponQuery, conds *Conds) (*ent.OrderCouponQuery, error) {
	q.Where(entordercoupon.DeletedAt(0))
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
			q.Where(entordercoupon.EntID(id))
		case cruder.NEQ:
			q.Where(entordercoupon.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid ordercoupon field")
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
				q.Where(entordercoupon.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid ordercoupon field")
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
			q.Where(entordercoupon.ID(id))
		default:
			return nil, fmt.Errorf("invalid ordercoupon field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entordercoupon.AppID(id))
		default:
			return nil, fmt.Errorf("invalid ordercoupon field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entordercoupon.UserID(id))
		default:
			return nil, fmt.Errorf("invalid ordercoupon field")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entordercoupon.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid ordercoupon field")
		}
	}
	if conds.CouponType != nil {
		_type, ok := conds.CouponType.Val.(basetypes.OrderCouponType)
		if !ok {
			return nil, fmt.Errorf("invalid coupontype")
		}
		switch conds.CouponType.Op {
		case cruder.EQ:
			q.Where(entordercoupon.CouponType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid ordercoupon field")
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
				q.Where(entordercoupon.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid ordercoupon field")
			}
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
				q.Where(entordercoupon.OrderIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid ordercoupon field")
			}
		}
	}
	return q, nil
}
