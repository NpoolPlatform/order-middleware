package compensate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	ordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"

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

const secondsPerDay = 24 * 60 * 60
const secondsPerHours = 60 * 60

var (
	now   = uint32(time.Now().Unix())
	order = ordermwpb.Order{
		ID:                   uuid.NewString(),
		AppID:                uuid.NewString(),
		UserID:               uuid.NewString(),
		GoodID:               uuid.NewString(),
		AppGoodID:            uuid.NewString(),
		ParentOrderID:        uuid.NewString(),
		Units:                "10001.000000000000000000",
		GoodValue:            "1007.000000000000000000",
		GoodValueUSD:         "1007.000000000000000000",
		PaymentAmount:        "1007.000000000000000000",
		DiscountAmount:       "10.000000000000000000",
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
		TransferAmount:       "1011.000000000000000000",
		BalanceAmount:        "110.000000000000000000",
		CoinUSDCurrency:      "1004.000000000000000000",
		LocalCoinUSDCurrency: "1005.000000000000000000",
		LiveCoinUSDCurrency:  "1006.000000000000000000",

		PaymentAccountID:   uuid.NewString(),
		PaymentStartAmount: "1010.000000000000000000",

		OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
		OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
		StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
		StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
		StartAt:              now - 5*secondsPerDay,
		EndAt:                now + 5*secondsPerDay,
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
	ret = npool.Compensate{
		ID:           uuid.NewString(),
		OrderID:      order.ID,
		AppID:        order.AppID,
		UserID:       order.UserID,
		GoodID:       order.GoodID,
		AppGoodID:    order.AppGoodID,
		Units:        order.Units,
		OrderStartAt: order.StartAt,
		OrderEndAt:   order.EndAt,
		StartAt:      now - secondsPerDay,
		EndAt:        now - secondsPerDay + 2*secondsPerHours,
		Title:        "title " + uuid.NewString(),
		Message:      "message " + uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := order1.NewHandler(
		context.Background(),
		order1.WithID(&order.ID, false),
		order1.WithAppID(&order.AppID, true),
		order1.WithUserID(&order.UserID, true),
		order1.WithGoodID(&order.GoodID, true),
		order1.WithAppGoodID(&order.AppGoodID, true),
		order1.WithParentOrderID(&order.ParentOrderID, false),
		order1.WithUnits(&order.Units, true),
		order1.WithGoodValue(&order.GoodValue, true),
		order1.WithGoodValueUSD(&order.GoodValueUSD, true),
		order1.WithPaymentAmount(&order.PaymentAmount, true),
		order1.WithDiscountAmount(&order.DiscountAmount, false),
		order1.WithPromotionID(&order.PromotionID, false),
		order1.WithDurationDays(&order.DurationDays, true),
		order1.WithOrderType(&order.OrderType, true),
		order1.WithInvestmentType(&order.InvestmentType, true),
		order1.WithCouponIDs(order.CouponIDs, false),
		order1.WithPaymentType(&order.PaymentType, true),
		order1.WithCoinTypeID(&order.CoinTypeID, true),
		order1.WithPaymentCoinTypeID(&order.PaymentCoinTypeID, true),
		order1.WithTransferAmount(&order.TransferAmount, true),
		order1.WithBalanceAmount(&order.BalanceAmount, true),
		order1.WithCoinUSDCurrency(&order.CoinUSDCurrency, true),
		order1.WithLocalCoinUSDCurrency(&order.LocalCoinUSDCurrency, true),
		order1.WithLiveCoinUSDCurrency(&order.LiveCoinUSDCurrency, true),
		order1.WithPaymentAccountID(&order.PaymentAccountID, false),
		order1.WithPaymentStartAmount(&order.PaymentStartAmount, false),
		order1.WithStartMode(&order.StartMode, true),
		order1.WithStartAt(&order.StartAt, true),
		order1.WithEndAt(&order.EndAt, true),
		order1.WithPaymentState(&order.PaymentState, false),
	)
	assert.Nil(t, err)

	_, err = h1.CreateOrder(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteOrder(context.Background())
	}
}

func createCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOrderID(&ret.OrderID, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithTitle(&ret.Title, true),
		WithMessage(&ret.Message, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateCompensate(context.Background())
		if assert.Nil(t, err) {
			ret.OrderEndAt = info.OrderEndAt
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateCompensate(t *testing.T) {
	ret.Message = "change message " + uuid.NewString()
	ret.StartAt = now - secondsPerDay + 2*secondsPerHours
	ret.EndAt = now - secondsPerDay + 6*secondsPerHours
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOrderID(&ret.OrderID, false),
		WithStartAt(&ret.StartAt, false),
		WithEndAt(&ret.EndAt, false),
		WithMessage(&ret.Message, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateCompensate(context.Background())
		if assert.Nil(t, err) {
			ret.OrderEndAt = info.OrderEndAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetCompensate(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getCompensates(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetCompensates(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteCompensate(context.Background())
		if assert.Nil(t, err) {
			ret.OrderEndAt = info.OrderEndAt
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetCompensate(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestCompensate(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createCompensate", createCompensate)
	t.Run("updateCompensate", updateCompensate)
	t.Run("getCompensate", getCompensate)
	t.Run("getCompensates", getCompensates)
	t.Run("deleteCompensate", deleteCompensate)
}
