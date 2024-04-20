package feeorder

import (
	"context"
	"time"

	feeordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/fee"
	feeorderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/fee/state"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*feeOrderQueryHandler
	now uint32
}

func (h *deleteHandler) deleteOrderBase(ctx context.Context, tx *ent.Tx) error {
	_, err := orderbasecrud.UpdateSet(
		tx.OrderBase.UpdateOneID(h._ent.OrderBaseID()),
		&orderbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *deleteHandler) deleteOrderStateBase(ctx context.Context, tx *ent.Tx) error {
	_, err := orderstatebasecrud.UpdateSet(
		tx.OrderStateBase.UpdateOneID(h._ent.OrderStateBaseID()),
		&orderstatebasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *deleteHandler) deleteFeeOrder(ctx context.Context, tx *ent.Tx) error {
	_, err := feeordercrud.UpdateSet(
		tx.FeeOrder.UpdateOneID(h._ent.FeeOrderID()),
		&feeordercrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *deleteHandler) deleteFeeOrderState(ctx context.Context, tx *ent.Tx) error {
	_, err := feeorderstatecrud.UpdateSet(
		tx.FeeOrderState.UpdateOneID(h._ent.FeeOrderStateID()),
		&feeorderstatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *Handler) DeleteFeeOrderWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		feeOrderQueryHandler: &feeOrderQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getFeeOrder(ctx); err != nil {
		return err
	}
	if !handler._ent.Exist() {
		return nil
	}

	if err := handler.deleteOrderBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.deleteOrderStateBase(ctx, tx); err != nil {
		return err
	}
	if err := handler.deleteFeeOrder(ctx, tx); err != nil {
		return err
	}
	return handler.deleteFeeOrderState(ctx, tx)
}

func (h *Handler) DeleteFeeOrder(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeleteFeeOrderWithTx(_ctx, tx)
	})
}
