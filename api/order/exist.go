package order

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
)

func (s *Server) ExistOrder(ctx context.Context, in *npool.ExistOrderRequest) (*npool.ExistOrderResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrder",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrder",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOrderResponse{
		Info: exist,
	}, nil
}

//nolint:dupl
func (s *Server) ExistOrderConds(ctx context.Context, in *npool.ExistOrderCondsRequest) (*npool.ExistOrderCondsResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistOrderConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOrderCondsResponse{
		Info: exist,
	}, nil
}
