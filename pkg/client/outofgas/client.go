package outofgas

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"

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

func CreateOutOfGas(ctx context.Context, in *npool.OutOfGasReq) (*npool.OutOfGas, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateOutOfGas(ctx, &npool.CreateOutOfGasRequest{
			Info: in,
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

func UpdateOutOfGas(ctx context.Context, in *npool.OutOfGasReq) (*npool.OutOfGas, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateOutOfGas(ctx, &npool.UpdateOutOfGasRequest{
			Info: in,
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

func GetOutOfGas(ctx context.Context, id string) (*npool.OutOfGas, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOutOfGas(ctx, &npool.GetOutOfGasRequest{
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
	return info.(*npool.OutOfGas), nil
}

func GetOutOfGass(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.OutOfGas, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOutOfGass(ctx, &npool.GetOutOfGassRequest{
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
	return infos.([]*npool.OutOfGas), total, nil
}

func GetOutOfGasOnly(ctx context.Context, conds *npool.Conds) (*npool.OutOfGas, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetOutOfGass(ctx, &npool.GetOutOfGassRequest{
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
	if len(infos.([]*npool.OutOfGas)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.OutOfGas)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.OutOfGas)[0], nil
}

func DeleteOutOfGas(ctx context.Context, in *npool.OutOfGasReq) (*npool.OutOfGas, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteOutOfGas(ctx, &npool.DeleteOutOfGasRequest{
			Info: in,
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

func ExistOutOfGas(ctx context.Context, id string) (bool, error) {
	exist, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistOutOfGas(ctx, &npool.ExistOutOfGasRequest{
			ID: id,
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
	exist, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
