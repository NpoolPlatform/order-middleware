package outofgas

import (
	"context"
	"fmt"

	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	powerrentalstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/powerrental/state"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	outofgas1 "github.com/NpoolPlatform/order-middleware/pkg/mw/outofgas"
	powerrentalstate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental/state"
)

type createHandler struct {
	*powerRentalStateQueryHandler
	sqlOutOfGas         string
	sqlPowerRentalState string
}

func (h *createHandler) constructOutOfGasSQL(ctx context.Context, req *outofgascrud.Req) {
	handler, _ := outofgas1.NewHandler(ctx)
	handler.Req = *req
	h.sqlOutOfGas = handler.ConstructCreateSQL()
}

func (h *createHandler) constructPowerRentalStateSQL(ctx context.Context, req *powerrentalstatecrud.Req) (err error) {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *req
	h.sqlPowerRentalState, err = handler.ConstructUpdateSQL()
	return err
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create outofgas: %v", err)
	}
	return nil
}

func (h *createHandler) createOutOfGas(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOutOfGas)
}

func (h *createHandler) updatePowerRentalState(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlPowerRentalState)
}

func (h *Handler) CreateOutOfGas(ctx context.Context) error {
	if *h.EndAt <= *h.StartAt {
		return fmt.Errorf("invalid duration")
	}

	handler := &createHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requirePowerRentalState(ctx); err != nil {
		return err
	}

	handler.constructOutOfGasSQL(ctx, &h.Req)
	if err := handler.constructPowerRentalStateSQL(ctx, &powerrentalstatecrud.Req{
		OrderID:         h.OrderID,
		OutOfGasSeconds: func() *uint32 { u := *h.EndAt - *h.StartAt + handler._ent.OutOfGasSeconds(); return &u }(),
	}); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createOutOfGas(_ctx, tx); err != nil {
			return err
		}
		return handler.updatePowerRentalState(_ctx, tx)
	})
}
