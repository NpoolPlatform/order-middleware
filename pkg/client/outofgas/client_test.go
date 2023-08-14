package outofgas

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

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/order-middleware/pkg/testinit"

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
	id  = uuid.NewString()
	ret = npool.OutOfGas{
		ID:      id,
		OrderID: uuid.NewString(),
		Start:   10002,
		End:     10003,
	}
)

func createOutOfGas(t *testing.T) {
	var (
		req = npool.OutOfGasReq{
			ID:      &ret.ID,
			OrderID: &ret.OrderID,
			Start:   &ret.Start,
			End:     &ret.End,
		}
	)

	info, err := CreateOutOfGas(context.Background(), &req)
	if assert.Nil(t, err) {
		if id != info.ID {
			ret.ID = info.ID
		}
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateOutOfGas(t *testing.T) {
	if ret.ID == id {
		var (
			start = uint32(10006)
			end   = uint32(10007)

			req = npool.OutOfGasReq{
				ID:    &ret.ID,
				Start: &start,
				End:   &end,
			}
		)
		info, err := UpdateOutOfGas(context.Background(), &req)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getOutOfGas(t *testing.T) {
	info, err := GetOutOfGas(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getOutOfGass(t *testing.T) {
	infos, _, err := GetOutOfGass(context.Background(), &npool.Conds{
		OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.NotEqual(t, len(infos), 0)
	}
}

func deleteOutOfGas(t *testing.T) {
	if ret.ID == id {
		info, err := DeleteOutOfGas(context.Background(), &npool.OutOfGasReq{
			ID: &ret.ID,
		})
		if assert.Nil(t, err) {
			assert.Equal(t, info, &ret)
		}

		info, err = GetOutOfGas(context.Background(), ret.ID)
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestOutOfGas(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createOutOfGas", createOutOfGas)
	t.Run("updateOutOfGas", updateOutOfGas)
	t.Run("getOutOfGas", getOutOfGas)
	t.Run("getOutOfGass", getOutOfGass)
	t.Run("deleteOutOfGas", deleteOutOfGas)
}
