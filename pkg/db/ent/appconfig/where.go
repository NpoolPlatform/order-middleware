// Code generated by ent, DO NOT EDIT.

package appconfig

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// EnableSimulateOrder applies equality check predicate on the "enable_simulate_order" field. It's identical to EnableSimulateOrderEQ.
func EnableSimulateOrder(v bool) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableSimulateOrder), v))
	})
}

// SimulateOrderCouponMode applies equality check predicate on the "simulate_order_coupon_mode" field. It's identical to SimulateOrderCouponModeEQ.
func SimulateOrderCouponMode(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponProbability applies equality check predicate on the "simulate_order_coupon_probability" field. It's identical to SimulateOrderCouponProbabilityEQ.
func SimulateOrderCouponProbability(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCashableProfitProbability applies equality check predicate on the "simulate_order_cashable_profit_probability" field. It's identical to SimulateOrderCashableProfitProbabilityEQ.
func SimulateOrderCashableProfitProbability(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// MaxUnpaidOrders applies equality check predicate on the "max_unpaid_orders" field. It's identical to MaxUnpaidOrdersEQ.
func MaxUnpaidOrders(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxTypedCouponsPerOrder applies equality check predicate on the "max_typed_coupons_per_order" field. It's identical to MaxTypedCouponsPerOrderEQ.
func MaxTypedCouponsPerOrder(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// EnableSimulateOrderEQ applies the EQ predicate on the "enable_simulate_order" field.
func EnableSimulateOrderEQ(v bool) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableSimulateOrder), v))
	})
}

// EnableSimulateOrderNEQ applies the NEQ predicate on the "enable_simulate_order" field.
func EnableSimulateOrderNEQ(v bool) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnableSimulateOrder), v))
	})
}

// EnableSimulateOrderIsNil applies the IsNil predicate on the "enable_simulate_order" field.
func EnableSimulateOrderIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEnableSimulateOrder)))
	})
}

// EnableSimulateOrderNotNil applies the NotNil predicate on the "enable_simulate_order" field.
func EnableSimulateOrderNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEnableSimulateOrder)))
	})
}

// SimulateOrderCouponModeEQ applies the EQ predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeNEQ applies the NEQ predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeNEQ(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeIn applies the In predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeIn(vs ...string) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSimulateOrderCouponMode), v...))
	})
}

// SimulateOrderCouponModeNotIn applies the NotIn predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeNotIn(vs ...string) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSimulateOrderCouponMode), v...))
	})
}

// SimulateOrderCouponModeGT applies the GT predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeGT(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeGTE applies the GTE predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeGTE(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeLT applies the LT predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeLT(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeLTE applies the LTE predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeLTE(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeContains applies the Contains predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeContains(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeHasPrefix applies the HasPrefix predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeHasPrefix(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeHasSuffix applies the HasSuffix predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeHasSuffix(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeIsNil applies the IsNil predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSimulateOrderCouponMode)))
	})
}

// SimulateOrderCouponModeNotNil applies the NotNil predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSimulateOrderCouponMode)))
	})
}

// SimulateOrderCouponModeEqualFold applies the EqualFold predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeEqualFold(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponModeContainsFold applies the ContainsFold predicate on the "simulate_order_coupon_mode" field.
func SimulateOrderCouponModeContainsFold(v string) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSimulateOrderCouponMode), v))
	})
}

// SimulateOrderCouponProbabilityEQ applies the EQ predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCouponProbabilityNEQ applies the NEQ predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityNEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCouponProbabilityIn applies the In predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityIn(vs ...decimal.Decimal) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSimulateOrderCouponProbability), v...))
	})
}

// SimulateOrderCouponProbabilityNotIn applies the NotIn predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityNotIn(vs ...decimal.Decimal) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSimulateOrderCouponProbability), v...))
	})
}

// SimulateOrderCouponProbabilityGT applies the GT predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityGT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCouponProbabilityGTE applies the GTE predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityGTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCouponProbabilityLT applies the LT predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityLT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCouponProbabilityLTE applies the LTE predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityLTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSimulateOrderCouponProbability), v))
	})
}

// SimulateOrderCouponProbabilityIsNil applies the IsNil predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSimulateOrderCouponProbability)))
	})
}

// SimulateOrderCouponProbabilityNotNil applies the NotNil predicate on the "simulate_order_coupon_probability" field.
func SimulateOrderCouponProbabilityNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSimulateOrderCouponProbability)))
	})
}

// SimulateOrderCashableProfitProbabilityEQ applies the EQ predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// SimulateOrderCashableProfitProbabilityNEQ applies the NEQ predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityNEQ(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// SimulateOrderCashableProfitProbabilityIn applies the In predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityIn(vs ...decimal.Decimal) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSimulateOrderCashableProfitProbability), v...))
	})
}

// SimulateOrderCashableProfitProbabilityNotIn applies the NotIn predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityNotIn(vs ...decimal.Decimal) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSimulateOrderCashableProfitProbability), v...))
	})
}

// SimulateOrderCashableProfitProbabilityGT applies the GT predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityGT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// SimulateOrderCashableProfitProbabilityGTE applies the GTE predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityGTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// SimulateOrderCashableProfitProbabilityLT applies the LT predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityLT(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// SimulateOrderCashableProfitProbabilityLTE applies the LTE predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityLTE(v decimal.Decimal) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSimulateOrderCashableProfitProbability), v))
	})
}

// SimulateOrderCashableProfitProbabilityIsNil applies the IsNil predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSimulateOrderCashableProfitProbability)))
	})
}

// SimulateOrderCashableProfitProbabilityNotNil applies the NotNil predicate on the "simulate_order_cashable_profit_probability" field.
func SimulateOrderCashableProfitProbabilityNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSimulateOrderCashableProfitProbability)))
	})
}

// MaxUnpaidOrdersEQ applies the EQ predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxUnpaidOrdersNEQ applies the NEQ predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxUnpaidOrdersIn applies the In predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxUnpaidOrders), v...))
	})
}

// MaxUnpaidOrdersNotIn applies the NotIn predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersNotIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxUnpaidOrders), v...))
	})
}

// MaxUnpaidOrdersGT applies the GT predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxUnpaidOrdersGTE applies the GTE predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxUnpaidOrdersLT applies the LT predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxUnpaidOrdersLTE applies the LTE predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxUnpaidOrders), v))
	})
}

// MaxUnpaidOrdersIsNil applies the IsNil predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMaxUnpaidOrders)))
	})
}

// MaxUnpaidOrdersNotNil applies the NotNil predicate on the "max_unpaid_orders" field.
func MaxUnpaidOrdersNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMaxUnpaidOrders)))
	})
}

// MaxTypedCouponsPerOrderEQ applies the EQ predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderNEQ applies the NEQ predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNEQ(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderIn applies the In predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxTypedCouponsPerOrder), v...))
	})
}

// MaxTypedCouponsPerOrderNotIn applies the NotIn predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNotIn(vs ...uint32) predicate.AppConfig {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxTypedCouponsPerOrder), v...))
	})
}

// MaxTypedCouponsPerOrderGT applies the GT predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderGT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderGTE applies the GTE predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderGTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderLT applies the LT predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderLT(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderLTE applies the LTE predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderLTE(v uint32) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxTypedCouponsPerOrder), v))
	})
}

// MaxTypedCouponsPerOrderIsNil applies the IsNil predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderIsNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMaxTypedCouponsPerOrder)))
	})
}

// MaxTypedCouponsPerOrderNotNil applies the NotNil predicate on the "max_typed_coupons_per_order" field.
func MaxTypedCouponsPerOrderNotNil() predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMaxTypedCouponsPerOrder)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AppConfig) predicate.AppConfig {
	return predicate.AppConfig(func(s *sql.Selector) {
		p(s.Not())
	})
}
