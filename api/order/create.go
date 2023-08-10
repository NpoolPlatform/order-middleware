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

func (s *Server) CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	req := in.GetInfo()
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(req.ID),
		order1.WithGoodID(req.GoodID),
		order1.WithAppID(req.AppID),
		order1.WithUserID(req.UserID),
		order1.WithParentOrderID(req.ParentOrderID),
		order1.WithPayWithParent(req.PayWithParent),
		order1.WithUnits(req.Units),
		order1.WithPromotionID(req.PromotionID),
		order1.WithDiscountCouponID(req.DiscountID),
		order1.WithUserSpecialReductionID(req.SpecialOfferID),
		order1.WithStartAt(req.Start),
		order1.WithEndAt(req.End),
		order1.WithFixAmountCouponID(req.FixAmountID),
		order1.WithCouponIDs(req.CouponIDs),
		order1.WithPaymentID(req.PaymentID),
		order1.WithPaymentAccountID(req.PaymentAccountID),
		order1.WithPaymentAccountStartAmount(req.PaymentAccountStartAmount),
		order1.WithPaymentAmount(req.PaymentAmount),
		order1.WithPayWithBalanceAmount(req.PayWithBalanceAmount),
		order1.WithPaymentCoinUSDCurrency(req.PaymentCoinUSDCurrency),
		order1.WithPaymentLocalUSDCurrency(req.PaymentLocalUSDCurrency),
		order1.WithPaymentLiveUSDCurrency(req.PaymentLiveUSDCurrency),
		order1.WithPaymentCoinID(req.PaymentCoinID),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrder",
			"Req", req,
			"error", err,
		)
		return &npool.CreateOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreateOrder(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrder",
			"Req", req,
			"error", err,
		)
		return &npool.CreateOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateOrderResponse{
		Info: info,
	}, nil
}
