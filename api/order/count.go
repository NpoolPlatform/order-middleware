package order

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
)

func (s *Server) CountOrders(ctx context.Context, in *npool.CountOrdersRequest) (*npool.CountOrdersResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.CountOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountOrdersResponse{
		Info: exist,
	}, nil
}
