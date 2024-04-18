package ordercoupon

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
	sql string
}

func (h *createHandler) constructSQL() {
	h.sql = h.ConstructCreateSQL()
}

func (h *createHandler) createOrderCoupon(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail create ordercoupon: %v", err)
	}
	return nil
}

func (h *Handler) CreateOrderCoupon(ctx context.Context) error {
	if h.EntID == nil {
		h.EntID = func() *uuid.UUID { uid := uuid.New(); return &uid }()
	}
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSQL()
	return db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		return handler.createOrderCoupon(ctx, tx)
	})
}
