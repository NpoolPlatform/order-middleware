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
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpaymentbalancelock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalancelock"
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
	if h.OrderBaseConds.EntID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldEntID),
				h.OrderBaseConds.EntID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.EntIDs != nil {
		s.OnP(
			sql.In(
				t.C(entorderbase.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range h.OrderBaseConds.EntIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
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
	if h.OrderBaseConds.GoodIDs != nil {
		s.OnP(
			sql.In(
				t.C(entorderbase.FieldGoodID),
				func() (_uids []interface{}) {
					for _, uid := range h.OrderBaseConds.GoodIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
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
	if h.OrderBaseConds.AppGoodIDs != nil {
		s.OnP(
			sql.In(
				t.C(entorderbase.FieldAppGoodID),
				func() (_uids []interface{}) {
					for _, uid := range h.OrderBaseConds.AppGoodIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	if h.OrderBaseConds.ParentOrderID != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldParentOrderID),
				h.OrderBaseConds.ParentOrderID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderBaseConds.ParentOrderIDs != nil {
		s.OnP(
			sql.In(
				t.C(entorderbase.FieldParentOrderID),
				func() (_uids []interface{}) {
					for _, uid := range h.OrderBaseConds.ParentOrderIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	if h.OrderBaseConds.OrderType != nil {
		s.OnP(
			sql.EQ(
				t.C(entorderbase.FieldOrderType),
				h.OrderBaseConds.OrderType.Val.(types.OrderType).String(),
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
	if h.FeeOrderConds.ID != nil {
		s.OnP(
			sql.EQ(
				t.C(entfeeorder.FieldID),
				h.FeeOrderConds.ID.Val.(uint32),
			),
		)
	}
	if h.FeeOrderConds.IDs != nil {
		s.OnP(
			sql.In(
				t.C(entfeeorder.FieldID),
				func() (_ids []interface{}) {
					for _, id := range h.FeeOrderConds.IDs.Val.([]uint32) {
						_ids = append(_ids, interface{}(id))
					}
					return _ids
				}()...,
			),
		)
	}
	if h.FeeOrderConds.EntID != nil {
		s.OnP(
			sql.EQ(
				t.C(entfeeorder.FieldEntID),
				h.FeeOrderConds.EntID.Val.(uuid.UUID),
			),
		)
	}
	if h.FeeOrderConds.EntIDs != nil {
		s.OnP(
			sql.In(
				t.C(entfeeorder.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range h.FeeOrderConds.EntIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	if h.FeeOrderConds.OrderID != nil {
		s.OnP(
			sql.EQ(
				t.C(entfeeorder.FieldOrderID),
				h.FeeOrderConds.OrderID.Val.(uuid.UUID),
			),
		)
	}
	if h.FeeOrderConds.OrderIDs != nil {
		s.OnP(
			sql.In(
				t.C(entfeeorder.FieldOrderID),
				func() (_uids []interface{}) {
					for _, uid := range h.FeeOrderConds.OrderIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
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
	if h.OrderStateBaseConds.PaymentTypes != nil {
		s.OnP(
			sql.In(
				t.C(entorderstatebase.FieldPaymentType),
				func() (_types []interface{}) {
					for _, _type := range h.OrderStateBaseConds.PaymentTypes.Val.([]types.PaymentType) {
						_types = append(_types, interface{}(_type.String()))
					}
					return _types
				}()...,
			),
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
	if h.OrderStateBaseConds.OrderStates != nil {
		s.OnP(
			sql.In(
				t.C(entorderstatebase.FieldOrderState),
				func() (_types []interface{}) {
					for _, _type := range h.OrderStateBaseConds.OrderStates.Val.([]types.OrderState) {
						_types = append(_types, interface{}(_type.String()))
					}
					return _types
				}()...,
			),
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
	if h.FeeOrderStateConds.PaymentStates != nil {
		s.OnP(
			sql.In(
				t.C(entfeeorderstate.FieldPaymentState),
				func() (_types []interface{}) {
					for _, _type := range h.FeeOrderStateConds.PaymentStates.Val.([]types.PaymentState) {
						_types = append(_types, interface{}(_type.String()))
					}
					return
				}()...,
			),
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

func (h *baseQueryHandler) queryJoinPaymentBase(s *sql.Selector) {
	t1 := sql.Table(entpaymentbase.Table)
	t2 := sql.Table(entpaymentbalancelock.Table)
	t3 := sql.Table(entorderlock.Table)

	s.LeftJoin(t1).
		On(
			s.C(entorderbase.FieldEntID),
			t1.C(entpaymentbase.FieldOrderID),
		).
		OnP(
			sql.EQ(t1.C(entpaymentbase.FieldObseleteState), types.PaymentObseleteState_PaymentObseleteNone.String()),
		).
		LeftJoin(t2).
		On(
			t1.C(entpaymentbase.FieldEntID),
			t2.C(entpaymentbalancelock.FieldPaymentID),
		).
		LeftJoin(t3).
		On(
			t2.C(entpaymentbalancelock.FieldLedgerLockID),
			t3.C(entorderlock.FieldEntID),
		)
	s.AppendSelect(
		sql.As(t1.C(entpaymentbase.FieldEntID), "payment_id"),
		sql.As(t3.C(entorderlock.FieldEntID), "ledger_lock_id"),
	)
}

func (h *baseQueryHandler) queryJoinOrderCoupon(s *sql.Selector) {
	t := sql.Table(entordercoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entordercoupon.FieldOrderID),
		)
	if h.OrderCouponConds.OrderID != nil {
		s.OnP(
			sql.EQ(
				t.C(entordercoupon.FieldOrderID),
				h.OrderCouponConds.OrderID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderCouponConds.OrderIDs != nil {
		s.OnP(
			sql.In(
				t.C(entordercoupon.FieldOrderID),
				func() (_uids []interface{}) {
					for _, uid := range h.OrderCouponConds.OrderIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	if h.OrderCouponConds.CouponID != nil {
		s.OnP(
			sql.EQ(
				t.C(entordercoupon.FieldCouponID),
				h.OrderCouponConds.CouponID.Val.(uuid.UUID),
			),
		)
	}
	if h.OrderCouponConds.CouponIDs != nil {
		s.OnP(
			sql.In(
				t.C(entordercoupon.FieldCouponID),
				func() (_uids []interface{}) {
					for _, uid := range h.OrderCouponConds.CouponIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}

}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinFeeOrder(s)
		h.queryJoinOrderStateBase(s)
		h.queryJoinFeeOrderState(s)
		h.queryJoinPaymentBase(s)
		h.queryJoinOrderCoupon(s)
	})
}
