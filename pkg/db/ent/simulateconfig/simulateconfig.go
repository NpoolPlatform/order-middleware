// Code generated by ent, DO NOT EDIT.

package simulateconfig

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the simulateconfig type in the database.
	Label = "simulate_config"
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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldSendCouponMode holds the string denoting the send_coupon_mode field in the database.
	FieldSendCouponMode = "send_coupon_mode"
	// FieldSendCouponProbability holds the string denoting the send_coupon_probability field in the database.
	FieldSendCouponProbability = "send_coupon_probability"
	// FieldCashableProfitProbability holds the string denoting the cashable_profit_probability field in the database.
	FieldCashableProfitProbability = "cashable_profit_probability"
	// FieldEnabled holds the string denoting the enabled field in the database.
	FieldEnabled = "enabled"
	// Table holds the table name of the simulateconfig in the database.
	Table = "simulate_configs"
)

// Columns holds all SQL columns for simulateconfig fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldSendCouponMode,
	FieldSendCouponProbability,
	FieldCashableProfitProbability,
	FieldEnabled,
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
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultSendCouponMode holds the default value on creation for the "send_coupon_mode" field.
	DefaultSendCouponMode string
	// DefaultSendCouponProbability holds the default value on creation for the "send_coupon_probability" field.
	DefaultSendCouponProbability decimal.Decimal
	// DefaultCashableProfitProbability holds the default value on creation for the "cashable_profit_probability" field.
	DefaultCashableProfitProbability decimal.Decimal
	// DefaultEnabled holds the default value on creation for the "enabled" field.
	DefaultEnabled bool
)
