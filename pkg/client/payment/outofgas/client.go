package outofgas

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/outofgas"
	servicename "github.com/NpoolPlatform/order-middleware/pkg/servicename"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateOutOfGas(ctx context.Context, in *npool.OutOfGasReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateOutOfGas(ctx, &npool.CreateOutOfGasRequest{
			Info: in,
		})
	})
	return err
}

func UpdateOutOfGas(ctx context.Context, in *npool.OutOfGasReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateOutOfGas(ctx, &npool.UpdateOutOfGasRequest{
			Info: in,
		})
	})
	return err
}

func DeleteOutOfGas(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteOutOfGas(ctx, &npool.DeleteOutOfGasRequest{
			Info: &npool.OutOfGasReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
