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
	promotionID = uuid.NewString()
	ret         = npool.Order{
		ID:                      uuid.NewString(),
		AppID:                   uuid.NewString(),
		UserID:                  uuid.NewString(),
		GoodID:                  uuid.NewString(),
		Units:                   "10001.000000000000000000",
		OrderTypeStr:            ordertypes.OrderType_Normal.String(),
		OrderType:               ordertypes.OrderType_Normal,
		OrderStateStr:           ordertypes.OrderState_OrderStateWaitPayment.String(),
		OrderState:              ordertypes.OrderState_OrderStateWaitPayment,
		ParentOrderID:           uuid.NewString(),
		ParentOrderGoodID:       "",
		Start:                   10002,
		End:                     10003,
		PaymentCoinTypeID:       uuid.NewString(),
		PaymentCoinUSDCurrency:  "1004.000000000000000000",
		PaymentLocalUSDCurrency: "1005.000000000000000000",
		PaymentLiveUSDCurrency:  "1006.000000000000000000",
		PaymentID:               uuid.NewString(),
		PaymentAccountID:        uuid.NewString(),
		PaymentAmount:           "1007.000000000000000000",
		PaymentStateStr:         ordertypes.PaymentState_PaymentStateWait.String(),
		PaymentState:            ordertypes.PaymentState_PaymentStateWait,
		PayWithBalanceAmount:    "1008.000000000000000000",
		PaidAt:                  1009,
		PaymentStartAmount:      "1010.000000000000000000",
		PaymentFinishAmount:     "0.000000000000000000",
		SpecialOfferID:          uuid.NewString(),
		CreatedAt:               0,
		UserCanceled:            false,
		PayWithParent:           false,
		InvestmentType:          ordertypes.InvestmentType_FullPayment,
		InvestmentTypeStr:       ordertypes.InvestmentType_FullPayment.String(),
	}
)

func createOrder(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, false),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithUnits(&ret.Units, true),
		WithType(&ret.OrderType, true),
		WithParentOrderID(&ret.ParentOrderID, true),
		WithPayWithParent(&ret.PayWithParent, true),
		WithPaymentCoinID(&ret.PaymentCoinTypeID, true),
		WithPayWithBalanceAmount(&ret.PayWithBalanceAmount, true),
		WithPaymentAccountID(&ret.PaymentAccountID, true),
		WithPaymentAmount(&ret.PaymentAmount, true),
		WithPaymentAccountStartAmount(&ret.PaymentStartAmount, true),
		WithPaymentCoinUSDCurrency(&ret.PaymentCoinUSDCurrency, true),
		WithPaymentLiveUSDCurrency(&ret.PaymentLiveUSDCurrency, true),
		WithPaymentLocalUSDCurrency(&ret.PaymentLocalUSDCurrency, true),
		WithUserSpecialReductionID(&ret.SpecialOfferID, true),
		WithStartAt(&ret.Start, true),
		WithEndAt(&ret.End, true),
		WithPromotionID(&promotionID, true),
		WithPaymentID(&ret.PaymentID, false),
		WithPaymentUserSetCanceled(&ret.UserCanceled, true),
		WithInvestmentType(&ret.InvestmentType, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateOrder(context.Background())
		if assert.Nil(t, err) {
			ret.PaidAt = info.PaidAt
			ret.CouponIDs = info.CouponIDs
			ret.CouponIDsStr = info.CouponIDsStr
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateOrder(t *testing.T) {
	ret.UserCanceled = true
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithPaymentID(&ret.PaymentID, true),
		WithPaymentUserSetCanceled(&ret.UserCanceled, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateOrder(context.Background())
		if assert.Nil(t, err) {
			ret.PaidAt = info.PaidAt
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
			ret.PaidAt = info.PaidAt
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
				ret.PaidAt = infos[0].PaidAt
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
