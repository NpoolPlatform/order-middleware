//nolint:nolintlint,dupl,gocyclo
package order

import (
	"context"
	"fmt"

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

func ValidateConds(conds *npool.Conds) error {
	if conds == nil {
		return nil
	}

	if conds.ID != nil {
		if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "ID", conds.GetID().GetValue(), "error", err)
			return err
		}
	}
	if conds.GoodID != nil {
		if _, err := uuid.Parse(conds.GetGoodID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "GoodID", conds.GetGoodID().GetValue(), "error", err)
			return err
		}
	}
	if conds.AppID != nil {
		if _, err := uuid.Parse(conds.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "AppID", conds.GetAppID().GetValue(), "error", err)
			return err
		}
	}
	if conds.UserID != nil {
		if _, err := uuid.Parse(conds.GetUserID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "UserID", conds.GetUserID().GetValue(), "error", err)
			return err
		}
	}
	if conds.Type != nil {
		switch ordermgrpb.OrderType(conds.Type.GetValue()) {
		case ordermgrpb.OrderType_Normal:
		case ordermgrpb.OrderType_Offline:
		case ordermgrpb.OrderType_Airdrop:
		default:
			logger.Sugar().Errorw("validate", "OrderType", conds.GetType().GetValue())
			return fmt.Errorf("ordertype is invalid")
		}
	}
	if conds.State != nil {
		switch ordermgrpb.OrderState(conds.State.GetValue()) {
		case ordermgrpb.OrderState_WaitPayment:
		case ordermgrpb.OrderState_Paid:
		case ordermgrpb.OrderState_PaymentTimeout:
		case ordermgrpb.OrderState_Canceled:
		case ordermgrpb.OrderState_InService:
		case ordermgrpb.OrderState_Expired:
		default:
			logger.Sugar().Errorw("validate", "OrderState", conds.GetState().GetValue())
			return fmt.Errorf("orderstate is invalid")
		}
	}
	if conds.CouponID != nil {
		if _, err := uuid.Parse(conds.GetCouponID().GetValue()); err != nil {
			logger.Sugar().Errorw("GetOrders", "CouponID", conds.GetCouponID().GetValue(), "error", err)
			return err
		}
	}
	for _, id := range conds.GetCouponIDs().GetValue() {
		if _, err := uuid.Parse(id); err != nil {
			logger.Sugar().Errorw("GetOrders", "error", err)
			return err
		}
	}

	return nil
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

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetOrderOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "order", "middleware", "GetOrderOnly")

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	info, err := order1.GetOrderOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetOrderOnly", "error", err)
		return &npool.GetOrderOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetOrderOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) CountOrders(ctx context.Context, in *npool.CountOrdersRequest) (*npool.CountOrdersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountOrders")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceInvoker(span, "order", "middleware", "CountOrders")

	if err := ValidateConds(in.GetConds()); err != nil {
		return &npool.CountOrdersResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	count, err := order1.CountOrders(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountOrders", "error", err)
		return &npool.CountOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountOrdersResponse{
		Info: count,
	}, nil
}
