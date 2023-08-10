//nolint:nolintlint,dupl
package payment

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	payment1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
)

func (s *Server) ExistPayment(ctx context.Context, in *npool.ExistPaymentRequest) (*npool.ExistPaymentResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(&in.ID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPayment",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPayment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPayment",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPaymentResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistPaymentConds(ctx context.Context, in *npool.ExistPaymentCondsRequest) (*npool.ExistPaymentCondsResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPaymentConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPaymentCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistPaymentConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistPaymentConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistPaymentCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistPaymentCondsResponse{
		Info: exist,
	}, nil
}
