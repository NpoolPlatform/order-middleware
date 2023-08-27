package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	orderstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	order1 "github.com/NpoolPlatform/order-middleware/pkg/mw/order"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
)

type updateHandler struct {
	*Handler
	OldStartAt           *uint32
	OldEndAt             *uint32
	finalEndAt           uint32
	finalCompensateHours uint32
}

//nolint:gocyclo
func (h *updateHandler) validateCompensate(ctx context.Context) error {
	if h.StartAt == nil && h.EndAt == nil {
		return nil
	}

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
		if compensates[key].ID == h.ID.String() {
			continue
		}
		if h.StartAt != nil && compensates[key].StartAt < *h.StartAt && *h.StartAt < compensates[key].EndAt {
			return fmt.Errorf("invalid startat")
		}
		if h.EndAt != nil && compensates[key].StartAt < *h.EndAt && *h.EndAt < compensates[key].EndAt {
			return fmt.Errorf("invalid endat")
		}
	}

	orderID := h.OrderID.String()
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

	orderStartAt := order.StartAt
	orderEndAt := order.EndAt
	compStartAt := *h.OldStartAt
	compEndAt := *h.OldEndAt
	if h.StartAt != nil {
		if *h.StartAt < orderStartAt || orderEndAt < *h.StartAt {
			return fmt.Errorf("invalid startat")
		}
		compStartAt = *h.StartAt
	}
	if h.EndAt != nil {
		if *h.EndAt < orderStartAt || orderEndAt < *h.EndAt {
			return fmt.Errorf("invalid endat")
		}
		compEndAt = *h.EndAt
	}
	if compEndAt < compStartAt {
		return fmt.Errorf("invalid startat")
	}

	oldDuration := *h.OldEndAt - *h.OldStartAt
	duration := compEndAt - compStartAt
	h.finalEndAt = orderEndAt - oldDuration + duration
	h.finalCompensateHours = order.CompensateHours - oldDuration + duration

	return nil
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
	if h.StartAt == nil && h.EndAt == nil {
		return nil
	}
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

func (h *Handler) UpdateCompensate(ctx context.Context) (*npool.Compensate, error) {
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	handler := &updateHandler{
		Handler:    h,
		OldStartAt: &info.StartAt,
		OldEndAt:   &info.EndAt,
	}
	if h.OrderID == nil {
		_id, err := uuid.Parse(info.OrderID)
		if err != nil {
			return nil, err
		}
		h.OrderID = &_id
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.validateCompensate(ctx); err != nil {
			return err
		}
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
