package outofgas

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	outofgas1 "github.com/NpoolPlatform/order-middleware/pkg/mw/outofgas"
)

type createHandler struct {
	*Handler
	sqlOutOfGas string
}

func (h *createHandler) constructOutOfGasSQL(ctx context.Context, req *outofgascrud.Req) {
	handler, _ := outofgas1.NewHandler(ctx)
	handler.Req = *req
	h.sqlOutOfGas = handler.ConstructCreateSQL()
}

func (h *createHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return wlog.Errorf("fail create outofgas: %v", err)
	}
	return nil
}

func (h *createHandler) createOutOfGas(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlOutOfGas)
}

func (h *Handler) CreateOutOfGas(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	handler.constructOutOfGasSQL(ctx, &h.Req)
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createOutOfGas(_ctx, tx)
	})
}
