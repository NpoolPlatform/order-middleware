package compensate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	goodtypes "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	orderbase1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderbase"

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

var ret = npool.Compensate{
	EntID:             uuid.NewString(),
	AppID:             uuid.NewString(),
	UserID:            uuid.NewString(),
	GoodID:            uuid.NewString(),
	AppGoodID:         uuid.NewString(),
	OrderID:           uuid.NewString(),
	CompensateFromID:  uuid.NewString(),
	CompensateType:    types.CompensateType_CompensateMalfunction,
	CompensateSeconds: 120,
}

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	ret.CompensateTypeStr = ret.CompensateType.String()

	h1, err := orderbase1.NewHandler(
		context.Background(),
		orderbase1.WithEntID(&ret.OrderID, false),
		orderbase1.WithAppID(&ret.AppID, true),
		orderbase1.WithUserID(&ret.UserID, true),
		orderbase1.WithGoodID(&ret.GoodID, true),
		orderbase1.WithGoodType(func() *goodtypes.GoodType { e := goodtypes.GoodType_PowerRental; return &e }(), true),
		orderbase1.WithAppGoodID(&ret.AppGoodID, true),
		orderbase1.WithOrderType(func() *types.OrderType { e := types.OrderType_Offline; return &e }(), true),
		orderbase1.WithCreateMethod(func() *types.OrderCreateMethod { e := types.OrderCreateMethod_OrderCreatedByAdmin; return &e }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateOrderBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteOrderBase(context.Background())
	}
}

func createCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithOrderID(&ret.OrderID, true),
		WithCompensateFromID(&ret.CompensateFromID, true),
		WithCompensateType(&ret.CompensateType, true),
		WithCompensateSeconds(&ret.CompensateSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.CreateCompensate(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetCompensate(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func updateCompensate(t *testing.T) {
	ret.CompensateSeconds = 180
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, false),
		WithCompensateSeconds(&ret.CompensateSeconds, true),
	)
	if assert.Nil(t, err) {
		err = handler.UpdateCompensate(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetCompensate(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func getCompensate(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
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
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		err = handler.DeleteCompensate(context.Background())
		assert.Nil(t, err)

		info, err := handler.GetCompensate(context.Background())
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
