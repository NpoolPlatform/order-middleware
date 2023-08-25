package order

import (
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
)

type OrderReq struct {
	OrderReq      *ordercrud.Req
	OrderStateReq *orderstatecrud.Req
	PaymentReq    *paymentcrud.Req
}
