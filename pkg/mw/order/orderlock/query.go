package orderlock

import (
	"context"
	"fmt"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order/orderlock"
	orderlockcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderlock"
)

type queryHandler struct {
	*Handler
	stm   *ent.OrderLockSelect
	infos []*npool.OrderLock
	total uint32
}

func (h *queryHandler) selectOrderLock(stm *ent.OrderLockQuery) {
	h.stm = stm.Select(
		entorderlock.FieldID,
		entorderlock.FieldAppID,
		entorderlock.FieldUserID,
		entorderlock.FieldOrderID,
		entorderlock.FieldLockType,
		entorderlock.FieldCreatedAt,
		entorderlock.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryOrderLock(cli *ent.Client) {
	h.selectOrderLock(
		cli.OrderLock.
			Query().
			Where(
				entorderlock.ID(*h.ID),
				entorderlock.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryOrderLocks(ctx context.Context, cli *ent.Client) error {
	stm, err := orderlockcrud.SetQueryConds(cli.OrderLock.Query(), h.Conds)
	if err != nil {
		return err
	}

	_total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(_total)
	h.selectOrderLock(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.LockType = basetypes.OrderLockType(basetypes.OrderLockType_value[info.LockTypeStr])
	}
}

func (h *Handler) GetOrderLock(ctx context.Context) (*npool.OrderLock, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOrderLock(cli)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetOrderLocks(ctx context.Context) ([]*npool.OrderLock, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrderLocks(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit)).
			Order(ent.Desc(entorderlock.FieldCreatedAt))
		if err := handler.scan(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
