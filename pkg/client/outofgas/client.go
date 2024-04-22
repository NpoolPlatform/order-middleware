package outofgas

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
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

func GetOutOfGas(ctx context.Context, id string) (*npool.OutOfGas, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOutOfGas(ctx, &npool.GetOutOfGasRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.OutOfGas), nil
}

func GetOutOfGases(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.OutOfGas, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOutOfGases(ctx, &npool.GetOutOfGasesRequest{
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
	return _infos.([]*npool.OutOfGas), total, nil
}

func GetOutOfGasOnly(ctx context.Context, conds *npool.Conds) (*npool.OutOfGas, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetOutOfGases(ctx, &npool.GetOutOfGasesRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.OutOfGas)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.OutOfGas)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.OutOfGas)[0], nil
}

func ExistOutOfGas(ctx context.Context, entID string) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistOutOfGas(ctx, &npool.ExistOutOfGasRequest{
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

func ExistOutOfGasConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistOutOfGasConds(ctx, &npool.ExistOutOfGasCondsRequest{
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
