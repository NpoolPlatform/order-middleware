package compensate

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

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

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
	id  = uuid.NewString()
	ret = npool.Compensate{
		EntID:        id,
		OrderID:      order.EntID,
		AppID:        order.AppID,
		UserID:       order.UserID,
		GoodID:       order.GoodID,
		AppGoodID:    order.AppGoodID,
		Units:        order.Units,
		OrderStartAt: order.StartAt,
		OrderEndAt:   order.EndAt,
		StartAt:      now + 6*seconds,
		EndAt:        now + 7*seconds,
		Title:        "Title " + uuid.NewString(),
		Message:      "Message " + uuid.NewString(),
	}
)

//nolint:funlen
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

	order.OrderState = ordertypes.OrderState_OrderStateWaitPayment
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

	order.OrderState = ordertypes.OrderState_OrderStatePaymentTransferBookKeeping
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentSpendBalance
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateTransferGoodStockLocked
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateAddCommission
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateAchievementBookKeeping
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStatePaymentUnlockAccount
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateUpdatePaidChilds
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

	order.OrderState = ordertypes.OrderState_OrderStateTransferGoodStockWaitStart
	_, err = order1.UpdateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:         &order.ID,
		OrderState: &order.OrderState,
	})
	assert.Nil(t, err)

	order.OrderState = ordertypes.OrderState_OrderStateUpdateInServiceChilds
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

func createCompensate(t *testing.T) {
	var (
		req = npool.CompensateReq{
			ID:      &ret.ID,
			OrderID: &ret.OrderID,
			StartAt: &ret.StartAt,
			EndAt:   &ret.EndAt,
			Title:   &ret.Title,
			Message: &ret.Message,
		}
	)

	info, err := CreateCompensate(context.Background(), &req)
	if assert.Nil(t, err) {
		if id != info.EntID {
			ret.EntID = info.EntID
		}
		ret.OrderEndAt = info.OrderEndAt
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateCompensate(t *testing.T) {
	if ret.EntID == id {
		var (
			startAt = now + 7*seconds
			endAt   = now + 8*seconds
			message = "Message update" + uuid.NewString()

			req = npool.CompensateReq{
				ID:      &ret.ID,
				StartAt: &startAt,
				EndAt:   &endAt,
				Message: &message,
			}
		)
		info, err := UpdateCompensate(context.Background(), &req)
		if assert.Nil(t, err) {
			ret.StartAt = info.StartAt
			ret.EndAt = info.EndAt
			ret.OrderEndAt = info.OrderEndAt
			ret.Message = info.Message
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getCompensate(t *testing.T) {
	info, err := GetCompensate(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getCompensates(t *testing.T) {
	infos, _, err := GetCompensates(context.Background(), &npool.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteCompensate(t *testing.T) {
	if ret.EntID == id {
		info, err := DeleteCompensate(context.Background(), &npool.CompensateReq{
			ID: &ret.ID,
		})
		if assert.Nil(t, err) {
			ret.OrderEndAt = info.OrderEndAt
			assert.Equal(t, info, &ret)
		}

		info, err = GetCompensate(context.Background(), ret.EntID)
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestCompensate(t *testing.T) {
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
	t.Run("createCompensate", createCompensate)
	t.Run("updateCompensate", updateCompensate)
	t.Run("getCompensate", getCompensate)
	t.Run("getCompensates", getCompensates)
	t.Run("deleteCompensate", deleteCompensate)
}
