package outofgas

import (
	outofgas "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/outofgas"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	outofgas.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	outofgas.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
