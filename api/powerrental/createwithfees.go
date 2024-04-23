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
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithEntID(in.GetPowerRentalOrder().EntID, false),
		powerrental1.WithAppID(in.GetPowerRentalOrder().AppID, true),
		powerrental1.WithUserID(in.GetPowerRentalOrder().UserID, true),
		powerrental1.WithGoodID(in.GetPowerRentalOrder().GoodID, true),
		powerrental1.WithGoodType(in.GetPowerRentalOrder().GoodType, true),
		powerrental1.WithAppGoodID(in.GetPowerRentalOrder().AppGoodID, true),
		powerrental1.WithOrderID(in.GetPowerRentalOrder().OrderID, false),
		powerrental1.WithOrderType(in.GetPowerRentalOrder().OrderType, true),
		powerrental1.WithPaymentType(in.GetPowerRentalOrder().PaymentType, true),
		powerrental1.WithSimulate(in.GetPowerRentalOrder().Simulate, false),
		powerrental1.WithCreateMethod(in.GetPowerRentalOrder().CreateMethod, true),

		powerrental1.WithAppGoodStockID(in.GetPowerRentalOrder().AppGoodStockID, false),
		powerrental1.WithUnits(in.GetPowerRentalOrder().Units, true),
		powerrental1.WithGoodValueUSD(in.GetPowerRentalOrder().GoodValueUSD, true),
		powerrental1.WithPaymentAmountUSD(in.GetPowerRentalOrder().PaymentAmountUSD, false),
		powerrental1.WithDiscountAmountUSD(in.GetPowerRentalOrder().DiscountAmountUSD, false),
		powerrental1.WithPromotionID(in.GetPowerRentalOrder().PromotionID, false),
		powerrental1.WithDurationSeconds(in.GetPowerRentalOrder().DurationSeconds, true),
		powerrental1.WithInvestmentType(in.GetPowerRentalOrder().InvestmentType, false),

		powerrental1.WithStartMode(in.GetPowerRentalOrder().StartMode, true),
		powerrental1.WithStartAt(in.GetPowerRentalOrder().StartAt, true),
		powerrental1.WithAppGoodStockLockID(in.GetPowerRentalOrder().AppGoodStockLockID, true),
		powerrental1.WithLedgerLockID(in.GetPowerRentalOrder().LedgerLockID, false),
		powerrental1.WithPaymentID(in.GetPowerRentalOrder().PaymentID, false),
		powerrental1.WithCouponIDs(in.GetPowerRentalOrder().CouponIDs, false),
		powerrental1.WithPaymentBalances(in.GetPowerRentalOrder().PaymentBalances, false),
		powerrental1.WithPaymentTransfers(in.GetPowerRentalOrder().PaymentTransfers, false),

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
		return &npool.CreatePowerRentalOrderWithFeesResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePowerRentalOrderWithFeesResponse{}, nil
}
