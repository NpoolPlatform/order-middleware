package config

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	config1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) DeleteSimulateConfig(ctx context.Context, in *npool.DeleteSimulateConfigRequest) (*npool.DeleteSimulateConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteSimulateConfig",
			"In", in,
		)
		return &npool.DeleteSimulateConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSimulateConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteSimulateConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteSimulateConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteSimulateConfigResponse{
		Info: info,
	}, nil
}
