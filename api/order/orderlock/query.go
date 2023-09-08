package orderlock

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"

	orderlock1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/orderlock"
)

func (s *Server) GetOrderLock(ctx context.Context, in *npool.GetOrderLockRequest) (*npool.GetOrderLockResponse, error) {
	handler, err := orderlock1.NewHandler(
		ctx,
		orderlock1.WithID(&in.ID, true),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderLock",
			"In", in,
			"error", err,
		)
		return &npool.GetOrderLockResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := handler.GetOrderLock(ctx)
	if err != nil {
		return &npool.GetOrderLockResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderLockResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOrderLocks(ctx context.Context, in *npool.GetOrderLocksRequest) (*npool.GetOrderLocksResponse, error) {
	handler, err := orderlock1.NewHandler(
		ctx,
		orderlock1.WithConds(in.GetConds()),
		orderlock1.WithOffset(in.GetOffset()),
		orderlock1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderLocks",
			"In", in,
			"error", err,
		)
		return &npool.GetOrderLocksResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetOrderLocks(ctx)
	if err != nil {
		return &npool.GetOrderLocksResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderLocksResponse{
		Infos: infos,
		Total: total,
	}, nil
}
