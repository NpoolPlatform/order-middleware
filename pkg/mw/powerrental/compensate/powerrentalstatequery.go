package compensate

import (
	"context"

	"entgo.io/ent/dialect/sql"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpowerrentalstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"
)

type powerRentalStateQueryHandler struct {
	*Handler
	offset int32
	limit  int32
	_ent   powerRentalStates
}

func (h *powerRentalStateQueryHandler) getPowerRentalStates(ctx context.Context, cli *ent.Client) error {
	conds := &orderbasecrud.Conds{}
	if h.OrderID != nil {
		conds.EntID = &cruder.Cond{Op: cruder.EQ, Val: *h.OrderID}
	}
	if h.GoodID != nil {
		conds.GoodID = &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID}
	}
	stm, err := orderbasecrud.SetQueryConds(cli.OrderBase.Query(), conds)
	if err != nil {
		return wlog.WrapError(err)
	}
	stmSelect := stm.Select(
		entorderbase.FieldEntID,
	).Modify(func(s *sql.Selector) {
		t1 := sql.Table(entorderstatebase.Table)
		t2 := sql.Table(entpowerrentalstate.Table)
		s.Join(t1).
			On(
				s.C(entorderbase.FieldEntID),
				t1.C(entorderstatebase.FieldOrderID),
			).
			OnP(
				sql.EQ(
					t1.C(entorderstatebase.FieldOrderState),
					types.OrderState_OrderStateInService.String(),
				),
			).
			Join(t2).
			On(
				s.C(entorderbase.FieldEntID),
				t2.C(entpowerrentalstate.FieldOrderID),
			).
			AppendSelect(
				t2.C(entpowerrentalstate.FieldID),
			)
	})
	stmSelect.Offset(int(h.offset))
	if h.limit == 0 {
		h.limit = 2
	}
	stmSelect.Limit(int(h.limit))
	return wlog.WrapError(stmSelect.Scan(ctx, &h._ent.powerRentalStates))
}

func (h *powerRentalStateQueryHandler) _getPowerRentalStates(ctx context.Context, cli *ent.Client) error {
	if h.OrderID == nil && h.GoodID == nil {
		return wlog.Errorf("invalid id")
	}
	h._ent.Drain()
	return h.getPowerRentalStates(ctx, cli)
}

func (h *powerRentalStateQueryHandler) requirePowerRentalStates(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		return h._getPowerRentalStates(_ctx, cli)
	})
}

func (h *powerRentalStateQueryHandler) requirePowerRentalStatesWithTx(ctx context.Context, tx *ent.Tx) error {
	return h._getPowerRentalStates(ctx, tx.Client())
}
