package order

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
)

type existHandler struct {
	*baseQueryHandler
	stmExist *ent.OrderSelect
}

func (h *existHandler) queryOrder(cli *ent.Client) {
	h.stmExist = h.SelectOrder(
		cli.Order.
			Query().
			Where(
				entorder.ID(*h.ID),
				entorder.DeletedAt(0),
			),
	)
}

func (h *existHandler) queryJoin() error {
	var err error
	h.stmExist.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	return err
}

func (h *Handler) ExistOrder(ctx context.Context) (bool, error) {
	handler := &existHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOrder(cli)
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_exist, err := handler.stmExist.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist

		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistOrderConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	exist := false
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmExist, err = handler.QueryOrders(cli)
		if err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_exist, err := handler.stmExist.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
