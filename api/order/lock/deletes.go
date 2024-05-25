package orderlock

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/lock"

	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/lock"
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
	multiHandler := &orderlock1.MultiHandler{}
	for _, req := range reqs {
		handler, err := orderlock1.NewHandler(
			ctx,
			orderlock1.WithID(req.ID, false),
			orderlock1.WithEntID(req.EntID, false),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"DeleteOrderLocks",
				"Req", req,
				"error", err,
			)
			return &npool.DeleteOrderLocksResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}
	if err := multiHandler.DeleteOrderLocks(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteOrderLocks",
			"Reqs", reqs,
			"error", err,
		)
		return &npool.DeleteOrderLocksResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteOrderLocksResponse{}, nil
}
