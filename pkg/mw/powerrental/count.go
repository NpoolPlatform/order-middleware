package powerrental

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type countHandler struct {
	*baseQueryHandler
}

func (h *Handler) CountPowerRentals(ctx context.Context) (count uint32, err error) {
	handler := &countHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		_count, err := handler.stmSelect.Count(_ctx)
		count = uint32(_count)
		return wlog.WrapError(err)
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
