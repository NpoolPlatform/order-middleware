package feeorder

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"

	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) GetFeeOrder(ctx context.Context, in *npool.GetFeeOrderRequest) (*npool.GetFeeOrderResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeOrder",
			"In", in,
			"error", err,
		)
		return &npool.GetFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetFeeOrder(ctx)
	if err != nil {
		return &npool.GetFeeOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeeOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetFeeOrders(ctx context.Context, in *npool.GetFeeOrdersRequest) (*npool.GetFeeOrdersResponse, error) {
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithConds(in.GetConds()),
		feeorder1.WithOffset(in.GetOffset()),
		feeorder1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetFeeOrders",
			"In", in,
			"error", err,
		)
		return &npool.GetFeeOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetFeeOrders(ctx)
	if err != nil {
		return &npool.GetFeeOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFeeOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
