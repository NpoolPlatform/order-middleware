package outofgas

import (
	"context"
	"fmt"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createOutOfGas(ctx context.Context, tx *ent.Tx) error {
	if *h.EndAt <= *h.StartAt {
		return fmt.Errorf("invalid startend")
	}

	if _, err := outofgascrud.CreateSet(
		tx.OutOfGas.Create(),
		&outofgascrud.Req{
			EntID:   h.EntID,
			OrderID: h.OrderID,
			StartAt: h.StartAt,
			EndAt:   h.EndAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateOrder(ctx context.Context, tx *ent.Tx) error {
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

func (h *Handler) CreateOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	if _, err := h.checkOutOfGas(ctx, true); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOutOfGas(ctx, tx); err != nil {
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
