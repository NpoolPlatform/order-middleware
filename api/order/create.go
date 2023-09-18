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
	if req == nil {
		logger.Sugar().Errorw(
			"CreateOrder",
			"In", in,
		)
		return &npool.CreateOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := order1.NewHandler(
		ctx,
		order1.WithID(req.ID, false),
		order1.WithAppID(req.AppID, true),
		order1.WithUserID(req.UserID, true),
		order1.WithGoodID(req.GoodID, true),
		order1.WithAppGoodID(req.AppGoodID, true),
		order1.WithParentOrderID(req.ParentOrderID, false),
		order1.WithUnits(req.Units, true),
		order1.WithGoodValue(req.GoodValue, true),
		order1.WithGoodValueUSD(req.GoodValueUSD, true),
		order1.WithPaymentAmount(req.PaymentAmount, false),
		order1.WithDiscountAmount(req.DiscountAmount, false),
		order1.WithPromotionID(req.PromotionID, false),
		order1.WithDurationDays(req.DurationDays, true),
		order1.WithOrderType(req.OrderType, true),
		order1.WithInvestmentType(req.InvestmentType, true),
		order1.WithCouponIDs(req.CouponIDs, false),
		order1.WithPaymentType(req.PaymentType, true),
		order1.WithCoinTypeID(req.CoinTypeID, true),
		order1.WithPaymentCoinTypeID(req.PaymentCoinTypeID, true),
		order1.WithTransferAmount(req.TransferAmount, false),
		order1.WithBalanceAmount(req.BalanceAmount, false),
		order1.WithCoinUSDCurrency(req.CoinUSDCurrency, true),
		order1.WithLocalCoinUSDCurrency(req.LocalCoinUSDCurrency, false),
		order1.WithLiveCoinUSDCurrency(req.LiveCoinUSDCurrency, true),
		order1.WithPaymentAccountID(req.PaymentAccountID, false),
		order1.WithPaymentStartAmount(req.PaymentStartAmount, false),
		order1.WithStartMode(req.StartMode, false),
		order1.WithStartAt(req.StartAt, true),
		order1.WithEndAt(req.EndAt, true),
		order1.WithPaymentState(req.PaymentState, false),
		order1.WithAppGoodStockLockID(req.AppGoodStockLockID, false),
		order1.WithLedgerLockID(req.LedgerLockID, false),
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

func (s *Server) CreateOrders(ctx context.Context, in *npool.CreateOrdersRequest) (*npool.CreateOrdersResponse, error) {
	handler, err := order1.NewHandler(
		ctx,
		order1.WithReqs(in.GetInfos(), true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateOrders(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrders",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateOrdersResponse{
		Infos: infos,
	}, nil
}
