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
		entoutofgas.FieldOrderID,
		entoutofgas.FieldStart,
		entoutofgas.FieldEnd,
		entoutofgas.FieldCreatedAt,
		entoutofgas.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOutOfGas(cli *ent.Client) {
	h.selectOutOfGas(
		cli.OutOfGas.
			Query().
			Where(
				entoutofgas.ID(*h.ID),
				entoutofgas.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryOutOfGass(ctx context.Context, cli *ent.Client) error {
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
		handler.queryOutOfGas(cli)
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

func (h *Handler) GetOutOfGass(ctx context.Context) ([]*npool.OutOfGas, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOutOfGass(_ctx, cli); err != nil {
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
