//nolint:nolintlint,dupl,gocyclo
package compensate

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"

	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/compensate"
)

func (s *Server) GetCompensate(ctx context.Context, in *npool.GetCompensateRequest) (*npool.GetCompensateResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCompensate",
			"In", in,
			"error", err,
		)
		return &npool.GetCompensateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetCompensate(ctx)
	if err != nil {
		return &npool.GetCompensateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCompensateResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCompensates(ctx context.Context, in *npool.GetCompensatesRequest) (*npool.GetCompensatesResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithConds(in.GetConds()),
		compensate1.WithOffset(in.GetOffset()),
		compensate1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetCompensates",
			"In", in,
			"error", err,
		)
		return &npool.GetCompensatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetCompensates(ctx)
	if err != nil {
		return &npool.GetCompensatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCompensatesResponse{
		Infos: infos,
		Total: total,
	}, nil
}
