//nolint:nolintlint,dupl
package outofgas

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	outofgas1 "github.com/NpoolPlatform/order-middleware/pkg/mw/outofgas"
)

func (s *Server) ExistOutOfGas(ctx context.Context, in *npool.ExistOutOfGasRequest) (*npool.ExistOutOfGasResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistOutOfGas(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOutOfGasResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistOutOfGasConds(ctx context.Context, in *npool.ExistOutOfGasCondsRequest) (*npool.ExistOutOfGasCondsResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOutOfGasConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOutOfGasCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistOutOfGasConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOutOfGasConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOutOfGasCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOutOfGasCondsResponse{
		Info: exist,
	}, nil
}
