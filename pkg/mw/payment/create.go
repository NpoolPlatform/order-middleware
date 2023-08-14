package payment

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

//nolint:gocyclo
func (h *createHandler) validate() error {
	if h.AppID == nil {
		return fmt.Errorf("invalid appid")
	}
	if h.GoodID == nil {
		return fmt.Errorf("invalid goodid")
	}
	if h.UserID == nil {
		return fmt.Errorf("invalid userid")
	}
	if h.OrderID == nil {
		return fmt.Errorf("invalid orderid")
	}
	if h.AccountID == nil {
		return fmt.Errorf("invalid accountid")
	}
	if h.StartAmount == nil {
		return fmt.Errorf("invalid startamount")
	}
	if h.StartAmount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("startamount is less than or equal to 0")
	}
	if h.Amount == nil {
		return fmt.Errorf("invalid amount")
	}
	if h.Amount.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("amount is less than or equal to 0")
	}
	if h.CoinUsdCurrency == nil {
		return fmt.Errorf("invalid coinusdcurrency")
	}
	if h.CoinUsdCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("coinusdcurrency is less than or equal to 0")
	}
	if h.LocalCoinUsdCurrency == nil {
		return fmt.Errorf("invalid localcoinusdcurrency")
	}
	if h.LocalCoinUsdCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("localcoinusdcurrency is less than or equal to 0")
	}
	if h.LiveCoinUsdCurrency == nil {
		return fmt.Errorf("invalid livecoinusdcurrency")
	}
	if h.LiveCoinUsdCurrency.Cmp(decimal.NewFromInt(0)) <= 0 {
		return fmt.Errorf("livecoinusdcurrency is less than or equal to 0")
	}
	if h.CoinInfoID == nil {
		return fmt.Errorf("invalid coininfoid")
	}
	switch *h.State {
	case basetypes.PaymentState_PaymentStateWait:
	default:
		return fmt.Errorf("invalid state")
	}
	return nil
}

func (h *Handler) CreatePayment(ctx context.Context) (*npool.Payment, error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.OrderID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		stm, err := paymentcrud.SetQueryConds(
			tx.Payment.Query(),
			&paymentcrud.Conds{
				AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
				UserID:    &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
				OrderID:   &cruder.Cond{Op: cruder.EQ, Val: *h.OrderID},
				AccountID: &cruder.Cond{Op: cruder.EQ, Val: *h.AccountID},
			},
		)
		if err != nil {
			return err
		}

		info, err := stm.Only(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return err
			}
		}
		if info != nil {
			return fmt.Errorf("payment exist")
		}
		if _, err := paymentcrud.CreateSet(
			tx.Payment.Create(),
			&paymentcrud.Req{
				ID:                   h.ID,
				AppID:                h.AppID,
				GoodID:               h.GoodID,
				UserID:               h.UserID,
				OrderID:              h.OrderID,
				AccountID:            h.AccountID,
				StartAmount:          h.StartAmount,
				Amount:               h.Amount,
				PayWithBalanceAmount: h.PayWithBalanceAmount,
				FinishAmount:         h.FinishAmount,
				CoinUsdCurrency:      h.CoinUsdCurrency,
				LocalCoinUsdCurrency: h.LocalCoinUsdCurrency,
				LiveCoinUsdCurrency:  h.LiveCoinUsdCurrency,
				CoinInfoID:           h.CoinInfoID,
				State:                h.State,
				ChainTransactionID:   h.ChainTransactionID,
				UserSetPaid:          h.UserSetPaid,
				UserSetCanceled:      h.UserSetCanceled,
				FakePayment:          h.FakePayment,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetPayment(ctx)
}
