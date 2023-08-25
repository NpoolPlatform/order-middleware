package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

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

var ret = npool.Order{
	ID:                uuid.NewString(),
	AppID:             uuid.NewString(),
	UserID:            uuid.NewString(),
	GoodID:            uuid.NewString(),
	AppGoodID:         uuid.NewString(),
	ParentOrderID:     uuid.NewString(),
	Units:             "10001.000000000000000000",
	GoodValue:         "1007.000000000000000000",
	PaymentAmount:     "1007.000000000000000000",
	DiscountAmount:    "10.000000000000000000",
	PromotionID:       uuid.NewString(),
	DurationDays:      10006,
	OrderTypeStr:      ordertypes.OrderType_Normal.String(),
	OrderType:         ordertypes.OrderType_Normal,
	InvestmentType:    ordertypes.InvestmentType_FullPayment,
	InvestmentTypeStr: ordertypes.InvestmentType_FullPayment.String(),
	PaymentTypeStr:    ordertypes.PaymentType_PayWithTransferAndBalance.String(),
	PaymentType:       ordertypes.PaymentType_PayWithTransferAndBalance,

	PaymentAccountID:            uuid.NewString(),
	PaymentCoinTypeID:           uuid.NewString(),
	PaymentStartAmount:          "1010.000000000000000000",
	PaymentTransferAmount:       "1011.000000000000000000",
	PaymentBalanceAmount:        "110.000000000000000000",
	PaymentCoinUSDCurrency:      "1004.000000000000000000",
	PaymentLocalCoinUSDCurrency: "1005.000000000000000000",
	PaymentLiveCoinUSDCurrency:  "1006.000000000000000000",

	OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
	OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
	StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
	StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
	StartAt:              10002,
	EndAt:                10003,
	LastBenefitAt:        10005,
	BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
	BenefitState:         ordertypes.BenefitState_BenefitWait,
	UserSetPaid:          false,
	UserSetCanceled:      false,
	PaymentTransactionID: "PaymentTransactionID" + uuid.NewString(),
	PaymentFinishAmount:  "0.000000000000000000",
	PaymentStateStr:      ordertypes.PaymentState_PaymentStateWait.String(),
	PaymentState:         ordertypes.PaymentState_PaymentStateWait,
	OutOfGasHours:        0,
	CompensateHours:      0,
}

var (
	req = npool.OrderReq{
		ID:                          &ret.ID,
		AppID:                       &ret.AppID,
		UserID:                      &ret.UserID,
		GoodID:                      &ret.GoodID,
		AppGoodID:                   &ret.AppGoodID,
		ParentOrderID:               &ret.ParentOrderID,
		Units:                       &ret.Units,
		GoodValue:                   &ret.GoodValue,
		PaymentAmount:               &ret.PaymentAmount,
		DiscountAmount:              &ret.DiscountAmount,
		PromotionID:                 &ret.PromotionID,
		DurationDays:                &ret.DurationDays,
		OrderType:                   &ret.OrderType,
		InvestmentType:              &ret.InvestmentType,
		PaymentType:                 &ret.PaymentType,
		PaymentAccountID:            &ret.PaymentAccountID,
		PaymentCoinTypeID:           &ret.PaymentCoinTypeID,
		PaymentStartAmount:          &ret.PaymentStartAmount,
		PaymentTransferAmount:       &ret.PaymentTransferAmount,
		PaymentBalanceAmount:        &ret.PaymentBalanceAmount,
		PaymentCoinUSDCurrency:      &ret.PaymentCoinUSDCurrency,
		PaymentLocalCoinUSDCurrency: &ret.PaymentLocalCoinUSDCurrency,
		PaymentLiveCoinUSDCurrency:  &ret.PaymentLiveCoinUSDCurrency,
		OrderState:                  &ret.OrderState,
		StartMode:                   &ret.StartMode,
		StartAt:                     &ret.StartAt,
		EndAt:                       &ret.EndAt,
		LastBenefitAt:               &ret.LastBenefitAt,
		BenefitState:                &ret.BenefitState,
		UserSetPaid:                 &ret.UserSetPaid,
		UserSetCanceled:             &ret.UserSetCanceled,
		PaymentTransactionID:        &ret.PaymentTransactionID,
		PaymentFinishAmount:         &ret.PaymentFinishAmount,
		PaymentState:                &ret.PaymentState,
		OutOfGasHours:               &ret.OutOfGasHours,
		CompensateHours:             &ret.CompensateHours,
	}
)

var info *npool.Order

func create(t *testing.T) {
	var err error
	info, err = CreateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.CouponIDs = info.CouponIDs
		ret.CouponIDsStr = info.CouponIDsStr
		assert.Equal(t, info, &ret)
	}
}

func update(t *testing.T) {
	var err error
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
	t.Run("create", create)
	t.Run("update", update)
	t.Run("getOrder", getOrder)
	t.Run("getOrders", getOrders)
	t.Run("deleteOrder", deleteOrder)
}
