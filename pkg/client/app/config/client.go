package config

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

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

func CreateSimulateConfig(ctx context.Context, in *npool.SimulateConfigReq) (*npool.SimulateConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateSimulateConfig(ctx, &npool.CreateSimulateConfigRequest{
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
	return info.(*npool.SimulateConfig), nil
}

func UpdateSimulateConfig(ctx context.Context, in *npool.SimulateConfigReq) (*npool.SimulateConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateSimulateConfig(ctx, &npool.UpdateSimulateConfigRequest{
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
	return info.(*npool.SimulateConfig), nil
}

func GetSimulateConfig(ctx context.Context, id string) (*npool.SimulateConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSimulateConfig(ctx, &npool.GetSimulateConfigRequest{
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
	return info.(*npool.SimulateConfig), nil
}

func GetSimulateConfigs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.SimulateConfig, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSimulateConfigs(ctx, &npool.GetSimulateConfigsRequest{
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
	return infos.([]*npool.SimulateConfig), total, nil
}

func GetSimulateConfigOnly(ctx context.Context, conds *npool.Conds) (*npool.SimulateConfig, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetSimulateConfigs(ctx, &npool.GetSimulateConfigsRequest{
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
	if len(infos.([]*npool.SimulateConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.SimulateConfig)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.SimulateConfig)[0], nil
}

func DeleteSimulateConfig(ctx context.Context, id uint32) (*npool.SimulateConfig, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteSimulateConfig(ctx, &npool.DeleteSimulateConfigRequest{
			Info: &npool.SimulateConfigReq{
				ID: &id,
			},
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.SimulateConfig), nil
}

func ExistSimulateConfig(ctx context.Context, id string) (bool, error) {
	exist, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistSimulateConfig(ctx, &npool.ExistSimulateConfigRequest{
			EntID: id,
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

func ExistSimulateConfigConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ExistSimulateConfigConds(ctx, &npool.ExistSimulateConfigCondsRequest{
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
