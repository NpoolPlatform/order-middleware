package order

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
	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	testinit "github.com/NpoolPlatform/order-middleware/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
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

var (
	now = uint32(time.Now().Unix())
	ret = npool.Order{
		ID:                   uuid.NewString(),
		AppID:                uuid.NewString(),
		UserID:               uuid.NewString(),
		GoodID:               uuid.NewString(),
		AppGoodID:            uuid.NewString(),
		Units:                "100.000000000000000000",
		GoodValue:            "1007.000000000000000000",
		GoodValueUSD:         "1007.000000000000000000",
		PaymentAmount:        "1011.000000000000000000",
		DiscountAmount:       "10.000000000000000000",
		PromotionID:          uuid.NewString(),
		DurationDays:         6 * secondsPerDay,
		OrderTypeStr:         ordertypes.OrderType_Normal.String(),
		OrderType:            ordertypes.OrderType_Normal,
		InvestmentType:       ordertypes.InvestmentType_FullPayment,
		InvestmentTypeStr:    ordertypes.InvestmentType_FullPayment.String(),
		PaymentTypeStr:       ordertypes.PaymentType_PayWithTransferOnly.String(),
		PaymentType:          ordertypes.PaymentType_PayWithTransferOnly,
		CoinTypeID:           uuid.NewString(),
		PaymentCoinTypeID:    uuid.NewString(),
		TransferAmount:       "1011.000000000000000000",
		BalanceAmount:        "0.000000000000000000",
		CoinUSDCurrency:      "1004.000000000000000000",
		LocalCoinUSDCurrency: "1005.000000000000000000",
		LiveCoinUSDCurrency:  "1006.000000000000000000",

		PaymentAccountID:   uuid.NewString(),
		PaymentStartAmount: "1010.000000000000000000",

		OrderStateStr:        ordertypes.OrderState_OrderStateCreated.String(),
		OrderState:           ordertypes.OrderState_OrderStateCreated,
		CancelStateStr:       ordertypes.OrderState_DefaultOrderState.String(),
		CancelState:          ordertypes.OrderState_DefaultOrderState,
		StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
		StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
		StartAt:              now + secondsPerDay,
		EndAt:                now + 6*secondsPerDay,
		LastBenefitAt:        0,
		BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
		BenefitState:         ordertypes.BenefitState_BenefitWait,
		UserSetPaid:          false,
		UserSetCanceled:      false,
		AdminSetCanceled:     false,
		PaymentTransactionID: "",
		PaymentFinishAmount:  "0.000000000000000000",
		PaymentStateStr:      ordertypes.PaymentState_PaymentStateWait.String(),
		PaymentState:         ordertypes.PaymentState_PaymentStateWait,
		OutOfGasHours:        0,
		CompensateHours:      0,
	}

	req = npool.OrderReq{
		ID:                   &ret.ID,
		AppID:                &ret.AppID,
		UserID:               &ret.UserID,
		GoodID:               &ret.GoodID,
		AppGoodID:            &ret.AppGoodID,
		Units:                &ret.Units,
		GoodValue:            &ret.GoodValue,
		GoodValueUSD:         &ret.GoodValueUSD,
		PaymentAmount:        &ret.PaymentAmount,
		DiscountAmount:       &ret.DiscountAmount,
		PromotionID:          &ret.PromotionID,
		DurationDays:         &ret.DurationDays,
		OrderType:            &ret.OrderType,
		InvestmentType:       &ret.InvestmentType,
		PaymentType:          &ret.PaymentType,
		CoinTypeID:           &ret.CoinTypeID,
		PaymentCoinTypeID:    &ret.PaymentCoinTypeID,
		TransferAmount:       &ret.TransferAmount,
		BalanceAmount:        &ret.BalanceAmount,
		CoinUSDCurrency:      &ret.CoinUSDCurrency,
		LocalCoinUSDCurrency: &ret.LocalCoinUSDCurrency,
		LiveCoinUSDCurrency:  &ret.LiveCoinUSDCurrency,
		PaymentAccountID:     &ret.PaymentAccountID,
		PaymentStartAmount:   &ret.PaymentStartAmount,
		OrderState:           &ret.OrderState,
		StartMode:            &ret.StartMode,
		StartAt:              &ret.StartAt,
		EndAt:                &ret.EndAt,
		LastBenefitAt:        &ret.LastBenefitAt,
		BenefitState:         &ret.BenefitState,
		UserSetPaid:          &ret.UserSetPaid,
		UserSetCanceled:      &ret.UserSetCanceled,
		AdminSetCanceled:     &ret.AdminSetCanceled,
		PaymentTransactionID: &ret.PaymentTransactionID,
		PaymentFinishAmount:  &ret.PaymentFinishAmount,
		PaymentState:         &ret.PaymentState,
		OutOfGasHours:        &ret.OutOfGasHours,
		CompensateHours:      &ret.CompensateHours,
	}
)

var info *npool.Order

func create(t *testing.T) {
	var err error
	info, err = CreateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.CouponIDs = info.CouponIDs
		ret.CouponIDsStr = info.CouponIDsStr
		ret.PaymentID = info.PaymentID
		ret.ParentOrderID = info.ParentOrderID
		assert.Equal(t, info, &ret)
	}
}

func update(t *testing.T) {
	var err error
	var (
		req = npool.OrderReq{
			ID: &ret.ID,
		}
	)
	info, err = UpdateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}
func getOrder(t *testing.T) {
	var err error
	info, err = GetOrder(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func getOrders(t *testing.T) {
	infos, _, err := GetOrders(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func deleteOrder(t *testing.T) {
	var err error
	info, err := DeleteOrder(context.Background(), &npool.OrderReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}

	info, err = GetOrder(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDetail(t *testing.T) {
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

	t.Run("create", create)
	t.Run("update", update)
	t.Run("getOrder", getOrder)
	t.Run("getOrders", getOrders)
	t.Run("deleteOrder", deleteOrder)
}
