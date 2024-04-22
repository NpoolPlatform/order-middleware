package api

import (
	"context"

	order "github.com/NpoolPlatform/message/npool/order/mw/v1"
	appconfig "github.com/NpoolPlatform/order-middleware/api/app/config"
	"github.com/NpoolPlatform/order-middleware/api/compensate"
	feeorder1 "github.com/NpoolPlatform/order-middleware/api/fee"
	order1 "github.com/NpoolPlatform/order-middleware/api/order"
	"github.com/NpoolPlatform/order-middleware/api/outofgas"
	powerrental1 "github.com/NpoolPlatform/order-middleware/api/powerrental"
	powerrentalcompensate1 "github.com/NpoolPlatform/order-middleware/api/powerrental/compensate"
	powerrentaloutofgas1 "github.com/NpoolPlatform/order-middleware/api/powerrental/outofgas"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	order.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	order.RegisterMiddlewareServer(server, &Server{})
	order1.Register(server)
	feeorder1.Register(server)
	powerrental1.Register(server)
	powerrentalcompensate1.Register(server)
	powerrentaloutofgas1.Register(server)
	compensate.Register(server)
	outofgas.Register(server)
	appconfig.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := order.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := compensate.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := outofgas.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := order1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := feeorder1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrental1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrentalcompensate1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := powerrentaloutofgas1.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appconfig.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
