package config

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"

	config1 "github.com/NpoolPlatform/order-middleware/pkg/mw/simulate/config"
)

func (s *Server) CreateSimulateConfig(ctx context.Context, in *npool.CreateSimulateConfigRequest) (*npool.CreateSimulateConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateSimulateConfig",
			"In", in,
		)
		return &npool.CreateSimulateConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(req.EntID, false),
		config1.WithAppID(req.AppID, true),
		config1.WithEnabledProfitTx(req.EnabledProfitTx, true),
		config1.WithProfitTxProbability(req.ProfitTxProbability, true),
		config1.WithSendCouponMode(req.SendCouponMode, true),
		config1.WithSendCouponProbability(req.SendCouponProbability, true),
		config1.WithEnabled(req.Enabled, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulateConfig",
			"Req", req,
			"error", err,
		)
		return &npool.CreateSimulateConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateSimulateConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateSimulateConfig",
			"Req", req,
			"error", err,
		)
		return &npool.CreateSimulateConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateSimulateConfigResponse{
		Info: info,
	}, nil
}
