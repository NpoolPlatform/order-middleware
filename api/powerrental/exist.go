package powerrental

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) ExistPowerRentalOrder(ctx context.Context, in *npool.ExistPowerRentalOrderRequest) (*npool.ExistPowerRentalOrderResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithOrderID(&in.OrderID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPowerRental(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPowerRentalOrderResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistPowerRentalOrderConds(ctx context.Context, in *npool.ExistPowerRentalOrderCondsRequest) (*npool.ExistPowerRentalOrderCondsResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalOrderConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalOrderCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPowerRentalConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPowerRentalOrderConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPowerRentalOrderCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPowerRentalOrderCondsResponse{
		Info: exist,
	}, nil
}
