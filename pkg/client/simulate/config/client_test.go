package config

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

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"

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

var (
	ret = npool.SimulateConfig{
		EntID:                     uuid.NewString(),
		AppID:                     uuid.NewString(),
		CashableProfitProbability: "1",
		SendCouponModeStr:         ordertypes.SendCouponMode_WithoutCoupon.String(),
		SendCouponMode:            ordertypes.SendCouponMode_WithoutCoupon,
		SendCouponProbability:     "1",
		Enabled:                   false,
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createSimulateConfig(t *testing.T) {
	var (
		req = npool.SimulateConfigReq{
			EntID:                     &ret.EntID,
			AppID:                     &ret.AppID,
			CashableProfitProbability: &ret.CashableProfitProbability,
			SendCouponMode:            &ret.SendCouponMode,
			SendCouponProbability:     &ret.SendCouponProbability,
			Enabled:                   &ret.Enabled,
		}
	)

	info, err := CreateSimulateConfig(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, info, &ret)
	}
}

func updateSimulateConfig(t *testing.T) {
	ret.CashableProfitProbability = "0.5"
	ret.SendCouponMode = ordertypes.SendCouponMode_RandomBenifit
	ret.SendCouponModeStr = ordertypes.SendCouponMode_RandomBenifit.String()
	ret.SendCouponProbability = "0.5"
	var (
		req = npool.SimulateConfigReq{
			ID:                        &ret.ID,
			CashableProfitProbability: &ret.CashableProfitProbability,
			SendCouponMode:            &ret.SendCouponMode,
			SendCouponProbability:     &ret.SendCouponProbability,
		}
	)
	info, err := UpdateSimulateConfig(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getSimulateConfig(t *testing.T) {
	info, err := GetSimulateConfig(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getSimulateConfigs(t *testing.T) {
	infos, _, err := GetSimulateConfigs(context.Background(), &npool.Conds{
		ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		SendCouponMode: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SendCouponMode)},
		Enabled:        &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteSimulateConfig(t *testing.T) {
	info, err := DeleteSimulateConfig(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = GetSimulateConfig(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestSimulateConfig(t *testing.T) {
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

	time.Sleep(10 * time.Second)
	t.Run("createSimulateConfig", createSimulateConfig)
	t.Run("updateSimulateConfig", updateSimulateConfig)
	t.Run("getSimulateConfig", getSimulateConfig)
	t.Run("getSimulateConfigs", getSimulateConfigs)
	t.Run("deleteSimulateConfig", deleteSimulateConfig)
}
