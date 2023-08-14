package payment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/shopspring/decimal"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/order-middleware/pkg/testinit"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	id  = uuid.NewString()
	ret = npool.Payment{
		ID:                   id,
		OrderID:              uuid.NewString(),
		AppID:                uuid.NewString(),
		UserID:               uuid.NewString(),
		GoodID:               uuid.NewString(),
		AccountID:            uuid.NewString(),
		StartAmount:          decimal.Zero.String(),
		Amount:               decimal.Zero.String(),
		PayWithBalanceAmount: decimal.Zero.String(),
		FinishAmount:         decimal.Zero.String(),
		CoinUsdCurrency:      decimal.Zero.String(),
		LocalCoinUsdCurrency: decimal.Zero.String(),
		LiveCoinUsdCurrency:  decimal.Zero.String(),
		CoinInfoID:           uuid.NewString(),
		PaymentStateStr:      basetypes.PaymentState_DefaultPaymentState.String(),
		State:                basetypes.PaymentState_DefaultPaymentState,
		ChainTransactionID:   uuid.NewString(),
		UserSetPaid:          false,
		UserSetCanceled:      false,
		FakePayment:          false,
	}
)

func createPayment(t *testing.T) {
	var (
		req = npool.PaymentReq{
			ID:                   &ret.ID,
			AppID:                &ret.AppID,
			GoodID:               &ret.GoodID,
			UserID:               &ret.UserID,
			OrderID:              &ret.OrderID,
			AccountID:            &ret.AccountID,
			StartAmount:          &ret.StartAmount,
			Amount:               &ret.Amount,
			PayWithBalanceAmount: &ret.PayWithBalanceAmount,
			FinishAmount:         &ret.FinishAmount,
			CoinUsdCurrency:      &ret.CoinUsdCurrency,
			LocalCoinUsdCurrency: &ret.LocalCoinUsdCurrency,
			LiveCoinUsdCurrency:  &ret.LiveCoinUsdCurrency,
			CoinInfoID:           &ret.CoinInfoID,
			State:                &ret.State,
			ChainTransactionID:   &ret.ChainTransactionID,
			UserSetPaid:          &ret.UserSetPaid,
			UserSetCanceled:      &ret.UserSetCanceled,
			FakePayment:          &ret.FakePayment,
		}
	)

	info, err := CreatePayment(context.Background(), &req)
	if assert.Nil(t, err) {
		if id != info.ID {
			ret.ID = info.ID
		}
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updatePayment(t *testing.T) {
	if ret.ID == id {
		var (
			state = basetypes.PaymentState_PaymentStateDone

			req = npool.PaymentReq{
				ID:    &ret.ID,
				State: &state,
			}
		)
		info, err := UpdatePayment(context.Background(), &req)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getPayment(t *testing.T) {
	info, err := GetPayment(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getPayments(t *testing.T) {
	infos, _, err := GetPayments(context.Background(), &npool.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deletePayment(t *testing.T) {
	if ret.ID == id {
		info, err := DeletePayment(context.Background(), &npool.PaymentReq{
			ID: &ret.ID,
		})
		if assert.Nil(t, err) {
			assert.Equal(t, info, &ret)
		}

		info, err = GetPayment(context.Background(), ret.ID)
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestPayment(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createPayment", createPayment)
	t.Run("updatePayment", updatePayment)
	t.Run("getPayment", getPayment)
	t.Run("getPayments", getPayments)
	t.Run("deletePayment", deletePayment)
}
