package appconfig

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"
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

func CreateAppConfig(ctx context.Context, in *npool.AppConfigReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateAppConfig(ctx, &npool.CreateAppConfigRequest{
			Info: in,
		})
	})
	return wlog.WrapError(err)
}

func UpdateAppConfig(ctx context.Context, in *npool.AppConfigReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateAppConfig(ctx, &npool.UpdateAppConfigRequest{
			Info: in,
		})
	})
	return wlog.WrapError(err)
}

func GetAppConfig(ctx context.Context, appID string) (*npool.AppConfig, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppConfig(ctx, &npool.GetAppConfigRequest{
			AppID: appID,
		})
		if err != nil {
			return nil, wlog.WrapError(err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return info.(*npool.AppConfig), nil
}

func GetAppConfigs(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.AppConfig, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppConfigs(ctx, &npool.GetAppConfigsRequest{
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
	return _infos.([]*npool.AppConfig), total, nil
}

func GetAppConfigOnly(ctx context.Context, conds *npool.Conds) (*npool.AppConfig, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetAppConfigs(ctx, &npool.GetAppConfigsRequest{
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
	if len(infos.([]*npool.AppConfig)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.AppConfig)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.AppConfig)[0], nil
}

func DeleteAppConfig(ctx context.Context, id *uint32, entID, appID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteAppConfig(ctx, &npool.DeleteAppConfigRequest{
			Info: &npool.AppConfigReq{
				ID:    id,
				EntID: entID,
				AppID: appID,
			},
		})
	})
	return wlog.WrapError(err)
}

func ExistAppConfig(ctx context.Context, appID string) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistAppConfig(ctx, &npool.ExistAppConfigRequest{
			AppID: appID,
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

func ExistAppConfigConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	exist, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistAppConfigConds(ctx, &npool.ExistAppConfigCondsRequest{
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
