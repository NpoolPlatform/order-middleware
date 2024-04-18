package compensate

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CompensateSelect
	stmCount  *ent.CompensateSelect
	infos     []*npool.Compensate
	total     uint32
}

func (h *queryHandler) selectCompensate(stm *ent.CompensateQuery) *ent.CompensateSelect {
	return stm.Select(entcompensate.FieldID)
}

func (h *queryHandler) queryCompensate(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Compensate.Query().Where(entcompensate.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcompensate.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcompensate.EntID(*h.EntID))
	}
	h.stmSelect = h.selectCompensate(stm)
	return nil
}

func (h *queryHandler) queryCompensates(cli *ent.Client) (*ent.CompensateSelect, error) {
	stm, err := compensatecrud.SetQueryConds(cli.Compensate.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectCompensate(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcompensate.Table)
	s.AppendSelect(
		t.C(entcompensate.FieldID),
		t.C(entcompensate.FieldEntID),
		t.C(entcompensate.FieldOrderID),
		t.C(entcompensate.FieldStartAt),
		t.C(entcompensate.FieldEndAt),
		t.C(entcompensate.FieldTitle),
		t.C(entcompensate.FieldMessage),
		t.C(entcompensate.FieldCreatedAt),
		t.C(entcompensate.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoinOrder(s *sql.Selector) error { //nolint
	t := sql.Table(entorder.Table)
	s.LeftJoin(t).
		On(
			s.C(entcompensate.FieldOrderID),
			t.C(entorder.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entorder.FieldDeletedAt), 0),
		)

	s.AppendSelect(
		sql.As(t.C(entorder.FieldAppID), "app_id"),
		sql.As(t.C(entorder.FieldUserID), "user_id"),
		sql.As(t.C(entorder.FieldGoodID), "good_id"),
		sql.As(t.C(entorder.FieldAppGoodID), "app_good_id"),
		sql.As(t.C(entorder.FieldUnitsV1), "units_v1"),
	)
	return nil
}

func (h *queryHandler) queryJoinOrderState(s *sql.Selector) error { //nolint
	t := sql.Table(entorderstate.Table)
	s.LeftJoin(t).
		On(
			s.C(entcompensate.FieldOrderID),
			t.C(entorderstate.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entorderstate.FieldDeletedAt), 0),
		)

	s.AppendSelect(
		sql.As(t.C(entorderstate.FieldStartAt), "order_start_at"),
		sql.As(t.C(entorderstate.FieldEndAt), "order_end_at"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinOrder(s)
		err = h.queryJoinOrderState(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinOrder(s)
		err = h.queryJoinOrderState(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		amount, err := decimal.NewFromString(info.Units)
		if err != nil {
			info.Units = decimal.NewFromInt(0).String()
		} else {
			info.Units = amount.String()
		}
	}
}

func (h *Handler) GetCompensate(ctx context.Context) (*npool.Compensate, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCompensate(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
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

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetCompensates(ctx context.Context) ([]*npool.Compensate, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryCompensates(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryCompensates(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entcompensate.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
