// Code generated by ent, DO NOT EDIT.

package feeorderstate

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// PaymentID applies equality check predicate on the "payment_id" field. It's identical to PaymentIDEQ.
func PaymentID(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentID), v))
	})
}

// PaidAt applies equality check predicate on the "paid_at" field. It's identical to PaidAtEQ.
func PaidAt(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaidAt), v))
	})
}

// UserSetPaid applies equality check predicate on the "user_set_paid" field. It's identical to UserSetPaidEQ.
func UserSetPaid(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserSetPaid), v))
	})
}

// UserSetCanceled applies equality check predicate on the "user_set_canceled" field. It's identical to UserSetCanceledEQ.
func UserSetCanceled(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserSetCanceled), v))
	})
}

// AdminSetCanceled applies equality check predicate on the "admin_set_canceled" field. It's identical to AdminSetCanceledEQ.
func AdminSetCanceled(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdminSetCanceled), v))
	})
}

// PaymentState applies equality check predicate on the "payment_state" field. It's identical to PaymentStateEQ.
func PaymentState(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentState), v))
	})
}

// CancelState applies equality check predicate on the "cancel_state" field. It's identical to CancelStateEQ.
func CancelState(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCancelState), v))
	})
}

// CanceledAt applies equality check predicate on the "canceled_at" field. It's identical to CanceledAtEQ.
func CanceledAt(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCanceledAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderID), v))
	})
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderID), v))
	})
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...uuid.UUID) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderID), v...))
	})
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...uuid.UUID) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderID), v...))
	})
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderID), v))
	})
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderID), v))
	})
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderID), v))
	})
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderID), v))
	})
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderID)))
	})
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderID)))
	})
}

// PaymentIDEQ applies the EQ predicate on the "payment_id" field.
func PaymentIDEQ(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentID), v))
	})
}

// PaymentIDNEQ applies the NEQ predicate on the "payment_id" field.
func PaymentIDNEQ(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentID), v))
	})
}

// PaymentIDIn applies the In predicate on the "payment_id" field.
func PaymentIDIn(vs ...uuid.UUID) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaymentID), v...))
	})
}

// PaymentIDNotIn applies the NotIn predicate on the "payment_id" field.
func PaymentIDNotIn(vs ...uuid.UUID) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaymentID), v...))
	})
}

// PaymentIDGT applies the GT predicate on the "payment_id" field.
func PaymentIDGT(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentID), v))
	})
}

// PaymentIDGTE applies the GTE predicate on the "payment_id" field.
func PaymentIDGTE(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentID), v))
	})
}

// PaymentIDLT applies the LT predicate on the "payment_id" field.
func PaymentIDLT(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentID), v))
	})
}

// PaymentIDLTE applies the LTE predicate on the "payment_id" field.
func PaymentIDLTE(v uuid.UUID) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentID), v))
	})
}

// PaymentIDIsNil applies the IsNil predicate on the "payment_id" field.
func PaymentIDIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentID)))
	})
}

// PaymentIDNotNil applies the NotNil predicate on the "payment_id" field.
func PaymentIDNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentID)))
	})
}

// PaidAtEQ applies the EQ predicate on the "paid_at" field.
func PaidAtEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaidAt), v))
	})
}

// PaidAtNEQ applies the NEQ predicate on the "paid_at" field.
func PaidAtNEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaidAt), v))
	})
}

// PaidAtIn applies the In predicate on the "paid_at" field.
func PaidAtIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaidAt), v...))
	})
}

// PaidAtNotIn applies the NotIn predicate on the "paid_at" field.
func PaidAtNotIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaidAt), v...))
	})
}

// PaidAtGT applies the GT predicate on the "paid_at" field.
func PaidAtGT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaidAt), v))
	})
}

// PaidAtGTE applies the GTE predicate on the "paid_at" field.
func PaidAtGTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaidAt), v))
	})
}

// PaidAtLT applies the LT predicate on the "paid_at" field.
func PaidAtLT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaidAt), v))
	})
}

// PaidAtLTE applies the LTE predicate on the "paid_at" field.
func PaidAtLTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaidAt), v))
	})
}

// PaidAtIsNil applies the IsNil predicate on the "paid_at" field.
func PaidAtIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaidAt)))
	})
}

// PaidAtNotNil applies the NotNil predicate on the "paid_at" field.
func PaidAtNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaidAt)))
	})
}

// UserSetPaidEQ applies the EQ predicate on the "user_set_paid" field.
func UserSetPaidEQ(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserSetPaid), v))
	})
}

// UserSetPaidNEQ applies the NEQ predicate on the "user_set_paid" field.
func UserSetPaidNEQ(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserSetPaid), v))
	})
}

// UserSetPaidIsNil applies the IsNil predicate on the "user_set_paid" field.
func UserSetPaidIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserSetPaid)))
	})
}

// UserSetPaidNotNil applies the NotNil predicate on the "user_set_paid" field.
func UserSetPaidNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserSetPaid)))
	})
}

// UserSetCanceledEQ applies the EQ predicate on the "user_set_canceled" field.
func UserSetCanceledEQ(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserSetCanceled), v))
	})
}

// UserSetCanceledNEQ applies the NEQ predicate on the "user_set_canceled" field.
func UserSetCanceledNEQ(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserSetCanceled), v))
	})
}

// UserSetCanceledIsNil applies the IsNil predicate on the "user_set_canceled" field.
func UserSetCanceledIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserSetCanceled)))
	})
}

// UserSetCanceledNotNil applies the NotNil predicate on the "user_set_canceled" field.
func UserSetCanceledNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserSetCanceled)))
	})
}

// AdminSetCanceledEQ applies the EQ predicate on the "admin_set_canceled" field.
func AdminSetCanceledEQ(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdminSetCanceled), v))
	})
}

// AdminSetCanceledNEQ applies the NEQ predicate on the "admin_set_canceled" field.
func AdminSetCanceledNEQ(v bool) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdminSetCanceled), v))
	})
}

// AdminSetCanceledIsNil applies the IsNil predicate on the "admin_set_canceled" field.
func AdminSetCanceledIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAdminSetCanceled)))
	})
}

// AdminSetCanceledNotNil applies the NotNil predicate on the "admin_set_canceled" field.
func AdminSetCanceledNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAdminSetCanceled)))
	})
}

// PaymentStateEQ applies the EQ predicate on the "payment_state" field.
func PaymentStateEQ(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentState), v))
	})
}

// PaymentStateNEQ applies the NEQ predicate on the "payment_state" field.
func PaymentStateNEQ(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentState), v))
	})
}

// PaymentStateIn applies the In predicate on the "payment_state" field.
func PaymentStateIn(vs ...string) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPaymentState), v...))
	})
}

// PaymentStateNotIn applies the NotIn predicate on the "payment_state" field.
func PaymentStateNotIn(vs ...string) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPaymentState), v...))
	})
}

// PaymentStateGT applies the GT predicate on the "payment_state" field.
func PaymentStateGT(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentState), v))
	})
}

// PaymentStateGTE applies the GTE predicate on the "payment_state" field.
func PaymentStateGTE(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentState), v))
	})
}

// PaymentStateLT applies the LT predicate on the "payment_state" field.
func PaymentStateLT(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentState), v))
	})
}

// PaymentStateLTE applies the LTE predicate on the "payment_state" field.
func PaymentStateLTE(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentState), v))
	})
}

// PaymentStateContains applies the Contains predicate on the "payment_state" field.
func PaymentStateContains(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaymentState), v))
	})
}

// PaymentStateHasPrefix applies the HasPrefix predicate on the "payment_state" field.
func PaymentStateHasPrefix(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaymentState), v))
	})
}

// PaymentStateHasSuffix applies the HasSuffix predicate on the "payment_state" field.
func PaymentStateHasSuffix(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaymentState), v))
	})
}

// PaymentStateIsNil applies the IsNil predicate on the "payment_state" field.
func PaymentStateIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentState)))
	})
}

// PaymentStateNotNil applies the NotNil predicate on the "payment_state" field.
func PaymentStateNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentState)))
	})
}

// PaymentStateEqualFold applies the EqualFold predicate on the "payment_state" field.
func PaymentStateEqualFold(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaymentState), v))
	})
}

// PaymentStateContainsFold applies the ContainsFold predicate on the "payment_state" field.
func PaymentStateContainsFold(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaymentState), v))
	})
}

// CancelStateEQ applies the EQ predicate on the "cancel_state" field.
func CancelStateEQ(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCancelState), v))
	})
}

// CancelStateNEQ applies the NEQ predicate on the "cancel_state" field.
func CancelStateNEQ(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCancelState), v))
	})
}

// CancelStateIn applies the In predicate on the "cancel_state" field.
func CancelStateIn(vs ...string) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCancelState), v...))
	})
}

// CancelStateNotIn applies the NotIn predicate on the "cancel_state" field.
func CancelStateNotIn(vs ...string) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCancelState), v...))
	})
}

// CancelStateGT applies the GT predicate on the "cancel_state" field.
func CancelStateGT(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCancelState), v))
	})
}

// CancelStateGTE applies the GTE predicate on the "cancel_state" field.
func CancelStateGTE(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCancelState), v))
	})
}

// CancelStateLT applies the LT predicate on the "cancel_state" field.
func CancelStateLT(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCancelState), v))
	})
}

// CancelStateLTE applies the LTE predicate on the "cancel_state" field.
func CancelStateLTE(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCancelState), v))
	})
}

// CancelStateContains applies the Contains predicate on the "cancel_state" field.
func CancelStateContains(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCancelState), v))
	})
}

// CancelStateHasPrefix applies the HasPrefix predicate on the "cancel_state" field.
func CancelStateHasPrefix(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCancelState), v))
	})
}

// CancelStateHasSuffix applies the HasSuffix predicate on the "cancel_state" field.
func CancelStateHasSuffix(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCancelState), v))
	})
}

// CancelStateIsNil applies the IsNil predicate on the "cancel_state" field.
func CancelStateIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCancelState)))
	})
}

// CancelStateNotNil applies the NotNil predicate on the "cancel_state" field.
func CancelStateNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCancelState)))
	})
}

// CancelStateEqualFold applies the EqualFold predicate on the "cancel_state" field.
func CancelStateEqualFold(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCancelState), v))
	})
}

// CancelStateContainsFold applies the ContainsFold predicate on the "cancel_state" field.
func CancelStateContainsFold(v string) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCancelState), v))
	})
}

// CanceledAtEQ applies the EQ predicate on the "canceled_at" field.
func CanceledAtEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCanceledAt), v))
	})
}

// CanceledAtNEQ applies the NEQ predicate on the "canceled_at" field.
func CanceledAtNEQ(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCanceledAt), v))
	})
}

// CanceledAtIn applies the In predicate on the "canceled_at" field.
func CanceledAtIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCanceledAt), v...))
	})
}

// CanceledAtNotIn applies the NotIn predicate on the "canceled_at" field.
func CanceledAtNotIn(vs ...uint32) predicate.FeeOrderState {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCanceledAt), v...))
	})
}

// CanceledAtGT applies the GT predicate on the "canceled_at" field.
func CanceledAtGT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCanceledAt), v))
	})
}

// CanceledAtGTE applies the GTE predicate on the "canceled_at" field.
func CanceledAtGTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCanceledAt), v))
	})
}

// CanceledAtLT applies the LT predicate on the "canceled_at" field.
func CanceledAtLT(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCanceledAt), v))
	})
}

// CanceledAtLTE applies the LTE predicate on the "canceled_at" field.
func CanceledAtLTE(v uint32) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCanceledAt), v))
	})
}

// CanceledAtIsNil applies the IsNil predicate on the "canceled_at" field.
func CanceledAtIsNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCanceledAt)))
	})
}

// CanceledAtNotNil applies the NotNil predicate on the "canceled_at" field.
func CanceledAtNotNil() predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCanceledAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FeeOrderState) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FeeOrderState) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
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
func Not(p predicate.FeeOrderState) predicate.FeeOrderState {
	return predicate.FeeOrderState(func(s *sql.Selector) {
		p(s.Not())
	})
}
