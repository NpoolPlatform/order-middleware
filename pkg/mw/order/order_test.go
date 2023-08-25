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

	appID         = uuid.NewString()
	userID        = uuid.NewString()
	parentOrderID = uuid.NewString()
	rets          = []npool.Order{
		{
			ID:                parentOrderID,
			AppID:             appID,
			UserID:            userID,
			GoodID:            uuid.NewString(),
			AppGoodID:         uuid.NewString(),
			Units:             "1001.000000000000000000",
			GoodValue:         "1002.000000000000000000",
			PaymentAmount:     "1003.000000000000000000",
			DiscountAmount:    "101.000000000000000000",
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
			PaymentStartAmount:          "1011.000000000000000000",
			PaymentTransferAmount:       "1012.000000000000000000",
			PaymentBalanceAmount:        "111.000000000000000000",
			PaymentCoinUSDCurrency:      "1003.000000000000000000",
			PaymentLocalCoinUSDCurrency: "1004.000000000000000000",
			PaymentLiveCoinUSDCurrency:  "1005.000000000000000000",

			OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
			OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
			StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
			StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
			StartAt:              10001,
			EndAt:                10002,
			LastBenefitAt:        10003,
			BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
			BenefitState:         ordertypes.BenefitState_BenefitWait,
			UserSetPaid:          false,
			UserSetCanceled:      false,
			PaymentTransactionID: "PaymentTransactionID1 " + uuid.NewString(),
			PaymentFinishAmount:  "0.000000000000000000",
			PaymentStateStr:      ordertypes.PaymentState_PaymentStateWait.String(),
			PaymentState:         ordertypes.PaymentState_PaymentStateWait,
			OutOfGasHours:        0,
			CompensateHours:      0,
			CouponIDs:            []string{uuid.NewString(), uuid.NewString()},
		},
		{
			ID:                uuid.NewString(),
			AppID:             appID,
			UserID:            userID,
			GoodID:            uuid.NewString(),
			AppGoodID:         uuid.NewString(),
			ParentOrderID:     parentOrderID,
			Units:             "1011.000000000000000000",
			GoodValue:         "1012.000000000000000000",
			PaymentAmount:     "1013.000000000000000000",
			DiscountAmount:    "102.000000000000000000",
			PromotionID:       uuid.NewString(),
			DurationDays:      10008,
			OrderTypeStr:      ordertypes.OrderType_Normal.String(),
			OrderType:         ordertypes.OrderType_Normal,
			InvestmentType:    ordertypes.InvestmentType_FullPayment,
			InvestmentTypeStr: ordertypes.InvestmentType_FullPayment.String(),
			PaymentTypeStr:    ordertypes.PaymentType_PayWithTransferAndBalance.String(),
			PaymentType:       ordertypes.PaymentType_PayWithTransferAndBalance,

			OrderStateStr:        ordertypes.OrderState_OrderStateWaitPayment.String(),
			OrderState:           ordertypes.OrderState_OrderStateWaitPayment,
			StartModeStr:         ordertypes.OrderStartMode_OrderStartConfirmed.String(),
			StartMode:            ordertypes.OrderStartMode_OrderStartConfirmed,
			StartAt:              10011,
			EndAt:                10012,
			LastBenefitAt:        10013,
			BenefitStateStr:      ordertypes.BenefitState_BenefitWait.String(),
			BenefitState:         ordertypes.BenefitState_BenefitWait,
			UserSetPaid:          false,
			UserSetCanceled:      false,
			PaymentTransactionID: "PaymentTransactionID2 " + uuid.NewString(),
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
			PaymentTransactionID: &rets[key].PaymentTransactionID,
			PaymentFinishAmount:  &rets[key].PaymentFinishAmount,
			PaymentState:         &rets[key].PaymentState,
			OutOfGasHours:        &rets[key].OutOfGasHours,
			CompensateHours:      &rets[key].CompensateHours,
			CouponIDs:            rets[key].CouponIDs,
		}
		if rets[key].PaymentTransferAmount != "" {
			retReq.PaymentAccountID = &rets[key].PaymentAccountID
			retReq.PaymentCoinTypeID = &rets[key].PaymentCoinTypeID
			retReq.PaymentStartAmount = &rets[key].PaymentStartAmount
			retReq.PaymentTransferAmount = &rets[key].PaymentTransferAmount
			retReq.PaymentBalanceAmount = &rets[key].PaymentBalanceAmount
			retReq.PaymentCoinUSDCurrency = &rets[key].PaymentCoinUSDCurrency
			retReq.PaymentLocalCoinUSDCurrency = &rets[key].PaymentLocalCoinUSDCurrency
			retReq.PaymentLiveCoinUSDCurrency = &rets[key].PaymentLiveCoinUSDCurrency
		} else {
			retReq.ParentOrderID = &rets[key].ParentOrderID
		}
		retsReq = append(retsReq, &retReq)
	}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(retsReq),
	)
	if assert.Nil(t, err) {
		infos, _, err := handler.CreateOrders(context.Background())
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
