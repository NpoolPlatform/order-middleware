package outofgas

import (
	"context"
	"fmt"
	"time"

	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entoutofgas "github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
)

func (h *Handler) DeleteOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}
	info, err := h.GetOutOfGas(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	now := uint32(time.Now().Unix())

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
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}