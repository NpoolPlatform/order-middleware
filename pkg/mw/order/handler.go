package order

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"

	"github.com/google/uuid"
)

type Handler struct {
	ID                  *uint32
	EntID               *uuid.UUID
	OrderBaseConds      *orderbasecrud.Conds
	OrderStateBaseConds *orderstatebasecrud.Conds
	Offset              int32
	Limit               int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		OrderBaseConds:      &orderbasecrud.Conds{},
		OrderStateBaseConds: &orderstatebasecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
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
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

//nolint:gocyclo,funlen
func (h *Handler) withOrderBaseConds(conds *npool.Conds) error {
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.EntIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetEntIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetEntIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	if conds.UserID != nil {
		id, err := uuid.Parse(conds.GetUserID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.UserID = &cruder.Cond{
			Op:  conds.GetUserID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.ParentOrderID != nil {
		id, err := uuid.Parse(conds.GetParentOrderID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.OrderBaseConds.ParentOrderID = &cruder.Cond{
			Op:  conds.GetParentOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.ParentOrderIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetParentOrderIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.OrderBaseConds.ParentOrderIDs = &cruder.Cond{
			Op:  conds.GetParentOrderIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.OrderType != nil {
		h.OrderBaseConds.OrderType = &cruder.Cond{
			Op:  conds.GetOrderType().GetOp(),
			Val: types.OrderType(conds.GetOrderType().GetValue()),
		}
	}
	if conds.CreatedAt != nil {
		h.OrderBaseConds.CreatedAt = &cruder.Cond{
			Op:  conds.GetCreatedAt().GetOp(),
			Val: conds.GetCreatedAt().GetValue(),
		}
	}
	return nil
}

func (h *Handler) withOrderStateBaseConds(conds *npool.Conds) error {
	if conds.OrderState != nil {
		h.OrderStateBaseConds.OrderState = &cruder.Cond{
			Op:  conds.GetOrderState().GetOp(),
			Val: types.OrderState(conds.GetOrderState().GetValue()),
		}
	}
	if conds.OrderStates != nil {
		_types := []types.OrderState{}
		for _, _type := range conds.GetOrderStates().GetValue() {
			_types = append(_types, types.OrderState(_type))
		}
		h.OrderStateBaseConds.OrderStates = &cruder.Cond{
			Op:  conds.GetOrderStates().GetOp(),
			Val: _types,
		}
	}
	if conds.PaymentType != nil {
		h.OrderStateBaseConds.PaymentType = &cruder.Cond{
			Op:  conds.GetPaymentType().GetOp(),
			Val: types.PaymentType(conds.GetPaymentType().GetValue()),
		}
	}
	if conds.PaymentTypes != nil {
		_types := []types.PaymentType{}
		for _, _type := range conds.GetPaymentTypes().GetValue() {
			_types = append(_types, types.PaymentType(_type))
		}
		h.OrderStateBaseConds.PaymentTypes = &cruder.Cond{
			Op:  conds.GetPaymentTypes().GetOp(),
			Val: _types,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withOrderBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withOrderStateBaseConds(conds)
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
