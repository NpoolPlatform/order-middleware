package feeorder

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
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
		return wlog.Errorf("invalid id")
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
			return wlog.Errorf("invalid paymenttype")
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
			return wlog.Errorf("invalid orderstate")
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
			return wlog.Errorf("invalid paymentstate")
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
		t.C(entfeeorderstate.FieldCanceledAt),
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

func (h *baseQueryHandler) queryJoinOrderCoupon(s *sql.Selector) error {
	t := sql.Table(entordercoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entordercoupon.FieldOrderID),
		)
	if h.OrderCouponConds.OrderID != nil {
		uid, ok := h.OrderCouponConds.OrderID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderid")
		}
		s.OnP(
			sql.EQ(t.C(entordercoupon.FieldOrderID), uid),
		)
	}
	if h.OrderCouponConds.OrderIDs != nil {
		uids, ok := h.OrderCouponConds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderids")
		}
		s.OnP(
			sql.In(
				t.C(entordercoupon.FieldOrderID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	if h.OrderCouponConds.CouponID != nil {
		uid, ok := h.OrderCouponConds.CouponID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponid")
		}
		s.OnP(
			sql.EQ(t.C(entordercoupon.FieldCouponID), uid),
		)
	}
	if h.OrderCouponConds.CouponIDs != nil {
		uids, ok := h.OrderCouponConds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponids")
		}
		s.OnP(
			sql.In(
				t.C(entordercoupon.FieldCouponID),
				func() (_uids []interface{}) {
					for _, uid := range uids {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	return nil
}

func (h *baseQueryHandler) queryJoinParentOrder(s *sql.Selector) {
	t1 := sql.Table(entorderbase.Table)
	s.Join(t1).
		On(
			s.C(entorderbase.FieldParentOrderID),
			t1.C(entorderbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entorderbase.FieldAppGoodID), "parent_app_good_id"),
			sql.As(t1.C(entorderbase.FieldGoodType), "parent_good_type"),
		)
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinFeeOrder(s)
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
		if err := h.queryJoinFeeOrderState(s); err != nil {
			logger.Sugar().Errorw("queryJoinFeeOrderState", "Error", err)
		}
		h.queryJoinPaymentBase(s)
		if err := h.queryJoinOrderCoupon(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderCoupon", "Error", err)
		}
		h.queryJoinParentOrder(s)
	})
}
