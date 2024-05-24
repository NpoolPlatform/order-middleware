package orderlock

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/lock"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderLockSelect
}

func (h *baseQueryHandler) selectOrderLock(stm *ent.OrderLockQuery) *ent.OrderLockSelect {
	return stm.Select(entorderlock.FieldID)
}

func (h *baseQueryHandler) queryOrderLock(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.OrderLock.Query().Where(entorderlock.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entorderlock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entorderlock.EntID(*h.EntID))
	}
	h.stmSelect = h.selectOrderLock(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderLocks(cli *ent.Client) (*ent.OrderLockSelect, error) {
	stm, err := orderlockcrud.SetQueryConds(cli.OrderLock.Query(), h.OrderLockConds)
	if err != nil {
		return nil, err
	}
	return h.selectOrderLock(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entorderlock.Table)
	s.Join(t1).On(
		s.C(entorderlock.FieldID),
		t1.C(entorderlock.FieldID),
	).AppendSelect(
		t1.C(entorderlock.FieldEntID),
		t1.C(entorderlock.FieldUserID),
		t1.C(entorderlock.FieldOrderID),
		t1.C(entorderlock.FieldLockType),
		t1.C(entorderlock.FieldCreatedAt),
		t1.C(entorderlock.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
}
