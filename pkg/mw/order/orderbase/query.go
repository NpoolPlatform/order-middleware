package orderbase

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
)

func (h *Handler) GetOrderBase(ctx context.Context) (OrderBase, error) {
	var _orderBase *ent.OrderBase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.OrderBase.Query()
		if h.ID != nil {
			stm.Where(entorderbase.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entorderbase.EntID(*h.EntID))
		}
		_orderBase, err = stm.Only(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	return &orderBase{
		_ent: _orderBase,
	}, nil
}
