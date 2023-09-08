package orderlock

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
	ordertypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"

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

var (
	ret = npool.OrderLock{
		ID:          uuid.NewString(),
		AppID:       uuid.NewString(),
		UserID:      uuid.NewString(),
		OrderID:     uuid.NewString(),
		LockTypeStr: ordertypes.OrderLockType_LockCommission.String(),
		LockType:    ordertypes.OrderLockType_LockCommission,
	}

	reqs = []*npool.OrderLockReq{
		{
			ID:       &ret.ID,
			AppID:    &ret.AppID,
			UserID:   &ret.UserID,
			OrderID:  &ret.OrderID,
			LockType: &ret.LockType,
		},
	}
)

func creates(t *testing.T) {
	var err error
	infos, err := CreateOrderLocks(context.Background(), reqs)
	fmt.Println("err: z", err)
	if assert.Nil(t, err) {
		ret.CreatedAt = infos[0].CreatedAt
		ret.UpdatedAt = infos[0].UpdatedAt
		assert.Equal(t, infos[0], &ret)
	}
}

func getOrderLock(t *testing.T) {
	var err error
	info, err := GetOrderLock(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func getOrderLocks(t *testing.T) {
	infos, _, err := GetOrderLocks(context.Background(), &npool.Conds{
		ID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: ret.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, infos[0].String(), ret.String())
	}
}

func deleteOrderLocks(t *testing.T) {
	var err error
	infos, err := DeleteOrderLocks(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, 0, len(infos))
	}

	info, err := GetOrderLock(context.Background(), ret.ID)
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

	t.Run("creates", creates)
	t.Run("getOrderLock", getOrderLock)
	t.Run("getOrderLocks", getOrderLocks)
	t.Run("deleteOrderLocks", deleteOrderLocks)
}
