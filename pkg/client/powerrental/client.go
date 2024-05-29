package powerrental

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	feeordermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
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

func CreatePowerRentalOrder(ctx context.Context, in *npool.PowerRentalOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreatePowerRentalOrder(ctx, &npool.CreatePowerRentalOrderRequest{
			Info: in,
		})
	})
	return wlog.WrapError(err)
}

func CreatePowerRentalOrderWithFees(ctx context.Context, powerRentalOrder *npool.PowerRentalOrderReq, feeOrders []*feeordermwpb.FeeOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreatePowerRentalOrderWithFees(ctx, &npool.CreatePowerRentalOrderWithFeesRequest{
			PowerRentalOrder: powerRentalOrder,
			FeeOrders:        feeOrders,
		})
	})
	return wlog.WrapError(err)
}

func UpdatePowerRentalOrder(ctx context.Context, in *npool.PowerRentalOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdatePowerRentalOrder(ctx, &npool.UpdatePowerRentalOrderRequest{
			Info: in,
		})
	})
	return wlog.WrapError(err)
}

func UpdatePowerRentalOrders(ctx context.Context, in []*npool.PowerRentalOrderReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdatePowerRentalOrders(ctx, &npool.UpdatePowerRentalOrdersRequest{
			Infos: in,
		})
	})
	return wlog.WrapError(err)
}

func GetPowerRentalOrder(ctx context.Context, orderID string) (*npool.PowerRentalOrder, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPowerRentalOrder(ctx, &npool.GetPowerRentalOrderRequest{
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
	return info.(*npool.PowerRentalOrder), nil
}

func GetPowerRentalOrders(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.PowerRentalOrder, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPowerRentalOrders(ctx, &npool.GetPowerRentalOrdersRequest{
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
	return _infos.([]*npool.PowerRentalOrder), total, nil
}

func CountPowerRentalOrders(ctx context.Context, conds *npool.Conds) (count uint32, err error) {
	_info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.CountPowerRentalOrders(ctx, &npool.CountPowerRentalOrdersRequest{
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

func GetPowerRentalOrderOnly(ctx context.Context, conds *npool.Conds) (*npool.PowerRentalOrder, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPowerRentalOrders(ctx, &npool.GetPowerRentalOrdersRequest{
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
	if len(infos.([]*npool.PowerRentalOrder)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.PowerRentalOrder)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.PowerRentalOrder)[0], nil
}

func ExistPowerRentalOrder(ctx context.Context, orderID string) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistPowerRentalOrder(ctx, &npool.ExistPowerRentalOrderRequest{
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

func ExistPowerRentalOrderConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistPowerRentalOrderConds(ctx, &npool.ExistPowerRentalOrderCondsRequest{
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

func DeletePowerRentalOrder(ctx context.Context, id *uint32, entID, orderID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeletePowerRentalOrder(ctx, &npool.DeletePowerRentalOrderRequest{
			Info: &npool.PowerRentalOrderReq{
				ID:      id,
				EntID:   entID,
				OrderID: orderID,
			},
		})
	})
	return wlog.WrapError(err)
}
