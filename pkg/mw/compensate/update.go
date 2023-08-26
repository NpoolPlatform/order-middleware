package compensate

import (
	"context"
	"fmt"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
)

type updateHandler struct {
	*Handler
	OldStartAt *uint32
	OldEndAt   *uint32
}

func (h *updateHandler) updateCompensate(ctx context.Context, tx *ent.Tx) error {
	compensate, err := tx.Compensate.
		Query().
		Where(
			entcompensate.ID(*h.ID),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if compensate == nil {
		return fmt.Errorf("invalid compensate")
	}

	if h.OrderID == nil {
		h.OrderID = &compensate.OrderID
	}
	h.OldStartAt = &compensate.StartAt
	h.OldEndAt = &compensate.EndAt

	if _, err := compensatecrud.UpdateSet(
		compensate.Update(),
		&compensatecrud.Req{
			StartAt: h.StartAt,
			EndAt:   h.EndAt,
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
	if orderstate == nil {
		return fmt.Errorf("invalid order")
	}

	orderStartAt := orderstate.StartAt
	orderEndAt := orderstate.EndAt
	compStartAt := *h.OldStartAt
	compEndAt := *h.OldEndAt
	if h.StartAt != nil {
		if *h.StartAt < orderStartAt || *h.StartAt > orderEndAt {
			return fmt.Errorf("invalid startat")
		}
		compStartAt = *h.StartAt
	}
	if h.EndAt != nil {
		if *h.EndAt > orderEndAt || *h.EndAt < orderStartAt {
			return fmt.Errorf("invalid endat")
		}
		compEndAt = *h.EndAt
	}
	if compStartAt > compEndAt {
		return fmt.Errorf("invalid startat")
	}

	oldDuration := *h.OldEndAt - *h.OldStartAt
	duration := compEndAt - compStartAt
	endAt := orderEndAt - oldDuration + duration

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

func (h *Handler) UpdateCompensate(ctx context.Context) (*npool.Compensate, error) {
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &updateHandler{
		Handler: h,
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
