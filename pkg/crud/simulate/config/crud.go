package config

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entconfig "github.com/NpoolPlatform/order-middleware/pkg/db/ent/simulateconfig"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                     *uuid.UUID
	AppID                     *uuid.UUID
	Units                     *decimal.Decimal
	Duration                  *uint32
	SendCouponMode            *basetypes.SendCouponMode
	SendCouponProbability     *decimal.Decimal
	CashableProfitProbability *decimal.Decimal
	Enabled                   *bool
	CreatedAt                 *uint32
	DeletedAt                 *uint32
}

func CreateSet(c *ent.SimulateConfigCreate, req *Req) *ent.SimulateConfigCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.SendCouponMode != nil {
		c.SetSendCouponMode(req.SendCouponMode.String())
	}
	if req.SendCouponProbability != nil {
		c.SetSendCouponProbability(*req.SendCouponProbability)
	}
	if req.CashableProfitProbability != nil {
		c.SetCashableProfitProbability(*req.CashableProfitProbability)
	}
	if req.Enabled != nil {
		c.SetEnabled(*req.Enabled)
	}
	if req.CreatedAt != nil {
		c.SetCreatedAt(*req.CreatedAt)
	}
	return c
}

func UpdateSet(u *ent.SimulateConfigUpdateOne, req *Req) *ent.SimulateConfigUpdateOne {
	if req.SendCouponMode != nil {
		u.SetSendCouponMode(req.SendCouponMode.String())
	}
	if req.SendCouponProbability != nil {
		u.SetSendCouponProbability(*req.SendCouponProbability)
	}
	if req.CashableProfitProbability != nil {
		u.SetCashableProfitProbability(*req.CashableProfitProbability)
	}
	if req.Enabled != nil {
		u.SetEnabled(*req.Enabled)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	EntID          *cruder.Cond
	ID             *cruder.Cond
	AppID          *cruder.Cond
	SendCouponMode *cruder.Cond
	Enabled        *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.SimulateConfigQuery, conds *Conds) (*ent.SimulateConfigQuery, error) {
	q.Where(entconfig.DeletedAt(0))
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
			q.Where(entconfig.EntID(id))
		case cruder.NEQ:
			q.Where(entconfig.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid simulateconfig field")
		}
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entconfig.ID(id))
		default:
			return nil, fmt.Errorf("invalid simulateconfig field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entconfig.AppID(id))
		default:
			return nil, fmt.Errorf("invalid simulateconfig field")
		}
	}
	if conds.SendCouponMode != nil {
		sendcouponmode, ok := conds.SendCouponMode.Val.(basetypes.SendCouponMode)
		if !ok {
			return nil, fmt.Errorf("invalid sendcouponmode")
		}
		switch conds.SendCouponMode.Op {
		case cruder.EQ:
			q.Where(entconfig.SendCouponMode(sendcouponmode.String()))
		default:
			return nil, fmt.Errorf("invalid simulateconfig field")
		}
	}
	if conds.Enabled != nil {
		enabled, ok := conds.Enabled.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid enabled")
		}
		switch conds.Enabled.Op {
		case cruder.EQ:
			q.Where(entconfig.Enabled(enabled))
		default:
			return nil, fmt.Errorf("invalid simulateconfig field")
		}
	}

	return q, nil
}
