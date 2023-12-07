package outofgas

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entoutofgas "github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
)

type queryHandler struct {
	*Handler
	stm   *ent.OutOfGasSelect
	infos []*npool.OutOfGas
	total uint32
}

func (h *queryHandler) selectOutOfGas(stm *ent.OutOfGasQuery) {
	h.stm = stm.Select(
		entoutofgas.FieldID,
		entoutofgas.FieldEntID,
		entoutofgas.FieldOrderID,
		entoutofgas.FieldStartAt,
		entoutofgas.FieldEndAt,
		entoutofgas.FieldCreatedAt,
		entoutofgas.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOutOfGas(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.OutOfGas.Query().Where(entoutofgas.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entoutofgas.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entoutofgas.EntID(*h.EntID))
	}
	h.selectOutOfGas(stm)
	return nil
}

func (h *queryHandler) queryOutOfGases(ctx context.Context, cli *ent.Client) error {
	stm, err := outofgascrud.SetQueryConds(cli.OutOfGas.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectOutOfGas(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOutOfGas(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetOutOfGases(ctx context.Context) ([]*npool.OutOfGas, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOutOfGases(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entoutofgas.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
