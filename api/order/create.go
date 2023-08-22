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
		order1.WithID(req.ID, false),
		order1.WithGoodID(req.GoodID, true),
		order1.WithAppID(req.AppID, true),
		order1.WithUserID(req.UserID, true),
		order1.WithParentOrderID(req.ParentOrderID, false),
		order1.WithPayWithParent(req.PayWithParent, false),
		order1.WithUnits(req.Units, false),
		order1.WithPromotionID(req.PromotionID, false),
		order1.WithUserSpecialReductionID(req.SpecialOfferID, false),
		order1.WithStartAt(req.Start, false),
		order1.WithEndAt(req.End, false),
		order1.WithCouponIDs(req.CouponIDs, false),
		order1.WithPaymentID(req.PaymentID, false),
		order1.WithPaymentAccountID(req.PaymentAccountID, true),
		order1.WithPaymentAccountStartAmount(req.PaymentAccountStartAmount, true),
		order1.WithPaymentAmount(req.PaymentAmount, true),
		order1.WithPayWithBalanceAmount(req.PayWithBalanceAmount, false),
		order1.WithPaymentCoinUSDCurrency(req.PaymentCoinUSDCurrency, true),
		order1.WithPaymentLocalUSDCurrency(req.PaymentLocalUSDCurrency, true),
		order1.WithPaymentLiveUSDCurrency(req.PaymentLiveUSDCurrency, true),
		order1.WithPaymentCoinID(req.PaymentCoinID, true),
		order1.WithType(req.OrderType, true),
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
