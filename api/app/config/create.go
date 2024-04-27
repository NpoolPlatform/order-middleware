package appconfig

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	appconfig1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
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
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithEntID(req.EntID, false),
		appconfig1.WithAppID(req.AppID, true),
		appconfig1.WithEnableSimulateOrder(req.EnableSimulateOrder, false),
		appconfig1.WithSimulateOrderUnits(req.SimulateOrderUnits, false),
		appconfig1.WithSimulateOrderDurationSeconds(req.SimulateOrderDurationSeconds, false),
		appconfig1.WithSimulateOrderCouponMode(req.SimulateOrderCouponMode, false),
		appconfig1.WithSimulateOrderCouponProbability(req.SimulateOrderCouponProbability, false),
		appconfig1.WithSimulateOrderCashableProfitProbability(req.SimulateOrderCashableProfitProbability, false),
		appconfig1.WithMaxUnpaidOrders(req.MaxUnpaidOrders, false),
		appconfig1.WithMaxTypedCouponsPerOrder(req.MaxTypedCouponsPerOrder, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"Req", req,
			"error", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.CreateAppConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateAppConfig",
			"Req", req,
			"error", err,
		)
		return &npool.CreateAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateAppConfigResponse{}, nil
}
