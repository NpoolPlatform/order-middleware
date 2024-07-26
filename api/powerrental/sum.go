//nolint:dupl
package powerrental

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"
	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) SumPowerRentalOrderUnits(ctx context.Context, in *npool.SumPowerRentalOrderUnitsRequest) (*npool.SumPowerRentalOrderUnitsResponse, error) {
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SumPowerRentalOrderUnits",
			"In", in,
			"Error", err,
		)
		return &npool.SumPowerRentalOrderUnitsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.SumPowerRentalUnits(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SumPowerRentalOrderUnits",
			"In", in,
			"Error", err,
		)
		return &npool.SumPowerRentalOrderUnitsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.SumPowerRentalOrderUnitsResponse{
		Info: exist,
	}, nil
}
