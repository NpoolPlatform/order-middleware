// Code generated by ent, DO NOT EDIT.

package powerrental

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// Units applies equality check predicate on the "units" field. It's identical to UnitsEQ.
func Units(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnits), v))
	})
}

// GoodValue applies equality check predicate on the "good_value" field. It's identical to GoodValueEQ.
func GoodValue(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodValue), v))
	})
}

// GoodValueUsd applies equality check predicate on the "good_value_usd" field. It's identical to GoodValueUsdEQ.
func GoodValueUsd(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodValueUsd), v))
	})
}

// PaymentAmount applies equality check predicate on the "payment_amount" field. It's identical to PaymentAmountEQ.
func PaymentAmount(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentAmount), v))
	})
}

// DiscountAmount applies equality check predicate on the "discount_amount" field. It's identical to DiscountAmountEQ.
func DiscountAmount(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountAmount), v))
	})
}

// PromotionID applies equality check predicate on the "promotion_id" field. It's identical to PromotionIDEQ.
func PromotionID(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPromotionID), v))
	})
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// InvestmentType applies equality check predicate on the "investment_type" field. It's identical to InvestmentTypeEQ.
func InvestmentType(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvestmentType), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// UnitsEQ applies the EQ predicate on the "units" field.
func UnitsEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnits), v))
	})
}

// UnitsNEQ applies the NEQ predicate on the "units" field.
func UnitsNEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnits), v))
	})
}

// UnitsIn applies the In predicate on the "units" field.
func UnitsIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnits), v...))
	})
}

// UnitsNotIn applies the NotIn predicate on the "units" field.
func UnitsNotIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnits), v...))
	})
}

// UnitsGT applies the GT predicate on the "units" field.
func UnitsGT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnits), v))
	})
}

// UnitsGTE applies the GTE predicate on the "units" field.
func UnitsGTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnits), v))
	})
}

// UnitsLT applies the LT predicate on the "units" field.
func UnitsLT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnits), v))
	})
}

// UnitsLTE applies the LTE predicate on the "units" field.
func UnitsLTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnits), v))
	})
}

// UnitsIsNil applies the IsNil predicate on the "units" field.
func UnitsIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnits)))
	})
}

// UnitsNotNil applies the NotNil predicate on the "units" field.
func UnitsNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnits)))
	})
}

// GoodValueEQ applies the EQ predicate on the "good_value" field.
func GoodValueEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodValue), v))
	})
}

// GoodValueNEQ applies the NEQ predicate on the "good_value" field.
func GoodValueNEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodValue), v))
	})
}

// GoodValueIn applies the In predicate on the "good_value" field.
func GoodValueIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodValue), v...))
	})
}

// GoodValueNotIn applies the NotIn predicate on the "good_value" field.
func GoodValueNotIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodValue), v...))
	})
}

// GoodValueGT applies the GT predicate on the "good_value" field.
func GoodValueGT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodValue), v))
	})
}

// GoodValueGTE applies the GTE predicate on the "good_value" field.
func GoodValueGTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodValue), v))
	})
}

// GoodValueLT applies the LT predicate on the "good_value" field.
func GoodValueLT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodValue), v))
	})
}

// GoodValueLTE applies the LTE predicate on the "good_value" field.
func GoodValueLTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodValue), v))
	})
}

// GoodValueIsNil applies the IsNil predicate on the "good_value" field.
func GoodValueIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodValue)))
	})
}

// GoodValueNotNil applies the NotNil predicate on the "good_value" field.
func GoodValueNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodValue)))
	})
}

// GoodValueUsdEQ applies the EQ predicate on the "good_value_usd" field.
func GoodValueUsdEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodValueUsd), v))
	})
}

// GoodValueUsdNEQ applies the NEQ predicate on the "good_value_usd" field.
func GoodValueUsdNEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodValueUsd), v))
	})
}

// GoodValueUsdIn applies the In predicate on the "good_value_usd" field.
func GoodValueUsdIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodValueUsd), v...))
	})
}

// GoodValueUsdNotIn applies the NotIn predicate on the "good_value_usd" field.
func GoodValueUsdNotIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodValueUsd), v...))
	})
}

// GoodValueUsdGT applies the GT predicate on the "good_value_usd" field.
func GoodValueUsdGT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodValueUsd), v))
	})
}

// GoodValueUsdGTE applies the GTE predicate on the "good_value_usd" field.
func GoodValueUsdGTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodValueUsd), v))
	})
}

// GoodValueUsdLT applies the LT predicate on the "good_value_usd" field.
func GoodValueUsdLT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodValueUsd), v))
	})
}

// GoodValueUsdLTE applies the LTE predicate on the "good_value_usd" field.
func GoodValueUsdLTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodValueUsd), v))
	})
}

// GoodValueUsdIsNil applies the IsNil predicate on the "good_value_usd" field.
func GoodValueUsdIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodValueUsd)))
	})
}

// GoodValueUsdNotNil applies the NotNil predicate on the "good_value_usd" field.
func GoodValueUsdNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodValueUsd)))
	})
}

// PaymentAmountEQ applies the EQ predicate on the "payment_amount" field.
func PaymentAmountEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentAmount), v))
	})
}

// PaymentAmountNEQ applies the NEQ predicate on the "payment_amount" field.
func PaymentAmountNEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentAmount), v))
	})
}

// PaymentAmountIn applies the In predicate on the "payment_amount" field.
func PaymentAmountIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaymentAmount), v...))
	})
}

// PaymentAmountNotIn applies the NotIn predicate on the "payment_amount" field.
func PaymentAmountNotIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaymentAmount), v...))
	})
}

// PaymentAmountGT applies the GT predicate on the "payment_amount" field.
func PaymentAmountGT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentAmount), v))
	})
}

// PaymentAmountGTE applies the GTE predicate on the "payment_amount" field.
func PaymentAmountGTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentAmount), v))
	})
}

// PaymentAmountLT applies the LT predicate on the "payment_amount" field.
func PaymentAmountLT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentAmount), v))
	})
}

// PaymentAmountLTE applies the LTE predicate on the "payment_amount" field.
func PaymentAmountLTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentAmount), v))
	})
}

// PaymentAmountIsNil applies the IsNil predicate on the "payment_amount" field.
func PaymentAmountIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentAmount)))
	})
}

// PaymentAmountNotNil applies the NotNil predicate on the "payment_amount" field.
func PaymentAmountNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentAmount)))
	})
}

// DiscountAmountEQ applies the EQ predicate on the "discount_amount" field.
func DiscountAmountEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDiscountAmount), v))
	})
}

// DiscountAmountNEQ applies the NEQ predicate on the "discount_amount" field.
func DiscountAmountNEQ(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDiscountAmount), v))
	})
}

// DiscountAmountIn applies the In predicate on the "discount_amount" field.
func DiscountAmountIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDiscountAmount), v...))
	})
}

// DiscountAmountNotIn applies the NotIn predicate on the "discount_amount" field.
func DiscountAmountNotIn(vs ...decimal.Decimal) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDiscountAmount), v...))
	})
}

// DiscountAmountGT applies the GT predicate on the "discount_amount" field.
func DiscountAmountGT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDiscountAmount), v))
	})
}

// DiscountAmountGTE applies the GTE predicate on the "discount_amount" field.
func DiscountAmountGTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDiscountAmount), v))
	})
}

// DiscountAmountLT applies the LT predicate on the "discount_amount" field.
func DiscountAmountLT(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDiscountAmount), v))
	})
}

// DiscountAmountLTE applies the LTE predicate on the "discount_amount" field.
func DiscountAmountLTE(v decimal.Decimal) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDiscountAmount), v))
	})
}

// DiscountAmountIsNil applies the IsNil predicate on the "discount_amount" field.
func DiscountAmountIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDiscountAmount)))
	})
}

// DiscountAmountNotNil applies the NotNil predicate on the "discount_amount" field.
func DiscountAmountNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDiscountAmount)))
	})
}

// PromotionIDEQ applies the EQ predicate on the "promotion_id" field.
func PromotionIDEQ(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPromotionID), v))
	})
}

// PromotionIDNEQ applies the NEQ predicate on the "promotion_id" field.
func PromotionIDNEQ(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPromotionID), v))
	})
}

// PromotionIDIn applies the In predicate on the "promotion_id" field.
func PromotionIDIn(vs ...uuid.UUID) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPromotionID), v...))
	})
}

// PromotionIDNotIn applies the NotIn predicate on the "promotion_id" field.
func PromotionIDNotIn(vs ...uuid.UUID) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPromotionID), v...))
	})
}

// PromotionIDGT applies the GT predicate on the "promotion_id" field.
func PromotionIDGT(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPromotionID), v))
	})
}

// PromotionIDGTE applies the GTE predicate on the "promotion_id" field.
func PromotionIDGTE(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPromotionID), v))
	})
}

// PromotionIDLT applies the LT predicate on the "promotion_id" field.
func PromotionIDLT(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPromotionID), v))
	})
}

// PromotionIDLTE applies the LTE predicate on the "promotion_id" field.
func PromotionIDLTE(v uuid.UUID) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPromotionID), v))
	})
}

// PromotionIDIsNil applies the IsNil predicate on the "promotion_id" field.
func PromotionIDIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPromotionID)))
	})
}

// PromotionIDNotNil applies the NotNil predicate on the "promotion_id" field.
func PromotionIDNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPromotionID)))
	})
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDuration), v))
	})
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDuration), v))
	})
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDuration), v...))
	})
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...uint32) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDuration), v...))
	})
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDuration), v))
	})
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDuration), v))
	})
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDuration), v))
	})
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v uint32) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDuration), v))
	})
}

// DurationIsNil applies the IsNil predicate on the "duration" field.
func DurationIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDuration)))
	})
}

// DurationNotNil applies the NotNil predicate on the "duration" field.
func DurationNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDuration)))
	})
}

// InvestmentTypeEQ applies the EQ predicate on the "investment_type" field.
func InvestmentTypeEQ(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeNEQ applies the NEQ predicate on the "investment_type" field.
func InvestmentTypeNEQ(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeIn applies the In predicate on the "investment_type" field.
func InvestmentTypeIn(vs ...string) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInvestmentType), v...))
	})
}

// InvestmentTypeNotIn applies the NotIn predicate on the "investment_type" field.
func InvestmentTypeNotIn(vs ...string) predicate.PowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInvestmentType), v...))
	})
}

// InvestmentTypeGT applies the GT predicate on the "investment_type" field.
func InvestmentTypeGT(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeGTE applies the GTE predicate on the "investment_type" field.
func InvestmentTypeGTE(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeLT applies the LT predicate on the "investment_type" field.
func InvestmentTypeLT(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeLTE applies the LTE predicate on the "investment_type" field.
func InvestmentTypeLTE(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeContains applies the Contains predicate on the "investment_type" field.
func InvestmentTypeContains(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeHasPrefix applies the HasPrefix predicate on the "investment_type" field.
func InvestmentTypeHasPrefix(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeHasSuffix applies the HasSuffix predicate on the "investment_type" field.
func InvestmentTypeHasSuffix(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeIsNil applies the IsNil predicate on the "investment_type" field.
func InvestmentTypeIsNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInvestmentType)))
	})
}

// InvestmentTypeNotNil applies the NotNil predicate on the "investment_type" field.
func InvestmentTypeNotNil() predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInvestmentType)))
	})
}

// InvestmentTypeEqualFold applies the EqualFold predicate on the "investment_type" field.
func InvestmentTypeEqualFold(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldInvestmentType), v))
	})
}

// InvestmentTypeContainsFold applies the ContainsFold predicate on the "investment_type" field.
func InvestmentTypeContainsFold(v string) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldInvestmentType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PowerRental) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PowerRental) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
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
func Not(p predicate.PowerRental) predicate.PowerRental {
	return predicate.PowerRental(func(s *sql.Selector) {
		p(s.Not())
	})
}
