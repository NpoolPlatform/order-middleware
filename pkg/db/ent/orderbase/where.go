// Code generated by ent, DO NOT EDIT.

package orderbase

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// ParentOrderID applies equality check predicate on the "parent_order_id" field. It's identical to ParentOrderIDEQ.
func ParentOrderID(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentOrderID), v))
	})
}

// OrderType applies equality check predicate on the "order_type" field. It's identical to OrderTypeEQ.
func OrderType(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderType), v))
	})
}

// PaymentType applies equality check predicate on the "payment_type" field. It's identical to PaymentTypeEQ.
func PaymentType(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentType), v))
	})
}

// CreateMethod applies equality check predicate on the "create_method" field. It's identical to CreateMethodEQ.
func CreateMethod(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateMethod), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppGoodID)))
	})
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppGoodID)))
	})
}

// ParentOrderIDEQ applies the EQ predicate on the "parent_order_id" field.
func ParentOrderIDEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldParentOrderID), v))
	})
}

// ParentOrderIDNEQ applies the NEQ predicate on the "parent_order_id" field.
func ParentOrderIDNEQ(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldParentOrderID), v))
	})
}

// ParentOrderIDIn applies the In predicate on the "parent_order_id" field.
func ParentOrderIDIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldParentOrderID), v...))
	})
}

// ParentOrderIDNotIn applies the NotIn predicate on the "parent_order_id" field.
func ParentOrderIDNotIn(vs ...uuid.UUID) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldParentOrderID), v...))
	})
}

// ParentOrderIDGT applies the GT predicate on the "parent_order_id" field.
func ParentOrderIDGT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldParentOrderID), v))
	})
}

// ParentOrderIDGTE applies the GTE predicate on the "parent_order_id" field.
func ParentOrderIDGTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldParentOrderID), v))
	})
}

// ParentOrderIDLT applies the LT predicate on the "parent_order_id" field.
func ParentOrderIDLT(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldParentOrderID), v))
	})
}

// ParentOrderIDLTE applies the LTE predicate on the "parent_order_id" field.
func ParentOrderIDLTE(v uuid.UUID) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldParentOrderID), v))
	})
}

// ParentOrderIDIsNil applies the IsNil predicate on the "parent_order_id" field.
func ParentOrderIDIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldParentOrderID)))
	})
}

// ParentOrderIDNotNil applies the NotNil predicate on the "parent_order_id" field.
func ParentOrderIDNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldParentOrderID)))
	})
}

// OrderTypeEQ applies the EQ predicate on the "order_type" field.
func OrderTypeEQ(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderType), v))
	})
}

// OrderTypeNEQ applies the NEQ predicate on the "order_type" field.
func OrderTypeNEQ(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderType), v))
	})
}

// OrderTypeIn applies the In predicate on the "order_type" field.
func OrderTypeIn(vs ...string) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderType), v...))
	})
}

// OrderTypeNotIn applies the NotIn predicate on the "order_type" field.
func OrderTypeNotIn(vs ...string) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderType), v...))
	})
}

// OrderTypeGT applies the GT predicate on the "order_type" field.
func OrderTypeGT(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderType), v))
	})
}

// OrderTypeGTE applies the GTE predicate on the "order_type" field.
func OrderTypeGTE(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderType), v))
	})
}

// OrderTypeLT applies the LT predicate on the "order_type" field.
func OrderTypeLT(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderType), v))
	})
}

// OrderTypeLTE applies the LTE predicate on the "order_type" field.
func OrderTypeLTE(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderType), v))
	})
}

// OrderTypeContains applies the Contains predicate on the "order_type" field.
func OrderTypeContains(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOrderType), v))
	})
}

// OrderTypeHasPrefix applies the HasPrefix predicate on the "order_type" field.
func OrderTypeHasPrefix(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOrderType), v))
	})
}

// OrderTypeHasSuffix applies the HasSuffix predicate on the "order_type" field.
func OrderTypeHasSuffix(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOrderType), v))
	})
}

// OrderTypeIsNil applies the IsNil predicate on the "order_type" field.
func OrderTypeIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderType)))
	})
}

// OrderTypeNotNil applies the NotNil predicate on the "order_type" field.
func OrderTypeNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderType)))
	})
}

// OrderTypeEqualFold applies the EqualFold predicate on the "order_type" field.
func OrderTypeEqualFold(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOrderType), v))
	})
}

// OrderTypeContainsFold applies the ContainsFold predicate on the "order_type" field.
func OrderTypeContainsFold(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOrderType), v))
	})
}

// PaymentTypeEQ applies the EQ predicate on the "payment_type" field.
func PaymentTypeEQ(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeNEQ applies the NEQ predicate on the "payment_type" field.
func PaymentTypeNEQ(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeIn applies the In predicate on the "payment_type" field.
func PaymentTypeIn(vs ...string) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaymentType), v...))
	})
}

// PaymentTypeNotIn applies the NotIn predicate on the "payment_type" field.
func PaymentTypeNotIn(vs ...string) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaymentType), v...))
	})
}

// PaymentTypeGT applies the GT predicate on the "payment_type" field.
func PaymentTypeGT(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeGTE applies the GTE predicate on the "payment_type" field.
func PaymentTypeGTE(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeLT applies the LT predicate on the "payment_type" field.
func PaymentTypeLT(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeLTE applies the LTE predicate on the "payment_type" field.
func PaymentTypeLTE(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeContains applies the Contains predicate on the "payment_type" field.
func PaymentTypeContains(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeHasPrefix applies the HasPrefix predicate on the "payment_type" field.
func PaymentTypeHasPrefix(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeHasSuffix applies the HasSuffix predicate on the "payment_type" field.
func PaymentTypeHasSuffix(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeIsNil applies the IsNil predicate on the "payment_type" field.
func PaymentTypeIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentType)))
	})
}

// PaymentTypeNotNil applies the NotNil predicate on the "payment_type" field.
func PaymentTypeNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentType)))
	})
}

// PaymentTypeEqualFold applies the EqualFold predicate on the "payment_type" field.
func PaymentTypeEqualFold(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeContainsFold applies the ContainsFold predicate on the "payment_type" field.
func PaymentTypeContainsFold(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaymentType), v))
	})
}

// CreateMethodEQ applies the EQ predicate on the "create_method" field.
func CreateMethodEQ(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodNEQ applies the NEQ predicate on the "create_method" field.
func CreateMethodNEQ(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodIn applies the In predicate on the "create_method" field.
func CreateMethodIn(vs ...string) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateMethod), v...))
	})
}

// CreateMethodNotIn applies the NotIn predicate on the "create_method" field.
func CreateMethodNotIn(vs ...string) predicate.OrderBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateMethod), v...))
	})
}

// CreateMethodGT applies the GT predicate on the "create_method" field.
func CreateMethodGT(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodGTE applies the GTE predicate on the "create_method" field.
func CreateMethodGTE(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodLT applies the LT predicate on the "create_method" field.
func CreateMethodLT(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodLTE applies the LTE predicate on the "create_method" field.
func CreateMethodLTE(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodContains applies the Contains predicate on the "create_method" field.
func CreateMethodContains(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodHasPrefix applies the HasPrefix predicate on the "create_method" field.
func CreateMethodHasPrefix(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodHasSuffix applies the HasSuffix predicate on the "create_method" field.
func CreateMethodHasSuffix(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodIsNil applies the IsNil predicate on the "create_method" field.
func CreateMethodIsNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreateMethod)))
	})
}

// CreateMethodNotNil applies the NotNil predicate on the "create_method" field.
func CreateMethodNotNil() predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreateMethod)))
	})
}

// CreateMethodEqualFold applies the EqualFold predicate on the "create_method" field.
func CreateMethodEqualFold(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCreateMethod), v))
	})
}

// CreateMethodContainsFold applies the ContainsFold predicate on the "create_method" field.
func CreateMethodContainsFold(v string) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCreateMethod), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderBase) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderBase) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
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
func Not(p predicate.OrderBase) predicate.OrderBase {
	return predicate.OrderBase(func(s *sql.Selector) {
		p(s.Not())
	})
}
