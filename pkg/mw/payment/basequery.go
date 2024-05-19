package payment

import (
	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	paymentbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpaymentbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbase"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.PaymentBaseSelect
}

func (h *baseQueryHandler) selectPaymentBase(stm *ent.PaymentBaseQuery) *ent.PaymentBaseSelect {
	return stm.Select(entpaymentbase.FieldID)
}

func (h *baseQueryHandler) queryPaymentBase(cli *ent.Client) error {
	if h.EntID == nil || h.ID == nil {
		return wlog.Errorf("invalid entid")
	}
	stm := cli.PaymentBase.
		Query().
		Where(
			entpaymentbase.DeletedAt(0),
		)
	if h.ID != nil {
		stm.Where(entpaymentbase.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entpaymentbase.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPaymentBase(stm)
	return nil
}

func (h *baseQueryHandler) queryPaymentBases(cli *ent.Client) (*ent.PaymentBaseSelect, error) {
	stm, err := paymentbasecrud.SetQueryConds(cli.PaymentBase.Query(), h.PaymentBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return h.selectPaymentBase(stm), nil
}

//nolint:gocyclo,funlen
func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entpaymentbase.Table)
	s.Join(t).
		On(
			s.C(entpaymentbase.FieldID),
			t.C(entpaymentbase.FieldID),
		).
		AppendSelect(
			t.C(entpaymentbase.FieldEntID),
			t.C(entpaymentbase.FieldOrderID),
			t.C(entpaymentbase.FieldObseleteState),
			t.C(entpaymentbase.FieldCreatedAt),
			t.C(entpaymentbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		// TODO: join order lock
	})
}
