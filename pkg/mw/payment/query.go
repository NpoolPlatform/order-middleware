package payment

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/payment"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.PaymentBaseSelect
	infos    []*npool.Payment
	total    uint32
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {

}

func (h *Handler) GetPayments(ctx context.Context) (infos []*npool.Payment, total uint32, err error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryPaymentBases(cli); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stmCount, err = handler.queryPaymentBases(cli); err != nil {
			return wlog.WrapError(err)
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		// TODO: get payment transfers
		// TODO: get payment balances
		return nil
	}); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()
	return handler.infos, handler.total, nil
}
