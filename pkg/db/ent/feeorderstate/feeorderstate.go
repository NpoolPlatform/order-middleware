// Code generated by ent, DO NOT EDIT.

package feeorderstate

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the feeorderstate type in the database.
	Label = "fee_order_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldPaymentID holds the string denoting the payment_id field in the database.
	FieldPaymentID = "payment_id"
	// FieldPaidAt holds the string denoting the paid_at field in the database.
	FieldPaidAt = "paid_at"
	// FieldUserSetPaid holds the string denoting the user_set_paid field in the database.
	FieldUserSetPaid = "user_set_paid"
	// FieldUserSetCanceled holds the string denoting the user_set_canceled field in the database.
	FieldUserSetCanceled = "user_set_canceled"
	// FieldAdminSetCanceled holds the string denoting the admin_set_canceled field in the database.
	FieldAdminSetCanceled = "admin_set_canceled"
	// FieldPaymentState holds the string denoting the payment_state field in the database.
	FieldPaymentState = "payment_state"
	// FieldCancelState holds the string denoting the cancel_state field in the database.
	FieldCancelState = "cancel_state"
	// FieldCanceledAt holds the string denoting the canceled_at field in the database.
	FieldCanceledAt = "canceled_at"
	// Table holds the table name of the feeorderstate in the database.
	Table = "fee_order_states"
)

// Columns holds all SQL columns for feeorderstate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldOrderID,
	FieldPaymentID,
	FieldPaidAt,
	FieldUserSetPaid,
	FieldUserSetCanceled,
	FieldAdminSetCanceled,
	FieldPaymentState,
	FieldCancelState,
	FieldCanceledAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/order-middleware/pkg/db/ent/runtime"
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultPaymentID holds the default value on creation for the "payment_id" field.
	DefaultPaymentID func() uuid.UUID
	// DefaultPaidAt holds the default value on creation for the "paid_at" field.
	DefaultPaidAt uint32
	// DefaultUserSetPaid holds the default value on creation for the "user_set_paid" field.
	DefaultUserSetPaid bool
	// DefaultUserSetCanceled holds the default value on creation for the "user_set_canceled" field.
	DefaultUserSetCanceled bool
	// DefaultAdminSetCanceled holds the default value on creation for the "admin_set_canceled" field.
	DefaultAdminSetCanceled bool
	// DefaultPaymentState holds the default value on creation for the "payment_state" field.
	DefaultPaymentState string
	// DefaultCancelState holds the default value on creation for the "cancel_state" field.
	DefaultCancelState string
	// DefaultCanceledAt holds the default value on creation for the "canceled_at" field.
	DefaultCanceledAt uint32
)
