package powerrental

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) CountPowerRentalOrders(ctx context.Context, in *npool.CountPowerRentalOrdersRequest) (*npool.CountPowerRentalOrdersResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CountPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountPowerRentalOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.CountPowerRentals(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CountPowerRentalOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CountPowerRentalOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CountPowerRentalOrdersResponse{
		Info: exist,
	}, nil
}
