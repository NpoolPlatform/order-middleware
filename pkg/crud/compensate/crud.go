package compensate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	OrderID   *uuid.UUID
	Start     *uint32
	End       *uint32
	Message   *string
	CreatedAt *uint32
	DeletedAt *uint32
}

func CreateSet(c *ent.CompensateCreate, req *Req) *ent.CompensateCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.Start != nil {
		c.SetStart(*req.Start)
	}
	if req.End != nil {
		c.SetEnd(*req.End)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.CompensateUpdateOne, req *Req) *ent.CompensateUpdateOne {
	if req.Start != nil {
		u.SetStart(*req.Start)
	}
	if req.End != nil {
		u.SetEnd(*req.End)
	}
	if req.Message != nil {
		u.SetMessage(*req.Message)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	IDs     *cruder.Cond
	OrderID *cruder.Cond
	Start   *cruder.Cond
	End     *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.CompensateQuery, conds *Conds) (*ent.CompensateQuery, error) {
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
			q.Where(entcompensate.ID(id))
		default:
			return nil, fmt.Errorf("invalid compensate field")
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
				q.Where(entcompensate.IDIn(ids...))
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
	if conds.Start != nil {
		start, ok := conds.Start.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid start")
		}
		switch conds.Start.Op {
		case cruder.LTE:
			q.Where(entcompensate.StartLTE(start))
		case cruder.GTE:
			q.Where(entcompensate.StartGTE(start))
		default:
			return nil, fmt.Errorf("invalid compensate field")
		}
	}
	if conds.End != nil {
		end, ok := conds.End.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid end")
		}
		switch conds.End.Op {
		case cruder.LTE:
			q.Where(entcompensate.EndLTE(end))
		case cruder.GTE:
			q.Where(entcompensate.EndGTE(end))
		default:
			return nil, fmt.Errorf("invalid compensate field")
		}
	}
	q.Where(entcompensate.DeletedAt(0))
	return q, nil
}
