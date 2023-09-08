package orderlock

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
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"

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
	rets = []npool.OrderLock{
		{
			ID:          uuid.NewString(),
			OrderID:     uuid.NewString(),
			AppID:       uuid.NewString(),
			UserID:      uuid.NewString(),
			LockType:    ordertypes.OrderLockType_LockCommission,
			LockTypeStr: ordertypes.OrderLockType_LockCommission.String(),
		},
	}
)

func createOrderLocks(t *testing.T) {
	retsReq := []*npool.OrderLockReq{}
	for key := range rets {
		retReq := npool.OrderLockReq{
			ID:       &rets[key].ID,
			AppID:    &rets[key].AppID,
			UserID:   &rets[key].UserID,
			OrderID:  &rets[key].OrderID,
			LockType: &rets[key].LockType,
		}
		retsReq = append(retsReq, &retReq)
	}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(retsReq),
	)
	if assert.Nil(t, err) {
		infos, err := handler.CreateOrderLocks(context.Background())
		if assert.Nil(t, err) {
			rets[0].CreatedAt = infos[0].CreatedAt
			rets[0].UpdatedAt = infos[0].UpdatedAt
			assert.Equal(t, len(rets), len(infos))
		}
	}
}

func getOrderLock(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&rets[0].ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetOrderLock(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &rets[0], info)
		}
	}
}

func getOrderLocks(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: rets[0].OrderID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetOrderLocks(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &rets[0], infos[0])
			}
		}
	}
}

func deleteOrderLocks(t *testing.T) {
	retsReq := []*npool.OrderLockReq{}
	for key := range rets {
		retReq := npool.OrderLockReq{
			ID:       &rets[key].ID,
			AppID:    &rets[key].AppID,
			UserID:   &rets[key].UserID,
			OrderID:  &rets[key].OrderID,
			LockType: &rets[key].LockType,
		}
		retsReq = append(retsReq, &retReq)
	}
	handler, err := NewHandler(
		context.Background(),
		WithReqs(retsReq),
	)
	if assert.Nil(t, err) {
		infos, err := handler.DeleteOrderLocks(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, len(infos), 0)
		}
	}
}

func TestOrderLock(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	time.Sleep(10 * time.Second)
	t.Run("createOrderLocks", createOrderLocks)
	t.Run("getOrderLock", getOrderLock)
	t.Run("getOrderLocks", getOrderLocks)
	t.Run("deleteOrderLocks", deleteOrderLocks)
}
