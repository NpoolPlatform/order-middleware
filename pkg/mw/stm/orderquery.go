package orderstm

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
)

type orderQueryHandler struct {
	*Handler
	_ent order
}

func (h *orderQueryHandler) getOrderBaseEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.OrderBase.Query()
	if h.OrderID != nil {
		stm.Where(entorderbase.EntID(*h.OrderID))
	}
	if h._ent.entOrderBase, err = stm.Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return err
	}
	return nil
}

func (h *orderQueryHandler) getOrderStateBase(ctx context.Context, cli *ent.Client) (err error) {
	h._ent.entOrderStateBase, err = cli.
		OrderStateBase.
		Query().
		Where(
			entorderstatebase.OrderID(h._ent.entOrderBase.EntID),
			entorderstatebase.DeletedAt(0),
		).
		Only(ctx)
	return err
}

func (h *orderQueryHandler) _getOrder(ctx context.Context, must bool) error {
	if h.OrderID == nil {
		return fmt.Errorf("invalid orderid")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getOrderBaseEnt(_ctx, cli, must); err != nil {
			return err
		}
		if h._ent.entOrderBase == nil {
			return nil
		}
		return h.getOrderStateBase(_ctx, cli)
	})
}

func (h *orderQueryHandler) requireOrder(ctx context.Context) error {
	return h._getOrder(ctx, true)
}
