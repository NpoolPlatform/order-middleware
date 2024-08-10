package powerrental

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental"

	powerrental1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental"
)

func (s *Server) CreatePowerRentalOrder(ctx context.Context, in *npool.CreatePowerRentalOrderRequest) (*npool.CreatePowerRentalOrderResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrder",
			"In", in,
		)
		return &npool.CreatePowerRentalOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithEntID(req.EntID, false),
		powerrental1.WithAppID(req.AppID, true),
		powerrental1.WithUserID(req.UserID, true),
		powerrental1.WithGoodID(req.GoodID, true),
		powerrental1.WithGoodType(req.GoodType, true),
		powerrental1.WithAppGoodID(req.AppGoodID, true),
		powerrental1.WithOrderID(req.OrderID, false),
		powerrental1.WithOrderType(req.OrderType, true),
		powerrental1.WithPaymentType(req.PaymentType, false),
		powerrental1.WithSimulate(req.Simulate, false),
		powerrental1.WithCreateMethod(req.CreateMethod, true),

		powerrental1.WithAppGoodStockID(req.AppGoodStockID, false),
		powerrental1.WithUnits(req.Units, true),
		powerrental1.WithGoodValueUSD(req.GoodValueUSD, true),
		powerrental1.WithPaymentAmountUSD(req.PaymentAmountUSD, false),
		powerrental1.WithDiscountAmountUSD(req.DiscountAmountUSD, false),
		powerrental1.WithPromotionID(req.PromotionID, false),
		powerrental1.WithDurationSeconds(req.DurationSeconds, true),
		powerrental1.WithInvestmentType(req.InvestmentType, false),

		powerrental1.WithStartMode(req.StartMode, true),
		powerrental1.WithStartAt(req.StartAt, true),
		powerrental1.WithAppGoodStockLockID(req.AppGoodStockLockID, false),
		powerrental1.WithLedgerLockID(req.LedgerLockID, false),
		powerrental1.WithPaymentID(req.PaymentID, false),
		powerrental1.WithCouponIDs(req.CouponIDs, false),
		powerrental1.WithPaymentBalances(req.PaymentBalances, false),
		powerrental1.WithPaymentTransfers(req.PaymentTransfers, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrder",
			"Req", req,
			"error", err,
		)
		return &npool.CreatePowerRentalOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := handler.CreatePowerRental(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreatePowerRentalOrder",
			"Req", req,
			"error", err,
		)
		return &npool.CreatePowerRentalOrderResponse{}, status.Error(codes.Aborted, "internal error")
	}

	return &npool.CreatePowerRentalOrderResponse{}, nil
}
