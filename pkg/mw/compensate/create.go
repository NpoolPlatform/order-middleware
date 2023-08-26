package compensate

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/compensate"
	compensatecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/compensate"

	"github.com/google/uuid"
)

func (h *Handler) CreateCompensate(ctx context.Context) (*npool.Compensate, error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := compensatecrud.CreateSet(
			tx.Compensate.Create(),
			&compensatecrud.Req{
				ID:      h.ID,
				OrderID: h.OrderID,
				StartAt: h.StartAt,
				EndAt:   h.EndAt,
				Message: h.Message,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetCompensate(ctx)
}
