package order

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	paymentmgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	testinit "github.com/NpoolPlatform/order-middleware/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	mgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var deviceInfo = npool.Order{
	ID:                      uuid.NewString(),
	AppID:                   uuid.NewString(),
	UserID:                  uuid.NewString(),
	GoodID:                  uuid.NewString(),
	Units:                   10001,
	OrderTypeStr:            mgrpb.OrderType_Normal.String(),
	OrderType:               mgrpb.OrderType_Normal,
	OrderStateStr:           mgrpb.OrderState_WaitPayment.String(),
	OrderState:              mgrpb.OrderState_WaitPayment,
	ParentOrderID:           uuid.NewString(),
	Start:                   10002,
	End:                     10003,
	PaymentCoinTypeID:       uuid.NewString(),
	PaymentCoinUSDCurrency:  "1004.000000000000000000",
	PaymentLocalUSDCurrency: "1005.000000000000000000",
	PaymentLiveUSDCurrency:  "1006.000000000000000000",
	PaymentID:               uuid.NewString(),
	PaymentAccountID:        uuid.NewString(),
	PaymentAmount:           "1007.000000000000000000",
	PaymentStateStr:         paymentmgrpb.PaymentState_Wait.String(),
	PaymentState:            paymentmgrpb.PaymentState_Wait,
	PayWithBalanceAmount:    "1008.000000000000000000",
	PaidAt:                  1009,
	PaymentStartAmount:      "1010.000000000000000000",
	PaymentFinishAmount:     "0.000000000000000000",
	FixAmountID:             uuid.NewString(),
	DiscountID:              uuid.NewString(),
	SpecialOfferID:          uuid.NewString(),
	UserCanceled:            false,
	PayWithParent:           false,
}

var (
	promotionID = uuid.NewString()
	req         = npool.OrderReq{
		AppID:                     &deviceInfo.AppID,
		UserID:                    &deviceInfo.UserID,
		GoodID:                    &deviceInfo.GoodID,
		Units:                     &deviceInfo.Units,
		OrderType:                 &deviceInfo.OrderType,
		ParentOrderID:             &deviceInfo.ParentOrderID,
		PayWithParent:             &deviceInfo.PayWithParent,
		PaymentCoinID:             &deviceInfo.PaymentCoinTypeID,
		PayWithBalanceAmount:      &deviceInfo.PayWithBalanceAmount,
		PaymentAccountID:          &deviceInfo.PaymentAccountID,
		PaymentAmount:             &deviceInfo.PaymentAmount,
		PaymentAccountStartAmount: &deviceInfo.PaymentStartAmount,
		PaymentCoinUSDCurrency:    &deviceInfo.PaymentCoinUSDCurrency,
		PaymentLiveUSDCurrency:    &deviceInfo.PaymentLiveUSDCurrency,
		PaymentLocalUSDCurrency:   &deviceInfo.PaymentLocalUSDCurrency,
		FixAmountID:               &deviceInfo.FixAmountID,
		DiscountID:                &deviceInfo.DiscountID,
		SpecialOfferID:            &deviceInfo.SpecialOfferID,
		Start:                     &deviceInfo.Start,
		End:                       &deviceInfo.End,
		PromotionID:               &promotionID,
		ID:                        &deviceInfo.ID,
		PaymentID:                 &deviceInfo.PaymentID,
		Canceled:                  &deviceInfo.UserCanceled,
	}
)

var info *npool.Order

func create(t *testing.T) {
	var err error
	info, err = CreateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		deviceInfo.CreatedAt = info.CreatedAt
		deviceInfo.PaidAt = info.PaidAt
		assert.Equal(t, info, &deviceInfo)
	}
}

func update(t *testing.T) {
	var err error
	info, err = UpdateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		deviceInfo.CreatedAt = info.CreatedAt
		deviceInfo.PaidAt = info.PaidAt
		assert.Equal(t, info.String(), deviceInfo.String())
	}
}
func getOrder(t *testing.T) {
	var err error
	info, err = GetOrder(context.Background(), deviceInfo.ID)
	if assert.Nil(t, err) {
		deviceInfo.CreatedAt = info.CreatedAt
		deviceInfo.PaidAt = info.PaidAt
		assert.Equal(t, info.String(), deviceInfo.String())
	}
}

func getAppOrders(t *testing.T) {
	infos, _, err := GetAppOrders(context.Background(), deviceInfo.AppID, 0, 1)
	if assert.Nil(t, err) {
		deviceInfo.CreatedAt = info.CreatedAt
		deviceInfo.PaidAt = info.PaidAt
		assert.Equal(t, infos[0].String(), deviceInfo.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("update", update)
	t.Run("getOrder", getOrder)
	t.Run("getAppOrders", getAppOrders)
}
