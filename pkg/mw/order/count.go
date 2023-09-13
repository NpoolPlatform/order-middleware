package order

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type countHandler struct {
	*baseQueryHandler
	stmCount *ent.OrderSelect
}

func (h *countHandler) queryJoin() error {
	var err error
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})

	return err
}

func (h *Handler) CountOrders(ctx context.Context) (uint32, error) {
	count := uint32(0)
	handler := &countHandler{
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

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		count = uint32(_total)
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, err
}
