//nolint:nolintlint,dupl
package outofgas

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"

	outofgas1 "github.com/NpoolPlatform/order-middleware/pkg/mw/outofgas"
)

func (s *Server) DeleteOutOfGas(ctx context.Context, in *npool.DeleteOutOfGasRequest) (*npool.DeleteOutOfGasResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteOutOfGas",
			"In", in,
		)
		return &npool.DeleteOutOfGasResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteOutOfGas(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteOutOfGasResponse{
		Info: info,
	}, nil
}
