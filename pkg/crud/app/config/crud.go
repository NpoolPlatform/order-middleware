package appconfig

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entappconfig "github.com/NpoolPlatform/order-middleware/pkg/db/ent/appconfig"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	EntID                                  *uuid.UUID
	AppID                                  *uuid.UUID
	EnableSimulateOrder                    *bool
	SimulateOrderUnits                     *decimal.Decimal
	SimulateOrderDurationSeconds           *uint32
	SimulateOrderCouponMode                *types.SimulateOrderCouponMode
	SimulateOrderCouponProbability         *decimal.Decimal
	SimulateOrderCashableProfitProbability *decimal.Decimal
	MaxUnpaidOrders                        *uint32
	DeletedAt                              *uint32
}

func CreateSet(c *ent.AppConfigCreate, req *Req) *ent.AppConfigCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.EnableSimulateOrder != nil {
		c.SetEnableSimulateOrder(*req.EnableSimulateOrder)
	}
	if req.SimulateOrderUnits != nil {
		c.SetSimulateOrderUnits(*req.SimulateOrderUnits)
	}
	if req.SimulateOrderDurationSeconds != nil {
		c.SetSimulateOrderDurationSeconds(*req.SimulateOrderDurationSeconds)
	}
	if req.SimulateOrderCouponMode != nil {
		c.SetSimulateOrderCouponMode(req.SimulateOrderCouponMode.String())
	}
	if req.SimulateOrderCouponProbability != nil {
		c.SetSimulateOrderCouponProbability(*req.SimulateOrderCouponProbability)
	}
	if req.SimulateOrderCashableProfitProbability != nil {
		c.SetSimulateOrderCashableProfitProbability(*req.SimulateOrderCashableProfitProbability)
	}
	if req.MaxUnpaidOrders != nil {
		c.SetMaxUnpaidOrders(*req.MaxUnpaidOrders)
	}
	return c
}

func UpdateSet(u *ent.AppConfigUpdateOne, req *Req) *ent.AppConfigUpdateOne {
	if req.EnableSimulateOrder != nil {
		u.SetEnableSimulateOrder(*req.EnableSimulateOrder)
	}
	if req.SimulateOrderUnits != nil {
		u.SetSimulateOrderUnits(*req.SimulateOrderUnits)
	}
	if req.SimulateOrderDurationSeconds != nil {
		u.SetSimulateOrderDurationSeconds(*req.SimulateOrderDurationSeconds)
	}
	if req.SimulateOrderCouponMode != nil {
		u.SetSimulateOrderCouponMode(req.SimulateOrderCouponMode.String())
	}
	if req.SimulateOrderCouponProbability != nil {
		u.SetSimulateOrderCouponProbability(*req.SimulateOrderCouponProbability)
	}
	if req.SimulateOrderCashableProfitProbability != nil {
		u.SetSimulateOrderCashableProfitProbability(*req.SimulateOrderCashableProfitProbability)
	}
	if req.MaxUnpaidOrders != nil {
		u.SetMaxUnpaidOrders(*req.MaxUnpaidOrders)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID     *cruder.Cond
	IDs    *cruder.Cond
	EntID  *cruder.Cond
	EntIDs *cruder.Cond
	AppID  *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppConfigQuery, conds *Conds) (*ent.AppConfigQuery, error) {
	q.Where(entappconfig.DeletedAt(0))
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
			q.Where(entappconfig.ID(id))
		case cruder.NEQ:
			q.Where(entappconfig.IDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid appconfig field")
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
				q.Where(entappconfig.IDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid appconfig field")
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
			q.Where(entappconfig.EntID(id))
		case cruder.NEQ:
			q.Where(entappconfig.EntIDNEQ(id))
		default:
			return nil, fmt.Errorf("invalid appconfig field")
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
				q.Where(entappconfig.EntIDIn(ids...))
			default:
				return nil, fmt.Errorf("invalid appconfig field")
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
			q.Where(entappconfig.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appconfig field")
		}
	}
	return q, nil
}
