package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entpowerrentalstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"
)

type powerRentalStateQueryHandler struct {
	*Handler
	_ent powerRentalState
}

func (h *powerRentalStateQueryHandler) getPowerRentalStateEnt(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h._ent.entPowerRentalState, err = cli.
		PowerRentalState.
		Query().
		Where(
			entpowerrentalstate.OrderID(*h.OrderID),
		).
		Only(ctx); err != nil {
		if ent.IsNotFound(err) && !must {
			return nil
		}
		return err
	}
	return nil
}

func (h *powerRentalStateQueryHandler) _getPowerRentalState(ctx context.Context, must bool) error {
	if h.OrderID == nil {
		return wlog.Errorf("invalid orderid")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h.getPowerRentalStateEnt(ctx, cli, must)
	})
}

func (h *powerRentalStateQueryHandler) requirePowerRentalState(ctx context.Context) error {
	return h._getPowerRentalState(ctx, true)
}
