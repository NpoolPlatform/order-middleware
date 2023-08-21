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

func (s *Server) UpdateCompensate(ctx context.Context, in *npool.UpdateCompensateRequest) (*npool.UpdateCompensateResponse, error) {
	req := in.GetInfo()
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(req.ID, true),
		compensate1.WithStart(req.Start, false),
		compensate1.WithEnd(req.End, false),
		compensate1.WithMessage(req.Message, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateCompensate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateCompensateResponse{
		Info: info,
	}, nil
}
