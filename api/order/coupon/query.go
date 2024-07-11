//nolint:nolintlint,dupl,gocyclo
package ordercoupon

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"
	ordercoupon1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order/coupon"
)

func (s *Server) GetOrderCoupons(ctx context.Context, in *npool.GetOrderCouponsRequest) (*npool.GetOrderCouponsResponse, error) {
	handler, err := ordercoupon1.NewHandler(
		ctx,
		ordercoupon1.WithConds(in.GetConds()),
		ordercoupon1.WithOffset(in.GetOffset()),
		ordercoupon1.WithLimit(in.GetLimit()),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"GetOrderCoupons",
			"In", in,
			"error", err,
		)
		return &npool.GetOrderCouponsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	infos, total, err := handler.GetOrderCoupons(ctx)
	if err != nil {
		return &npool.GetOrderCouponsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderCouponsResponse{
		Infos: infos,
		Total: total,
	}, nil
}
