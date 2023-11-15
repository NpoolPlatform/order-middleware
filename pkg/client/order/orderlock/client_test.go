package orderlock

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

	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/client/order"
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
const seconds = 1

var (
	now   = uint32(time.Now().Unix())
	order = ordermwpb.Order{
		EntID:                uuid.NewString(),
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
	ret = npool.OrderLock{
		EntID:       uuid.NewString(),
		AppID:       order.AppID,
		UserID:      order.UserID,
		OrderID:     order.EntID,
		LockTypeStr: ordertypes.OrderLockType_LockCommission.String(),
		LockType:    ordertypes.OrderLockType_LockCommission,
	}

	reqs = []*npool.OrderLockReq{
		{
			EntID:    &ret.EntID,
			AppID:    &ret.AppID,
			UserID:   &ret.UserID,
			OrderID:  &ret.OrderID,
			LockType: &ret.LockType,
		},
	}
)

func setup(t *testing.T) func(*testing.T) {
	info, err := order1.CreateOrder(context.Background(), &ordermwpb.OrderReq{
		EntID:                &order.EntID,
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
	order.ID = info.ID

	return func(*testing.T) {
		_, _ = order1.DeleteOrder(context.Background(), &ordermwpb.OrderReq{
			ID: &order.ID,
		})
	}
}

func creates(t *testing.T) {
	var err error
	infos, err := CreateOrderLocks(context.Background(), reqs)
	if assert.Nil(t, err) {
		ret.CreatedAt = infos[0].CreatedAt
		ret.UpdatedAt = infos[0].UpdatedAt
		ret.ID = infos[0].ID
		assert.Equal(t, infos[0], &ret)
	}
}

func getOrderLock(t *testing.T) {
	var err error
	info, err := GetOrderLock(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func getOrderLocks(t *testing.T) {
	infos, _, err := GetOrderLocks(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func deleteOrderLocks(t *testing.T) {
	var err error
	reqs[0].ID = &ret.ID
	infos, err := DeleteOrderLocks(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
	}

	info, err := GetOrderLock(context.Background(), ret.EntID)
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

	teardown := setup(t)
	defer teardown(t)

	t.Run("creates", creates)
	t.Run("getOrderLock", getOrderLock)
	t.Run("getOrderLocks", getOrderLocks)
	t.Run("deleteOrderLocks", deleteOrderLocks)
}
