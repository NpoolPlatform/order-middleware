package outofgas

import (
	"context"
	"fmt"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	orderstatebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
)

type updateHandler struct {
	*Handler
	outOfGasSeconds uint32
}

func (h *updateHandler) updateOutOfGas(ctx context.Context, tx *ent.Tx) error {
	if _, err := outofgascrud.UpdateSet(
		tx.OutOfGas.UpdateOneID(*h.ID),
		&outofgascrud.Req{
			StartAt: h.StartAt,
			EndAt:   h.EndAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateOrder(ctx context.Context, tx *ent.Tx) error {
	orderstate, err := tx.OrderState.
		Query().
		Where(
			entorderstate.OrderID(*h.OrderID),
			entorderstate.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if orderstate.OrderState != types.OrderState_OrderStateInService.String() {
		return fmt.Errorf("permission denied")
	}
	if *h.StartAt < orderstate.StartAt || orderstate.EndAt < *h.EndAt {
		return fmt.Errorf("invalid outofgas")
	}
	outOfGasHours := orderstate.OutofgasHours + (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour
	if outOfGasHours < h.outOfGasSeconds/timedef.SecondsPerHour {
		return fmt.Errorf("invalid outofgas")
	}
	outOfGasHours -= h.outOfGasSeconds / timedef.SecondsPerHour

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			OutOfGasHours: &outOfGasHours,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	seconds, err := h.checkOutOfGas(ctx, false)
	if err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler:         h,
		outOfGasSeconds: seconds,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateOutOfGas(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateOrder(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOutOfGas(ctx)
}
