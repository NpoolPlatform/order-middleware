package outofgas

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

//nolint:gocyclo
func (h *Handler) checkOutOfGas(ctx context.Context, newOutOfGas bool) (uint32, error) {
	outOfGasSeconds := uint32(0)

	if !newOutOfGas {
		info, err := h.GetOutOfGas(ctx)
		if err != nil {
			return 0, err
		}
		if info == nil {
			return 0, fmt.Errorf("invalid outofgas")
		}

		outOfGasSeconds = info.EndAt - info.StartAt

		orderID, err := uuid.Parse(info.OrderID)
		if err != nil {
			return 0, err
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
		return 0, fmt.Errorf("invalid duration")
	}

	if *h.EndAt < *h.StartAt {
		return 0, fmt.Errorf("invalid outofgas")
	}

	now := uint32(time.Now().Unix())
	if now < *h.EndAt {
		return 0, fmt.Errorf("invalid outofgas")
	}

	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		conds := &outofgascrud.Conds{
			OrderID:  &cruder.Cond{Op: cruder.EQ, Val: *h.OrderID},
			StartEnd: &cruder.Cond{Op: cruder.OVERLAP, Val: []uint32{*h.StartAt, *h.EndAt}},
		}
		if h.ID != nil {
			conds.ID = &cruder.Cond{Op: cruder.NEQ, Val: *h.ID}
		}
		stm, err := outofgascrud.SetQueryConds(cli.OutOfGas.Query(), conds)
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
	}); err != nil {
		return 0, err
	}
	return outOfGasSeconds, nil
}
