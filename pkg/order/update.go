package order

import (
	"context"
	"fmt"

	paymentcrud "github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"

	orderpb "github.com/NpoolPlatform/message/npool/cloud-hashing-order"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
)

func UpdateOrder(ctx context.Context, in *npool.OrderReq) (info *npool.Order, err error) {
	p, err := paymentcrud.Get(ctx, &orderpb.GetPaymentRequest{
		ID: in.GetPaymentID(),
	})
	if err != nil {
		return nil, err
	}
	if p.Info.OrderID != in.GetID() {
		return nil, fmt.Errorf("invalid order")
	}

	if in.Canceled != nil {
		p.Info.UserSetCanceled = in.GetCanceled()
	}
	_, err = paymentcrud.Update(ctx, &orderpb.UpdatePaymentRequest{
		Info: p.Info,
	})
	if err != nil {
		return nil, err
	}

	return GetOrder(ctx, in.GetID())
}
