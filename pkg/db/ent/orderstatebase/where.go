// Code generated by ent, DO NOT EDIT.

package orderstatebase

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderState applies equality check predicate on the "order_state" field. It's identical to OrderStateEQ.
func OrderState(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderState), v))
	})
}

// StartMode applies equality check predicate on the "start_mode" field. It's identical to StartModeEQ.
func StartMode(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartMode), v))
	})
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// LastBenefitAt applies equality check predicate on the "last_benefit_at" field. It's identical to LastBenefitAtEQ.
func LastBenefitAt(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastBenefitAt), v))
	})
}

// BenefitState applies equality check predicate on the "benefit_state" field. It's identical to BenefitStateEQ.
func BenefitState(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitState), v))
	})
}

// PaymentType applies equality check predicate on the "payment_type" field. It's identical to PaymentTypeEQ.
func PaymentType(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentType), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// OrderStateEQ applies the EQ predicate on the "order_state" field.
func OrderStateEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderState), v))
	})
}

// OrderStateNEQ applies the NEQ predicate on the "order_state" field.
func OrderStateNEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderState), v))
	})
}

// OrderStateIn applies the In predicate on the "order_state" field.
func OrderStateIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderState), v...))
	})
}

// OrderStateNotIn applies the NotIn predicate on the "order_state" field.
func OrderStateNotIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderState), v...))
	})
}

// OrderStateGT applies the GT predicate on the "order_state" field.
func OrderStateGT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderState), v))
	})
}

// OrderStateGTE applies the GTE predicate on the "order_state" field.
func OrderStateGTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderState), v))
	})
}

// OrderStateLT applies the LT predicate on the "order_state" field.
func OrderStateLT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderState), v))
	})
}

// OrderStateLTE applies the LTE predicate on the "order_state" field.
func OrderStateLTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderState), v))
	})
}

// OrderStateContains applies the Contains predicate on the "order_state" field.
func OrderStateContains(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderState), v))
	})
}

// OrderStateHasPrefix applies the HasPrefix predicate on the "order_state" field.
func OrderStateHasPrefix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderState), v))
	})
}

// OrderStateHasSuffix applies the HasSuffix predicate on the "order_state" field.
func OrderStateHasSuffix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderState), v))
	})
}

// OrderStateIsNil applies the IsNil predicate on the "order_state" field.
func OrderStateIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderState)))
	})
}

// OrderStateNotNil applies the NotNil predicate on the "order_state" field.
func OrderStateNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderState)))
	})
}

// OrderStateEqualFold applies the EqualFold predicate on the "order_state" field.
func OrderStateEqualFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderState), v))
	})
}

// OrderStateContainsFold applies the ContainsFold predicate on the "order_state" field.
func OrderStateContainsFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderState), v))
	})
}

// StartModeEQ applies the EQ predicate on the "start_mode" field.
func StartModeEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartMode), v))
	})
}

// StartModeNEQ applies the NEQ predicate on the "start_mode" field.
func StartModeNEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartMode), v))
	})
}

// StartModeIn applies the In predicate on the "start_mode" field.
func StartModeIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartMode), v...))
	})
}

// StartModeNotIn applies the NotIn predicate on the "start_mode" field.
func StartModeNotIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartMode), v...))
	})
}

// StartModeGT applies the GT predicate on the "start_mode" field.
func StartModeGT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartMode), v))
	})
}

// StartModeGTE applies the GTE predicate on the "start_mode" field.
func StartModeGTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartMode), v))
	})
}

// StartModeLT applies the LT predicate on the "start_mode" field.
func StartModeLT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartMode), v))
	})
}

// StartModeLTE applies the LTE predicate on the "start_mode" field.
func StartModeLTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartMode), v))
	})
}

// StartModeContains applies the Contains predicate on the "start_mode" field.
func StartModeContains(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStartMode), v))
	})
}

// StartModeHasPrefix applies the HasPrefix predicate on the "start_mode" field.
func StartModeHasPrefix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStartMode), v))
	})
}

// StartModeHasSuffix applies the HasSuffix predicate on the "start_mode" field.
func StartModeHasSuffix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStartMode), v))
	})
}

// StartModeIsNil applies the IsNil predicate on the "start_mode" field.
func StartModeIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartMode)))
	})
}

// StartModeNotNil applies the NotNil predicate on the "start_mode" field.
func StartModeNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartMode)))
	})
}

// StartModeEqualFold applies the EqualFold predicate on the "start_mode" field.
func StartModeEqualFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStartMode), v))
	})
}

// StartModeContainsFold applies the ContainsFold predicate on the "start_mode" field.
func StartModeContainsFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStartMode), v))
	})
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartAt), v))
	})
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartAt), v))
	})
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartAt), v...))
	})
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartAt), v...))
	})
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartAt), v))
	})
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartAt), v))
	})
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartAt), v))
	})
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartAt), v))
	})
}

// StartAtIsNil applies the IsNil predicate on the "start_at" field.
func StartAtIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartAt)))
	})
}

// StartAtNotNil applies the NotNil predicate on the "start_at" field.
func StartAtNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartAt)))
	})
}

// LastBenefitAtEQ applies the EQ predicate on the "last_benefit_at" field.
func LastBenefitAtEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastBenefitAt), v))
	})
}

// LastBenefitAtNEQ applies the NEQ predicate on the "last_benefit_at" field.
func LastBenefitAtNEQ(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastBenefitAt), v))
	})
}

// LastBenefitAtIn applies the In predicate on the "last_benefit_at" field.
func LastBenefitAtIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastBenefitAt), v...))
	})
}

// LastBenefitAtNotIn applies the NotIn predicate on the "last_benefit_at" field.
func LastBenefitAtNotIn(vs ...uint32) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastBenefitAt), v...))
	})
}

// LastBenefitAtGT applies the GT predicate on the "last_benefit_at" field.
func LastBenefitAtGT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastBenefitAt), v))
	})
}

// LastBenefitAtGTE applies the GTE predicate on the "last_benefit_at" field.
func LastBenefitAtGTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastBenefitAt), v))
	})
}

// LastBenefitAtLT applies the LT predicate on the "last_benefit_at" field.
func LastBenefitAtLT(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastBenefitAt), v))
	})
}

// LastBenefitAtLTE applies the LTE predicate on the "last_benefit_at" field.
func LastBenefitAtLTE(v uint32) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastBenefitAt), v))
	})
}

// LastBenefitAtIsNil applies the IsNil predicate on the "last_benefit_at" field.
func LastBenefitAtIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLastBenefitAt)))
	})
}

// LastBenefitAtNotNil applies the NotNil predicate on the "last_benefit_at" field.
func LastBenefitAtNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLastBenefitAt)))
	})
}

// BenefitStateEQ applies the EQ predicate on the "benefit_state" field.
func BenefitStateEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitState), v))
	})
}

// BenefitStateNEQ applies the NEQ predicate on the "benefit_state" field.
func BenefitStateNEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefitState), v))
	})
}

// BenefitStateIn applies the In predicate on the "benefit_state" field.
func BenefitStateIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBenefitState), v...))
	})
}

// BenefitStateNotIn applies the NotIn predicate on the "benefit_state" field.
func BenefitStateNotIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBenefitState), v...))
	})
}

// BenefitStateGT applies the GT predicate on the "benefit_state" field.
func BenefitStateGT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefitState), v))
	})
}

// BenefitStateGTE applies the GTE predicate on the "benefit_state" field.
func BenefitStateGTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefitState), v))
	})
}

// BenefitStateLT applies the LT predicate on the "benefit_state" field.
func BenefitStateLT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefitState), v))
	})
}

// BenefitStateLTE applies the LTE predicate on the "benefit_state" field.
func BenefitStateLTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefitState), v))
	})
}

// BenefitStateContains applies the Contains predicate on the "benefit_state" field.
func BenefitStateContains(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBenefitState), v))
	})
}

// BenefitStateHasPrefix applies the HasPrefix predicate on the "benefit_state" field.
func BenefitStateHasPrefix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBenefitState), v))
	})
}

// BenefitStateHasSuffix applies the HasSuffix predicate on the "benefit_state" field.
func BenefitStateHasSuffix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBenefitState), v))
	})
}

// BenefitStateIsNil applies the IsNil predicate on the "benefit_state" field.
func BenefitStateIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBenefitState)))
	})
}

// BenefitStateNotNil applies the NotNil predicate on the "benefit_state" field.
func BenefitStateNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBenefitState)))
	})
}

// BenefitStateEqualFold applies the EqualFold predicate on the "benefit_state" field.
func BenefitStateEqualFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBenefitState), v))
	})
}

// BenefitStateContainsFold applies the ContainsFold predicate on the "benefit_state" field.
func BenefitStateContainsFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBenefitState), v))
	})
}

// PaymentTypeEQ applies the EQ predicate on the "payment_type" field.
func PaymentTypeEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeNEQ applies the NEQ predicate on the "payment_type" field.
func PaymentTypeNEQ(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeIn applies the In predicate on the "payment_type" field.
func PaymentTypeIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaymentType), v...))
	})
}

// PaymentTypeNotIn applies the NotIn predicate on the "payment_type" field.
func PaymentTypeNotIn(vs ...string) predicate.OrderStateBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaymentType), v...))
	})
}

// PaymentTypeGT applies the GT predicate on the "payment_type" field.
func PaymentTypeGT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeGTE applies the GTE predicate on the "payment_type" field.
func PaymentTypeGTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeLT applies the LT predicate on the "payment_type" field.
func PaymentTypeLT(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeLTE applies the LTE predicate on the "payment_type" field.
func PaymentTypeLTE(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeContains applies the Contains predicate on the "payment_type" field.
func PaymentTypeContains(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeHasPrefix applies the HasPrefix predicate on the "payment_type" field.
func PaymentTypeHasPrefix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeHasSuffix applies the HasSuffix predicate on the "payment_type" field.
func PaymentTypeHasSuffix(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeIsNil applies the IsNil predicate on the "payment_type" field.
func PaymentTypeIsNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentType)))
	})
}

// PaymentTypeNotNil applies the NotNil predicate on the "payment_type" field.
func PaymentTypeNotNil() predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentType)))
	})
}

// PaymentTypeEqualFold applies the EqualFold predicate on the "payment_type" field.
func PaymentTypeEqualFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeContainsFold applies the ContainsFold predicate on the "payment_type" field.
func PaymentTypeContainsFold(v string) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaymentType), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderStateBase) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderStateBase) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
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
func Not(p predicate.OrderStateBase) predicate.OrderStateBase {
	return predicate.OrderStateBase(func(s *sql.Selector) {
		p(s.Not())
	})
}
