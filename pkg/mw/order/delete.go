package order

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
	deletedAt uint32
}

func (h *deleteHandler) deleteOrder(ctx context.Context, tx *ent.Tx, id uint32) error {
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

	if info.OrderState != types.OrderState_OrderStateCreated {
		return nil, fmt.Errorf("permission denied")
	}
	entID := uuid.MustParse(info.EntID)
	h.EntID = &entID

	handler := &deleteHandler{
		Handler:   h,
		deletedAt: uint32(time.Now().Unix()),
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOrder(ctx, tx, *h.ID); err != nil {
			return err
		}
		if err := handler.deleteOrderState(ctx, tx, *h.EntID); err != nil {
			return err
		}
		if info.PaymentID != uuid.Nil.String() {
			if err := handler.deletePayment(ctx, tx, *h.EntID); err != nil {
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

	ids := []uint32{}
	for _, req := range h.Reqs {
		if req.ID == nil {
			return nil, fmt.Errorf("invalid id")
		}
		ids = append(ids, *req.ID)
	}
	h.Conds = &ordercrud.Conds{
		IDs: &cruder.Cond{Op: cruder.IN, Val: ids},
	}
	infos, _, err := h.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	if len(infos) == 0 {
		return nil, nil
	}

	for _, info := range infos {
		if info.OrderState != types.OrderState_OrderStateCreated {
			return nil, fmt.Errorf("permission denied")
		}
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		for _, req := range h.Reqs {
			if err := handler.deleteOrder(ctx, tx, *req.ID); err != nil {
				return err
			}
			if err := handler.deleteOrderState(ctx, tx, *req.EntID); err != nil {
				return err
			}
			if err := handler.deletePayment(ctx, tx, *req.EntID); err != nil {
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
