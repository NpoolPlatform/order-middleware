package feeorder

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"
	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) ExistFeeOrder(ctx context.Context, in *npool.ExistFeeOrderRequest) (*npool.ExistFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistFeeOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistFeeOrderResponse{
		Info: exist,
	}, nil
}

//nolint:dupl
func (s *Server) ExistFeeOrderConds(ctx context.Context, in *npool.ExistFeeOrderCondsRequest) (*npool.ExistFeeOrderCondsResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFeeOrderConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeOrderCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistFeeOrderConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistFeeOrderConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistFeeOrderCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistFeeOrderCondsResponse{
		Info: exist,
	}, nil
}
