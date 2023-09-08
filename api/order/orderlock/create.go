package orderlock

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"

	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderlock"
)

func (s *Server) CreateOrderLocks(ctx context.Context, in *npool.CreateOrderLocksRequest) (*npool.CreateOrderLocksResponse, error) {
	handler, err := orderlock1.NewHandler(
		ctx,
		orderlock1.WithReqs(in.GetInfos()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderLocks",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderLocksResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.CreateOrderLocks(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateOrderLocks",
			"In", in,
			"Error", err,
		)
		return &npool.CreateOrderLocksResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.CreateOrderLocksResponse{
		Infos: infos,
	}, nil
}
