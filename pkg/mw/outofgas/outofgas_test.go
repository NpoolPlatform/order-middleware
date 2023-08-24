package outofgas

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"

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

var ret = npool.OutOfGas{
	ID:      uuid.NewString(),
	OrderID: uuid.NewString(),
	StartAt: 10002,
	EndAt:   10003,
}

func createOutOfGas(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOrderID(&ret.OrderID, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateOutOfGas(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateOutOfGas(t *testing.T) {
	ret.StartAt = uint32(10007)
	ret.EndAt = uint32(10008)
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithOrderID(&ret.OrderID, false),
		WithStartAt(&ret.StartAt, false),
		WithEndAt(&ret.EndAt, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateOutOfGas(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getOutOfGas(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetOutOfGas(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getOutOfGass(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetOutOfGass(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteOutOfGas(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteOutOfGas(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetOutOfGas(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestOutOfGas(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createOutOfGas", createOutOfGas)
	t.Run("updateOutOfGas", updateOutOfGas)
	t.Run("getOutOfGas", getOutOfGas)
	t.Run("getOutOfGass", getOutOfGass)
	t.Run("deleteOutOfGas", deleteOutOfGas)
}
