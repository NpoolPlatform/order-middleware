package api

import (
	"context"

	order "github.com/NpoolPlatform/message/npool/order/mw/v1"
	"github.com/NpoolPlatform/order-middleware/api/compensate"
	order1 "github.com/NpoolPlatform/order-middleware/api/order"
	"github.com/NpoolPlatform/order-middleware/api/outofgas"
	"github.com/NpoolPlatform/order-middleware/api/payment"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	order.RegisterMiddlewareServer(server, &Server{})
	order1.Register(server)
	compensate.Register(server)
	outofgas.Register(server)
	payment.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := order.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
