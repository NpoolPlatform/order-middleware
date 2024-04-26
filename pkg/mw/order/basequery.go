package order

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	types "github.com/NpoolPlatform/message/npool/basetypes/order/v1"
	orderbasecrud "github.com/NpoolPlatform/order-middleware/pkg/crud/order/orderbase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent"
	entorderbase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	entorderstatebase "github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.OrderBaseSelect
}

func (h *baseQueryHandler) selectOrderBase(stm *ent.OrderBaseQuery) *ent.OrderBaseSelect {
	return stm.Select(entorderbase.FieldID)
}

func (h *baseQueryHandler) queryOrderBase(cli *ent.Client) {
	h.stmSelect = h.selectOrderBase(
		cli.OrderBase.
			Query().
			Where(
				entorderbase.EntID(*h.EntID),
				entorderbase.DeletedAt(0),
			),
	)
}

func (h *baseQueryHandler) queryOrderBases(cli *ent.Client) (*ent.OrderBaseSelect, error) {
	stm, err := orderbasecrud.SetQueryConds(cli.OrderBase.Query(), h.OrderBaseConds)
	if err != nil {
		return nil, err
	}
	return h.selectOrderBase(stm), nil
}

//nolint:gocyclo,funlen
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
		t.C(entorderbase.FieldSimulate),
		t.C(entorderbase.FieldCreatedAt),
		t.C(entorderbase.FieldUpdatedAt),
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
		t.C(entorderstatebase.FieldStartMode),
		t.C(entorderstatebase.FieldStartAt),
		t.C(entorderstatebase.FieldLastBenefitAt),
		t.C(entorderstatebase.FieldBenefitState),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinOrderStateBase(s); err != nil {
			logger.Sugar().Errorw("queryJoinOrderStateBase", "Error", err)
		}
	})
}
