//nolint:dupl
package outofgas

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/outofgas"

	outofgas1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental/outofgas"
)

func (s *Server) UpdateOutOfGas(ctx context.Context, in *npool.UpdateOutOfGasRequest) (*npool.UpdateOutOfGasResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateOutOfGas",
			"In", in,
		)
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithID(req.ID, false),
		outofgas1.WithEntID(req.EntID, false),
		outofgas1.WithEndAt(req.EndAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOutOfGas",
			"Req", req,
			"error", err,
		)
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.UpdateOutOfGas(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateOutOfGas",
			"Req", req,
			"error", err,
		)
		return &npool.UpdateOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateOutOfGasResponse{}, nil
}
