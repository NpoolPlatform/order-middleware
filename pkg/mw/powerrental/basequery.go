package powerrental

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
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
		return nil, wlog.WrapError(err)
	}
	return h.selectOrderBase(stm), nil
}

//nolint:funlen
func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) error {
	t := sql.Table(entorderbase.Table)
	s.Join(t).
		On(
			s.C(entorderbase.FieldID),
			t.C(entorderbase.FieldID),
		)
	if h.OrderBaseConds.EntID != nil {
		id, ok := h.OrderBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldEntID), id),
		)
	}
	if h.OrderBaseConds.EntIDs != nil {
		uids, ok := h.OrderBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t.C(entorderbase.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.OrderBaseConds.AppID != nil {
		id, ok := h.OrderBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldAppID), id),
		)
	}
	if h.OrderBaseConds.UserID != nil {
		id, ok := h.OrderBaseConds.UserID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid userid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldUserID), id),
		)
	}
	if h.OrderBaseConds.GoodID != nil {
		id, ok := h.OrderBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldGoodID), id),
		)
	}
	if h.OrderBaseConds.GoodIDs != nil {
		uids, ok := h.OrderBaseConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t.C(entorderbase.FieldGoodID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.OrderBaseConds.AppGoodID != nil {
		id, ok := h.OrderBaseConds.AppGoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldAppGoodID), id),
		)
	}
	if h.OrderBaseConds.AppGoodIDs != nil {
		uids, ok := h.OrderBaseConds.AppGoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(sql.In(t.C(entorderbase.FieldAppGoodID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.OrderBaseConds.OrderType != nil {
		_type, ok := h.OrderBaseConds.OrderType.Val.(types.OrderType)
		if !ok {
			return wlog.Errorf("invalid ordertype")
		}
		s.OnP(
			sql.EQ(t.C(entorderbase.FieldOrderType), _type),
		)
	}
	s.AppendSelect(
		t.C(entorderbase.FieldAppID),
		t.C(entorderbase.FieldUserID),
		t.C(entorderbase.FieldGoodID),
		t.C(entorderbase.FieldGoodType),
		t.C(entorderbase.FieldAppGoodID),
		t.C(entorderbase.FieldOrderType),
		t.C(entorderbase.FieldCreateMethod),
		t.C(entorderbase.FieldCreatedAt),
		t.C(entorderbase.FieldUpdatedAt),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinPowerRental(s *sql.Selector) error {
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
		id, ok := h.PowerRentalConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrental.FieldID), id),
		)
	}
	if h.PowerRentalConds.IDs != nil {
		ids, ok := h.PowerRentalConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t.C(entpowerrental.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, interface{}(id))
			}
			return _ids
		}()...))
	}
	if h.PowerRentalConds.EntID != nil {
		id, ok := h.PowerRentalConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrental.FieldEntID), id),
		)
	}
	if h.PowerRentalConds.EntIDs != nil {
		uids, ok := h.PowerRentalConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t.C(entpowerrental.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.PowerRentalConds.OrderID != nil {
		id, ok := h.PowerRentalConds.OrderID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderid")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrental.FieldOrderID), id),
		)
	}
	if h.PowerRentalConds.OrderIDs != nil {
		uids, ok := h.PowerRentalConds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderids")
		}
		s.OnP(sql.In(t.C(entpowerrental.FieldOrderID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
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
	return nil
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
		_types, ok := h.OrderStateBaseConds.PaymentTypes.Val.([]types.PaymentType)
		if !ok {
			return wlog.Errorf("invalid paymenttypes")
		}
		s.OnP(sql.In(t.C(entorderstatebase.FieldPaymentType), func() (__types []interface{}) {
			for _, _type := range _types {
				__types = append(__types, interface{}(_type.String()))
			}
			return __types
		}()...))
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
		states, ok := h.OrderStateBaseConds.OrderStates.Val.([]types.OrderState)
		if !ok {
			return wlog.Errorf("invalid orderstate")
		}
		s.OnP(sql.In(t.C(entorderstatebase.FieldOrderState), func() (_states []interface{}) {
			for _, _state := range states {
				_states = append(_states, interface{}(_state.String()))
			}
			return _states
		}()...))
	}
	s.AppendSelect(
		t.C(entorderstatebase.FieldPaymentType),
		t.C(entorderstatebase.FieldOrderState),
		t.C(entorderstatebase.FieldStartMode),
		t.C(entorderstatebase.FieldStartAt),
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
			return wlog.Errorf("invalid paymentstate")
		}
		s.OnP(
			sql.EQ(t.C(entpowerrentalstate.FieldPaymentState), _state.String()),
		)
	}
	if h.PowerRentalStateConds.PaymentStates != nil {
		states, ok := h.PowerRentalStateConds.PaymentStates.Val.([]types.PaymentState)
		if !ok {
			return wlog.Errorf("invalid paymentstates")
		}
		s.OnP(sql.In(t.C(entpowerrentalstate.FieldPaymentState), func() (_states []interface{}) {
			for _, _state := range states {
				_states = append(_states, interface{}(_state.String()))
			}
			return
		}()...))
	}
	s.AppendSelect(
		t.C(entpowerrentalstate.FieldPaymentID),
		t.C(entpowerrentalstate.FieldPaidAt),
		t.C(entpowerrentalstate.FieldUserSetPaid),
		t.C(entpowerrentalstate.FieldUserSetCanceled),
		t.C(entpowerrentalstate.FieldAdminSetCanceled),
		t.C(entpowerrentalstate.FieldPaymentState),
		t.C(entpowerrentalstate.FieldOutofgasSeconds),
		t.C(entpowerrentalstate.FieldCompensateSeconds),
		t.C(entpowerrentalstate.FieldCancelState),
		t.C(entpowerrentalstate.FieldCanceledAt),
		t.C(entpowerrentalstate.FieldDurationSeconds),
		t.C(entpowerrentalstate.FieldRenewState),
		t.C(entpowerrentalstate.FieldRenewNotifyAt),
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
		LeftJoin(t2).
		On(
			t1.C(entpaymentbase.FieldEntID),
			t2.C(entpaymentbalancelock.FieldPaymentID),
		).
		OnP(
			sql.EQ(t1.C(entpaymentbase.FieldObseleteState), types.PaymentObseleteState_PaymentObseleteNone.String()),
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
			sql.And(
				sql.EQ(t.C(entorderlock.FieldLockType), types.OrderLockType_LockStock.String()),
				sql.EQ(t.C(entorderlock.FieldDeletedAt), 0),
			),
		)
	s.AppendSelect(
		sql.As(t.C(entorderlock.FieldEntID), "app_good_stock_lock_id"),
	)
}

func (h *baseQueryHandler) queryJoinOrderCoupon(s *sql.Selector) error {
	t := sql.Table(entordercoupon.Table)
	s.LeftJoin(t).
		On(
			s.C(entorderbase.FieldEntID),
			t.C(entordercoupon.FieldOrderID),
		).
		Distinct()
	if h.OrderCouponConds.OrderID != nil {
		id, ok := h.OrderCouponConds.OrderID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderid")
		}
		s.OnP(
			sql.EQ(t.C(entordercoupon.FieldOrderID), id),
		)
		s.Where(
			sql.EQ(t.C(entordercoupon.FieldOrderID), id),
		)
	}
	if h.OrderCouponConds.OrderIDs != nil {
		uids, ok := h.OrderCouponConds.OrderIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid orderids")
		}
		s.OnP(sql.In(t.C(entordercoupon.FieldOrderID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
		s.Where(sql.In(t.C(entordercoupon.FieldOrderID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	if h.OrderCouponConds.CouponID != nil {
		id, ok := h.OrderCouponConds.CouponID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponid")
		}
		s.OnP(
			sql.EQ(t.C(entordercoupon.FieldCouponID), id),
		)
		s.Where(
			sql.EQ(t.C(entordercoupon.FieldCouponID), id),
		)
	}
	if h.OrderCouponConds.CouponIDs != nil {
		uids, ok := h.OrderCouponConds.CouponIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid couponids")
		}
		s.OnP(sql.In(t.C(entordercoupon.FieldCouponID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
		s.Where(sql.In(t.C(entordercoupon.FieldCouponID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return _uids
		}()...))
	}
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRental", "Error", err)
		}
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
		if err := h.queryJoinPowerRentalState(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRentalState", "Error", err)
		}
		h.queryJoinPaymentBase(s)
		h.queryJoinStockLock(s)
		if err := h.queryJoinOrderCoupon(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderCoupon", "Error", err)
		}
	})
}
