package compensate

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/compensate"
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

func CreateCompensate(ctx context.Context, in *npool.CompensateReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateCompensate(ctx, &npool.CreateCompensateRequest{
			Info: in,
		})
	})
	return err
}

func DeleteCompensate(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteCompensate(ctx, &npool.DeleteCompensateRequest{
			Info: &npool.CompensateReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
