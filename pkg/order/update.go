package order

import (
	"context"
	"fmt"

	paymentcrud "github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/order-middleware/pkg/db"

	orderpb "github.com/NpoolPlatform/message/npool/cloud-hashing-order"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	"github.com/google/uuid"
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

	err = db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		_, err = cli.
			Payment.
			UpdateOneID(uuid.MustParse(in.GetPaymentID())).
			SetUserSetCanceled(in.GetCanceled()).
			Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return GetOrder(ctx, in.GetID())
}
