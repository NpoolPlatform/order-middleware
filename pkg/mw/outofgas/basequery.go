package outofgas

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entoutofgas "github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OutOfGasSelect
}

func (h *baseQueryHandler) selectOutOfGas(stm *ent.OutOfGasQuery) *ent.OutOfGasSelect {
	return stm.Select(entoutofgas.FieldID)
}

func (h *baseQueryHandler) queryOutOfGas(cli *ent.Client) error {
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
	h.stmSelect = h.selectOutOfGas(stm)
	return nil
}

func (h *baseQueryHandler) queryOutOfGases(cli *ent.Client) (*ent.OutOfGasSelect, error) {
	stm, err := outofgascrud.SetQueryConds(cli.OutOfGas.Query(), h.OutOfGasConds)
	if err != nil {
		return nil, err
	}
	return h.selectOutOfGas(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entoutofgas.Table)
	s.AppendSelect(
		t.C(entoutofgas.FieldEntID),
		t.C(entoutofgas.FieldOrderID),
		t.C(entoutofgas.FieldStartAt),
		t.C(entoutofgas.FieldEndAt),
		t.C(entoutofgas.FieldCreatedAt),
		t.C(entoutofgas.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinOrder(s *sql.Selector) error { //nolint
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entoutofgas.FieldOrderID),
			t.C(entorderbase.FieldEntID),
		).
		OnP(
			sql.EQ(t.C(entorderbase.FieldDeletedAt), 0),
		)
	if h.OrderBaseConds.AppID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldAppID),
				h.OrderBaseConds.AppID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.UserID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldUserID),
				h.OrderBaseConds.UserID.Val.(uuid.UUID),
			),
		)
	}
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldGoodType),
		t.C(entorderbase.FieldAppGoodID),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinOrder(s)
	})
}
