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
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateOrder",
			"In", in,
		)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(req.ID, true),
		order1.WithAppID(req.AppID, false),
		order1.WithParentOrderID(req.ParentOrderID, false),
		order1.WithOrderState(req.OrderState, false),
		order1.WithCancelState(req.CancelState, false),
		order1.WithStartMode(req.StartMode, false),
		order1.WithStartAt(req.StartAt, false),
		order1.WithEndAt(req.EndAt, false),
		order1.WithLastBenefitAt(req.LastBenefitAt, false),
		order1.WithBenefitState(req.BenefitState, false),
		order1.WithUserSetPaid(req.UserSetPaid, false),
		order1.WithUserSetCanceled(req.UserSetCanceled, false),
		order1.WithAdminSetCanceled(req.AdminSetCanceled, false),
		order1.WithPaymentTransactionID(req.PaymentTransactionID, false),
		order1.WithPaymentFinishAmount(req.PaymentFinishAmount, false),
		order1.WithPaymentState(req.PaymentState, false),
		order1.WithOutOfGasHours(req.OutOfGasHours, false),
		order1.WithCompensateHours(req.CompensateHours, false),
		order1.WithCommissionLockID(req.CommissionLockID, false),
		order1.WithRollback(req.Rollback, false),
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
		order1.WithReqs(in.GetInfos(), false),
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
