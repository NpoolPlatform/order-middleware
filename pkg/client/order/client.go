package order

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
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

func GetOrders(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Order, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOrders(ctx, &npool.GetOrdersRequest{
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
	return _infos.([]*npool.Order), total, nil
}

func ExistOrder(ctx context.Context, entID string) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistOrder(ctx, &npool.ExistOrderRequest{
			EntID: entID,
		})
		if err != nil {
			return false, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return exist.(bool), err
}

func ExistOrderConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistOrderConds(ctx, &npool.ExistOrderCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return false, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return exist.(bool), err
}

func CountOrders(ctx context.Context, conds *npool.Conds) (uint32, error) {
	count, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CountOrders(ctx, &npool.CountOrdersRequest{
			Conds: conds,
		})
		if err != nil {
			return false, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, err
	}
	return count.(uint32), err
}

func GetOrder(ctx context.Context, entID string) (info *npool.Order, err error) {
	_info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOrder(ctx, &npool.GetOrderRequest{
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
	return _info.(*npool.Order), nil
}

func GetOrdersOnly(ctx context.Context, conds *npool.Conds) (info *npool.Order, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOrders(ctx, &npool.GetOrdersRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(_infos.([]*npool.Order)) == 0 {
		return nil, wlog.Errorf("invalid order")
	}
	if len(_infos.([]*npool.Order)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return _infos.([]*npool.Order)[0], nil
}
