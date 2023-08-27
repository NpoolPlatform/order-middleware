package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func (h *Handler) checkCompensate(ctx context.Context, newCompensate bool) error {
	if !newCompensate {
		info, err := h.GetCompensate(ctx)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("invalid compensate")
		}

		orderID, err := uuid.Parse(info.OrderID)
		if err != nil {
			return err
		}
		h.OrderID = &orderID

		if h.StartAt == nil {
			h.StartAt = &info.StartAt
		}
		if h.EndAt == nil {
			h.EndAt = &info.EndAt
		}
	}

	if h.StartAt == nil || h.EndAt == nil {
		return fmt.Errorf("invalid duration")
	}

	if *h.EndAt < *h.StartAt {
		return fmt.Errorf("invalid compensate")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		conds := &compensatecrud.Conds{
			OrderID:  &cruder.Cond{Op: cruder.EQ, Val: *h.OrderID},
			StartEnd: &cruder.Cond{Op: cruder.OVERLAP, Val: []uint32{*h.StartAt, *h.EndAt}},
		}
		if h.ID != nil {
			conds.ID = &cruder.Cond{Op: cruder.NEQ, Val: *h.ID}
		}
		stm, err := compensatecrud.SetQueryConds(cli.Compensate.Query(), conds)
		if err != nil {
			return err
		}
		exist, err := stm.Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("time overlap")
		}
		return nil
	})
}
