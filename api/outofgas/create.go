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

func (s *Server) CreateOutOfGas(ctx context.Context, in *npool.CreateOutOfGasRequest) (*npool.CreateOutOfGasResponse, error) {
	req := in.GetInfo()
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithID(req.ID, true),
		outofgas1.WithOrderID(req.OrderID, true),
		outofgas1.WithStartAt(req.StartAt, true),
		outofgas1.WithEndAt(req.EndAt, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOutOfGas",
			"Req", req,
			"error", err,
		)
		return &npool.CreateOutOfGasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateOutOfGas(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOutOfGas",
			"Req", req,
			"error", err,
		)
		return &npool.CreateOutOfGasResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateOutOfGasResponse{
		Info: info,
	}, nil
}
