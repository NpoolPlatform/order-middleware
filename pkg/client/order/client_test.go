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
	orderbasetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
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
	ID:                      uuid.NewString(),
	AppID:                   uuid.NewString(),
	UserID:                  uuid.NewString(),
	GoodID:                  uuid.NewString(),
	Units:                   "10001.000000000000000000",
	OrderTypeStr:            orderbasetypes.OrderType_Normal.String(),
	OrderType:               orderbasetypes.OrderType_Normal,
	OrderStateStr:           orderbasetypes.OrderState_OrderStateWaitPayment.String(),
	OrderState:              orderbasetypes.OrderState_OrderStateWaitPayment,
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
	PaymentStateStr:         orderbasetypes.PaymentState_PaymentStateWait.String(),
	PaymentState:            orderbasetypes.PaymentState_PaymentStateWait,
	PayWithBalanceAmount:    "1008.000000000000000000",
	PaidAt:                  1009,
	PaymentStartAmount:      "1010.000000000000000000",
	PaymentFinishAmount:     "0.000000000000000000",
	SpecialOfferID:          uuid.NewString(),
	CreatedAt:               0,
	UserCanceled:            false,
	PayWithParent:           false,
}

var (
	promotionID = uuid.NewString()
	req         = npool.OrderReq{
		AppID:                     &ret.AppID,
		UserID:                    &ret.UserID,
		GoodID:                    &ret.GoodID,
		Units:                     &ret.Units,
		OrderType:                 &ret.OrderType,
		ParentOrderID:             &ret.ParentOrderID,
		PayWithParent:             &ret.PayWithParent,
		PaymentCoinID:             &ret.PaymentCoinTypeID,
		PayWithBalanceAmount:      &ret.PayWithBalanceAmount,
		PaymentAccountID:          &ret.PaymentAccountID,
		PaymentAmount:             &ret.PaymentAmount,
		PaymentAccountStartAmount: &ret.PaymentStartAmount,
		PaymentCoinUSDCurrency:    &ret.PaymentCoinUSDCurrency,
		PaymentLiveUSDCurrency:    &ret.PaymentLiveUSDCurrency,
		PaymentLocalUSDCurrency:   &ret.PaymentLocalUSDCurrency,
		SpecialOfferID:            &ret.SpecialOfferID,
		Start:                     &ret.Start,
		End:                       &ret.End,
		PromotionID:               &promotionID,
		ID:                        &ret.ID,
		PaymentID:                 &ret.PaymentID,
		Canceled:                  &ret.UserCanceled,
	}
)

var info *npool.Order

func create(t *testing.T) {
	var err error
	info, err = CreateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.PaidAt = info.PaidAt
		ret.CouponIDs = info.CouponIDs
		ret.CouponIDsStr = info.CouponIDsStr
		assert.Equal(t, info, &ret)
	}
}

func update(t *testing.T) {
	var err error
	info, err = UpdateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.PaidAt = info.PaidAt
		assert.Equal(t, info.String(), ret.String())
	}
}
func getOrder(t *testing.T) {
	var err error
	info, err = GetOrder(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.PaidAt = info.PaidAt
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
		ret.CreatedAt = infos[0].CreatedAt
		ret.PaidAt = infos[0].PaidAt
		assert.Equal(t, infos[0].String(), ret.String())
	}
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
}
