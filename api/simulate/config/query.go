package config

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"

	config1 "github.com/NpoolPlatform/order-middleware/pkg/mw/simulate/config"
)

func (s *Server) GetSimulateConfig(ctx context.Context, in *npool.GetSimulateConfigRequest) (*npool.GetSimulateConfigResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulateConfig",
			"In", in,
			"error", err,
		)
		return &npool.GetSimulateConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetSimulateConfig(ctx)
	if err != nil {
		return &npool.GetSimulateConfigResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSimulateConfigResponse{
		Info: info,
	}, nil
}

func (s *Server) GetSimulateConfigs(ctx context.Context, in *npool.GetSimulateConfigsRequest) (*npool.GetSimulateConfigsResponse, error) {
	handler, err := config1.NewHandler(
		ctx,
		config1.WithConds(in.GetConds()),
		config1.WithOffset(in.GetOffset()),
		config1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetSimulateConfigs",
			"In", in,
			"error", err,
		)
		return &npool.GetSimulateConfigsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetSimulateConfigs(ctx)
	if err != nil {
		return &npool.GetSimulateConfigsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSimulateConfigsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
