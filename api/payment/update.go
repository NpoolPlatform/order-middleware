package payment

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	payment1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
)

func (s *Server) UpdatePayment(ctx context.Context, in *npool.UpdatePaymentRequest) (*npool.UpdatePaymentResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdatePayment",
			"In", in,
		)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(req.ID, false),
		payment1.WithEntID(req.EntID, false),
		payment1.WithObseleteState(req.ObseleteState, false),
		payment1.WithPaymentTransfers(req.PaymentTransfers, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePayment",
			"Req", req,
			"error", err,
		)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.UpdatePayment(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdatePayment",
			"Req", req,
			"error", err,
		)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdatePaymentResponse{}, nil
}
