//nolint:dupl
package powerrental

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"

	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) UpdatePowerRentalOrder(ctx context.Context, in *npool.UpdatePowerRentalOrderRequest) (*npool.UpdatePowerRentalOrderResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdatePowerRentalOrder",
			"In", in,
		)
		return &npool.UpdatePowerRentalOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithID(req.ID, false),
		powerrental1.WithEntID(req.EntID, false),
		powerrental1.WithOrderID(req.OrderID, false),
		powerrental1.WithPaymentType(req.PaymentType, false),

		powerrental1.WithOrderState(req.OrderState, false),
		powerrental1.WithStartMode(req.StartMode, false),
		powerrental1.WithStartAt(req.StartAt, false),
		powerrental1.WithLastBenefitAt(req.LastBenefitAt, false),
		powerrental1.WithBenefitState(req.BenefitState, false),
		powerrental1.WithUserSetPaid(req.UserSetPaid, false),
		powerrental1.WithUserSetCanceled(req.UserSetCanceled, false),
		powerrental1.WithAdminSetCanceled(req.AdminSetCanceled, false),
		powerrental1.WithPaymentState(req.PaymentState, false),
		powerrental1.WithRenewState(req.RenewState, false),
		powerrental1.WithRenewNotifyAt(req.RenewNotifyAt, false),

		powerrental1.WithLedgerLockID(req.LedgerLockID, false),
		powerrental1.WithPaymentID(req.PaymentID, false),
		powerrental1.WithCouponIDs(req.CouponIDs, false),
		powerrental1.WithPaymentBalances(req.PaymentBalances, false),
		powerrental1.WithPaymentTransfers(req.PaymentTransfers, false),

		powerrental1.WithRollback(req.Rollback, false),
		powerrental1.WithPoolOrderUserID(req.PoolOrderUserID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRentalOrder",
			"Req", req,
			"error", err,
		)
		return &npool.UpdatePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, "internal error")
	}
	if err := handler.UpdatePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRentalOrder",
			"Req", req,
			"error", err,
		)
		return &npool.UpdatePowerRentalOrderResponse{}, status.Error(codes.Aborted, "internal error")
	}

	return &npool.UpdatePowerRentalOrderResponse{}, nil
}

func (s *Server) UpdatePowerRentalOrders(ctx context.Context, in *npool.UpdatePowerRentalOrdersRequest) (*npool.UpdatePowerRentalOrdersResponse, error) {
	multiHandler := &powerrental1.MultiHandler{}
	for _, req := range in.GetInfos() {
		handler, err := powerrental1.NewHandler(
			ctx,
			powerrental1.WithID(req.ID, false),
			powerrental1.WithEntID(req.EntID, false),
			powerrental1.WithOrderID(req.OrderID, false),
			powerrental1.WithPaymentType(req.PaymentType, false),

			powerrental1.WithOrderState(req.OrderState, false),
			powerrental1.WithStartMode(req.StartMode, false),
			powerrental1.WithStartAt(req.StartAt, false),
			powerrental1.WithLastBenefitAt(req.LastBenefitAt, false),
			powerrental1.WithBenefitState(req.BenefitState, false),
			powerrental1.WithUserSetPaid(req.UserSetPaid, false),
			powerrental1.WithUserSetCanceled(req.UserSetCanceled, false),
			powerrental1.WithAdminSetCanceled(req.AdminSetCanceled, false),
			powerrental1.WithPaymentState(req.PaymentState, false),
			powerrental1.WithRenewState(req.RenewState, false),
			powerrental1.WithRenewNotifyAt(req.RenewNotifyAt, false),

			powerrental1.WithLedgerLockID(req.LedgerLockID, false),
			powerrental1.WithPaymentID(req.PaymentID, false),
			powerrental1.WithCouponIDs(req.CouponIDs, false),
			powerrental1.WithPaymentBalances(req.PaymentBalances, false),
			powerrental1.WithPaymentTransfers(req.PaymentTransfers, false),

			powerrental1.WithRollback(req.Rollback, false),
			powerrental1.WithPoolOrderUserID(req.PoolOrderUserID, false),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"UpdatePowerRentalOrders",
				"Req", req,
				"error", err,
			)
			return &npool.UpdatePowerRentalOrdersResponse{}, status.Error(codes.InvalidArgument, "internal error")
		}
		multiHandler.AppendHandler(handler)
	}
	if err := multiHandler.UpdatePowerRentals(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdatePowerRentalOrders",
			"Req", in,
			"error", err,
		)
		return &npool.UpdatePowerRentalOrdersResponse{}, status.Error(codes.Aborted, "internal error")
	}

	return &npool.UpdatePowerRentalOrdersResponse{}, nil
}
