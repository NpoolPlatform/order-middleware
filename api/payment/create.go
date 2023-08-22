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
		payment1.WithID(req.ID, false),
		payment1.WithAppID(req.AppID, true),
		payment1.WithGoodID(req.GoodID, true),
		payment1.WithUserID(req.UserID, true),
		payment1.WithOrderID(req.OrderID, true),
		payment1.WithAccountID(req.AccountID, true),
		payment1.WithStartAmount(req.StartAmount, true),
		payment1.WithAmount(req.Amount, true),
		payment1.WithPayWithBalanceAmount(req.PayWithBalanceAmount, false),
		payment1.WithFinishAmount(req.FinishAmount, false),
		payment1.WithCoinUsdCurrency(req.CoinUsdCurrency, true),
		payment1.WithLocalCoinUsdCurrency(req.LocalCoinUsdCurrency, true),
		payment1.WithLiveCoinUsdCurrency(req.LiveCoinUsdCurrency, true),
		payment1.WithCoinInfoID(req.CoinInfoID, true),
		payment1.WithState(req.State, true),
		payment1.WithChainTransactionID(req.ChainTransactionID, false),
		payment1.WithUserSetPaid(req.UserSetPaid, false),
		payment1.WithUserSetCanceled(req.UserSetCanceled, false),
		payment1.WithFakePayment(req.FakePayment, false),
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
