package order

import (
	"context"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
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
	deletedAt uint32
}

//nolint:dupl
func (h *deleteHandler) deleteOrder(ctx context.Context, tx *ent.Tx, id uuid.UUID) error {
	order, err := tx.Order.
		Query().
		Where(
			entorder.ID(id),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err := ordercrud.UpdateSet(
		order.Update(),
		&ordercrud.Req{
			DeletedAt: &h.deletedAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deleteOrderState(ctx context.Context, tx *ent.Tx, orderID uuid.UUID) error {
	orderstate, err := tx.OrderState.
		Query().
		Where(
			entorderstate.OrderID(orderID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			DeletedAt: &h.deletedAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deletePayment(ctx context.Context, tx *ent.Tx, orderID uuid.UUID) error {
	payment, err := tx.Payment.
		Query().
		Where(
			entpayment.OrderID(orderID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil
		}
		return err
	}

	if _, err := paymentcrud.UpdateSet(
		payment.Update(),
		&paymentcrud.Req{
			DeletedAt: &h.deletedAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteOrder(ctx context.Context) (*npool.Order, error) {
	info, err := h.GetOrder(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &deleteHandler{
		Handler:   h,
		deletedAt: uint32(time.Now().Unix()),
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOrder(ctx, tx, *h.ID); err != nil {
			return err
		}
		if err := handler.deleteOrderState(ctx, tx, *h.ID); err != nil {
			return err
		}
		if info.PaymentID != uuid.Nil.String() {
			if err := handler.deletePayment(ctx, tx, *h.ID); err != nil {
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

func (h *Handler) DeleteOrders(ctx context.Context) ([]*npool.Order, error) {
	handler := &deleteHandler{
		Handler:   h,
		deletedAt: uint32(time.Now().Unix()),
	}

	ids := []uuid.UUID{}
	for _, req := range h.Reqs {
		ids = append(ids, *req.ID)
	}
	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	infos, _, err := h.GetOrders(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.deleteOrder(ctx, tx, *req.ID); err != nil {
				return err
			}
			if err := handler.deleteOrderState(ctx, tx, *req.ID); err != nil {
				return err
			}
			if err := handler.deletePayment(ctx, tx, *req.ID); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return infos, nil
}
