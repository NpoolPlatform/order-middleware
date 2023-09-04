package order

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/order/v1"

	ordercrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order"
	entorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	entorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	entpayment "github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"
)

type baseQueryHandler struct {
	*Handler
}

func (h *baseQueryHandler) SelectOrder(stm *ent.OrderQuery) *ent.OrderSelect {
	return stm.Select(entorder.FieldID)
}

func (h *baseQueryHandler) QueryOrders(cli *ent.Client) (*ent.OrderSelect, error) {
	stm, err := ordercrud.SetQueryConds(cli.Order.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.SelectOrder(stm), nil
}

func (h *baseQueryHandler) QueryJoinMyself(s *sql.Selector) {
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
		t.C(entorder.FieldGoodValueUsd),
		t.C(entorder.FieldPaymentAmount),
		t.C(entorder.FieldDiscountAmount),
		t.C(entorder.FieldPromotionID),
		t.C(entorder.FieldDurationDays),
		t.C(entorder.FieldOrderType),
		t.C(entorder.FieldInvestmentType),
		t.C(entorder.FieldCouponIds),
		t.C(entorder.FieldPaymentType),
		t.C(entorder.FieldCoinTypeID),
		t.C(entorder.FieldPaymentCoinTypeID),
		t.C(entorder.FieldTransferAmount),
		t.C(entorder.FieldBalanceAmount),
		t.C(entorder.FieldCoinUsdCurrency),
		t.C(entorder.FieldLocalCoinUsdCurrency),
		t.C(entorder.FieldLiveCoinUsdCurrency),
		t.C(entorder.FieldCreatedAt),
		t.C(entorder.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) QueryJoinPayment(s *sql.Selector) error {
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
		sql.As(t.C(entpayment.FieldStartAmount), "payment_start_amount"),
	)
	return nil
}

//nolint:gocyclo
func (h *baseQueryHandler) QueryJoinOrderState(s *sql.Selector) error {
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

	s.AppendSelect(
		sql.As(t.C(entorderstate.FieldOrderState), "order_state"),
		sql.As(t.C(entorderstate.FieldCancelState), "cancel_state"),
		sql.As(t.C(entorderstate.FieldStartMode), "start_mode"),
		sql.As(t.C(entorderstate.FieldStartAt), "start_at"),
		sql.As(t.C(entorderstate.FieldEndAt), "end_at"),
		sql.As(t.C(entorderstate.FieldLastBenefitAt), "last_benefit_at"),
		sql.As(t.C(entorderstate.FieldBenefitState), "benefit_state"),
		sql.As(t.C(entorderstate.FieldUserSetPaid), "user_set_paid"),
		sql.As(t.C(entorderstate.FieldAdminSetCanceled), "admin_set_canceled"),
		sql.As(t.C(entorderstate.FieldUserSetCanceled), "user_set_canceled"),
		sql.As(t.C(entorderstate.FieldPaymentTransactionID), "payment_transaction_id"),
		sql.As(t.C(entorderstate.FieldPaymentFinishAmount), "payment_finish_amount"),
		sql.As(t.C(entorderstate.FieldPaymentState), "payment_state"),
		sql.As(t.C(entorderstate.FieldOutofgasHours), "outofgas_hours"),
		sql.As(t.C(entorderstate.FieldCompensateHours), "compensate_hours"),
	)
	return nil
}