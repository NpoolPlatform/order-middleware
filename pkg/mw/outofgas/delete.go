package outofgas

import (
	"context"
	"fmt"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteOutOfGas(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := outofgascrud.UpdateSet(
		tx.OutOfGas.UpdateOneID(*h.ID),
		&outofgascrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) updateOrder(ctx context.Context, tx *ent.Tx) error {
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

	if orderstate.OutofgasHours < (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour {
		return fmt.Errorf("invalid outofgas")
	}
	outOfGasHours := orderstate.OutofgasHours - (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour

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

func (h *Handler) DeleteOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	info, err := h.GetOutOfGas(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if _, err := h.checkOutOfGas(ctx, false); err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteOutOfGas(ctx, tx); err != nil {
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

	return info, nil
}
