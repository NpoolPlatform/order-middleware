package feeorder

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"

	feeorder1 "github.com/NpoolPlatform/order-middleware/pkg/mw/fee"
)

func (s *Server) DeleteFeeOrder(ctx context.Context, in *npool.DeleteFeeOrderRequest) (*npool.DeleteFeeOrderResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeleteFeeOrder",
			"In", in,
		)
		return &npool.DeleteFeeOrderResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := feeorder1.NewHandler(
		ctx,
		feeorder1.WithID(req.ID, false),
		feeorder1.WithEntID(req.EntID, false),
		feeorder1.WithOrderID(req.OrderID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeleteFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFeeOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	if err := handler.DeleteFeeOrder(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeleteFeeOrder",
			"In", in,
			"Error", err,
		)
		return &npool.DeleteFeeOrderResponse{}, status.Error(codes.Aborted, err.Error())
	}
	return &npool.DeleteFeeOrderResponse{}, nil
}
