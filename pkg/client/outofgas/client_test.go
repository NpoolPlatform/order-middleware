package outofgas

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"

	order1 "github.com/NpoolPlatform/order-middleware/pkg/client/order"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/order-middleware/pkg/testinit"

	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
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

const secondsPerDay = 24 * 60 * 60
const seconds = 1

var (
	now   = uint32(time.Now().Unix())
	order = ordermwpb.Order{
		ID:                   uuid.NewString(),
		AppID:                uuid.NewString(),
		UserID:               uuid.NewString(),
		GoodID:               uuid.NewString(),
		AppGoodID:            uuid.NewString(),
		Units:                "100",
		GoodValue:            "1007",
		GoodValueUSD:         "1007",
		PaymentAmount:        "1121",
		DiscountAmount:       "10",
		PromotionID:          uuid.NewString(),
		DurationDays:         now + 5*secondsPerDay,
		OrderTypeStr:         ordertypes.OrderType_Normal.String(),
		OrderType:            ordertypes.OrderType_Normal,
		InvestmentType:       ordertypes.InvestmentType_FullPayment,
		InvestmentTypeStr:    ordertypes.InvestmentType_FullPayment.String(),
		PaymentTypeStr:       ordertypes.PaymentType_PayWithTransferAndBalance.String(),
		PaymentType:          ordertypes.PaymentType_PayWithTransferAndBalance,
		CoinTypeID:           uuid.NewString(),
		PaymentCoinTypeID:    uuid.NewString(),
		TransferAmount:       "1011",
		BalanceAmount:        "110",
		CoinUSDCurrency:      "1004",
		LocalCoinUSDCurrency: "1005",
		LiveCoinUSDCurrency:  "1006",

		PaymentAccountID:   uuid.NewString(),
		PaymentStartAmount: "1010",

		OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
		OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
		StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
		StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
		StartAt:              now + 5*seconds,
		EndAt:                now + 5*secondsPerDay,
		LastBenefitAt:        0,
		BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
		BenefitState:         ordertypes.BenefitState_BenefitWait,
		UserSetPaid:          false,
		UserSetCanceled:      false,
		AdminSetCanceled:     false,
		PaymentTransactionID: "",
		PaymentFinishAmount:  "0",
		PaymentStateStr:      ordertypes.PaymentState_PaymentStateWait.String(),
		PaymentState:         ordertypes.PaymentState_PaymentStateWait,
		OutOfGasHours:        0,
		CompensateHours:      0,
		AppGoodStockLockID:   uuid.NewString(),
		LedgerLockID:         uuid.NewString(),
	}
	id  = uuid.NewString()
	ret = npool.OutOfGas{
		ID:      id,
		OrderID: order.ID,
		StartAt: now + 6*seconds,
		EndAt:   now + 7*seconds,
	}
)

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	_, err := order1.CreateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:                   &order.ID,
		AppID:                &order.AppID,
		UserID:               &order.UserID,
		GoodID:               &order.GoodID,
		AppGoodID:            &order.AppGoodID,
		Units:                &order.Units,
		GoodValue:            &order.GoodValue,
		GoodValueUSD:         &order.GoodValueUSD,
		PaymentAmount:        &order.PaymentAmount,
		DiscountAmount:       &order.DiscountAmount,
		PromotionID:          &order.PromotionID,
		DurationDays:         &order.DurationDays,
		OrderType:            &order.OrderType,
		InvestmentType:       &order.InvestmentType,
		PaymentType:          &order.PaymentType,
		CoinTypeID:           &order.CoinTypeID,
		PaymentCoinTypeID:    &order.PaymentCoinTypeID,
		TransferAmount:       &order.TransferAmount,
		BalanceAmount:        &order.BalanceAmount,
		CoinUSDCurrency:      &order.CoinUSDCurrency,
		LocalCoinUSDCurrency: &order.LocalCoinUSDCurrency,
		LiveCoinUSDCurrency:  &order.LiveCoinUSDCurrency,
		PaymentAccountID:     &order.PaymentAccountID,
		PaymentStartAmount:   &order.PaymentStartAmount,
		StartMode:            &order.StartMode,
		StartAt:              &order.StartAt,
		EndAt:                &order.EndAt,
		PaymentState:         &order.PaymentState,
		AppGoodStockLockID:   &order.AppGoodStockLockID,
		LedgerLockID:         &order.LedgerLockID,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateWaitPayment
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateCheckPayment
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentTransferReceived
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentTransferReceivedCheck
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentTransferBookKept
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentTransferBookKeptCheck
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentBalanceSpent
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentBalanceSpentCheck
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateGoodStockTransferred
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateGoodStockTransferredCheck
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateCommissionAdded
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateCommissionAddedCheck
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateAchievementBookKept
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateAchievementBookKeptCheck
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaid
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateInService
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = order1.DeleteOrder(context.Background(), &ordermwpb.OrderReq{
			ID: &order.ID,
		})
	}
}

func createOutOfGas(t *testing.T) {
	var (
		req = npool.OutOfGasReq{
			ID:      &ret.ID,
			OrderID: &ret.OrderID,
			StartAt: &ret.StartAt,
			EndAt:   &ret.EndAt,
		}
	)

	info, err := CreateOutOfGas(context.Background(), &req)
	if assert.Nil(t, err) {
		if id != info.ID {
			ret.ID = info.ID
		}
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateOutOfGas(t *testing.T) {
	if ret.ID == id {
		var (
			startAt = now + 7*seconds
			endAt   = now + 8*seconds

			req = npool.OutOfGasReq{
				ID:      &ret.ID,
				StartAt: &startAt,
				EndAt:   &endAt,
			}
		)
		info, err := UpdateOutOfGas(context.Background(), &req)
		if assert.Nil(t, err) {
			ret.StartAt = info.StartAt
			ret.EndAt = info.EndAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getOutOfGas(t *testing.T) {
	info, err := GetOutOfGas(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getOutOfGases(t *testing.T) {
	infos, _, err := GetOutOfGases(context.Background(), &npool.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteOutOfGas(t *testing.T) {
	if ret.ID == id {
		info, err := DeleteOutOfGas(context.Background(), &npool.OutOfGasReq{
			ID: &ret.ID,
		})
		if assert.Nil(t, err) {
			assert.Equal(t, info, &ret)
		}

		info, err = GetOutOfGas(context.Background(), ret.ID)
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestOutOfGas(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setup(t)
	defer teardown(t)

	time.Sleep(10 * time.Second)
	t.Run("createOutOfGas", createOutOfGas)
	t.Run("updateOutOfGas", updateOutOfGas)
	t.Run("getOutOfGas", getOutOfGas)
	t.Run("getOutOfGases", getOutOfGases)
	t.Run("deleteOutOfGas", deleteOutOfGas)
}
