//nolint:nolintlint,dupl
package compensate

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"

	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/compensate"
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
		compensate1.WithOrderID(req.OrderID, true),
		compensate1.WithStartAt(req.StartAt, true),
		compensate1.WithEndAt(req.EndAt, true),
		compensate1.WithTitle(req.Title, true),
		compensate1.WithMessage(req.Message, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCompensate",
			"Req", req,
			"error", err,
		)
		return &npool.CreateCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateCompensate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateCompensate",
			"Req", req,
			"error", err,
		)
		return &npool.CreateCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateCompensateResponse{
		Info: info,
	}, nil
}
