package orderlock

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uint32
	EntID     *uuid.UUID
	AppID     *uuid.UUID
	UserID    *uuid.UUID
	OrderID   *uuid.UUID
	LockType  *basetypes.OrderLockType
	CreatedAt *uint32
	DeletedAt *uint32
}

func CreateSet(c *ent.OrderLockCreate, req *Req) *ent.OrderLockCreate {
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
	if req.LockType != nil {
		c.SetLockType(req.LockType.String())
	}

	return c
}

func UpdateSet(u *ent.OrderLockUpdateOne, req *Req) *ent.OrderLockUpdateOne {
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	ID       *cruder.Cond
	AppID    *cruder.Cond
	UserID   *cruder.Cond
	OrderID  *cruder.Cond
	LockType *cruder.Cond
	IDs      *cruder.Cond
	OrderIDs *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.OrderLockQuery, conds *Conds) (*ent.OrderLockQuery, error) {
	q.Where(entorderlock.DeletedAt(0))
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
			q.Where(entorderlock.EntID(id))
		case cruder.NEQ:
			q.Where(entorderlock.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid orderlock field")
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
				q.Where(entorderlock.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid orderlock field")
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
			q.Where(entorderlock.ID(id))
		default:
			return nil, fmt.Errorf("invalid orderlock field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entorderlock.AppID(id))
		default:
			return nil, fmt.Errorf("invalid orderlock field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid userid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entorderlock.UserID(id))
		default:
			return nil, fmt.Errorf("invalid orderlock field")
		}
	}
	if conds.OrderID != nil {
		id, ok := conds.OrderID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid orderid")
		}
		switch conds.OrderID.Op {
		case cruder.EQ:
			q.Where(entorderlock.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid orderlock field")
		}
	}
	if conds.LockType != nil {
		_type, ok := conds.LockType.Val.(basetypes.OrderLockType)
		if !ok {
			return nil, fmt.Errorf("invalid locktype")
		}
		switch conds.LockType.Op {
		case cruder.EQ:
			q.Where(entorderlock.LockType(_type.String()))
		default:
			return nil, fmt.Errorf("invalid orderlock field")
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
				q.Where(entorderlock.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid orderlock field")
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
				q.Where(entorderlock.OrderIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid orderlock field")
			}
		}
	}
	return q, nil
}
