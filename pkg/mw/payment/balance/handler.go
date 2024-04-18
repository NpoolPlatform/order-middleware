package paymentbalance

import (
	"context"
	"fmt"

	paymentbalancecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment/balance"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	paymentbalancecrud.Req
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

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithAmount(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid amount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid amount")
		}
		h.Amount = &amount
		return nil
	}
}

func WithCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid coinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid coinusdcurrency")
		}
		h.CoinUSDCurrency = &amount
		return nil
	}
}

//nolint:dupl
func WithLocalCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid localcoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid localcoinusdcurrency")
		}
		h.LocalCoinUSDCurrency = &amount
		return nil
	}
}

//nolint:dupl
func WithLiveCoinUSDCurrency(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid livecoinusdcurrency")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid livecoinusdcurrency")
		}
		h.LiveCoinUSDCurrency = &amount
		return nil
	}
}
