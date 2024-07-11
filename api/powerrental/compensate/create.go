package compensate

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/compensate"

	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental/compensate"
)

func (s *Server) CreateCompensate(ctx context.Context, in *npool.CreateCompensateRequest) (*npool.CreateCompensateResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateCompensate",
			"In", in,
		)
		return &npool.CreateCompensateResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithEntID(req.EntID, false),
		compensate1.WithGoodID(req.GoodID, false),
		compensate1.WithOrderID(req.OrderID, false),
		compensate1.WithCompensateFromID(req.CompensateFromID, true),
		compensate1.WithCompensateType(req.CompensateType, true),
		compensate1.WithCompensateSeconds(req.CompensateSeconds, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCompensate",
			"Req", req,
			"error", err,
		)
		return &npool.CreateCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.CreateCompensate(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateCompensate",
			"Req", req,
			"error", err,
		)
		return &npool.CreateCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCompensateResponse{}, nil
}
