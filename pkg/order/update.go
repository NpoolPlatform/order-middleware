package order

import (
	"context"
	"fmt"

	paymentcrud "github.com/NpoolPlatform/cloud-hashing-order/pkg/crud/payment"

	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent"
	"github.com/NpoolPlatform/cloud-hashing-order/pkg/db/ent/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"

	orderconst "github.com/NpoolPlatform/cloud-hashing-order/pkg/const"

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

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		info, err := tx.
			Payment.
			Query().
			Where(
				payment.ID(uuid.MustParse(in.GetPaymentID())),
			).
			ForUpdate().
			Only(ctx)
		if err != nil {
			return err
		}

		if info.State != orderconst.PaymentStateWait {
			return fmt.Errorf("not wait payment")
		}

		_, err = info.
			Update().
			SetUserSetCanceled(in.GetCanceled()).
			Save(ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return GetOrder(ctx, in.GetID())
}
