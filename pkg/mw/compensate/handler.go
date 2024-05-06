package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	compensatecrud.Req
	CompensateConds *compensatecrud.Conds
	OrderBaseConds  *orderbasecrud.Conds
	Offset          int32
	Limit           int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		CompensateConds: &compensatecrud.Conds{},
		OrderBaseConds:  &orderbasecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithOrderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid orderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.OrderID = &_id
		return nil
	}
}

func WithCompensateFromID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid compensatefromid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CompensateFromID = &_id
		return nil
	}
}

func WithCompensateType(e *types.CompensateType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid compensatetype")
			}
			return nil
		}
		switch *e {
		case types.CompensateType_CompensateMalfunction:
		case types.CompensateType_CompensateWalfare:
		case types.CompensateType_CompensateStarterDelay:
		default:
			return wlog.Errorf("invalid compensatetype")
		}
		h.CompensateType = e
		return nil
	}
}

func WithCompensateSeconds(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CompensateSeconds = startAt
		return nil
	}
}

func (h *Handler) withCompensateConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.CompensateConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.CompensateConds.EntID = &cruder.Cond{
			Op: conds.GetEntID().GetOp(), Val: id,
		}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return err
		}
		h.CompensateConds.OrderID = &cruder.Cond{
			Op:  conds.GetOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.CompensateConds.OrderIDs = &cruder.Cond{
			Op:  conds.GetOrderIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.CompensateFromID != nil {
		id, err := uuid.Parse(conds.GetCompensateFromID().GetValue())
		if err != nil {
			return err
		}
		h.CompensateConds.CompensateFromID = &cruder.Cond{
			Op:  conds.GetCompensateFromID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func (h *Handler) withOrderBaseConds(conds *npool.Conds) error {
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return err
		}
		h.OrderBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.OrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetOrderIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return err
		}
		h.OrderBaseConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return err
		}
		h.OrderBaseConds.UserID = &cruder.Cond{
			Op:  conds.GetUserID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withCompensateConds(conds); err != nil {
			return err
		}
		return h.withOrderBaseConds(conds)
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
