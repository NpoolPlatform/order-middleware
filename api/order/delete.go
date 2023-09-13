package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
)

func (s *Server) DeleteOrder(ctx context.Context, in *npool.DeleteOrderRequest) (*npool.DeleteOrderResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteOrder",
			"In", in,
		)
		return &npool.DeleteOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrder",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeleteOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrder",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) DeleteOrders(ctx context.Context, in *npool.DeleteOrdersRequest) (*npool.DeleteOrdersResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"DeleteOrders",
			"In", in,
		)
		return &npool.DeleteOrdersResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := order1.NewHandler(
		ctx,
		order1.WithReqs(reqs, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrders",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.DeleteOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrders",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteOrdersResponse{
		Infos: infos,
	}, nil
}
