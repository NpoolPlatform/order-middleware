package orderlock

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/lock"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount *ent.OrderLockSelect
	infos    []*npool.OrderLock
	total    uint32
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.LockType = basetypes.OrderLockType(basetypes.OrderLockType_value[info.LockTypeStr])
	}
}

func (h *Handler) GetOrderLock(ctx context.Context) (*npool.OrderLock, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderLock(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetOrderLocks(ctx context.Context) (infos []*npool.OrderLock, total uint32, err error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryOrderLocks(cli); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stmCount, err = handler.queryOrderLocks(cli); err != nil {
			return wlog.WrapError(err)
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(_total)

		handler.
			stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entorderlock.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, 0, wlog.WrapError(err)
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
