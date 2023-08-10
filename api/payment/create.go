//nolint:nolintlint,dupl
package payment

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	payment1 "github.com/NpoolPlatform/order-middleware/pkg/mw/payment"
)

func (s *Server) CreatePayment(ctx context.Context, in *npool.CreatePaymentRequest) (*npool.CreatePaymentResponse, error) {
	req := in.GetInfo()
	handler, err := payment1.NewHandler(
		ctx,
		payment1.WithID(req.ID),
		payment1.WithAppID(req.AppID),
		payment1.WithGoodID(req.GoodID),
		payment1.WithUserID(req.UserID),
		payment1.WithOrderID(req.OrderID),
		payment1.WithAccountID(req.AccountID),
		payment1.WithStartAmount(req.StartAmount),
		payment1.WithAmount(req.Amount),
		payment1.WithPayWithBalanceAmount(req.PayWithBalanceAmount),
		payment1.WithFinishAmount(req.FinishAmount),
		payment1.WithCoinUsdCurrency(req.CoinUsdCurrency),
		payment1.WithLocalCoinUsdCurrency(req.LocalCoinUsdCurrency),
		payment1.WithLiveCoinUsdCurrency(req.LiveCoinUsdCurrency),
		payment1.WithCoinInfoID(req.CoinInfoID),
		payment1.WithState(req.State),
		payment1.WithChainTransactionID(req.ChainTransactionID),
		payment1.WithUserSetPaid(req.UserSetPaid),
		payment1.WithUserSetCanceled(req.UserSetCanceled),
		payment1.WithFakePayment(req.FakePayment),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePayment",
			"Req", req,
			"error", err,
		)
		return &npool.CreatePaymentResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.CreatePayment(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreatePayment",
			"Req", req,
			"error", err,
		)
		return &npool.CreatePaymentResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreatePaymentResponse{
		Info: info,
	}, nil
}
