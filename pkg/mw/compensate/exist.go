package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
)

type existHandler struct {
	*Handler
	stm *ent.CompensateQuery
}

func (h *existHandler) queryCompensate(cli *ent.Client) {
	h.stm = cli.Compensate.
		Query().
		Where(
			entcompensate.ID(*h.ID),
			entcompensate.DeletedAt(0),
		)
}

func (h *existHandler) queryCompensates(cli *ent.Client) error {
	stm, err := compensatecrud.SetQueryConds(cli.Compensate.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *Handler) ExistCompensate(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCompensate(cli)
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

func (h *Handler) ExistCompensateConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCompensates(cli); err != nil {
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
