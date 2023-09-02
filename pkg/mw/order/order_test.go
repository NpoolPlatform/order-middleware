package order

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

const secondsPerDay = 24 * 60 * 60

var (
	now = uint32(time.Now().Unix())
	ret = npool.Order{
		ID:                   uuid.NewString(),
		AppID:                uuid.NewString(),
		UserID:               uuid.NewString(),
		GoodID:               uuid.NewString(),
		AppGoodID:            uuid.NewString(),
		ParentOrderID:        uuid.NewString(),
		Units:                "100.000000000000000000",
		GoodValue:            "1007.000000000000000000",
		GoodValueUSD:         "1007.000000000000000000",
		PaymentAmount:        "1121.000000000000000000",
		DiscountAmount:       "10.000000000000000000",
		PromotionID:          uuid.NewString(),
		DurationDays:         6 * secondsPerDay,
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

	appID         = uuid.NewString()
	userID        = uuid.NewString()
	parentOrderID = uuid.NewString()
	rets          = []npool.Order{
		{
			ID:                   parentOrderID,
			AppID:                appID,
			UserID:               userID,
			GoodID:               uuid.NewString(),
			AppGoodID:            uuid.NewString(),
			Units:                "1001.000000000000000000",
			GoodValue:            "1002.000000000000000000",
			GoodValueUSD:         "1002.000000000000000000",
			PaymentAmount:        "1115.000000000000000000",
			DiscountAmount:       "101.000000000000000000",
			PromotionID:          uuid.NewString(),
			DurationDays:         6 * secondsPerDay,
			OrderTypeStr:         ordertypes.OrderType_Normal.String(),
			OrderType:            ordertypes.OrderType_Normal,
			InvestmentType:       ordertypes.InvestmentType_FullPayment,
			InvestmentTypeStr:    ordertypes.InvestmentType_FullPayment.String(),
			PaymentTypeStr:       ordertypes.PaymentType_PayWithTransferAndBalance.String(),
			PaymentType:          ordertypes.PaymentType_PayWithTransferAndBalance,
			CoinTypeID:           uuid.NewString(),
			PaymentCoinTypeID:    uuid.NewString(),
			TransferAmount:       "1012.000000000000000000",
			BalanceAmount:        "103.000000000000000000",
			CoinUSDCurrency:      "11.000000000000000000",
			LocalCoinUSDCurrency: "12.000000000000000000",
			LiveCoinUSDCurrency:  "13.000000000000000000",

			PaymentAccountID:   uuid.NewString(),
			PaymentStartAmount: "1011.000000000000000000",

			OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
			OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
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
			CouponIDs:            []string{uuid.NewString(), uuid.NewString()},
		},
		{
			ID:                   uuid.NewString(),
			AppID:                appID,
			UserID:               userID,
			GoodID:               uuid.NewString(),
			AppGoodID:            uuid.NewString(),
			ParentOrderID:        parentOrderID,
			Units:                "1011.000000000000000000",
			GoodValue:            "1012.000000000000000000",
			GoodValueUSD:         "1012.000000000000000000",
			PaymentAmount:        "0.000000000000000000",
			DiscountAmount:       "0.000000000000000000",
			PromotionID:          uuid.NewString(),
			DurationDays:         3 * secondsPerDay,
			OrderTypeStr:         ordertypes.OrderType_Normal.String(),
			OrderType:            ordertypes.OrderType_Normal,
			InvestmentType:       ordertypes.InvestmentType_FullPayment,
			InvestmentTypeStr:    ordertypes.InvestmentType_FullPayment.String(),
			PaymentTypeStr:       ordertypes.PaymentType_PayWithParentOrder.String(),
			PaymentType:          ordertypes.PaymentType_PayWithParentOrder,
			CoinTypeID:           uuid.NewString(),
			PaymentCoinTypeID:    uuid.NewString(),
			TransferAmount:       "0.000000000000000000",
			BalanceAmount:        "0.000000000000000000",
			CoinUSDCurrency:      "1003.000000000000000000",
			LocalCoinUSDCurrency: "1004.000000000000000000",
			LiveCoinUSDCurrency:  "1005.000000000000000000",

			OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
			OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
			CancelStateStr:       ordertypes.OrderState_DefaultOrderState.String(),
			CancelState:          ordertypes.OrderState_DefaultOrderState,
			StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
			StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
			StartAt:              now + secondsPerDay,
			EndAt:                now + 3*secondsPerDay,
			LastBenefitAt:        0,
			BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
			BenefitState:         ordertypes.BenefitState_BenefitWait,
			UserSetPaid:          false,
			UserSetCanceled:      false,
			PaymentTransactionID: "",
			PaymentFinishAmount:  "0.000000000000000000",
			PaymentStateStr:      ordertypes.PaymentState_PaymentStateNoPayment.String(),
			PaymentState:         ordertypes.PaymentState_PaymentStateNoPayment,
			OutOfGasHours:        0,
			CompensateHours:      0,
		},
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
		WithGoodValueUSD(&ret.GoodValueUSD, true),
		WithPaymentAmount(&ret.PaymentAmount, true),
		WithDiscountAmount(&ret.DiscountAmount, false),
		WithPromotionID(&ret.PromotionID, false),
		WithDurationDays(&ret.DurationDays, true),
		WithOrderType(&ret.OrderType, true),
		WithInvestmentType(&ret.InvestmentType, true),
		WithCouponIDs(ret.CouponIDs, false),
		WithPaymentType(&ret.PaymentType, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithPaymentCoinTypeID(&ret.PaymentCoinTypeID, true),
		WithTransferAmount(&ret.TransferAmount, true),
		WithBalanceAmount(&ret.BalanceAmount, true),
		WithCoinUSDCurrency(&ret.CoinUSDCurrency, true),
		WithLocalCoinUSDCurrency(&ret.LocalCoinUSDCurrency, true),
		WithLiveCoinUSDCurrency(&ret.LiveCoinUSDCurrency, true),
		WithPaymentAccountID(&ret.PaymentAccountID, true),
		WithPaymentStartAmount(&ret.PaymentStartAmount, true),
		WithStartMode(&ret.StartMode, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithPaymentState(&ret.PaymentState, false),
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

func createOrders(t *testing.T) {
	retsReq := []*npool.OrderReq{}
	for key := range rets {
		retReq := npool.OrderReq{
			ID:                   &rets[key].ID,
			AppID:                &rets[key].AppID,
			UserID:               &rets[key].UserID,
			GoodID:               &rets[key].GoodID,
			AppGoodID:            &rets[key].AppGoodID,
			Units:                &rets[key].Units,
			GoodValue:            &rets[key].GoodValue,
			GoodValueUSD:         &rets[key].GoodValueUSD,
			PaymentAmount:        &rets[key].PaymentAmount,
			DiscountAmount:       &rets[key].DiscountAmount,
			PromotionID:          &rets[key].PromotionID,
			DurationDays:         &rets[key].DurationDays,
			OrderType:            &rets[key].OrderType,
			InvestmentType:       &rets[key].InvestmentType,
			PaymentType:          &rets[key].PaymentType,
			OrderState:           &rets[key].OrderState,
			StartMode:            &rets[key].StartMode,
			StartAt:              &rets[key].StartAt,
			EndAt:                &rets[key].EndAt,
			LastBenefitAt:        &rets[key].LastBenefitAt,
			BenefitState:         &rets[key].BenefitState,
			UserSetPaid:          &rets[key].UserSetPaid,
			UserSetCanceled:      &rets[key].UserSetCanceled,
			AdminSetCanceled:     &rets[key].AdminSetCanceled,
			PaymentTransactionID: &rets[key].PaymentTransactionID,
			PaymentFinishAmount:  &rets[key].PaymentFinishAmount,
			PaymentState:         &rets[key].PaymentState,
			OutOfGasHours:        &rets[key].OutOfGasHours,
			CompensateHours:      &rets[key].CompensateHours,
			CouponIDs:            rets[key].CouponIDs,
			CoinTypeID:           &rets[key].CoinTypeID,
			PaymentCoinTypeID:    &rets[key].PaymentCoinTypeID,
			TransferAmount:       &rets[key].TransferAmount,
			BalanceAmount:        &rets[key].BalanceAmount,
			CoinUSDCurrency:      &rets[key].CoinUSDCurrency,
			LocalCoinUSDCurrency: &rets[key].LocalCoinUSDCurrency,
			LiveCoinUSDCurrency:  &rets[key].LiveCoinUSDCurrency,
		}
		if rets[key].TransferAmount != "" {
			retReq.PaymentAccountID = &rets[key].PaymentAccountID
			retReq.PaymentStartAmount = &rets[key].PaymentStartAmount
		} else {
			retReq.ParentOrderID = &rets[key].ParentOrderID
		}
		retsReq = append(retsReq, &retReq)
	}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(retsReq, true),
	)
	if assert.Nil(t, err) {
		infos, err := handler.CreateOrders(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, len(rets), len(infos))
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
	for key := range rets {
		handler, err := NewHandler(
			context.Background(),
			WithID(&rets[key].ID, true),
		)
		assert.Nil(t, err)
		info, err := handler.DeleteOrder(context.Background())
		if assert.Nil(t, err) {
			assert.NotEqual(t, info, nil)
		}
	}
}

func TestOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createOrder", createOrder)
	t.Run("createOrders", createOrders)
	t.Run("updateOrder", updateOrder)
	t.Run("getOrder", getOrder)
	t.Run("getOrders", getOrders)
	t.Run("deleteOrder", deleteOrder)
}
