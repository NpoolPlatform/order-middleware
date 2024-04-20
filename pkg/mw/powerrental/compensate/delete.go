package compensate

import (
	"context"
	"time"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/compensate"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*powerRentalStateQueryHandler
	now uint32
}

func (h *deleteHandler) deleteCompensate(ctx context.Context, tx *ent.Tx) error {
	_, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx)
	return err
}

func (h *deleteHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.
		PowerRentalState.
		UpdateOneID(h._ent.PowerRentalStateID()).
		AddCompensateSeconds(0 - int32(*h.CompensateSeconds)).
		Save(ctx)
	return err
}

func (h *Handler) DeleteCompensate(ctx context.Context) error {
	h1, err := compensate1.NewHandler(
		ctx,
		compensate1.WithID(h.ID, false),
		compensate1.WithEntID(func() *string {
			if h.EntID == nil {
				return nil
			}
			s := h.EntID.String()
			return &s
		}(), false),
		compensate1.WithOrderID(func() *string {
			if h.OrderID == nil {
				return nil
			}
			s := h.OrderID.String()
			return &s
		}(), false),
	)
	if err != nil {
		return err
	}

	info, err := h1.GetCompensate(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	handler := &deleteHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	h.ID = &info.ID
	h.EntID = func() *uuid.UUID { uid := uuid.MustParse(info.EntID); return &uid }()
	h.OrderID = func() *uuid.UUID { uid := uuid.MustParse(info.OrderID); return &uid }()
	h.CompensateSeconds = &info.CompensateSeconds

	if err := handler.requirePowerRentalStates(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteCompensate(_ctx, tx); err != nil {
			return err
		}
		return handler.updatePowerRentalState(_ctx, tx)
	})
}
