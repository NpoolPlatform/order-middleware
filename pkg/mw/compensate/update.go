package compensate

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update compensates "
	if h.CompensateSeconds != nil {
		_sql += fmt.Sprintf("%vcompensate_seconds = '%v', ", set, *h.CompensateSeconds)
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where "
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v ", *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
		whereAnd = "and"
	}

	h.sql = _sql
	return nil
}

func (h *updateHandler) updateCompensate(ctx context.Context, tx *ent.Tx) error {
	if _, err := compensatecrud.UpdateSet(
		tx.Compensate.UpdateOneID(*h.ID),
		&compensatecrud.Req{
			CompensateSeconds: h.CompensateSeconds,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateCompensate(ctx context.Context) error {
	info, err := h.GetCompensate(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid compensate")
	}

	handler := &updateHandler{
		Handler: h,
	}
	h.ID = &info.ID
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateCompensate(ctx, tx)
	})
}
