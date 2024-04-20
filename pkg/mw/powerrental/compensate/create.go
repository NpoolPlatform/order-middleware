package compensate

import (
	"context"
	"fmt"

	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	powerrentalstatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/powerrental/state"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	compensate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/compensate"
	powerrentalstate1 "github.com/NpoolPlatform/order-middleware/pkg/mw/powerrental/state"
)

type createHandler struct {
	*powerRentalStateQueryHandler
	sqlCompensates       []string
	sqlPowerRentalStates []string
}

func (h *createHandler) constructCompensateSQL(ctx context.Context, req *compensatecrud.Req) {
	handler, _ := compensate1.NewHandler(ctx)
	handler.Req = *req
	h.sqlCompensates = append(h.sqlCompensates, handler.ConstructCreateSQL())
}

func (h *createHandler) constructPowerRentalStateSQL(ctx context.Context, req *powerrentalstatecrud.Req) error {
	handler, _ := powerrentalstate1.NewHandler(ctx)
	handler.Req = *req
	sql, err := handler.ConstructUpdateSQL()
	if err != nil {
		return err
	}
	h.sqlPowerRentalStates = append(h.sqlPowerRentalStates, sql)
	return nil
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create compensate: %v", err)
	}
	return nil
}

func (h *createHandler) createCompensates(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlCompensates {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) updatePowerRentalStates(ctx context.Context, tx *ent.Tx) error {
	for _, sql := range h.sqlPowerRentalStates {
		if err := h.execSQL(ctx, tx, sql); err != nil {
			return err
		}
	}
	return nil
}

func (h *createHandler) createGoodCompensates(ctx context.Context, tx *ent.Tx) error {
	return nil
}

func (h *createHandler) createOrderCompensate(ctx context.Context, tx *ent.Tx) error {
	if err := h.requirePowerRentalState(ctx); err != nil {
		return err
	}

	h.constructCompensateSQL(ctx, &h.Req)
	if err := h.constructPowerRentalStateSQL(ctx, &powerrentalstatecrud.Req{
		OrderID:           h.OrderID,
		CompensateSeconds: func() *uint32 { u := *h.CompensateSeconds + h._ent.CompensateSeconds(); return &u }(),
	}); err != nil {
		return err
	}

	if err := h.createCompensates(ctx, tx); err != nil {
		return err
	}
	return h.updatePowerRentalStates(ctx, tx)
}

func (h *Handler) CreateCompensate(ctx context.Context) error {
	if h.OrderID == nil && h.GoodID == nil {
		return fmt.Errorf("invalid compensate id")
	}

	handler := &createHandler{
		powerRentalStateQueryHandler: &powerRentalStateQueryHandler{
			Handler: h,
		},
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if h.GoodID != nil {
			return handler.createGoodCompensates(_ctx, tx)
		}
		return handler.createOrderCompensate(_ctx, tx)
	})
}
