package compensate

import (
	"context"
	"fmt"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatebasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/statebase"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
)

type updateHandler struct {
	*Handler
	compensateSeconds uint32
}

func (h *updateHandler) updateCompensate(ctx context.Context, tx *ent.Tx) error {
	if _, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
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
		return fmt.Errorf("invalid compensate")
	}
	endAt := orderstate.EndAt + *h.EndAt - *h.StartAt - h.compensateSeconds
	compensateHours := orderstate.CompensateHours + (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour
	if compensateHours < h.compensateSeconds/timedef.SecondsPerHour {
		return fmt.Errorf("invalid compensate")
	}
	compensateHours -= h.compensateSeconds / timedef.SecondsPerHour

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

func (h *Handler) UpdateCompensate(ctx context.Context) (*npool.Compensate, error) {
	seconds, err := h.checkCompensate(ctx, false)
	if err != nil {
		return nil, err
	}

	handler := &updateHandler{
		Handler:           h,
		compensateSeconds: seconds,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateCompensate(ctx, tx); err != nil {
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
