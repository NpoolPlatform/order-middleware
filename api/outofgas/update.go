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

func (s *Server) UpdateOutOfGas(ctx context.Context, in *npool.UpdateOutOfGasRequest) (*npool.UpdateOutOfGasResponse, error) {
	req := in.GetInfo()
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithID(req.ID, true),
		outofgas1.WithStartAt(req.StartAt, false),
		outofgas1.WithEndAt(req.EndAt, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateOutOfGas(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOutOfGas",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateOutOfGasResponse{
		Info: info,
	}, nil
}
