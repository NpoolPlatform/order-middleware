package appconfig

import (
	appconfig "github.com/NpoolPlatform/message/npool/order/mw/v1/app/config"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appconfig.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appconfig.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
