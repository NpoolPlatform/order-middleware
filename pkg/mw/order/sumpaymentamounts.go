package order

import (
	"context"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/shopspring/decimal"
)

type sumPaymentAmountsHandler struct {
	*baseQueryHandler
	stmCount  *ent.OrderSelect
	stmSelect *ent.OrderSelect
	infos     []*npool.Order
	amount    *decimal.Decimal
}

func (h *sumPaymentAmountsHandler) queryJoin() error {
	var err error
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	h.stmSelect.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	return err
}

func (h *sumPaymentAmountsHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *sumPaymentAmountsHandler) formalize() {
	total := decimal.RequireFromString("0")
	for _, info := range h.infos {
		_amount, err := decimal.NewFromString(info.PaymentAmount)
		if err != nil {
			continue
		}
		currency, err := decimal.NewFromString(info.CoinUSDCurrency)
		if err != nil {
			continue
		}
		total = total.Add(_amount.Mul(currency))
	}
	h.amount = &total
}

func (h *Handler) SumOrderPaymentAmounts(ctx context.Context) (string, error) {
	handler := &sumPaymentAmountsHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmCount, err = handler.QueryOrders(cli)
		if err != nil {
			return err
		}
		handler.stmSelect, err = handler.QueryOrders(cli)
		if err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}

		handler.stmSelect.Order(ent.Desc(entorder.FieldCreatedAt))
		return handler.scan(_ctx)
	})
	if err != nil {
		return "0", err
	}
	handler.formalize()
	return handler.amount.String(), nil
}
