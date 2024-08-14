package order

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

func (h *Handler) CountOrders(ctx context.Context) (count uint32, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_count, err := handler.stmSelect.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		count = uint32(_count)
		return nil
	})
	if err != nil {
		return 0, wlog.WrapError(err)
	}
	return count, nil
}
