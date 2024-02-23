package config

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entconfig "github.com/NpoolPlatform/order-middleware/pkg/db/ent/simulateconfig"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/simulate/config"
	configcrud "github.com/NpoolPlatform/order-middleware/pkg/crud/simulate/config"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.SimulateConfigSelect
	stmCount  *ent.SimulateConfigSelect
	infos     []*npool.SimulateConfig
	total     uint32
}

func (h *queryHandler) selectSimulateConfig(stm *ent.SimulateConfigQuery) *ent.SimulateConfigSelect {
	return stm.Select(entconfig.FieldID)
}

func (h *queryHandler) querySimulateConfig(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.SimulateConfig.Query().Where(entconfig.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entconfig.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entconfig.EntID(*h.EntID))
	}
	h.stmSelect = h.selectSimulateConfig(stm)
	return nil
}

func (h *queryHandler) querySimulateConfigs(cli *ent.Client) (*ent.SimulateConfigSelect, error) {
	stm, err := configcrud.SetQueryConds(cli.SimulateConfig.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectSimulateConfig(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entconfig.Table)
	s.AppendSelect(
		t.C(entconfig.FieldID),
		t.C(entconfig.FieldEntID),
		t.C(entconfig.FieldAppID),
		t.C(entconfig.FieldSendCouponMode),
		t.C(entconfig.FieldSendCouponProbability),
		t.C(entconfig.FieldEnabledProfitTx),
		t.C(entconfig.FieldProfitTxProbability),
		t.C(entconfig.FieldEnabled),
		t.C(entconfig.FieldCreatedAt),
		t.C(entconfig.FieldUpdatedAt),
	)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		probability, err := decimal.NewFromString(info.SendCouponProbability)
		if err != nil {
			info.SendCouponProbability = decimal.NewFromInt(0).String()
		} else {
			info.SendCouponProbability = probability.String()
		}
		txProbability, err := decimal.NewFromString(info.ProfitTxProbability)
		if err != nil {
			info.ProfitTxProbability = decimal.NewFromInt(0).String()
		} else {
			info.ProfitTxProbability = txProbability.String()
		}
		info.SendCouponMode = basetypes.SendCouponMode(basetypes.SendCouponMode_value[info.SendCouponModeStr])
	}
}

func (h *Handler) GetSimulateConfig(ctx context.Context) (*npool.SimulateConfig, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.querySimulateConfig(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetSimulateConfigs(ctx context.Context) ([]*npool.SimulateConfig, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.querySimulateConfigs(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.querySimulateConfigs(cli)
		if err != nil {
			return err
		}

		if err := handler.queryJoin(); err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(_total)

		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit)).
			Order(ent.Desc(entconfig.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
