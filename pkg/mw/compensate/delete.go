package compensate

import (
	"context"
	"fmt"
	"time"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
)

type deleteHandler struct {
	*Handler
	OldStartAt *uint32
	OldEndAt   *uint32
}

func (h *deleteHandler) deleteCompensate(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
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

	h.OrderID = &compensate.OrderID
	h.OldStartAt = &compensate.StartAt
	h.OldEndAt = &compensate.EndAt

	if _, err := compensatecrud.UpdateSet(
		compensate.Update(),
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
	if orderstate == nil {
		return fmt.Errorf("invalid order")
	}

	oldDuration := *h.OldEndAt - *h.OldStartAt
	startAt := orderstate.StartAt
	endAt := orderstate.EndAt - oldDuration
	if endAt < startAt {
		return fmt.Errorf("invalid endat")
	}

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

func (h *Handler) DeleteCompensate(ctx context.Context) (*npool.Compensate, error) {
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
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
