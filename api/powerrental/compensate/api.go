package compensate

import (
	compensate "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/compensate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	compensate.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	compensate.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
