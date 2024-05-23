package order

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/shopspring/decimal"
)

type sumHandler struct {
	*Handler
	sql string
}

func (h *sumHandler) constructSumOrderPaymentUSDSQL() {
	h.sql = `
	  select
	    sum(payment_amount_usd_total)
	  from (
	    select
	      sum(payment_amount_usd)
	    as
	      payment_amount_usd_total
	    from
	      power_rentals
	    union all
	    select
	      sum(payment_amount_usd)
	    as
	      payment_amount_usd_total
	    from
	      fee_orders
	  ) as tmp
	`
	// TODO: if we have other order type, just append
}

func (h *Handler) SumOrdersPaymentUSD(ctx context.Context) (amount string, err error) {
	handler := &sumHandler{
		Handler: h,
	}
	handler.constructSumOrderPaymentUSDSQL()
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		rows, err := cli.QueryContext(_ctx, handler.sql)
		if err != nil {
			return wlog.WrapError(err)
		}
		return wlog.WrapError(rows.Scan(&amount))
	}); err != nil {
		return decimal.NewFromInt(0).String(), wlog.WrapError(err)
	}
	return amount, wlog.WrapError(err)
}

func (h *Handler) SubOrdersValueUSD(ctx context.Context) (string, error) {
	return "", nil
}
