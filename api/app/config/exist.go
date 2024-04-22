package appconfig

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"
	appconfig1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) ExistAppConfig(ctx context.Context, in *npool.ExistAppConfigRequest) (*npool.ExistAppConfigResponse, error) {
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAppConfigResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAppConfigConds(ctx context.Context, in *npool.ExistAppConfigCondsRequest) (*npool.ExistAppConfigCondsResponse, error) {
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAppConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistAppConfigConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistAppConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistAppConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistAppConfigCondsResponse{
		Info: exist,
	}, nil
}
