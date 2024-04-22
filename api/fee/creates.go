package feeorder

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"

	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) CreateFeeOrders(ctx context.Context, in *npool.CreateFeeOrdersRequest) (*npool.CreateFeeOrdersResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"CreateFeeOrders",
			"In", in,
		)
		return &npool.CreateFeeOrdersResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	multiHandler := &feeorder1.MultiHandler{}
	for _, req := range reqs {
		handler, err := feeorder1.NewHandler(
			ctx,
			feeorder1.WithEntID(req.EntID, false),
			feeorder1.WithAppID(req.AppID, true),
			feeorder1.WithUserID(req.UserID, true),
			feeorder1.WithGoodID(req.GoodID, true),
			feeorder1.WithGoodType(req.GoodType, true),
			feeorder1.WithAppGoodID(req.AppGoodID, true),
			feeorder1.WithOrderID(req.OrderID, false),
			feeorder1.WithParentOrderID(req.ParentOrderID, false),
			feeorder1.WithOrderType(req.OrderType, true),
			feeorder1.WithPaymentType(req.PaymentType, true),
			feeorder1.WithCreateMethod(req.CreateMethod, true),

			feeorder1.WithGoodValueUSD(req.GoodValueUSD, true),
			feeorder1.WithPaymentAmountUSD(req.PaymentAmountUSD, false),
			feeorder1.WithDiscountAmountUSD(req.DiscountAmountUSD, false),
			feeorder1.WithPromotionID(req.PromotionID, false),
			feeorder1.WithDurationSeconds(req.DurationSeconds, true),
			feeorder1.WithLedgerLockID(req.LedgerLockID, false),
			feeorder1.WithPaymentID(req.PaymentID, false),
			feeorder1.WithCouponIDs(req.CouponIDs, false),
			feeorder1.WithPaymentBalances(req.PaymentBalances, false),
			feeorder1.WithPaymentTransfers(req.PaymentTransfers, false),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"CreateFeeOrders",
				"Req", req,
				"error", err,
			)
			return &npool.CreateFeeOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}
	if err := multiHandler.CreateFeeOrders(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateFeeOrders",
			"Reqs", reqs,
			"error", err,
		)
		return &npool.CreateFeeOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateFeeOrdersResponse{}, nil
}
