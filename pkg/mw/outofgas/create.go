package outofgas

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/outofgas"
	outofgascrud "github.com/NpoolPlatform/order-middleware/pkg/crud/outofgas"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"

	"github.com/google/uuid"
)

func (h *Handler) CreateOutOfGas(ctx context.Context) (*npool.OutOfGas, error) {
	key := fmt.Sprintf("%v:%v", basetypes.Prefix_PrefixCreateUserTransfer, *h.OrderID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	handler, err := NewHandler(
		ctx,
		WithConds(&npool.Conds{
			OrderID: &basetypes.StringVal{Op: cruder.EQ, Value: h.OrderID.String()},
		}),
	)
	if err != nil {
		return nil, err
	}
	exist, err := handler.ExistOutOfGasConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("outofgas exist")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithTx(ctx, func(ctx context.Context, tx *ent.Tx) error {
		if _, err := outofgascrud.CreateSet(
			tx.OutOfGas.Create(),
			&outofgascrud.Req{
				ID:      h.ID,
				OrderID: h.OrderID,
				Start:   h.Start,
				End:     h.End,
			},
		).Save(ctx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetOutOfGas(ctx)
}
