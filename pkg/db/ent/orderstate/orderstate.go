// Code generated by ent, DO NOT EDIT.

package orderstate

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the orderstate type in the database.
	Label = "order_state"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldOrderState holds the string denoting the order_state field in the database.
	FieldOrderState = "order_state"
	// FieldCancelState holds the string denoting the cancel_state field in the database.
	FieldCancelState = "cancel_state"
	// FieldStartMode holds the string denoting the start_mode field in the database.
	FieldStartMode = "start_mode"
	// FieldStartAt holds the string denoting the start_at field in the database.
	FieldStartAt = "start_at"
	// FieldEndAt holds the string denoting the end_at field in the database.
	FieldEndAt = "end_at"
	// FieldLastBenefitAt holds the string denoting the last_benefit_at field in the database.
	FieldLastBenefitAt = "last_benefit_at"
	// FieldBenefitState holds the string denoting the benefit_state field in the database.
	FieldBenefitState = "benefit_state"
	// FieldUserSetPaid holds the string denoting the user_set_paid field in the database.
	FieldUserSetPaid = "user_set_paid"
	// FieldUserSetCanceled holds the string denoting the user_set_canceled field in the database.
	FieldUserSetCanceled = "user_set_canceled"
	// FieldAdminSetCanceled holds the string denoting the admin_set_canceled field in the database.
	FieldAdminSetCanceled = "admin_set_canceled"
	// FieldPaymentTransactionID holds the string denoting the payment_transaction_id field in the database.
	FieldPaymentTransactionID = "payment_transaction_id"
	// FieldPaymentFinishAmount holds the string denoting the payment_finish_amount field in the database.
	FieldPaymentFinishAmount = "payment_finish_amount"
	// FieldPaymentState holds the string denoting the payment_state field in the database.
	FieldPaymentState = "payment_state"
	// FieldOutofgasHours holds the string denoting the outofgas_hours field in the database.
	FieldOutofgasHours = "outofgas_hours"
	// FieldCompensateHours holds the string denoting the compensate_hours field in the database.
	FieldCompensateHours = "compensate_hours"
	// Table holds the table name of the orderstate in the database.
	Table = "order_states"
)

// Columns holds all SQL columns for orderstate fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldOrderID,
	FieldOrderState,
	FieldCancelState,
	FieldStartMode,
	FieldStartAt,
	FieldEndAt,
	FieldLastBenefitAt,
	FieldBenefitState,
	FieldUserSetPaid,
	FieldUserSetCanceled,
	FieldAdminSetCanceled,
	FieldPaymentTransactionID,
	FieldPaymentFinishAmount,
	FieldPaymentState,
	FieldOutofgasHours,
	FieldCompensateHours,
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
//
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
	// DefaultOrderState holds the default value on creation for the "order_state" field.
	DefaultOrderState string
	// DefaultCancelState holds the default value on creation for the "cancel_state" field.
	DefaultCancelState string
	// DefaultStartMode holds the default value on creation for the "start_mode" field.
	DefaultStartMode string
	// DefaultStartAt holds the default value on creation for the "start_at" field.
	DefaultStartAt uint32
	// DefaultEndAt holds the default value on creation for the "end_at" field.
	DefaultEndAt uint32
	// DefaultLastBenefitAt holds the default value on creation for the "last_benefit_at" field.
	DefaultLastBenefitAt uint32
	// DefaultBenefitState holds the default value on creation for the "benefit_state" field.
	DefaultBenefitState string
	// DefaultUserSetPaid holds the default value on creation for the "user_set_paid" field.
	DefaultUserSetPaid bool
	// DefaultUserSetCanceled holds the default value on creation for the "user_set_canceled" field.
	DefaultUserSetCanceled bool
	// DefaultAdminSetCanceled holds the default value on creation for the "admin_set_canceled" field.
	DefaultAdminSetCanceled bool
	// DefaultPaymentTransactionID holds the default value on creation for the "payment_transaction_id" field.
	DefaultPaymentTransactionID string
	// DefaultPaymentFinishAmount holds the default value on creation for the "payment_finish_amount" field.
	DefaultPaymentFinishAmount decimal.Decimal
	// DefaultPaymentState holds the default value on creation for the "payment_state" field.
	DefaultPaymentState string
	// DefaultOutofgasHours holds the default value on creation for the "outofgas_hours" field.
	DefaultOutofgasHours uint32
	// DefaultCompensateHours holds the default value on creation for the "compensate_hours" field.
	DefaultCompensateHours uint32
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
