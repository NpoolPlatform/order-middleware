//nolint:dupl
package order

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
)

func (s *Server) SumOrdersPaymentUSD(ctx context.Context, in *npool.SumOrdersPaymentUSDRequest) (*npool.SumOrdersPaymentUSDResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"SumOrdersPaymentUSD",
			"In", in,
			"Error", err,
		)
		return &npool.SumOrdersPaymentUSDResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.SumOrdersPaymentUSD(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"SumOrdersPaymentUSD",
			"In", in,
			"Error", err,
		)
		return &npool.SumOrdersPaymentUSDResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.SumOrdersPaymentUSDResponse{
		Info: exist,
	}, nil
}
