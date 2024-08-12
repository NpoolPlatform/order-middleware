package poolorderuser

import (
	"context"
	"fmt"

	poolorderusermwpb "github.com/NpoolPlatform/message/npool/order/mw/v1/powerrental/poolorderuser"
	poolorderusercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/powerrental/poolorderuser"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	poolorderuserent "github.com/NpoolPlatform/order-middleware/pkg/db/ent/poolorderuser"
)

type queryHandler struct {
	*Handler
	stm   *ent.PoolOrderUserSelect
	infos []poolorderusermwpb.PoolOrderUser
	total uint32
}

func (h *queryHandler) selectPoolOrderUser(stm *ent.PoolOrderUserQuery) {
	h.stm = stm.Select(
		poolorderuserent.FieldID,
		poolorderuserent.FieldEntID,
		poolorderuserent.FieldOrderID,
		poolorderuserent.FieldPoolOrderUserID,
	)
}

func (h *queryHandler) queryPoolOrderUser(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.PoolOrderUser.Query().Where(poolorderuserent.DeletedAt(0))
	if h.ID != nil {
		stm.Where(poolorderuserent.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(poolorderuserent.EntID(*h.EntID))
	}
	h.selectPoolOrderUser(stm)
	return nil
}

func (h *queryHandler) queryPoolOrderUsers(ctx context.Context, cli *ent.Client) error {
	stm, err := poolorderusercrud.SetQueryConds(cli.PoolOrderUser.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectPoolOrderUser(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetPoolOrderUser(ctx context.Context) (*poolorderusermwpb.PoolOrderUser, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPoolOrderUser(cli); err != nil {
			return err
		}
		const singleRowLimit = 2
		handler.stm.Offset(0).Limit(singleRowLimit)
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many record")
	}

	return &handler.infos[0], nil
}

func (h *Handler) GetPoolOrderUsers(ctx context.Context) ([]poolorderusermwpb.PoolOrderUser, uint32, error) {
	if h.PoolOrderUserID == nil {
		return nil, 0, nil
	}
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPoolOrderUsers(_ctx, cli); err != nil {
			return err
		}
		handler.stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
