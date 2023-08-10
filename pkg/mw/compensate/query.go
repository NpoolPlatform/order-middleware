package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
)

type queryHandler struct {
	*Handler
	stm   *ent.CompensateSelect
	infos []*npool.Compensate
	total uint32
}

func (h *queryHandler) selectCompensate(stm *ent.CompensateQuery) {
	h.stm = stm.Select(
		entcompensate.FieldID,
		entcompensate.FieldOrderID,
		entcompensate.FieldStart,
		entcompensate.FieldEnd,
		entcompensate.FieldMessage,
		entcompensate.FieldCreatedAt,
		entcompensate.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCompensate(cli *ent.Client) {
	h.selectCompensate(
		cli.Compensate.
			Query().
			Where(
				entcompensate.ID(*h.ID),
				entcompensate.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryCompensates(ctx context.Context, cli *ent.Client) error {
	stm, err := compensatecrud.SetQueryConds(cli.Compensate.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectCompensate(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetCompensate(ctx context.Context) (*npool.Compensate, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryCompensate(cli)
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

func (h *Handler) GetCompensates(ctx context.Context) ([]*npool.Compensate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCompensates(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entcompensate.FieldCreatedAt))
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
