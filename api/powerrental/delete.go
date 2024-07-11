package powerrental

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"

	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) DeletePowerRentalOrder(ctx context.Context, in *npool.DeletePowerRentalOrderRequest) (*npool.DeletePowerRentalOrderResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeletePowerRentalOrder",
			"In", in,
		)
		return &npool.DeletePowerRentalOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(req.ID, false),
		powerrental1.WithEntID(req.EntID, false),
		powerrental1.WithOrderID(req.OrderID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePowerRentalOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	if err := handler.DeletePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeletePowerRentalOrder",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePowerRentalOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeletePowerRentalOrderResponse{}, nil
}
