package compensate

import (
	"context"
	"fmt"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
)

func (h *Handler) UpdateCompensate(ctx context.Context) (*npool.Compensate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		compensate, err := tx.Compensate.
			Query().
			Where(
				entcompensate.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if compensate == nil {
			return fmt.Errorf("invalid compensate")
		}

		if _, err := compensatecrud.UpdateSet(
			compensate.Update(),
			&compensatecrud.Req{
				Start:   h.Start,
				End:     h.End,
				Message: h.Message,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCompensate(ctx)
}
