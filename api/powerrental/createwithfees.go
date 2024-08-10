package powerrental

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"

	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) CreatePowerRentalOrderWithFees(ctx context.Context, in *npool.CreatePowerRentalOrderWithFeesRequest) (*npool.CreatePowerRentalOrderWithFeesResponse, error) {
	powerRentalOrder := in.GetPowerRentalOrder()
	if powerRentalOrder == nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrderWithFees",
			"Req", in,
		)
		return &npool.CreatePowerRentalOrderWithFeesResponse{}, status.Error(codes.InvalidArgument, "invalid info")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithEntID(powerRentalOrder.EntID, false),
		powerrental1.WithAppID(powerRentalOrder.AppID, true),
		powerrental1.WithUserID(powerRentalOrder.UserID, true),
		powerrental1.WithGoodID(powerRentalOrder.GoodID, true),
		powerrental1.WithGoodType(powerRentalOrder.GoodType, true),
		powerrental1.WithAppGoodID(powerRentalOrder.AppGoodID, true),
		powerrental1.WithOrderID(powerRentalOrder.OrderID, false),
		powerrental1.WithOrderType(powerRentalOrder.OrderType, true),
		powerrental1.WithPaymentType(powerRentalOrder.PaymentType, false),
		powerrental1.WithSimulate(powerRentalOrder.Simulate, false),
		powerrental1.WithCreateMethod(powerRentalOrder.CreateMethod, true),

		powerrental1.WithAppGoodStockID(powerRentalOrder.AppGoodStockID, false),
		powerrental1.WithUnits(powerRentalOrder.Units, true),
		powerrental1.WithGoodValueUSD(powerRentalOrder.GoodValueUSD, true),
		powerrental1.WithPaymentAmountUSD(powerRentalOrder.PaymentAmountUSD, false),
		powerrental1.WithDiscountAmountUSD(powerRentalOrder.DiscountAmountUSD, false),
		powerrental1.WithPromotionID(powerRentalOrder.PromotionID, false),
		powerrental1.WithDurationSeconds(powerRentalOrder.DurationSeconds, true),
		powerrental1.WithInvestmentType(powerRentalOrder.InvestmentType, false),

		powerrental1.WithStartMode(powerRentalOrder.StartMode, true),
		powerrental1.WithStartAt(powerRentalOrder.StartAt, true),
		powerrental1.WithAppGoodStockLockID(powerRentalOrder.AppGoodStockLockID, false),
		powerrental1.WithLedgerLockID(powerRentalOrder.LedgerLockID, false),
		powerrental1.WithPaymentID(powerRentalOrder.PaymentID, false),
		powerrental1.WithCouponIDs(powerRentalOrder.CouponIDs, false),
		powerrental1.WithPaymentBalances(powerRentalOrder.PaymentBalances, false),
		powerrental1.WithPaymentTransfers(powerRentalOrder.PaymentTransfers, false),

		powerrental1.WithFeeOrders(in.GetFeeOrders(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrderWithFees",
			"Req", in,
			"error", err,
		)
		return &npool.CreatePowerRentalOrderWithFeesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.CreatePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrderWithFees",
			"Req", in,
			"error", err,
		)
		return &npool.CreatePowerRentalOrderWithFeesResponse{}, status.Error(codes.Aborted, "internal error")
	}

	return &npool.CreatePowerRentalOrderWithFeesResponse{}, nil
}
