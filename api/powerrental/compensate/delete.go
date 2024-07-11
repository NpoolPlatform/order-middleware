package compensate

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/compensate"

	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental/compensate"
)

func (s *Server) DeleteCompensate(ctx context.Context, in *npool.DeleteCompensateRequest) (*npool.DeleteCompensateResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteCompensate",
			"In", in,
		)
		return &npool.DeleteCompensateResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(req.ID, false),
		compensate1.WithEntID(req.EntID, false),
		compensate1.WithOrderID(req.OrderID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	if err := handler.DeleteCompensate(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteCompensateResponse{}, nil
}
