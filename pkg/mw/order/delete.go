package order

import (
	"context"
	"fmt"
	"time"

	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
)

func (h *Handler) DeleteOrder(ctx context.Context) (*npool.Order, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetOrder(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		order, err := tx.Order.
			Query().
			Where(
				entorder.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if order == nil {
			return fmt.Errorf("invalid order")
		}

		payment, err := tx.Payment.
			Query().
			Where(
				entpayment.AppID(order.AppID),
				entpayment.UserID(order.UserID),
				entpayment.GoodID(order.GoodID),
				entpayment.OrderID(order.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if payment == nil {
			return fmt.Errorf("invalid payment")
		}

		if _, err := ordercrud.UpdateSet(
			order.Update(),
			&ordercrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}

		if _, err := paymentcrud.UpdateSet(
			payment.Update(),
			&paymentcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
