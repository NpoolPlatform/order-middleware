package outofgas

import (
	"context"
	"fmt"

	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entoutofgas "github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
)

func (h *Handler) UpdateOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	info, err := h.GetOutOfGas(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		outofgas, err := tx.OutOfGas.
			Query().
			Where(
				entoutofgas.ID(*h.ID),
			).
			ForUpdate().
			Only(_ctx)
		if err != nil {
			return err
		}
		if outofgas == nil {
			return fmt.Errorf("invalid outofgas")
		}

		if _, err := outofgascrud.UpdateSet(
			outofgas.Update(),
			&outofgascrud.Req{
				StartAt: h.StartAt,
				EndAt:   h.EndAt,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOutOfGas(ctx)
}
