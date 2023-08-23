package payment

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	paymentcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/payment"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.PaymentSelect
	stmCount  *ent.PaymentSelect
	infos     []*npool.Payment
	total     uint32
}

func (h *queryHandler) selectPayment(stm *ent.PaymentQuery) *ent.PaymentSelect {
	return stm.Select(entpayment.FieldID)
}

func (h *queryHandler) queryPayment(cli *ent.Client) {
	h.stmSelect = h.selectPayment(
		cli.Payment.
			Query().
			Where(
				entpayment.ID(*h.ID),
				entpayment.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryPayments(cli *ent.Client) (*ent.PaymentSelect, error) {
	stm, err := paymentcrud.SetQueryConds(cli.Payment.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectPayment(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entpayment.Table)
	s.AppendSelect(
		t.C(entpayment.FieldID),
		t.C(entpayment.FieldAppID),
		t.C(entpayment.FieldGoodID),
		t.C(entpayment.FieldUserID),
		t.C(entpayment.FieldOrderID),
		t.C(entpayment.FieldAccountID),
		t.C(entpayment.FieldStartAmount),
		t.C(entpayment.FieldAmount),
		t.C(entpayment.FieldPayWithBalanceAmount),
		t.C(entpayment.FieldFinishAmount),
		t.C(entpayment.FieldCoinUsdCurrency),
		t.C(entpayment.FieldLocalCoinUsdCurrency),
		t.C(entpayment.FieldLiveCoinUsdCurrency),
		t.C(entpayment.FieldCoinInfoID),
		sql.As(t.C(entpayment.FieldStateV1), "state"),
		t.C(entpayment.FieldChainTransactionID),
		t.C(entpayment.FieldUserSetPaid),
		t.C(entpayment.FieldUserSetCanceled),
		t.C(entpayment.FieldFakePayment),
		t.C(entpayment.FieldCreatedAt),
		t.C(entpayment.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.State = basetypes.PaymentState(basetypes.PaymentState_value[info.PaymentStateStr])
	}
}

func (h *Handler) GetPayment(ctx context.Context) (*npool.Payment, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryPayment(cli)
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

func (h *Handler) GetPayments(ctx context.Context) ([]*npool.Payment, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryPayments(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryPayments(cli)
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
			Order(ent.Desc(entpayment.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
