package feeorder

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entfeeorder "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorder"
	entfeeorderstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorderstate"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpaymentbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderBaseSelect
}

func (h *baseQueryHandler) selectOrderBase(stm *ent.OrderBaseQuery) *ent.OrderBaseSelect {
	return stm.Select(entorderbase.FieldID)
}

func (h *baseQueryHandler) queryOrderBase(cli *ent.Client) error {
	if h.OrderID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.OrderBase.Query().Where(entorderbase.DeletedAt(0))
	if h.OrderID != nil {
		stm.Where(entorderbase.EntID(*h.OrderID))
	}
	h.stmSelect = h.selectOrderBase(stm)
	return nil
}

func (h *baseQueryHandler) queryOrderBases(cli *ent.Client) (*ent.OrderBaseSelect, error) {
	stm, err := orderbasecrud.SetQueryConds(cli.OrderBase.Query(), h.OrderBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectOrderBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldID),
			t.C(entorderbase.FieldID),
		)
	if h.OrderBaseConds.AppID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldAppID),
				h.OrderBaseConds.AppID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.UserID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldUserID),
				h.OrderBaseConds.UserID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.GoodID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldGoodID),
				h.OrderBaseConds.GoodID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.AppGoodID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldAppGoodID),
				h.OrderBaseConds.AppGoodID.Val.(uuid.UUID),
			),
		)
	}
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldGoodType),
		t.C(entorderbase.FieldAppGoodID),
		t.C(entorderbase.FieldParentOrderID),
		t.C(entorderbase.FieldOrderType),
		t.C(entorderbase.FieldCreateMethod),
		t.C(entorderbase.FieldCreatedAt),
		t.C(entorderbase.FieldUpdatedAt),
	)
}

func (h *baseQueryHandler) queryJoinFeeOrder(s *sql.Selector) {
	t := sql.Table(entfeeorder.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entfeeorder.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entfeeorder.FieldDeletedAt), 0),
		)
	s.AppendSelect(
		t.C(entfeeorder.FieldID),
		t.C(entfeeorder.FieldEntID),
		t.C(entfeeorder.FieldOrderID),
		t.C(entfeeorder.FieldGoodValueUsd),
		t.C(entfeeorder.FieldPaymentAmountUsd),
		t.C(entfeeorder.FieldDiscountAmountUsd),
		t.C(entfeeorder.FieldPromotionID),
		t.C(entfeeorder.FieldDurationSeconds),
	)
}

func (h *baseQueryHandler) queryJoinOrderStateBase(s *sql.Selector) error {
	t := sql.Table(entorderstatebase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entorderstatebase.FieldOrderID),
		)
	if h.OrderStateBaseConds.PaymentType != nil {
		_type, ok := h.OrderStateBaseConds.PaymentType.Val.(types.PaymentType)
		if !ok {
			return fmt.Errorf("invalid paymenttype")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldPaymentType), _type.String()),
		)
	}
	if h.OrderStateBaseConds.OrderState != nil {
		_state, ok := h.OrderStateBaseConds.OrderState.Val.(types.OrderState)
		if !ok {
			return fmt.Errorf("invalid orderstate")
		}
		s.OnP(
			sql.EQ(t.C(entorderstatebase.FieldOrderState), _state.String()),
		)
	}
	s.AppendSelect(
		t.C(entorderstatebase.FieldPaymentType),
		t.C(entorderstatebase.FieldOrderState),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinFeeOrderState(s *sql.Selector) error {
	t := sql.Table(entfeeorderstate.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entfeeorderstate.FieldOrderID),
		)
	if h.FeeOrderStateConds.PaymentState != nil {
		_state, ok := h.FeeOrderStateConds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return fmt.Errorf("invalid paymentstate")
		}
		s.OnP(
			sql.EQ(t.C(entfeeorderstate.FieldPaymentState), _state.String()),
		)
	}
	s.AppendSelect(
		t.C(entfeeorderstate.FieldPaymentID),
		t.C(entfeeorderstate.FieldPaidAt),
		t.C(entfeeorderstate.FieldUserSetPaid),
		t.C(entfeeorderstate.FieldUserSetCanceled),
		t.C(entfeeorderstate.FieldAdminSetCanceled),
		t.C(entfeeorderstate.FieldPaymentState),
		t.C(entfeeorderstate.FieldCancelState),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinLedgerLock(s *sql.Selector) {
	t := sql.Table(entorderlock.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entorderlock.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entorderlock.FieldLockType), types.OrderLockType_LockBalance.String()),
		)
	s.AppendSelect(
		sql.As(t.C(entorderlock.FieldEntID), "ledger_lock_id"),
	)
}

func (h *baseQueryHandler) queryJoinPaymentBase(s *sql.Selector) {
	t := sql.Table(entpaymentbase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entpaymentbase.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entpaymentbase.FieldObseleteState), types.PaymentObseleteState_PaymentObseleteNone.String()),
		)
	s.AppendSelect(
		sql.As(t.C(entpaymentbase.FieldEntID), "payment_id"),
	)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinFeeOrder(s)
		h.queryJoinOrderStateBase(s)
		h.queryJoinFeeOrderState(s)
		h.queryJoinLedgerLock(s)
		h.queryJoinPaymentBase(s)
	})
}
