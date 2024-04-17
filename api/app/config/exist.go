package config

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) ExistSimulateConfig(ctx context.Context, in *npool.ExistSimulateConfigRequest) (*npool.ExistSimulateConfigResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSimulateConfig",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSimulateConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSimulateConfig",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSimulateConfigResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSimulateConfigConds(ctx context.Context, in *npool.ExistSimulateConfigCondsRequest) (*npool.ExistSimulateConfigCondsResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSimulateConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSimulateConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistSimulateConfigConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistSimulateConfigConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistSimulateConfigCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistSimulateConfigCondsResponse{
		Info: exist,
	}, nil
}
