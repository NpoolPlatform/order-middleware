package compensate

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"

	"github.com/google/uuid"
)

type Req struct {
	EntID     *uuid.UUID
	OrderID   *uuid.UUID
	StartAt   *uint32
	EndAt     *uint32
	Title     *string
	Message   *string
	DeletedAt *uint32
}

func CreateSet(c *ent.CompensateCreate, req *Req) *ent.CompensateCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.OrderID != nil {
		c.SetOrderID(*req.OrderID)
	}
	if req.StartAt != nil {
		c.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		c.SetEndAt(*req.EndAt)
	}
	if req.Title != nil {
		c.SetTitle(*req.Title)
	}
	if req.Message != nil {
		c.SetMessage(*req.Message)
	}
	return c
}

func UpdateSet(u *ent.CompensateUpdateOne, req *Req) *ent.CompensateUpdateOne {
	if req.StartAt != nil {
		u.SetStartAt(*req.StartAt)
	}
	if req.EndAt != nil {
		u.SetEndAt(*req.EndAt)
	}
	if req.Title != nil {
		u.SetTitle(*req.Title)
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
	ID       *cruder.Cond
	IDs      *cruder.Cond
	EntID    *cruder.Cond
	EntIDs   *cruder.Cond
	OrderID  *cruder.Cond
	StartAt  *cruder.Cond
	EndAt    *cruder.Cond
	StartEnd *cruder.Cond
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
	if conds.StartAt != nil {
		start, ok := conds.StartAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid startat")
		}
		switch conds.StartAt.Op {
		case cruder.LT:
			q.Where(entcompensate.StartAtLT(start))
		case cruder.LTE:
			q.Where(entcompensate.StartAtLTE(start))
		case cruder.GT:
			q.Where(entcompensate.StartAtGT(start))
		case cruder.GTE:
			q.Where(entcompensate.StartAtGTE(start))
		case cruder.EQ:
			q.Where(entcompensate.StartAt(start))
		default:
			return nil, fmt.Errorf("invalid compensate field")
		}
	}
	if conds.EndAt != nil {
		end, ok := conds.EndAt.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid endat")
		}
		switch conds.EndAt.Op {
		case cruder.LT:
			q.Where(entcompensate.EndAtLT(end))
		case cruder.LTE:
			q.Where(entcompensate.EndAtLTE(end))
		case cruder.GT:
			q.Where(entcompensate.EndAtGT(end))
		case cruder.GTE:
			q.Where(entcompensate.EndAtGTE(end))
		case cruder.EQ:
			q.Where(entcompensate.EndAtEQ(end))
		default:
			return nil, fmt.Errorf("invalid compensate field")
		}
	}
	if conds.StartEnd != nil {
		ats, ok := conds.StartEnd.Val.([]uint32)
		if !ok || len(ats) != 2 {
			return nil, fmt.Errorf("invalid startend")
		}
		switch conds.StartEnd.Op {
		case cruder.OVERLAP:
			q.Where(
				entcompensate.Or(
					entcompensate.And(
						entcompensate.StartAtLTE(ats[0]),
						entcompensate.EndAtGTE(ats[0]),
					),
					entcompensate.And(
						entcompensate.StartAtLTE(ats[1]),
						entcompensate.EndAtGTE(ats[1]),
					),
					entcompensate.And(
						entcompensate.StartAtGTE(ats[0]),
						entcompensate.EndAtLTE(ats[1]),
					),
				),
			)
		default:
			return nil, fmt.Errorf("invalid compensate field")
		}
	}
	return q, nil
}
