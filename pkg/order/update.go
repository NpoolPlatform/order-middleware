package order

import (
	"context"
	"fmt"

	paymentcrud "github.com/NpoolPlatform/order-manager/pkg/crud/payment"

	"github.com/NpoolPlatform/order-manager/pkg/db"
	"github.com/NpoolPlatform/order-manager/pkg/db/ent"

	"github.com/NpoolPlatform/order-manager/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"

	paymentpb "github.com/NpoolPlatform/message/npool/order/mgr/v1/payment"

	"github.com/google/uuid"
)

func UpdateOrder(ctx context.Context, in *npool.OrderReq) (info *npool.Order, err error) {
	p, err := paymentcrud.Row(ctx, uuid.MustParse(in.GetPaymentID()))
	if err != nil {
		return nil, err
	}
	if p.OrderID.String() != in.GetID() {
		return nil, fmt.Errorf("invalid order")
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
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

		if info.State != paymentpb.PaymentState_Wait.String() {
			if in.GetCanceled() {
				return fmt.Errorf("not wait payment")
			}
		}

		u1, err := paymentcrud.UpdateSet(
			info.Update(),
			&paymentpb.PaymentReq{
				UserSetCanceled: in.Canceled,
				State:           in.PaymentState,
				FinishAmount:    in.PaymentFinishAmount,
				FakePayment:     in.FakePayment,
			},
		)
		if err != nil {
			return err
		}

		_, err = u1.Save(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return GetOrder(ctx, in.GetID())
}
