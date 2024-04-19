package ordercoupon

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	ordercouponcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/coupon"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	ordercouponcrud.Req
	OrderCouponConds *ordercouponcrud.Conds
	OrderBaseConds   *orderbasecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		OrderCouponConds: &ordercouponcrud.Conds{},
		OrderBaseConds:   &orderbasecrud.Conds{},
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
				return fmt.Errorf("invalid id")
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
				return fmt.Errorf("invalid entid")
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
				return fmt.Errorf("invalid orderid")
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

func WithCouponID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid couponid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CouponID = &_id
		return nil
	}
}

func (h *Handler) withOrderCouponConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.OrderCouponConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.OrderCouponConds.EntID = &cruder.Cond{
			Op: conds.GetEntID().GetOp(), Val: id,
		}
	}
	if conds.OrderID != nil {
		id, err := uuid.Parse(conds.GetOrderID().GetValue())
		if err != nil {
			return err
		}
		h.OrderCouponConds.OrderID = &cruder.Cond{
			Op:  conds.GetOrderID().GetOp(),
			Val: id,
		}
	}
	if conds.CouponID != nil {
		id, err := uuid.Parse(conds.GetCouponID().GetValue())
		if err != nil {
			return err
		}
		h.OrderCouponConds.CouponID = &cruder.Cond{
			Op:  conds.GetCouponID().GetOp(),
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
		if err := h.withOrderCouponConds(conds); err != nil {
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