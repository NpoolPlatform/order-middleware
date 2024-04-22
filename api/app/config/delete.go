package appconfig

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	appconfig1 "github.com/NpoolPlatform/order-middleware/pkg/mw/app/config"
)

func (s *Server) DeleteAppConfig(ctx context.Context, in *npool.DeleteAppConfigRequest) (*npool.DeleteAppConfigResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteAppConfig",
			"In", in,
		)
		return &npool.DeleteAppConfigResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appconfig1.NewHandler(
		ctx,
		appconfig1.WithID(req.ID, false),
		appconfig1.WithEntID(req.EntID, false),
		appconfig1.WithAppID(req.AppID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	if err := handler.DeleteAppConfig(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteAppConfig",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteAppConfigResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteAppConfigResponse{}, nil
}
