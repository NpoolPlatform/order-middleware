package order

import (
	"context"
	"fmt"
	"time"

	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/google/uuid"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
)

type deleteHandler struct {
	*Handler
	deleteAt *uint32
}

//nolint:dupl
func (h *deleteHandler) deleteOrder(ctx context.Context, tx *ent.Tx, req *ordercrud.Req) error {
	order, err := tx.Order.
		Query().
		Where(
			entorder.ID(*req.ID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("invalid order")
	}

	if _, err := ordercrud.UpdateSet(
		order.Update(),
		&ordercrud.Req{
			DeletedAt: h.deleteAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deleteOrderState(ctx context.Context, tx *ent.Tx, req *orderstatecrud.Req) error {
	orderstate, err := tx.OrderState.
		Query().
		Where(
			entorderstate.OrderID(*req.OrderID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if orderstate == nil {
		return fmt.Errorf("invalid orderstate")
	}

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			DeletedAt: h.deleteAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deletePayment(ctx context.Context, tx *ent.Tx, req *paymentcrud.Req) error {
	payment, err := tx.Payment.
		Query().
		Where(
			entpayment.OrderID(*req.OrderID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if payment == nil {
		return fmt.Errorf("invalid payment")
	}

	if _, err := paymentcrud.UpdateSet(
		payment.Update(),
		&paymentcrud.Req{
			DeletedAt: h.deleteAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

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
	handler := &deleteHandler{
		Handler: h,
	}

	orderReq := &ordercrud.Req{
		ID: h.ID,
	}
	orderstateReq := &orderstatecrud.Req{
		OrderID: h.ID,
	}
	paymentReq := &paymentcrud.Req{
		OrderID: h.ID,
	}

	now := uint32(time.Now().Unix())
	handler.deleteAt = &now
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOrder(ctx, tx, orderReq); err != nil {
			return err
		}
		if err := handler.deleteOrderState(ctx, tx, orderstateReq); err != nil {
			return err
		}
		if info.PaymentID != uuid.Nil.String() {
			if err := handler.deletePayment(ctx, tx, paymentReq); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
