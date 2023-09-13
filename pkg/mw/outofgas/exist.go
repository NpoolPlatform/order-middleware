package outofgas

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	entoutofgas "github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
)

func (h *Handler) ExistOutOfGas(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			OutOfGas.
			Query().
			Where(
				entoutofgas.ID(*h.ID),
				entoutofgas.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistOutOfGasConds(ctx context.Context) (bool, error) {
	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := outofgascrud.SetQueryConds(cli.OutOfGas.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
