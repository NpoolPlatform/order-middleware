package order

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
)

type existHandler struct {
	*Handler
	stm *ent.OrderQuery
}

func (h *existHandler) queryOrder(cli *ent.Client) {
	h.stm = cli.Order.
		Query().
		Where(
			entorder.ID(*h.ID),
			entorder.DeletedAt(0),
		)
}

func (h *existHandler) queryOrders(cli *ent.Client) error {
	stm, err := ordercrud.SetQueryConds(cli.Order.Query(), h.Conds)
	if err != nil {
		return err
	}
	h.stm = stm
	return nil
}

func (h *existHandler) queryJoinOrder(s *sql.Selector) error { //nolint
	t := sql.Table(entorderstate.Table)
	s.LeftJoin(t).
		On(
			s.C(entorder.FieldID),
			t.C(entorderstate.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entorderstate.FieldDeletedAt), 0),
		)

	if h.Conds != nil && h.Conds.OrderState != nil {
		state, ok := h.Conds.OrderState.Val.(basetypes.OrderState)
		if !ok {
			return fmt.Errorf("invalid order orderstate")
		}
		s.Where(
			sql.EQ(t.C(entorderstate.FieldOrderState), state.String()),
		)
	}
	if h.Conds != nil && h.Conds.StartMode != nil {
		startMode, ok := h.Conds.StartMode.Val.(basetypes.OrderStartMode)
		if !ok {
			return fmt.Errorf("invalid order startmode")
		}
		s.Where(
			sql.EQ(t.C(entorderstate.FieldStartMode), startMode.String()),
		)
	}
	if h.Conds != nil && h.Conds.BenefitState != nil {
		benefitState, ok := h.Conds.BenefitState.Val.(basetypes.BenefitState)
		if !ok {
			return fmt.Errorf("invalid order benefitstate")
		}
		s.Where(
			sql.EQ(t.C(entorderstate.FieldBenefitState), benefitState.String()),
		)
	}
	if h.Conds != nil && h.Conds.PaymentState != nil {
		paymentState, ok := h.Conds.PaymentState.Val.(basetypes.PaymentState)
		if !ok {
			return fmt.Errorf("invalid order paymentstate")
		}
		s.Where(
			sql.EQ(t.C(entorderstate.FieldPaymentState), paymentState.String()),
		)
	}
	if h.Conds != nil && h.Conds.PaymentTransactionID != nil {
		paymentTransactionID, ok := h.Conds.PaymentTransactionID.Val.(string)
		if !ok {
			return fmt.Errorf("invalid order paymenttransactionid")
		}
		s.Where(
			sql.EQ(t.C(entorderstate.FieldPaymentTransactionID), paymentTransactionID),
		)
	}
	if h.Conds != nil && h.Conds.LastBenefitAt != nil {
		lastBenefitAt, ok := h.Conds.LastBenefitAt.Val.(uint32)
		if !ok {
			return fmt.Errorf("invalid order lastbenefitat")
		}
		s.Where(
			sql.EQ(t.C(entorderstate.FieldLastBenefitAt), lastBenefitAt),
		)
	}
	if h.Conds != nil && h.Conds.OrderStates != nil {
		states, ok := h.Conds.OrderStates.Val.([]string)
		if !ok {
			return fmt.Errorf("invalid order orderstates")
		}
		if len(states) > 0 {
			var valueInterfaces []interface{}
			for _, value := range states {
				valueInterfaces = append(valueInterfaces, value)
			}
			s.Where(
				sql.In(t.C(entorderstate.FieldOrderState), valueInterfaces...),
			)
		}
	}
	return nil
}

func (h *existHandler) queryJoin() error {
	var err error
	h.stm.Modify(func(s *sql.Selector) {
		err = h.queryJoinOrder(s)
	})
	return err
}

func (h *Handler) ExistOrder(ctx context.Context) (bool, error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryOrder(cli)
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistOrderConds(ctx context.Context) (bool, error) {
	handler := &existHandler{
		Handler: h,
	}

	exist := false

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryOrders(cli); err != nil {
			return err
		}
		if err := handler.queryJoin(); err != nil {
			return err
		}
		_exist, err := handler.stm.Exist(_ctx)
		if err != nil {
			return err
		}
		exist = _exist
		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
