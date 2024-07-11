package appconfig

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	appconfig1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) GetAppConfig(ctx context.Context, in *npool.GetAppConfigRequest) (*npool.GetAppConfigResponse, error) {
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithAppID(&in.AppID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfig",
			"In", in,
			"error", err,
		)
		return &npool.GetAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetAppConfig(ctx)
	if err != nil {
		return &npool.GetAppConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetAppConfigs(ctx context.Context, in *npool.GetAppConfigsRequest) (*npool.GetAppConfigsResponse, error) {
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithConds(in.GetConds()),
		appconfig1.WithOffset(in.GetOffset()),
		appconfig1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetAppConfigs",
			"In", in,
			"error", err,
		)
		return &npool.GetAppConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetAppConfigs(ctx)
	if err != nil {
		return &npool.GetAppConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAppConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
