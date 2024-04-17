package config

import (
	"github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	config.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	config.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
