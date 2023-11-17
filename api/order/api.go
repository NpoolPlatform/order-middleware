package order

import (
	"context"

	"github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	order.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return order.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
