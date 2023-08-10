package outofgas

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uuid.UUID
	OrderID *uuid.UUID
	Start   *uint32
	End     *uint32
	Reqs    []*npool.OutOfGasReq
	Conds   *outofgascrud.Conds
	Offset  int32
	Limit   int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithOrderID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
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

func WithStart(start *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if start == nil {
			return nil
		}
		h.Start = start
		return nil
	}
}

func WithEnd(end *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if end == nil {
			return nil
		}
		h.End = end
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &outofgascrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
		}
		if conds.Start != nil {
			h.Conds.Start = &cruder.Cond{Op: conds.GetStart().GetOp(), Val: basetypes.SignMethod(conds.GetStart().GetValue())}
		}
		if conds.End != nil {
			h.Conds.End = &cruder.Cond{Op: conds.GetEnd().GetOp(), Val: basetypes.SignMethod(conds.GetEnd().GetValue())}
		}

		return nil
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

func WithReqs(reqs []*npool.OutOfGasReq) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, req := range reqs {
			if _, err := uuid.Parse(*req.OrderID); err != nil {
				return err
			}
			if req.ID != nil {
				if _, err := uuid.Parse(*req.ID); err != nil {
					return err
				}
			}
		}
		h.Reqs = reqs
		return nil
	}
}
