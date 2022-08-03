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
)

func (s *Server) CreateOrder(ctx context.Context, in *npool.CreateOrderRequest) (*npool.CreateOrderResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateOrder")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		return &npool.CreateOrderResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "order", "middleware", "Create")

	info, err := order1.CreateOrder(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateOrder", "error", err)
		return &npool.CreateOrderResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateOrderResponse{
		Info: info,
	}, nil
}
