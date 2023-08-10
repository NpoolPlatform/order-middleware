package payment

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	constant "github.com/NpoolPlatform/order-middleware/pkg/const"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Handler struct {
	ID                   *uuid.UUID
	AppID                *uuid.UUID
	GoodID               *uuid.UUID
	UserID               *uuid.UUID
	OrderID              *uuid.UUID
	AccountID            *uuid.UUID
	StartAmount          *decimal.Decimal
	Amount               *decimal.Decimal
	PayWithBalanceAmount *decimal.Decimal
	FinishAmount         *decimal.Decimal
	CoinUsdCurrency      *decimal.Decimal
	LocalCoinUsdCurrency *decimal.Decimal
	LiveCoinUsdCurrency  *decimal.Decimal
	CoinInfoID           *uuid.UUID
	State                *basetypes.PaymentState
	ChainTransactionID   *string
	UserSetPaid          *bool
	UserSetCanceled      *bool
	FakePayment          *bool
	Reqs                 []*npool.PaymentReq
	Conds                *paymentcrud.Conds
	Offset               int32
	Limit                int32
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

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
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

func WithGoodID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithUserID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.UserID = &_id
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

func WithAccountID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AccountID = &_id
		return nil
	}
}

func WithStartAmount(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.StartAmount = &amount
		return nil
	}
}

func WithAmount(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.Amount = &amount
		return nil
	}
}

func WithPayWithBalanceAmount(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.PayWithBalanceAmount = &amount
		return nil
	}
}

func WithFinishAmount(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.FinishAmount = &amount
		return nil
	}
}

func WithCoinUsdCurrency(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.CoinUsdCurrency = &amount
		return nil
	}
}

func WithLocalCoinUsdCurrency(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.LocalCoinUsdCurrency = &amount
		return nil
	}
}

func WithLiveCoinUsdCurrency(value *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		h.LiveCoinUsdCurrency = &amount
		return nil
	}
}

func WithCoinInfoID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinInfoID = &_id
		return nil
	}
}

func WithChainTransactionID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		h.ChainTransactionID = id
		return nil
	}
}

func WithUserSetPaid(value *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		h.UserSetPaid = value
		return nil
	}
}

func WithUserSetCanceled(value *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		h.UserSetCanceled = value
		return nil
	}
}

func WithFakePayment(value *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			return nil
		}
		h.FakePayment = value
		return nil
	}
}

func WithState(state *basetypes.PaymentState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case basetypes.PaymentState_PaymentStateWait:
		case basetypes.PaymentState_PaymentStateCanceled:
		case basetypes.PaymentState_PaymentStateTimeOut:
		case basetypes.PaymentState_PaymentStateDone:
		default:
			return fmt.Errorf("invalid lockedby")
		}
		h.State = state
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &paymentcrud.Conds{}
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
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{Op: conds.GetAppID().GetOp(), Val: id}
		}
		if conds.UserID != nil {
			id, err := uuid.Parse(conds.GetUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.UserID = &cruder.Cond{Op: conds.GetUserID().GetOp(), Val: id}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{Op: conds.GetGoodID().GetOp(), Val: id}
		}
		if conds.OrderID != nil {
			id, err := uuid.Parse(conds.GetOrderID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.OrderID = &cruder.Cond{Op: conds.GetOrderID().GetOp(), Val: id}
		}
		if conds.AccountID != nil {
			id, err := uuid.Parse(conds.GetAccountID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AccountID = &cruder.Cond{Op: conds.GetAccountID().GetOp(), Val: id}
		}
		if conds.CoinInfoID != nil {
			id, err := uuid.Parse(conds.GetCoinInfoID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinInfoID = &cruder.Cond{Op: conds.GetCoinInfoID().GetOp(), Val: id}
		}
		if conds.State != nil {
			switch conds.GetState().GetValue() {
			case uint32(basetypes.PaymentState_PaymentStateWait):
			case uint32(basetypes.PaymentState_PaymentStateDone):
			case uint32(basetypes.PaymentState_PaymentStateCanceled):
			case uint32(basetypes.PaymentState_PaymentStateTimeOut):
			default:
				return fmt.Errorf("invalid state")
			}
			_state := conds.GetState().GetValue()
			h.Conds.State = &cruder.Cond{Op: conds.GetState().GetOp(), Val: basetypes.PaymentState(_state)}
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

//nolint:gocyclo
func WithReqs(reqs []*npool.PaymentReq) func(context.Context, *Handler) error {
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
			if req.AppID != nil {
				if _, err := uuid.Parse(*req.AppID); err != nil {
					return err
				}
			}
			if req.GoodID != nil {
				if _, err := uuid.Parse(*req.GoodID); err != nil {
					return err
				}
			}
			if req.UserID != nil {
				if _, err := uuid.Parse(*req.UserID); err != nil {
					return err
				}
			}
			if req.OrderID != nil {
				if _, err := uuid.Parse(*req.OrderID); err != nil {
					return err
				}
			}
			if req.AccountID != nil {
				if _, err := uuid.Parse(*req.AccountID); err != nil {
					return err
				}
			}
			if req.CoinInfoID != nil {
				if _, err := uuid.Parse(*req.CoinInfoID); err != nil {
					return err
				}
			}
			if req.State != nil {
				switch req.GetState() {
				case basetypes.PaymentState_PaymentStateWait:
				case basetypes.PaymentState_PaymentStateDone:
				case basetypes.PaymentState_PaymentStateCanceled:
				case basetypes.PaymentState_PaymentStateTimeOut:
				default:
					return fmt.Errorf("invalid State")
				}
			}
		}
		h.Reqs = reqs
		return nil
	}
}
