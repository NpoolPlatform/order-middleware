package orderlock

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/lock"

	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/lock"
)

func (s *Server) CreateOrderLocks(ctx context.Context, in *npool.CreateOrderLocksRequest) (*npool.CreateOrderLocksResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"CreateOrderLocks",
			"In", in,
		)
		return &npool.CreateOrderLocksResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	multiHandler := &orderlock1.MultiHandler{}
	for _, req := range reqs {
		if req.LockType == nil || *req.LockType != types.OrderLockType_LockCommission {
			logger.Sugar().Errorw(
				"CreateOrderLocks",
				"In", in,
			)
			return &npool.CreateOrderLocksResponse{}, status.Error(codes.Aborted, "invalid locktype")
		}
		handler, err := orderlock1.NewHandler(
			ctx,
			orderlock1.WithEntID(req.EntID, false),
			orderlock1.WithUserID(req.UserID, true),
			orderlock1.WithOrderID(req.OrderID, true),
			orderlock1.WithLockType(req.LockType, true),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"CreateOrderLocks",
				"Req", req,
				"error", err,
			)
			return &npool.CreateOrderLocksResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}
	if err := multiHandler.CreateOrderLocks(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateOrderLocks",
			"Reqs", reqs,
			"error", err,
		)
		return &npool.CreateOrderLocksResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateOrderLocksResponse{}, nil
}
