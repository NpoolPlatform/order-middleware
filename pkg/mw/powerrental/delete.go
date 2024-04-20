package powerrental

import (
	"context"
	"time"

	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	orderstatebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"
	powerrentalcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/powerrental"
	powerrentalstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/powerrental/state"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type deleteHandler struct {
	*powerRentalQueryHandler
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

func (h *deleteHandler) deletePowerRental(ctx context.Context, tx *ent.Tx) error {
	_, err := powerrentalcrud.UpdateSet(
		tx.PowerRental.UpdateOneID(h._ent.PowerRentalID()),
		&powerrentalcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *deleteHandler) deletePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	_, err := powerrentalstatecrud.UpdateSet(
		tx.PowerRentalState.UpdateOneID(h._ent.PowerRentalStateID()),
		&powerrentalstatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *Handler) DeletePowerRentalWithTx(ctx context.Context, tx *ent.Tx) error {
	handler := &deleteHandler{
		powerRentalQueryHandler: &powerRentalQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getPowerRental(ctx); err != nil {
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
	if err := handler.deletePowerRental(ctx, tx); err != nil {
		return err
	}
	return handler.deletePowerRentalState(ctx, tx)
}

func (h *Handler) DeletePowerRental(ctx context.Context) error {
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return h.DeletePowerRentalWithTx(_ctx, tx)
	})
}
