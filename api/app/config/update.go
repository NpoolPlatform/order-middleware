package appconfig

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	appconfig1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) UpdateAppConfig(ctx context.Context, in *npool.UpdateAppConfigRequest) (*npool.UpdateAppConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateAppConfig",
			"In", in,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithID(req.ID, false),
		appconfig1.WithEntID(req.EntID, false),
		appconfig1.WithAppID(req.AppID, false),
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
			"UpdateAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	if err := handler.UpdateAppConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateAppConfigResponse{}, nil
}
