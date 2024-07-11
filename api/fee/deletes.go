package feeorder

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"

	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) DeleteFeeOrders(ctx context.Context, in *npool.DeleteFeeOrdersRequest) (*npool.DeleteFeeOrdersResponse, error) {
	reqs := in.GetInfos()
	if len(reqs) == 0 {
		logger.Sugar().Errorw(
			"DeleteFeeOrders",
			"In", in,
		)
		return &npool.DeleteFeeOrdersResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	multiHandler := &feeorder1.MultiHandler{}
	for _, req := range reqs {
		handler, err := feeorder1.NewHandler(
			ctx,
			feeorder1.WithID(req.ID, false),
			feeorder1.WithEntID(req.EntID, false),
			feeorder1.WithOrderID(req.OrderID, false),
		)
		if err != nil {
			logger.Sugar().Errorw(
				"DeleteFeeOrders",
				"Req", req,
				"error", err,
			)
			return &npool.DeleteFeeOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
		multiHandler.AppendHandler(handler)
	}
	if err := multiHandler.DeleteFeeOrders(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteFeeOrders",
			"Reqs", reqs,
			"error", err,
		)
		return &npool.DeleteFeeOrdersResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeleteFeeOrdersResponse{}, nil
}
