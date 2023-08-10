//nolint:nolintlint,dupl
package compensate

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/compensate"
)

func (s *Server) ExistCompensate(ctx context.Context, in *npool.ExistCompensateRequest) (*npool.ExistCompensateResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCompensate(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCompensate",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCompensateResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCompensateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistCompensateConds(ctx context.Context, in *npool.ExistCompensateCondsRequest) (*npool.ExistCompensateCondsResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCompensateConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCompensateCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistCompensateConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistCompensateConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistCompensateCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistCompensateCondsResponse{
		Info: exist,
	}, nil
}
