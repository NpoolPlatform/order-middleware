package order

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

var ret = npool.FeeOrder{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UserID:            uuid.NewString(),
	GoodID:            uuid.NewString(),
	GoodType:          goodtypes.GoodType_TechniqueServiceFee,
	AppGoodID:         uuid.NewString(),
	OrderID:           uuid.NewString(),
	ParentOrderID:     uuid.NewString(),
	OrderType:         types.OrderType_Normal,
	PaymentType:       types.PaymentType_PayWithBalanceOnly,
	CreateMethod:      types.OrderCreateMethod_OrderCreatedByAdmin,
	GoodValueUSD:      decimal.NewFromInt(120).String(),
	PaymentAmountUSD:  decimal.NewFromInt(110).String(),
	DiscountAmountUSD: decimal.NewFromInt(10).String(),
	PromotionID:       uuid.NewString(),
	DurationSeconds:   100000,
	LedgerLockID:      uuid.NewString(),
	PaymentID:         uuid.NewString(),
	Coupons: []*ordercouponmwpb.OrderCouponInfo{
		{
			CouponID: uuid.NewString(),
		},
	},
	PaymentBalances: []*paymentmwpb.PaymentBalanceInfo{
		{
			CoinTypeID:           uuid.NewString(),
			Amount:               decimal.NewFromInt(110).String(),
			CoinUSDCurrency:      decimal.NewFromInt(1).String(),
			LocalCoinUSDCurrency: decimal.NewFromInt(1).String(),
			LiveCoinUSDCurrency:  decimal.NewFromInt(1).String(),
		},
	},
	OrderState:   types.OrderState_OrderStateCreated,
	PaymentState: types.PaymentState_PaymentStateWait,
}

func setup(t *testing.T) func(*testing.T) {
	for _, paymentBalance := range ret.PaymentBalances {
		paymentBalance.PaymentID = ret.PaymentID
	}
	for _, orderCoupon := range ret.Coupons {
		orderCoupon.OrderID = ret.OrderID
	}

	ret.GoodTypeStr = ret.GoodType.String()
	ret.OrderTypeStr = ret.OrderType.String()
	ret.PaymentTypeStr = ret.PaymentType.String()
	ret.CreateMethodStr = ret.CreateMethod.String()
	ret.OrderStateStr = ret.OrderState.String()
	ret.PaymentStateStr = ret.PaymentState.String()
	ret.CancelStateStr = ret.CancelState.String()

	return func(*testing.T) {
		_ = h1.DeleteOrderBase(context.Background())
	}
}

func createOrder(t *testing.T) {
	err := CreateOrder(context.Background(), &npool.FeeOrderReq{})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateOrder(t *testing.T) {
	var err error
	var (
		req = npool.OrderReq{
			ID: &ret.ID,
		}
	)
	info, err = UpdateOrder(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}
func getOrder(t *testing.T) {
	var err error
	info, err = GetOrder(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func getOrders(t *testing.T) {
	infos, _, err := GetOrders(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.EntID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func deleteOrder(t *testing.T) {
	var err error
	info, err := DeleteOrder(context.Background(), &npool.OrderReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}

	info, err = GetOrder(context.Background(), ret.EntID)
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

	t.Run("createOrder", createOrder)
	t.Run("updateOrder", updateOrder)
	t.Run("getOrder", getOrder)
	t.Run("getOrders", getOrders)
	t.Run("deleteOrder", deleteOrder)
}
