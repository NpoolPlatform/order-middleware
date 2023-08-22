package payment

import (
	"context"
	"fmt"

	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
)

func (h *Handler) UpdatePayment(ctx context.Context) (*npool.Payment, error) {
	info, err := h.GetPayment(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		payment, err := tx.Payment.
			Query().
			Where(
				entpayment.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if payment == nil {
			return fmt.Errorf("invalid payment")
		}

		if _, err := paymentcrud.UpdateSet(
			payment.Update(),
			&paymentcrud.Req{
				State:           h.State,
				UserSetPaid:     h.UserSetPaid,
				UserSetCanceled: h.UserSetCanceled,
				FakePayment:     h.FakePayment,
				FinishAmount:    h.FinishAmount,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetPayment(ctx)
}
