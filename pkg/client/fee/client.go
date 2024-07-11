package feeorder

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
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

func CreateFeeOrder(ctx context.Context, in *npool.FeeOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateFeeOrder(ctx, &npool.CreateFeeOrderRequest{
			Info: in,
		})
	})
	return wlog.WrapError(err)
}

func CreateFeeOrders(ctx context.Context, in []*npool.FeeOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateFeeOrders(ctx, &npool.CreateFeeOrdersRequest{
			Infos: in,
		})
	})
	return wlog.WrapError(err)
}

func UpdateFeeOrder(ctx context.Context, in *npool.FeeOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateFeeOrder(ctx, &npool.UpdateFeeOrderRequest{
			Info: in,
		})
	})
	return wlog.WrapError(err)
}

func GetFeeOrder(ctx context.Context, orderID string) (*npool.FeeOrder, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetFeeOrder(ctx, &npool.GetFeeOrderRequest{
			OrderID: orderID,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info.(*npool.FeeOrder), nil
}

func GetFeeOrders(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.FeeOrder, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetFeeOrders(ctx, &npool.GetFeeOrdersRequest{
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
	return _infos.([]*npool.FeeOrder), total, nil
}

func CountFeeOrders(ctx context.Context, conds *npool.Conds) (count uint32, err error) {
	_info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CountFeeOrders(ctx, &npool.CountFeeOrdersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, err
	}
	return _info.(uint32), nil
}

func GetFeeOrderOnly(ctx context.Context, conds *npool.Conds) (*npool.FeeOrder, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetFeeOrders(ctx, &npool.GetFeeOrdersRequest{
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
	if len(infos.([]*npool.FeeOrder)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.FeeOrder)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.FeeOrder)[0], nil
}

func ExistFeeOrder(ctx context.Context, orderID string) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistFeeOrder(ctx, &npool.ExistFeeOrderRequest{
			OrderID: orderID,
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

func ExistFeeOrderConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistFeeOrderConds(ctx, &npool.ExistFeeOrderCondsRequest{
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

func DeleteFeeOrder(ctx context.Context, id *uint32, entID, orderID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteFeeOrder(ctx, &npool.DeleteFeeOrderRequest{
			Info: &npool.FeeOrderReq{
				ID:      id,
				EntID:   entID,
				OrderID: orderID,
			},
		})
	})
	return wlog.WrapError(err)
}

func DeleteFeeOrders(ctx context.Context, in []*npool.FeeOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteFeeOrders(ctx, &npool.DeleteFeeOrdersRequest{
			Infos: in,
		})
	})
	return wlog.WrapError(err)
}
