package compensate

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entcompensate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.CompensateSelect
}

func (h *baseQueryHandler) selectCompensate(stm *ent.CompensateQuery) *ent.CompensateSelect {
	return stm.Select(entcompensate.FieldID)
}

func (h *baseQueryHandler) queryCompensate(cli *ent.Client) error {
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

func (h *baseQueryHandler) queryCompensates(cli *ent.Client) (*ent.CompensateSelect, error) {
	stm, err := compensatecrud.SetQueryConds(cli.Compensate.Query(), h.CompensateConds)
	if err != nil {
		return nil, err
	}
	return h.selectCompensate(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcompensate.Table)
	s.AppendSelect(
		t.C(entcompensate.FieldID),
		t.C(entcompensate.FieldEntID),
		t.C(entcompensate.FieldOrderID),
		t.C(entcompensate.FieldCompensateFromID),
		t.C(entcompensate.FieldCompensateType),
		t.C(entcompensate.FieldCompensateSeconds),
		t.C(entcompensate.FieldCreatedAt),
		t.C(entcompensate.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinOrder(s *sql.Selector) error { //nolint
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entcompensate.FieldOrderID),
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
		if err := h.queryJoinOrder(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrder", "Error", err)
		}
	})
}
