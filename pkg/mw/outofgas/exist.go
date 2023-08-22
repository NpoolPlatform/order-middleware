package outofgas

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	entoutofgas "github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
)

type existHandler struct {
	*Handler
	stm *ent.OutOfGasQuery
}

func (h *existHandler) queryOutOfGas(cli *ent.Client) {
	h.stm = cli.OutOfGas.
		Query().
		Where(
			entoutofgas.ID(*h.ID),
			entoutofgas.DeletedAt(0),
		)
}

func (h *existHandler) queryOutOfGass(cli *ent.Client) error {
	stm, err := outofgascrud.SetQueryConds(cli.OutOfGas.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistOutOfGas(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOutOfGas(cli)
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (h *Handler) ExistOutOfGasConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOutOfGass(cli); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
