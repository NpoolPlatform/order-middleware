package appconfig

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
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

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

var ret = npool.AppConfig{
	EntID:                                  uuid.NewString(),
	AppID:                                  uuid.NewString(),
	EnableSimulateOrder:                    false,
	SimulateOrderCouponMode:                ordertypes.SimulateOrderCouponMode_WithoutCoupon,
	SimulateOrderCouponProbability:         "1",
	SimulateOrderCashableProfitProbability: "0.1",
	MaxUnpaidOrders:                        10,
	MaxTypedCouponsPerOrder:                10,
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	ret.SimulateOrderCouponModeStr = ret.SimulateOrderCouponMode.String()
	return func(*testing.T) {}
}

func createAppConfig(t *testing.T) {
	err := CreateAppConfig(context.Background(), &npool.AppConfigReq{
		EntID:                                  &ret.EntID,
		AppID:                                  &ret.AppID,
		EnableSimulateOrder:                    &ret.EnableSimulateOrder,
		SimulateOrderCouponMode:                &ret.SimulateOrderCouponMode,
		SimulateOrderCouponProbability:         &ret.SimulateOrderCouponProbability,
		SimulateOrderCashableProfitProbability: &ret.SimulateOrderCashableProfitProbability,
		MaxUnpaidOrders:                        &ret.MaxUnpaidOrders,
		MaxTypedCouponsPerOrder:                &ret.MaxTypedCouponsPerOrder,
	})
	if assert.Nil(t, err) {
		info, err := GetAppConfig(context.Background(), ret.AppID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateAppConfig(t *testing.T) {
	ret.MaxUnpaidOrders = 20
	err := UpdateAppConfig(context.Background(), &npool.AppConfigReq{
		ID:                                     &ret.ID,
		EntID:                                  &ret.EntID,
		AppID:                                  &ret.AppID,
		EnableSimulateOrder:                    &ret.EnableSimulateOrder,
		SimulateOrderCouponMode:                &ret.SimulateOrderCouponMode,
		SimulateOrderCouponProbability:         &ret.SimulateOrderCouponProbability,
		SimulateOrderCashableProfitProbability: &ret.SimulateOrderCashableProfitProbability,
		MaxUnpaidOrders:                        &ret.MaxUnpaidOrders,
		MaxTypedCouponsPerOrder:                &ret.MaxTypedCouponsPerOrder,
	})
	if assert.Nil(t, err) {
		info, err := GetAppConfig(context.Background(), ret.AppID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getAppConfig(t *testing.T) {
	info, err := GetAppConfig(context.Background(), ret.AppID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getAppConfigs(t *testing.T) {
	infos, _, err := GetAppConfigs(context.Background(), &npool.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteAppConfig(t *testing.T) {
	err := DeleteAppConfig(context.Background(), &ret.ID, &ret.EntID, &ret.AppID)
	assert.Nil(t, err)

	info, err := GetAppConfig(context.Background(), ret.AppID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestAppConfig(t *testing.T) {
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

	t.Run("createAppConfig", createAppConfig)
	t.Run("updateAppConfig", updateAppConfig)
	t.Run("getAppConfig", getAppConfig)
	t.Run("getAppConfigs", getAppConfigs)
	t.Run("deleteAppConfig", deleteAppConfig)
}
