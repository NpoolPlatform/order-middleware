package powerrental

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entordercoupon "github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	entorderlock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	entpaymentbalancelock "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalancelock"
	entpaymentbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbase"
	entpowerrental "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrental"
	entpowerrentalstate "github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"

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

func (h *baseQueryHandler) queryJoinPowerRental(s *sql.Selector) {
	t := sql.Table(entpowerrental.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entpowerrental.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entpowerrental.FieldDeletedAt), 0),
		)
	if h.PowerRentalConds.ID != nil {
		s.OnP(
			sql.EQ(
				t.C(entpowerrental.FieldID),
				h.PowerRentalConds.ID.Val.(uint32),
			),
		)
	}
	if h.PowerRentalConds.IDs != nil {
		s.OnP(
			sql.In(
				t.C(entpowerrental.FieldID),
				func() (_ids []interface{}) {
					for _, id := range h.PowerRentalConds.IDs.Val.([]uint32) {
						_ids = append(_ids, interface{}(id))
					}
					return _ids
				}()...,
			),
		)
	}
	if h.PowerRentalConds.EntID != nil {
		s.OnP(
			sql.EQ(
				t.C(entpowerrental.FieldEntID),
				h.PowerRentalConds.EntID.Val.(uuid.UUID),
			),
		)
	}
	if h.PowerRentalConds.EntIDs != nil {
		s.OnP(
			sql.In(
				t.C(entpowerrental.FieldEntID),
				func() (_uids []interface{}) {
					for _, uid := range h.PowerRentalConds.EntIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	if h.PowerRentalConds.OrderID != nil {
		s.OnP(
			sql.EQ(
				t.C(entpowerrental.FieldOrderID),
				h.PowerRentalConds.OrderID.Val.(uuid.UUID),
			),
		)
	}
	if h.PowerRentalConds.OrderIDs != nil {
		s.OnP(
			sql.In(
				t.C(entpowerrental.FieldOrderID),
				func() (_uids []interface{}) {
					for _, uid := range h.PowerRentalConds.OrderIDs.Val.([]uuid.UUID) {
						_uids = append(_uids, interface{}(uid))
					}
					return _uids
				}()...,
			),
		)
	}
	s.AppendSelect(
		t.C(entpowerrental.FieldID),
		t.C(entpowerrental.FieldEntID),
		t.C(entpowerrental.FieldOrderID),
		t.C(entpowerrental.FieldAppGoodStockID),
		t.C(entpowerrental.FieldUnits),
		t.C(entpowerrental.FieldGoodValueUsd),
		t.C(entpowerrental.FieldPaymentAmountUsd),
		t.C(entpowerrental.FieldDiscountAmountUsd),
		t.C(entpowerrental.FieldPromotionID),
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

func (h *baseQueryHandler) queryJoinPowerRentalState(s *sql.Selector) error {
	t := sql.Table(entpowerrentalstate.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entpowerrentalstate.FieldOrderID),
		)
	if h.PowerRentalStateConds.PaymentState != nil {
		_state, ok := h.PowerRentalStateConds.PaymentState.Val.(types.PaymentState)
		if !ok {
			return fmt.Errorf("invalid paymentstate")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrentalstate.FieldPaymentState), _state.String()),
		)
	}
	if h.PowerRentalStateConds.PaymentStates != nil {
		s.OnP(
			sql.In(
				t.C(entpowerrentalstate.FieldPaymentState),
				func() (_types []interface{}) {
					for _, _type := range h.PowerRentalStateConds.PaymentStates.Val.([]types.PaymentState) {
						_types = append(_types, interface{}(_type.String()))
					}
					return
				}()...,
			),
		)
	}
	s.AppendSelect(
		t.C(entpowerrentalstate.FieldPaymentID),
		t.C(entpowerrentalstate.FieldPaidAt),
		t.C(entpowerrentalstate.FieldUserSetPaid),
		t.C(entpowerrentalstate.FieldUserSetCanceled),
		t.C(entpowerrentalstate.FieldAdminSetCanceled),
		t.C(entpowerrentalstate.FieldPaymentState),
		t.C(entpowerrentalstate.FieldCancelState),
		t.C(entpowerrentalstate.FieldDurationSeconds),
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

func (h *baseQueryHandler) queryJoinStockLock(s *sql.Selector) {
	t := sql.Table(entorderlock.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entorderlock.FieldOrderID),
		).
		OnP(
			sql.EQ(t.C(entorderlock.FieldLockType), types.OrderLockType_LockStock.String()),
		)
	s.AppendSelect(
		sql.As(t.C(entorderlock.FieldEntID), "app_good_stock_lock_id"),
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
		h.queryJoinPowerRental(s)
		h.queryJoinOrderStateBase(s)
		h.queryJoinPowerRentalState(s)
		h.queryJoinPaymentBase(s)
		h.queryJoinStockLock(s)
		h.queryJoinOrderCoupon(s)
	})
}
