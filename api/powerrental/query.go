package powerrental

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"

	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) GetPowerRentalOrder(ctx context.Context, in *npool.GetPowerRentalOrderRequest) (*npool.GetPowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentalOrder",
			"In", in,
			"error", err,
		)
		return &npool.GetPowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetPowerRental(ctx)
	if err != nil {
		return &npool.GetPowerRentalOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPowerRentalOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetPowerRentalOrders(ctx context.Context, in *npool.GetPowerRentalOrdersRequest) (*npool.GetPowerRentalOrdersResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithConds(in.GetConds()),
		powerrental1.WithOffset(in.GetOffset()),
		powerrental1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPowerRentalOrders",
			"In", in,
			"error", err,
		)
		return &npool.GetPowerRentalOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetPowerRentals(ctx)
	if err != nil {
		return &npool.GetPowerRentalOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPowerRentalOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
