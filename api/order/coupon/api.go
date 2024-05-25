package ordercoupon

import (
	ordercoupon "github.com/NpoolPlatform/message/npool/order/mw/v1/order/coupon"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	ordercoupon.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	ordercoupon.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
