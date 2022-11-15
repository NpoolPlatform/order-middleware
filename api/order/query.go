//nolint:nolintlint,dupl
package order

import (
	"context"

	order1 "github.com/NpoolPlatform/order-middleware/pkg/order"
	commontracer "github.com/NpoolPlatform/order-middleware/pkg/tracer"

	constant1 "github.com/NpoolPlatform/order-middleware/pkg/const"
	constant "github.com/NpoolPlatform/order-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
)

func (s *Server) GetOrder(ctx context.Context, in *npool.GetOrderRequest) (*npool.GetOrderResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOrder")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())
	span = commontracer.TraceInvoker(span, "order", "middleware", "GetOrder")

	if _, err := uuid.Parse(in.GetID()); err != nil {
		logger.Sugar().Errorw("GetOrder", "ID", in.GetID(), "error", err)
		return &npool.GetOrderResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := order1.GetOrder(ctx, in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetOrder", "error", err)
		return &npool.GetOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) GetOrders(ctx context.Context, in *npool.GetOrdersRequest) (*npool.GetOrdersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOrders")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "order", "middleware", "GetOrders")

	if in.GetConds().ID != nil {
		if _, err := uuid.Parse(in.GetConds().GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "ID", in.GetConds().GetID().GetValue(), "error", err)
			return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().GoodID != nil {
		if _, err := uuid.Parse(in.GetConds().GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().AppID != nil {
		if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
			return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().UserID != nil {
		if _, err := uuid.Parse(in.GetConds().GetUserID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "UserID", in.GetConds().GetUserID().GetValue(), "error", err)
			return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	limit := in.GetLimit()
	if limit == 0 {
		limit = constant1.DefaultLimitRows
	}

	infos, total, err := order1.GetOrders(ctx, in.Conds, in.GetOffset(), limit)
	if err != nil {
		logger.Sugar().Errorw("GetOrders", "error", err)
		return &npool.GetOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrdersResponse{
		Infos: infos,
		Total: total,
	}, nil
}
