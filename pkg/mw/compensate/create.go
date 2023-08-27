package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	finalEndAt           uint32
	finalCompensateHours uint32
}

func (h *createHandler) validateCompensate(ctx context.Context) error {
	if *h.EndAt < *h.StartAt {
		return fmt.Errorf("invalid startat")
	}
	orderID := h.OrderID.String()
	handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID.String()},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if err != nil {
		return err
	}
	compensates, _, err := handler.GetCompensates(ctx)
	if err != nil {
		return nil
	}
	for key := range compensates {
		if compensates[key].StartAt < *h.StartAt && *h.StartAt < compensates[key].EndAt {
			return fmt.Errorf("invalid startat")
		}
		if compensates[key].StartAt < *h.EndAt && *h.EndAt < compensates[key].EndAt {
			return fmt.Errorf("invalid endat")
		}
	}

	orderHandler, err := order1.NewHandler(
		ctx,
		order1.WithID(&orderID, true),
	)
	if err != nil {
		return err
	}
	order, err := orderHandler.GetOrder(ctx)
	if err != nil {
		return err
	}
	if order == nil {
		return fmt.Errorf("invalid order")
	}

	compensateHours := order.CompensateHours
	orderStartAt := order.StartAt
	orderEndAt := order.EndAt
	if *h.StartAt < orderStartAt || orderEndAt < *h.EndAt {
		return fmt.Errorf("invalid startat")
	}
	duration := *h.EndAt - *h.StartAt
	h.finalEndAt = orderEndAt + duration
	h.finalCompensateHours = compensateHours + duration

	return nil
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

//nolint:dupl
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

	if _, err := orderstatecrud.UpdateSet(
		orderstate.Update(),
		&orderstatecrud.Req{
			EndAt:           &h.finalEndAt,
			CompensateHours: &h.finalCompensateHours,
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
		if err := handler.validateCompensate(ctx); err != nil {
			return err
		}
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
