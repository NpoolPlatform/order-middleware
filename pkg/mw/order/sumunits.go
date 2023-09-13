package order

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/shopspring/decimal"
)

type sumUnitsHandler struct {
	*baseQueryHandler
	stmCount *ent.OrderSelect
	stmSum   *ent.OrderSelect
}

func (h *sumUnitsHandler) queryJoin() error {
	var err error
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	h.stmSum.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	return err
}

func (h *Handler) SumOrderUnits(ctx context.Context) (string, error) {
	sum := decimal.NewFromInt(0).String()
	handler := &sumUnitsHandler{
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
		handler.stmSum, err = handler.QueryOrders(cli)
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
		if _total == 0 {
			return nil
		}
		_sum, err := handler.stmSum.Modify(func(s *sql.Selector) {
			s.Select(sql.Sum(entorder.FieldUnitsV1))
		}).String(_ctx)
		if err != nil {
			return err
		}
		sum = _sum

		return nil
	})
	if err != nil {
		return sum, err
	}
	return sum, nil
}
