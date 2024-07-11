package compensate

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entpowerrentalstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"

	"github.com/google/uuid"
)

type powerRentalStateQueryHandler struct {
	*Handler
	offset int32
	limit  int32
	_ent   powerRentalStates
}

func (h *powerRentalStateQueryHandler) getGoodPowerRentalStateEnts(ctx context.Context, cli *ent.Client) error {
	orders, err := cli.
		OrderBase.
		Query().
		Where(
			entorderbase.GoodID(*h.GoodID),
		).
		Offset(int(h.offset)).
		Limit(int(h.limit)).
		All(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if len(orders) == 0 {
		return nil
	}

	h._ent.entPowerRentalStates, err = cli.
		PowerRentalState.
		Query().
		Where(
			entpowerrentalstate.OrderIDIn(func() (_uids []uuid.UUID) {
				for _, order := range orders {
					_uids = append(_uids, order.EntID)
				}
				return
			}()...),
		).
		All(ctx)
	return wlog.WrapError(err)
}

func (h *powerRentalStateQueryHandler) getOrderPowerRentalStateEnt(ctx context.Context, cli *ent.Client, must bool) error {
	_ent, err := cli.
		PowerRentalState.
		Query().
		Where(
			entpowerrentalstate.OrderID(*h.OrderID),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return wlog.WrapError(err)
	}
	h._ent.entPowerRentalStates = append(h._ent.entPowerRentalStates, _ent)
	return nil
}

func (h *powerRentalStateQueryHandler) _getPowerRentalStates(ctx context.Context, cli *ent.Client, must bool) error {
	if h.OrderID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}
	h._ent.entPowerRentalStates = []*ent.PowerRentalState{}
	if h.OrderID != nil {
		return h.getOrderPowerRentalStateEnt(ctx, cli, must)
	}
	return h.getGoodPowerRentalStateEnts(ctx, cli)
}

func (h *powerRentalStateQueryHandler) requirePowerRentalStates(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getPowerRentalStates(_ctx, cli, true)
	})
}

func (h *powerRentalStateQueryHandler) requirePowerRentalStatesWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getPowerRentalStates(ctx, tx.Client(), true)
}
