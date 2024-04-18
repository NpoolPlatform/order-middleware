package ordercoupon

import (
	"context"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type existHandler struct {
	*baseQueryHandler
}

func (h *Handler) ExistOrderCoupon(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderCoupon(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistOrderCouponConds(ctx context.Context) (exist bool, err error) {
	handler := &existHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrderCoupons(cli)
		if err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
