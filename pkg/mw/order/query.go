package order

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
)

type queryHandler struct {
	*baseQueryHandler
	stmSelect *ent.OrderSelect
	stmCount  *ent.OrderSelect
	infos     []*npool.Order
	total     uint32
}

func (h *queryHandler) queryOrder(cli *ent.Client) {
	h.stmSelect = h.SelectOrder(
		cli.Order.
			Query().
			Where(
				entorder.ID(*h.ID),
				entorder.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.QueryJoinMyself(s)
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.QueryJoinPayment(s)
		err = h.QueryJoinOrderState(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.OrderType = basetypes.OrderType(basetypes.OrderType_value[info.OrderTypeStr])
		info.OrderState = basetypes.OrderState(basetypes.OrderState_value[info.OrderStateStr])
		info.CancelState = basetypes.OrderState(basetypes.OrderState_value[info.CancelStateStr])
		info.PaymentState = basetypes.PaymentState(basetypes.PaymentState_value[info.PaymentStateStr])
		info.PaymentType = basetypes.PaymentType(basetypes.PaymentType_value[info.PaymentTypeStr])
		info.InvestmentType = basetypes.InvestmentType(basetypes.InvestmentType_value[info.InvestmentTypeStr])
		info.StartMode = basetypes.OrderStartMode(basetypes.OrderStartMode_value[info.StartModeStr])
		info.BenefitState = basetypes.BenefitState(basetypes.BenefitState_value[info.BenefitStateStr])
		_ = json.Unmarshal([]byte(info.CouponIDsStr), &info.CouponIDs)
	}
}

func (h *Handler) GetOrder(ctx context.Context) (*npool.Order, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOrder(cli)
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

func (h *Handler) GetOrders(ctx context.Context) ([]*npool.Order, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.QueryOrders(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.QueryOrders(cli)
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
			Order(ent.Desc(entorder.FieldCreatedAt))

		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
