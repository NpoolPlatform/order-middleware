//nolint:dupl
package ordercoupon

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	ordercoupon1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/coupon"
)

func (s *Server) ExistOrderCouponConds(ctx context.Context, in *npool.ExistOrderCouponCondsRequest) (*npool.ExistOrderCouponCondsResponse, error) {
	handler, err := ordercoupon1.NewHandler(
		ctx,
		ordercoupon1.WithConds(in.GetConds()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderCouponConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderCouponCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	exist, err := handler.ExistOrderCouponConds(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"ExistOrderCouponConds",
			"In", in,
			"Error", err,
		)
		return &npool.ExistOrderCouponCondsResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.ExistOrderCouponCondsResponse{
		Info: exist,
	}, nil
}
