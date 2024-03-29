package outofgas

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID      *uint32
	EntID   *uuid.UUID
	OrderID *uuid.UUID
	StartAt *uint32
	EndAt   *uint32
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

func WithStartAt(startAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if startAt == nil {
			if must {
				return fmt.Errorf("invalid start")
			}
			return nil
		}
		h.StartAt = startAt
		return nil
	}
}

func WithEndAt(endAt *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if endAt == nil {
			if must {
				return fmt.Errorf("invalid end")
			}
			return nil
		}
		h.EndAt = endAt
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
			h.Conds.ID = &cruder.Cond{
				Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
			}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{
				Op: conds.GetEntID().GetOp(), Val: id,
			}
		}
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
		}
		if conds.StartAt != nil {
			h.Conds.StartAt = &cruder.Cond{Op: conds.GetStartAt().GetOp(), Val: conds.GetStartAt().GetValue()}
		}
		if conds.EndAt != nil {
			h.Conds.EndAt = &cruder.Cond{Op: conds.GetEndAt().GetOp(), Val: conds.GetEndAt().GetValue()}
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
			if req.EntID != nil {
				if _, err := uuid.Parse(*req.EntID); err != nil {
					return err
				}
			}
		}
		h.Reqs = reqs
		return nil
	}
}
