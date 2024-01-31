//nolint:dupl
package config

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	configcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/simulate/config"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                    *uint32
	EntID                 *uuid.UUID
	AppID                 *uuid.UUID
	Units                 *decimal.Decimal
	Duration              *uint32
	SendCouponMode        *basetypes.SendCouponMode
	SendCouponProbability *decimal.Decimal
	Enabled               *bool
	Reqs                  []*npool.SimulateConfigReq
	Conds                 *configcrud.Conds
	Offset                int32
	Limit                 int32
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithUnits(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid units")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("units is less than or equal to 0")
		}
		h.Units = &amount
		return nil
	}
}

func WithDuration(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return fmt.Errorf("invalid duration")
			}
			return nil
		}
		h.Duration = duration
		return nil
	}
}

func WithSendCouponMode(sendcouponmode *basetypes.SendCouponMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if sendcouponmode == nil {
			if must {
				return fmt.Errorf("invalid sendcouponmode")
			}
			return nil
		}
		switch *sendcouponmode {
		case basetypes.SendCouponMode_WithoutCoupon:
		case basetypes.SendCouponMode_FirstBenifit:
		case basetypes.SendCouponMode_RandomBenifit:
		case basetypes.SendCouponMode_FirstAndRandomBenifit:
		default:
			return fmt.Errorf("invalid sendcouponmode")
		}
		h.SendCouponMode = sendcouponmode
		return nil
	}
}

func WithSendCouponProbability(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid sendcouponprobability")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("sendcouponprobability is less than 0")
		}
		h.SendCouponProbability = &amount
		return nil
	}
}

func WithEnabled(value *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid enabled")
			}
			return nil
		}
		h.Enabled = value
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &configcrud.Conds{}
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.SendCouponMode != nil {
			switch conds.GetSendCouponMode().GetValue() {
			case uint32(basetypes.SendCouponMode_WithoutCoupon):
			case uint32(basetypes.SendCouponMode_FirstBenifit):
			case uint32(basetypes.SendCouponMode_RandomBenifit):
			case uint32(basetypes.SendCouponMode_FirstAndRandomBenifit):
			default:
				return fmt.Errorf("invalid sendcouponmode")
			}
			_type := conds.GetSendCouponMode().GetValue()
			h.Conds.SendCouponMode = &cruder.Cond{Op: conds.GetSendCouponMode().GetOp(), Val: basetypes.SendCouponMode(_type)}
		}
		if conds.Enabled != nil {
			h.Conds.Enabled = &cruder.Cond{
				Op:  conds.GetEnabled().GetOp(),
				Val: conds.GetEnabled().GetValue(),
			}
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
