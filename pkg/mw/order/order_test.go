package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

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

var (
	ret = npool.Order{
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
)

func createOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, false),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithParentOrderID(&ret.ParentOrderID, false),
		WithUnits(&ret.Units, true),
		WithGoodValue(&ret.GoodValue, true),
		WithPaymentAmount(&ret.PaymentAmount, true),
		WithDiscountAmount(&ret.DiscountAmount, false),
		WithPromotionID(&ret.PromotionID, false),
		WithDurationDays(&ret.DurationDays, true),
		WithOrderType(&ret.OrderType, true),
		WithInvestmentType(&ret.InvestmentType, true),
		WithCouponIDs(ret.CouponIDs, false),
		WithPaymentType(&ret.PaymentType, true),
		WithPaymentAccountID(&ret.PaymentAccountID, false),
		WithPaymentCoinTypeID(&ret.PaymentCoinTypeID, false),
		WithPaymentStartAmount(&ret.PaymentStartAmount, false),
		WithPaymentTransferAmount(&ret.PaymentTransferAmount, false),
		WithPaymentBalanceAmount(&ret.PaymentBalanceAmount, false),
		WithPaymentCoinUSDCurrency(&ret.PaymentCoinUSDCurrency, false),
		WithPaymentLocalCoinUSDCurrency(&ret.PaymentLocalCoinUSDCurrency, false),
		WithPaymentLiveCoinUSDCurrency(&ret.PaymentLiveCoinUSDCurrency, false),
		WithStartMode(&ret.StartMode, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithLastBenefitAt(&ret.LastBenefitAt, false),
		WithBenefitState(&ret.BenefitState, false),
		WithUserSetPaid(&ret.UserSetPaid, false),
		WithUserSetCanceled(&ret.UserSetCanceled, false),
		WithPaymentTransactionID(&ret.PaymentTransactionID, false),
		WithPaymentFinishAmount(&ret.PaymentFinishAmount, false),
		WithOutOfGasHours(&ret.OutOfGasHours, false),
		WithCompensateHours(&ret.CompensateHours, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateOrder(context.Background())
		if assert.Nil(t, err) {
			ret.CouponIDs = info.CouponIDs
			ret.CouponIDsStr = info.CouponIDsStr
			ret.PaymentID = info.PaymentID
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateOrder(t *testing.T) {
	ret.UserSetCanceled = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithUserSetCanceled(&ret.UserSetCanceled, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateOrder(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetOrder(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getOrders(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetOrders(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteOrder(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetOrder(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createOrder", createOrder)
	t.Run("updateOrder", updateOrder)
	t.Run("getOrder", getOrder)
	t.Run("getOrders", getOrders)
	t.Run("deleteOrder", deleteOrder)
}
