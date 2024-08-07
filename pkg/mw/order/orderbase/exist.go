package orderbase

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
)

func (h *Handler) ExistOrderBase(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.OrderBase.Query()
		if h.ID != nil {
			stm.Where(entorderbase.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entorderbase.EntID(*h.EntID))
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
