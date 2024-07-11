package orderlock

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/lock"
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

func CreateOrderLocks(ctx context.Context, in []*npool.OrderLockReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateOrderLocks(ctx, &npool.CreateOrderLocksRequest{
			Infos: in,
		})
	})
	return wlog.WrapError(err)
}

func GetOrderLock(ctx context.Context, entID string) (*npool.OrderLock, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOrderLock(ctx, &npool.GetOrderLockRequest{
			EntID: entID,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info.(*npool.OrderLock), nil
}

func GetOrderLocks(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.OrderLock, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOrderLocks(ctx, &npool.GetOrderLocksRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		total = resp.Total
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return _infos.([]*npool.OrderLock), total, nil
}
func DeleteOrderLocks(ctx context.Context, in []*npool.OrderLockReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteOrderLocks(ctx, &npool.DeleteOrderLocksRequest{
			Infos: in,
		})
	})
	return wlog.WrapError(err)
}
