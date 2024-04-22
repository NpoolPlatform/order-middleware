package config

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	config1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) CreateAppConfig(ctx context.Context, in *npool.CreateAppConfigRequest) (*npool.CreateAppConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"In", in,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := config1.NewHandler(
		ctx,
		config1.WithEntID(req.EntID, false),
		config1.WithAppID(req.AppID, true),
		config1.WithEnableSimulateOrder(req.EnableSimulateOrder, false),
		config1.WithSimulateOrderCouponMode(req.SimulateOrderCouponMode, false),
		config1.WithSimulateOrderCouponProbability(req.SimulateOrderCouponProbability, false),
		config1.WithSimulateOrderCashableProfitProbability(req.SimulateOrderCashableProfitProbability, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"Req", req,
			"error", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateAppConfig(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"Req", req,
			"error", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAppConfigResponse{
		Info: info,
	}, nil
}
