package config

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
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"

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
	ret = npool.SimulateConfig{
		EntID:                 uuid.NewString(),
		AppID:                 uuid.NewString(),
		ProfitTxProbability:   "1",
		EnabledProfitTx:       true,
		SendCouponModeStr:     ordertypes.SendCouponMode_WithoutCoupon.String(),
		SendCouponMode:        ordertypes.SendCouponMode_WithoutCoupon,
		SendCouponProbability: "1",
		Enabled:               false,
	}
)

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createSimulateConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithEnabledProfitTx(&ret.EnabledProfitTx, true),
		WithProfitTxProbability(&ret.ProfitTxProbability, true),
		WithSendCouponMode(&ret.SendCouponMode, true),
		WithSendCouponProbability(&ret.SendCouponProbability, true),
		WithEnabled(&ret.Enabled, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateSimulateConfig(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateSimulateConfig(t *testing.T) {
	ret.EnabledProfitTx = false
	ret.ProfitTxProbability = "20"
	ret.SendCouponMode = ordertypes.SendCouponMode_RandomBenifit
	ret.SendCouponModeStr = ordertypes.SendCouponMode_RandomBenifit.String()
	ret.SendCouponProbability = "10"
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEnabledProfitTx(&ret.EnabledProfitTx, true),
		WithProfitTxProbability(&ret.ProfitTxProbability, true),
		WithSendCouponMode(&ret.SendCouponMode, true),
		WithSendCouponProbability(&ret.SendCouponProbability, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateSimulateConfig(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getSimulateConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetSimulateConfig(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getSimulateConfigs(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			SendCouponMode: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.SendCouponMode)},
			Enabled:        &basetypes.BoolVal{Op: cruder.EQ, Value: ret.Enabled},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetSimulateConfigs(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteSimulateConfig(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteSimulateConfig(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetSimulateConfig(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestSimulateConfig(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	time.Sleep(10 * time.Second)
	t.Run("createSimulateConfig", createSimulateConfig)
	t.Run("updateSimulateConfig", updateSimulateConfig)
	t.Run("getSimulateConfig", getSimulateConfig)
	t.Run("getSimulateConfigs", getSimulateConfigs)
	t.Run("deleteSimulateConfig", deleteSimulateConfig)
}
