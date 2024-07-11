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

func (s *Server) CountCompensates(ctx context.Context, in *npool.CountCompensatesRequest) (*npool.CountCompensatesResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountCompensates",
			"In", in,
			"Error", err,
		)
		return &npool.CountCompensatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.CountCompensates(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountCompensates",
			"In", in,
			"Error", err,
		)
		return &npool.CountCompensatesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountCompensatesResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountCompensateOrders(ctx context.Context, in *npool.CountCompensateOrdersRequest) (*npool.CountCompensateOrdersResponse, error) {
	handler, err := compensate1.NewHandler(
		ctx,
		compensate1.WithCompensateFromIDs(in.GetCompensateFromIDs(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountCompensateOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountCompensateOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	infos, err := handler.CountCompensateOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountCompensateOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountCompensateOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountCompensateOrdersResponse{
		Infos: infos,
	}, nil
}
