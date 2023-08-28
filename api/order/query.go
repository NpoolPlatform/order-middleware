//nolint:nolintlint,dupl,gocyclo
package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
)

func (s *Server) GetOrder(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrder",
			"In", in,
			"error", err,
		)
		return &npool.GetOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetOrder(ctx)
	if err != nil {
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOrders(ctx context.Context, in *npool.GetOrdersRequest) (*npool.GetOrdersResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithConds(in.GetConds()),
		order1.WithOffset(in.GetOffset()),
		order1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrders",
			"In", in,
			"error", err,
		)
		return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetOrders(ctx)
	if err != nil {
		return &npool.GetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}

func (s *Server) SumOrderUnits(ctx context.Context, in *npool.SumOrderUnitsRequest) (*npool.SumOrderUnitsResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SumOrderUnits",
			"In", in,
			"error", err,
		)
		return &npool.SumOrderUnitsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	count, err := handler.SumOrderUnits(ctx)
	if err != nil {
		return &npool.SumOrderUnitsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.SumOrderUnitsResponse{
		Info: count,
	}, nil
}
