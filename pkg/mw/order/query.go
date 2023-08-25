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
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
	"github.com/shopspring/decimal"

	npool "github.com/NpoolPlatform/message/npool/order/mw/v1/order"
	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.OrderSelect
	stmCount  *ent.OrderSelect
	infos     []*npool.Order
	total     uint32
}

func (h *queryHandler) selectOrder(stm *ent.OrderQuery) *ent.OrderSelect {
	return stm.Select(entorder.FieldID)
}

func (h *queryHandler) queryOrder(cli *ent.Client) {
	h.stmSelect = h.selectOrder(
		cli.Order.
			Query().
			Where(
				entorder.ID(*h.ID),
				entorder.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryOrders(cli *ent.Client) (*ent.OrderSelect, error) {
	stm, err := ordercrud.SetQueryConds(cli.Order.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectOrder(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorder.Table)
	s.AppendSelect(
		t.C(entorder.FieldID),
		t.C(entorder.FieldAppID),
		t.C(entorder.FieldUserID),
		t.C(entorder.FieldGoodID),
		t.C(entorder.FieldAppGoodID),
		t.C(entorder.FieldPaymentID),
		t.C(entorder.FieldParentOrderID),
		t.C(entorder.FieldUnitsV1),
		t.C(entorder.FieldGoodValue),
		t.C(entorder.FieldPaymentAmount),
		t.C(entorder.FieldDiscountAmount),
		t.C(entorder.FieldPromotionID),
		t.C(entorder.FieldDurationDays),
		t.C(entorder.FieldOrderType),
		t.C(entorder.FieldInvestmentType),
		t.C(entorder.FieldCouponIds),
		t.C(entorder.FieldPaymentType),
	)
}

func (h *queryHandler) queryJoinPayment(s *sql.Selector) error { //nolint
	t := sql.Table(entpayment.Table)
	s.LeftJoin(t).
		On(
			s.C(entorder.FieldID),
			t.C(entpayment.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entpayment.FieldDeletedAt), 0),
		)

	s.AppendSelect(
		sql.As(t.C(entpayment.FieldAccountID), "payment_account_id"),
		sql.As(t.C(entpayment.FieldCoinTypeID), "payment_coin_type_id"),
		sql.As(t.C(entpayment.FieldStartAmount), "payment_start_amount"),
		sql.As(t.C(entpayment.FieldTransferAmount), "payment_transfer_amount"),
		sql.As(t.C(entpayment.FieldBalanceAmount), "payment_balance_amount"),
		sql.As(t.C(entpayment.FieldCoinUsdCurrency), "payment_coin_usd_currency"),
		sql.As(t.C(entpayment.FieldLocalCoinUsdCurrency), "payment_local_coin_usd_currency"),
		sql.As(t.C(entpayment.FieldLiveCoinUsdCurrency), "payment_live_coin_usd_currency"),
	)
	return nil
}

func (h *queryHandler) queryJoinOrderState(s *sql.Selector) error { //nolint
	t := sql.Table(entorderstate.Table)
	s.LeftJoin(t).
		On(
			s.C(entorder.FieldID),
			t.C(entorderstate.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entorderstate.FieldDeletedAt), 0),
		)

	s.AppendSelect(
		sql.As(t.C(entorderstate.FieldOrderState), "order_state"),
		sql.As(t.C(entorderstate.FieldStartMode), "start_mode"),
		sql.As(t.C(entorderstate.FieldStartAt), "start_at"),
		sql.As(t.C(entorderstate.FieldEndAt), "end_at"),
		sql.As(t.C(entorderstate.FieldLastBenefitAt), "last_benefit_at"),
		sql.As(t.C(entorderstate.FieldBenefitState), "benefit_state"),
		sql.As(t.C(entorderstate.FieldUserSetPaid), "user_set_paid"),
		sql.As(t.C(entorderstate.FieldUserSetCanceled), "user_set_canceled"),
		sql.As(t.C(entorderstate.FieldPaymentTransactionID), "payment_transaction_id"),
		sql.As(t.C(entorderstate.FieldPaymentFinishAmount), "payment_finish_amount"),
		sql.As(t.C(entorderstate.FieldPaymentState), "payment_state"),
		sql.As(t.C(entorderstate.FieldOutofgasHours), "outofgas_hours"),
		sql.As(t.C(entorderstate.FieldCompensateHours), "compensate_hours"),
	)
	return nil
}

func (h *queryHandler) queryJoin() error {
	var err error
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinPayment(s)
		err = h.queryJoinOrderState(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinPayment(s)
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
		Handler: h,
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
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryOrders(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryOrders(cli)
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

func (h *Handler) CountOrders(ctx context.Context) (uint32, error) {
	count := uint32(0)
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmCount, err = handler.queryOrders(cli)
		if err != nil {
			return err
		}

		_total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		count = uint32(_total)

		return handler.scan(_ctx)
	})
	if err != nil {
		return 0, err
	}

	return count, err
}

func (h *Handler) SumOrderUnits(ctx context.Context) (string, error) {
	sum := decimal.NewFromInt(0).String()
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := ordercrud.SetQueryConds(cli.Order.Query(), &ordercrud.Conds{})
		if err != nil {
			return err
		}
		_count, err := stm.Count(ctx)
		if err != nil {
			return err
		}
		if _count == 0 {
			return nil
		}
		_sum, err := stm.
			Modify(func(s *sql.Selector) {
				s.Select(sql.Sum(entorder.FieldUnitsV1))
			}).
			String(ctx)
		if err != nil {
			return err
		}
		sum = _sum

		return nil
	})
	if err != nil {
		return sum, err
	}
	return sum, nil
}
