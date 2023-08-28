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
const secondsPerHours = 60 * 60

var (
	now   = uint32(time.Now().Unix())
	order = ordermwpb.Order{
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
		DurationDays:      now + 5*secondsPerDay,
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
		StartAt:              now - 5*secondsPerDay,
		EndAt:                now + 5*secondsPerDay,
		LastBenefitAt:        now + 3*secondsPerDay,
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
	id  = uuid.NewString()
	ret = npool.Compensate{
		ID:           id,
		OrderID:      order.ID,
		AppID:        order.AppID,
		UserID:       order.UserID,
		GoodID:       order.GoodID,
		AppGoodID:    order.AppGoodID,
		Units:        order.Units,
		OrderStartAt: order.StartAt,
		OrderEndAt:   order.EndAt,
		StartAt:      now + secondsPerDay,
		EndAt:        now + secondsPerDay + 2*secondsPerHours,
		Message:      "Message " + uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	_, err := order1.CreateOrder(context.Background(), &ordermwpb.OrderReq{
		ID:                          &order.ID,
		AppID:                       &order.AppID,
		UserID:                      &order.UserID,
		GoodID:                      &order.GoodID,
		AppGoodID:                   &order.AppGoodID,
		ParentOrderID:               &order.ParentOrderID,
		Units:                       &order.Units,
		GoodValue:                   &order.GoodValue,
		PaymentAmount:               &order.PaymentAmount,
		DiscountAmount:              &order.DiscountAmount,
		PromotionID:                 &order.PromotionID,
		DurationDays:                &order.DurationDays,
		OrderType:                   &order.OrderType,
		InvestmentType:              &order.InvestmentType,
		PaymentType:                 &order.PaymentType,
		PaymentAccountID:            &order.PaymentAccountID,
		PaymentCoinTypeID:           &order.PaymentCoinTypeID,
		PaymentStartAmount:          &order.PaymentStartAmount,
		PaymentTransferAmount:       &order.PaymentTransferAmount,
		PaymentBalanceAmount:        &order.PaymentBalanceAmount,
		PaymentCoinUSDCurrency:      &order.PaymentCoinUSDCurrency,
		PaymentLocalCoinUSDCurrency: &order.PaymentLocalCoinUSDCurrency,
		PaymentLiveCoinUSDCurrency:  &order.PaymentLiveCoinUSDCurrency,
		OrderState:                  &order.OrderState,
		StartMode:                   &order.StartMode,
		StartAt:                     &order.StartAt,
		EndAt:                       &order.EndAt,
		LastBenefitAt:               &order.LastBenefitAt,
		BenefitState:                &order.BenefitState,
		UserSetPaid:                 &order.UserSetPaid,
		UserSetCanceled:             &order.UserSetCanceled,
		PaymentTransactionID:        &order.PaymentTransactionID,
		PaymentFinishAmount:         &order.PaymentFinishAmount,
		PaymentState:                &order.PaymentState,
		OutOfGasHours:               &order.OutOfGasHours,
		CompensateHours:             &order.CompensateHours,
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
			Message: &ret.Message,
		}
	)

	info, err := CreateCompensate(context.Background(), &req)
	if assert.Nil(t, err) {
		if id != info.ID {
			ret.ID = info.ID
		}
		ret.OrderEndAt = info.OrderEndAt
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateCompensate(t *testing.T) {
	if ret.ID == id {
		var (
			startAt = now + secondsPerDay + 2*secondsPerHours
			endAt   = now + secondsPerDay + 6*secondsPerHours
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
	info, err := GetCompensate(context.Background(), ret.ID)
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
	if ret.ID == id {
		info, err := DeleteCompensate(context.Background(), &npool.CompensateReq{
			ID: &ret.ID,
		})
		if assert.Nil(t, err) {
			ret.OrderEndAt = info.OrderEndAt
			assert.Equal(t, info, &ret)
		}

		info, err = GetCompensate(context.Background(), ret.ID)
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

	t.Run("createCompensate", createCompensate)
	t.Run("updateCompensate", updateCompensate)
	t.Run("getCompensate", getCompensate)
	t.Run("getCompensates", getCompensates)
	t.Run("deleteCompensate", deleteCompensate)
}
