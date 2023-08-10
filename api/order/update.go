//nolint:nolintlint,dupl
package order

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
)

func (s *Server) UpdateOrder(ctx context.Context, in *npool.UpdateOrderRequest) (*npool.UpdateOrderResponse, error) {
	req := in.GetInfo()
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(req.ID),
		order1.WithAppID(req.AppID),
		order1.WithPaymentID(req.PaymentID),
		order1.WithStartAt(req.Start),
		order1.WithLastBenefitAt(req.LastBenefitAt),
		order1.WithPaymentUserSetCanceled(req.Canceled),
		order1.WithPaymentFinishAmount(req.PaymentFinishAmount),
		order1.WithPaymentFakePayment(req.FakePayment),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	info, err := handler.UpdateOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrder",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateOrders(ctx context.Context, in *npool.UpdateOrdersRequest) (*npool.UpdateOrdersResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrders",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.UpdateOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateOrders",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.UpdateOrdersResponse{
		Infos: infos,
	}, nil
}
