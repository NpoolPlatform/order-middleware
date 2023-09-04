package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createCompensate(ctx context.Context, tx *ent.Tx) error {
	if *h.EndAt <= *h.StartAt {
		return fmt.Errorf("invalid startend")
	}
	if _, err := compensatecrud.CreateSet(
		tx.Compensate.Create(),
		&compensatecrud.Req{
			ID:      h.ID,
			OrderID: h.OrderID,
			StartAt: h.StartAt,
			EndAt:   h.EndAt,
			Title:   h.Title,
			Message: h.Message,
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

	if *h.StartAt < orderstate.StartAt || orderstate.EndAt < *h.EndAt {
		return fmt.Errorf("invalid compensate")
	}
	endAt := orderstate.EndAt + *h.EndAt - *h.StartAt
	compensateHours := orderstate.CompensateHours + (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			EndAt:           &endAt,
			CompensateHours: &compensateHours,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCompensate(ctx context.Context) (*npool.Compensate, error) {
	if _, err := h.checkCompensate(ctx, true); err != nil {
		return nil, err
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if err := handler.createCompensate(ctx, tx); err != nil {
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

	return h.GetCompensate(ctx)
}
