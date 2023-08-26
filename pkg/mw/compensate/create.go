package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

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
	if _, err := compensatecrud.CreateSet(
		tx.Compensate.Create(),
		&compensatecrud.Req{
			ID:      h.ID,
			OrderID: h.OrderID,
			StartAt: h.StartAt,
			EndAt:   h.EndAt,
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
	if orderstate == nil {
		return fmt.Errorf("invalid order")
	}

	orderStartAt := orderstate.StartAt
	orderEndAt := orderstate.EndAt
	if *h.StartAt > *h.EndAt || *h.StartAt < orderStartAt || *h.StartAt > orderEndAt || *h.EndAt > orderEndAt {
		return fmt.Errorf("invalid startat")
	}
	duration := *h.EndAt - *h.StartAt
	endAt := orderEndAt + duration

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			EndAt: &endAt,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateCompensate(ctx context.Context) (*npool.Compensate, error) {
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
