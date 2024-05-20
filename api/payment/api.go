package payment

import (
	payment "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	payment.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	payment.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
