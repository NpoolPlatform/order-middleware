package payment

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	payment1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
)

func (s *Server) GetPayments(ctx context.Context, in *npool.GetPaymentsRequest) (*npool.GetPaymentsResponse, error) {
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithConds(in.GetConds()),
		payment1.WithOffset(in.GetOffset()),
		payment1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetPayments",
			"In", in,
			"error", err,
		)
		return &npool.GetPaymentsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetPayments(ctx)
	if err != nil {
		return &npool.GetPaymentsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetPaymentsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
