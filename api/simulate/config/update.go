package config

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"

	config1 "github.com/NpoolPlatform/order-middleware/pkg/mw/simulate/config"
)

func (s *Server) UpdateSimulateConfig(ctx context.Context, in *npool.UpdateSimulateConfigRequest) (*npool.UpdateSimulateConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateSimulateConfig",
			"In", in,
		)
		return &npool.UpdateSimulateConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithID(req.ID, true),
		config1.WithUnits(req.Units, false),
		config1.WithDuration(req.Duration, false),
		config1.WithSendCouponMode(req.SendCouponMode, false),
		config1.WithSendCouponProbability(req.SendCouponProbability, false),
		config1.WithEnabled(req.Enabled, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulateConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateSimulateConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateSimulateConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateSimulateConfigResponse{
		Info: info,
	}, nil
}
