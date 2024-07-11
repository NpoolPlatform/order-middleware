// Code generated by ent, DO NOT EDIT.

package compensate

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// CompensateFromID applies equality check predicate on the "compensate_from_id" field. It's identical to CompensateFromIDEQ.
func CompensateFromID(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensateFromID), v))
	})
}

// CompensateType applies equality check predicate on the "compensate_type" field. It's identical to CompensateTypeEQ.
func CompensateType(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensateType), v))
	})
}

// CompensateSeconds applies equality check predicate on the "compensate_seconds" field. It's identical to CompensateSecondsEQ.
func CompensateSeconds(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensateSeconds), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// CompensateFromIDEQ applies the EQ predicate on the "compensate_from_id" field.
func CompensateFromIDEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensateFromID), v))
	})
}

// CompensateFromIDNEQ applies the NEQ predicate on the "compensate_from_id" field.
func CompensateFromIDNEQ(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompensateFromID), v))
	})
}

// CompensateFromIDIn applies the In predicate on the "compensate_from_id" field.
func CompensateFromIDIn(vs ...uuid.UUID) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompensateFromID), v...))
	})
}

// CompensateFromIDNotIn applies the NotIn predicate on the "compensate_from_id" field.
func CompensateFromIDNotIn(vs ...uuid.UUID) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompensateFromID), v...))
	})
}

// CompensateFromIDGT applies the GT predicate on the "compensate_from_id" field.
func CompensateFromIDGT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompensateFromID), v))
	})
}

// CompensateFromIDGTE applies the GTE predicate on the "compensate_from_id" field.
func CompensateFromIDGTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompensateFromID), v))
	})
}

// CompensateFromIDLT applies the LT predicate on the "compensate_from_id" field.
func CompensateFromIDLT(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompensateFromID), v))
	})
}

// CompensateFromIDLTE applies the LTE predicate on the "compensate_from_id" field.
func CompensateFromIDLTE(v uuid.UUID) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompensateFromID), v))
	})
}

// CompensateFromIDIsNil applies the IsNil predicate on the "compensate_from_id" field.
func CompensateFromIDIsNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCompensateFromID)))
	})
}

// CompensateFromIDNotNil applies the NotNil predicate on the "compensate_from_id" field.
func CompensateFromIDNotNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCompensateFromID)))
	})
}

// CompensateTypeEQ applies the EQ predicate on the "compensate_type" field.
func CompensateTypeEQ(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeNEQ applies the NEQ predicate on the "compensate_type" field.
func CompensateTypeNEQ(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeIn applies the In predicate on the "compensate_type" field.
func CompensateTypeIn(vs ...string) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompensateType), v...))
	})
}

// CompensateTypeNotIn applies the NotIn predicate on the "compensate_type" field.
func CompensateTypeNotIn(vs ...string) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompensateType), v...))
	})
}

// CompensateTypeGT applies the GT predicate on the "compensate_type" field.
func CompensateTypeGT(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeGTE applies the GTE predicate on the "compensate_type" field.
func CompensateTypeGTE(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeLT applies the LT predicate on the "compensate_type" field.
func CompensateTypeLT(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeLTE applies the LTE predicate on the "compensate_type" field.
func CompensateTypeLTE(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeContains applies the Contains predicate on the "compensate_type" field.
func CompensateTypeContains(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeHasPrefix applies the HasPrefix predicate on the "compensate_type" field.
func CompensateTypeHasPrefix(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeHasSuffix applies the HasSuffix predicate on the "compensate_type" field.
func CompensateTypeHasSuffix(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeIsNil applies the IsNil predicate on the "compensate_type" field.
func CompensateTypeIsNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCompensateType)))
	})
}

// CompensateTypeNotNil applies the NotNil predicate on the "compensate_type" field.
func CompensateTypeNotNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCompensateType)))
	})
}

// CompensateTypeEqualFold applies the EqualFold predicate on the "compensate_type" field.
func CompensateTypeEqualFold(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCompensateType), v))
	})
}

// CompensateTypeContainsFold applies the ContainsFold predicate on the "compensate_type" field.
func CompensateTypeContainsFold(v string) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCompensateType), v))
	})
}

// CompensateSecondsEQ applies the EQ predicate on the "compensate_seconds" field.
func CompensateSecondsEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCompensateSeconds), v))
	})
}

// CompensateSecondsNEQ applies the NEQ predicate on the "compensate_seconds" field.
func CompensateSecondsNEQ(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCompensateSeconds), v))
	})
}

// CompensateSecondsIn applies the In predicate on the "compensate_seconds" field.
func CompensateSecondsIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCompensateSeconds), v...))
	})
}

// CompensateSecondsNotIn applies the NotIn predicate on the "compensate_seconds" field.
func CompensateSecondsNotIn(vs ...uint32) predicate.Compensate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCompensateSeconds), v...))
	})
}

// CompensateSecondsGT applies the GT predicate on the "compensate_seconds" field.
func CompensateSecondsGT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCompensateSeconds), v))
	})
}

// CompensateSecondsGTE applies the GTE predicate on the "compensate_seconds" field.
func CompensateSecondsGTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCompensateSeconds), v))
	})
}

// CompensateSecondsLT applies the LT predicate on the "compensate_seconds" field.
func CompensateSecondsLT(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCompensateSeconds), v))
	})
}

// CompensateSecondsLTE applies the LTE predicate on the "compensate_seconds" field.
func CompensateSecondsLTE(v uint32) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCompensateSeconds), v))
	})
}

// CompensateSecondsIsNil applies the IsNil predicate on the "compensate_seconds" field.
func CompensateSecondsIsNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCompensateSeconds)))
	})
}

// CompensateSecondsNotNil applies the NotNil predicate on the "compensate_seconds" field.
func CompensateSecondsNotNil() predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCompensateSeconds)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Compensate) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Compensate) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
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
func Not(p predicate.Compensate) predicate.Compensate {
	return predicate.Compensate(func(s *sql.Selector) {
		p(s.Not())
	})
}
