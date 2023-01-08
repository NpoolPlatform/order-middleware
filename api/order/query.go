//nolint:nolintlint,dupl,gocyclo
package order

import (
	"context"

	ordermgrpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/order"

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

	if in.Conds != nil {
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
		if in.GetConds().CouponID != nil {
			if _, err := uuid.Parse(in.GetConds().GetCouponID().GetValue()); err != nil {
				logger.Sugar().Errorw("GetOrders", "CouponID", in.GetConds().GetCouponID().GetValue(), "error", err)
				return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		for _, id := range in.GetConds().GetCouponIDs().GetValue() {
			if _, err := uuid.Parse(id); err != nil {
				logger.Sugar().Errorw("GetOrders", "error", err)
				return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
			}
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

func (s *Server) GetOrderOnly(ctx context.Context, in *npool.GetOrderOnlyRequest) (*npool.GetOrderOnlyResponse, error) {
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
			return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().GoodID != nil {
		if _, err := uuid.Parse(in.GetConds().GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "GoodID", in.GetConds().GetGoodID().GetValue(), "error", err)
			return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().AppID != nil {
		if _, err := uuid.Parse(in.GetConds().GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "AppID", in.GetConds().GetAppID().GetValue(), "error", err)
			return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().UserID != nil {
		if _, err := uuid.Parse(in.GetConds().GetUserID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "UserID", in.GetConds().GetUserID().GetValue(), "error", err)
			return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if in.GetConds().Type != nil {
		switch ordermgrpb.OrderType(in.GetConds().Type.GetValue()) {
		case ordermgrpb.OrderType_Normal:
		case ordermgrpb.OrderType_Offline:
		case ordermgrpb.OrderType_Airdrop:
		default:
			logger.Sugar().Errorw("validate", "OrderType", in.GetConds().Type)
			return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, "OrderType is invalid")
		}
	}
	if in.GetConds().State != nil {
		switch ordermgrpb.OrderState(in.GetConds().State.GetValue()) {
		case ordermgrpb.OrderState_WaitPayment:
		case ordermgrpb.OrderState_Paid:
		case ordermgrpb.OrderState_PaymentTimeout:
		case ordermgrpb.OrderState_Canceled:
		case ordermgrpb.OrderState_InService:
		case ordermgrpb.OrderState_Expired:
		default:
			logger.Sugar().Errorw("validate", "OrderType", in.GetConds().Type)
			return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, "OrderType is invalid")
		}
	}

	info, err := order1.GetOrderOnly(ctx, in.Conds)
	if err != nil {
		logger.Sugar().Errorw("GetOrders", "error", err)
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderOnlyResponse{
		Info: info,
	}, nil
}
