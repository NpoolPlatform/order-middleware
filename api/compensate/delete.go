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

func (s *Server) DeleteCompensate(ctx context.Context, in *npool.DeleteCompensateRequest) (*npool.DeleteCompensateResponse, error) {
	req := in.GetInfo()
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteCompensate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteCompensateResponse{
		Info: info,
	}, nil
}
