package compensate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"

	"github.com/google/uuid"
)

type Req struct {
	EntID             *uuid.UUID
	OrderID           *uuid.UUID
	CompensateFromID  *uuid.UUID
	CompensateType    *types.CompensateType
	CompensateSeconds *uint32
	DeletedAt         *uint32
}

func CreateSet(c *ent.CompensateCreate, req *Req) *ent.CompensateCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.CompensateFromID != nil {
		c.SetCompensateFromID(*req.CompensateFromID)
	}
	if req.CompensateType != nil {
		c.SetCompensateType(req.CompensateType.String())
	}
	if req.CompensateSeconds != nil {
		c.SetCompensateSeconds(*req.CompensateSeconds)
	}
	return c
}

func UpdateSet(u *ent.CompensateUpdateOne, req *Req) *ent.CompensateUpdateOne {
	if req.CompensateSeconds != nil {
		u.SetCompensateSeconds(*req.CompensateSeconds)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID               *cruder.Cond
	IDs              *cruder.Cond
	EntID            *cruder.Cond
	EntIDs           *cruder.Cond
	OrderID          *cruder.Cond
	OrderIDs         *cruder.Cond
	CompensateFromID *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CompensateQuery, conds *Conds) (*ent.CompensateQuery, error) {
	q.Where(entcompensate.DeletedAt(0))
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
			q.Where(entcompensate.ID(id))
		case cruder.NEQ:
			q.Where(entcompensate.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid compensate field")
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
				q.Where(entcompensate.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid compensate field")
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
			q.Where(entcompensate.EntID(id))
		case cruder.NEQ:
			q.Where(entcompensate.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid compensate field")
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
				q.Where(entcompensate.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid compensate field")
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
			q.Where(entcompensate.OrderID(id))
		default:
			return nil, fmt.Errorf("invalid compensate field")
		}
	}
	if conds.CompensateFromID != nil {
		id, ok := conds.CompensateFromID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid compensateid")
		}
		switch conds.CompensateFromID.Op {
		case cruder.EQ:
			q.Where(entcompensate.CompensateFromID(id))
		default:
			return nil, fmt.Errorf("invalid compensate field")
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
				q.Where(entcompensate.OrderIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid compensate field")
			}
		}
	}
	return q, nil
}
