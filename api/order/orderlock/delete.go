package orderlock

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"

	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderlock"
)

func (s *Server) DeleteOrderLocks(ctx context.Context, in *npool.DeleteOrderLocksRequest) (*npool.DeleteOrderLocksResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"DeleteOrderLocks",
			"In", in,
		)
		return &npool.DeleteOrderLocksResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := orderlock1.NewHandler(
		ctx,
		orderlock1.WithReqs(reqs),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderLocks",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderLocksResponse{}, status.Error(codes.Aborted, err.Error())
	}
	infos, err := handler.DeleteOrderLocks(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderLocks",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteOrderLocksResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteOrderLocksResponse{
		Infos: infos,
	}, nil
}
