//nolint:nolintlint,dupl
package payment

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	payment1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
)

func (s *Server) DeletePayment(ctx context.Context, in *npool.DeletePaymentRequest) (*npool.DeletePaymentResponse, error) {
	req := in.GetInfo()
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(req.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePayment",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.DeletePayment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePayment",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeletePaymentResponse{
		Info: info,
	}, nil
}
