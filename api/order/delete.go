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
