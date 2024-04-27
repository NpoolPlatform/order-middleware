package feeorder

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) CountFeeOrders(ctx context.Context, in *npool.CountFeeOrdersRequest) (*npool.CountFeeOrdersResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountFeeOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	count, err := handler.CountFeeOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountFeeOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountFeeOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountFeeOrdersResponse{
		Info: count,
	}, nil
}
