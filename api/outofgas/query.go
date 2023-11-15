//nolint:nolintlint,dupl,gocyclo
package outofgas

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"

	outofgas1 "github.com/NpoolPlatform/order-middleware/pkg/mw/outofgas"
)

func (s *Server) GetOutOfGas(ctx context.Context, in *npool.GetOutOfGasRequest) (*npool.GetOutOfGasResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithEntID(&in.EntID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOutOfGas",
			"In", in,
			"error", err,
		)
		return &npool.GetOutOfGasResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetOutOfGas(ctx)
	if err != nil {
		return &npool.GetOutOfGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOutOfGasResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOutOfGases(ctx context.Context, in *npool.GetOutOfGasesRequest) (*npool.GetOutOfGasesResponse, error) {
	handler, err := outofgas1.NewHandler(
		ctx,
		outofgas1.WithConds(in.GetConds()),
		outofgas1.WithOffset(in.GetOffset()),
		outofgas1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOutOfGases",
			"In", in,
			"error", err,
		)
		return &npool.GetOutOfGasesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetOutOfGases(ctx)
	if err != nil {
		return &npool.GetOutOfGasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOutOfGasesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
