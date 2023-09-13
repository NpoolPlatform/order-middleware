package orderlock

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"

	servicename "github.com/NpoolPlatform/order-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateOrderLocks(ctx context.Context, in []*npool.OrderLockReq) ([]*npool.OrderLock, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateOrderLocks(ctx, &npool.CreateOrderLocksRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.OrderLock), nil
}

func GetOrderLock(ctx context.Context, id string) (*npool.OrderLock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderLock(ctx, &npool.GetOrderLockRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.OrderLock), nil
}

func GetOrderLocks(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OrderLock, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderLocks(ctx, &npool.GetOrderLocksRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}

		total = resp.Total

		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*npool.OrderLock), total, nil
}

func GetOrderLockOnly(ctx context.Context, conds *npool.Conds) (*npool.OrderLock, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOrderLocks(ctx, &npool.GetOrderLocksRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.OrderLock)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.OrderLock)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.OrderLock)[0], nil
}

func DeleteOrderLocks(ctx context.Context, in []*npool.OrderLockReq) ([]*npool.OrderLock, error) {
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteOrderLocks(ctx, &npool.DeleteOrderLocksRequest{
			Infos: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*npool.OrderLock), nil
}
