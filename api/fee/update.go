package feeorder

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"

	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) UpdateFeeOrder(ctx context.Context, in *npool.UpdateFeeOrderRequest) (*npool.UpdateFeeOrderResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateFeeOrder",
			"In", in,
		)
		return &npool.UpdateFeeOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithID(req.ID, false),
		feeorder1.WithEntID(req.EntID, false),
		feeorder1.WithOrderID(req.OrderID, false),
		feeorder1.WithPaymentType(req.PaymentType, false),

		feeorder1.WithOrderState(req.OrderState, false),
		feeorder1.WithUserSetPaid(req.UserSetPaid, false),
		feeorder1.WithUserSetCanceled(req.UserSetCanceled, false),
		feeorder1.WithAdminSetCanceled(req.AdminSetCanceled, false),
		feeorder1.WithPaymentState(req.PaymentState, false),
		feeorder1.WithRollback(req.Rollback, false),
		feeorder1.WithLedgerLockID(req.LedgerLockID, false),
		feeorder1.WithPaymentID(req.PaymentID, false),
		feeorder1.WithPaymentBalances(req.PaymentBalances, false),
		feeorder1.WithPaymentTransfers(req.PaymentTransfers, false),

		feeorder1.WithMainOrder(func() *bool { b := true; return &b }(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFeeOrder",
			"Req", req,
			"error", err,
		)
		return &npool.UpdateFeeOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.UpdateFeeOrder(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateFeeOrder",
			"Req", req,
			"error", err,
		)
		return &npool.UpdateFeeOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateFeeOrderResponse{}, nil
}
