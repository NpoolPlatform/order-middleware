package payment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/order-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Payment{
	ID:                   uuid.NewString(),
	OrderID:              uuid.NewString(),
	AppID:                uuid.NewString(),
	UserID:               uuid.NewString(),
	GoodID:               uuid.NewString(),
	AccountID:            uuid.NewString(),
	StartAmount:          "1001.000000000000000000",
	Amount:               "1002.000000000000000000",
	PayWithBalanceAmount: "1003.000000000000000000",
	FinishAmount:         "1004.000000000000000000",
	CoinUsdCurrency:      "1005.000000000000000000",
	LocalCoinUsdCurrency: "1006.000000000000000000",
	LiveCoinUsdCurrency:  "1007.000000000000000000",
	CoinInfoID:           uuid.NewString(),
	PaymentStateStr:      ordertypes.PaymentState_PaymentStateWait.String(),
	State:                ordertypes.PaymentState_PaymentStateWait,
	ChainTransactionID:   uuid.NewString(),
	UserSetPaid:          false,
	UserSetCanceled:      false,
	FakePayment:          false,
}

func createPayment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithUserID(&ret.UserID, true),
		WithOrderID(&ret.OrderID, true),
		WithAccountID(&ret.AccountID, true),
		WithStartAmount(&ret.StartAmount, true),
		WithAmount(&ret.Amount, true),
		WithPayWithBalanceAmount(&ret.PayWithBalanceAmount, true),
		WithFinishAmount(&ret.FinishAmount, true),
		WithCoinUsdCurrency(&ret.CoinUsdCurrency, true),
		WithLocalCoinUsdCurrency(&ret.LocalCoinUsdCurrency, true),
		WithLiveCoinUsdCurrency(&ret.LiveCoinUsdCurrency, true),
		WithCoinInfoID(&ret.CoinInfoID, true),
		WithState(&ret.State, true),
		WithChainTransactionID(&ret.ChainTransactionID, true),
		WithUserSetPaid(&ret.UserSetPaid, true),
		WithUserSetCanceled(&ret.UserSetCanceled, true),
		WithFakePayment(&ret.FakePayment, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreatePayment(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updatePayment(t *testing.T) {
	ret.State = ordertypes.PaymentState_PaymentStateDone
	ret.PaymentStateStr = ordertypes.PaymentState_PaymentStateDone.String()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOrderID(&ret.ID, false),
		WithState(&ret.State, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdatePayment(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getPayment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetPayment(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getPayments(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetPayments(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deletePayment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeletePayment(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetPayment(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestPayment(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createPayment", createPayment)
	t.Run("updatePayment", updatePayment)
	t.Run("getPayment", getPayment)
	t.Run("getPayments", getPayments)
	t.Run("deletePayment", deletePayment)
}
