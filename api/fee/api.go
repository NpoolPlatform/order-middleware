package feeorder

import (
	feeorder "github.com/NpoolPlatform/message/npool/order/mw/v1/fee"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	feeorder.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	feeorder.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
