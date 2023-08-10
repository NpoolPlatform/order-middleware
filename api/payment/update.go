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

func (s *Server) UpdatePayment(ctx context.Context, in *npool.UpdatePaymentRequest) (*npool.UpdatePaymentResponse, error) {
	req := in.GetInfo()
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(req.ID),
		payment1.WithState(req.State),
		payment1.WithUserSetPaid(req.UserSetPaid),
		payment1.WithUserSetCanceled(req.UserSetCanceled),
		payment1.WithFakePayment(req.FakePayment),
		payment1.WithFinishAmount(req.FinishAmount),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePayment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdatePayment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePayment",
			"In", in,
			"Error", err,
		)
		return &npool.UpdatePaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdatePaymentResponse{
		Info: info,
	}, nil
}
