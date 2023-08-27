package compensate

import (
	"context"
	"fmt"
	"time"

	timedef "github.com/NpoolPlatform/go-service-framework/pkg/const/time"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteCompensate(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
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

	endAt := orderstate.EndAt - (*h.EndAt - *h.StartAt)
	if endAt <= orderstate.StartAt || endAt < uint32(time.Now().Unix()) {
		return fmt.Errorf("invalid compensate")
	}
	if orderstate.CompensateHours < (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour {
		return fmt.Errorf("invalid compensate")
	}
	compensateHours := orderstate.CompensateHours - (*h.EndAt-*h.StartAt)/timedef.SecondsPerHour

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

func (h *Handler) DeleteCompensate(ctx context.Context) (*npool.Compensate, error) {
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	if _, err := h.checkCompensate(ctx, false); err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteCompensate(ctx, tx); err != nil {
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
