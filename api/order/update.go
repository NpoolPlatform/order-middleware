//nolint:nolintlint,dupl
package order

import (
	"context"

	order1 "github.com/NpoolPlatform/order-middleware/pkg/order"
	commontracer "github.com/NpoolPlatform/order-middleware/pkg/tracer"
	tracer "github.com/NpoolPlatform/order-middleware/pkg/tracer/order"

	constant "github.com/NpoolPlatform/order-middleware/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
)

func (s *Server) UpdateOrder(ctx context.Context, in *npool.UpdateOrderRequest) (*npool.UpdateOrderResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateOrder")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateOrder", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateOrderResponse{}, err
	}
	if _, err := uuid.Parse(in.GetInfo().GetPaymentID()); err != nil {
		logger.Sugar().Errorw("UpdateOrder", "PaymentID", in.GetInfo().GetPaymentID(), "error", err)
		return &npool.UpdateOrderResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "order", "middleware", "Update")

	info, err := order1.UpdateOrder(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateOrder", "error", err)
		return &npool.UpdateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrderResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateOrders(ctx context.Context, in *npool.UpdateOrdersRequest) (*npool.UpdateOrdersResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateOrders")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.UpdateOrdersResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	for _, info := range in.GetInfos() {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			logger.Sugar().Errorw("UpdateOrders", "ID", info.GetID(), "error", err)
			return &npool.UpdateOrdersResponse{}, err
		}
		if _, err := uuid.Parse(info.GetPaymentID()); err != nil {
			logger.Sugar().Errorw("UpdateOrders", "PaymentID", info.GetPaymentID(), "error", err)
			return &npool.UpdateOrdersResponse{}, err
		}
	}

	span = commontracer.TraceInvoker(span, "order", "middleware", "Update")

	infos, err := order1.UpdateOrders(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("UpdateOrders", "error", err)
		return &npool.UpdateOrdersResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateOrdersResponse{
		Infos: infos,
	}, nil
}
