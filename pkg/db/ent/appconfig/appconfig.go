// Code generated by ent, DO NOT EDIT.

package appconfig

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the appconfig type in the database.
	Label = "app_config"
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
	// FieldSimulateOrderCouponMode holds the string denoting the simulate_order_coupon_mode field in the database.
	FieldSimulateOrderCouponMode = "simulate_order_coupon_mode"
	// FieldSimulateOrderCouponProbability holds the string denoting the simulate_order_coupon_probability field in the database.
	FieldSimulateOrderCouponProbability = "simulate_order_coupon_probability"
	// FieldSimulateOrderCashableProfitProbability holds the string denoting the simulate_order_cashable_profit_probability field in the database.
	FieldSimulateOrderCashableProfitProbability = "simulate_order_cashable_profit_probability"
	// FieldEnableSimulateOrder holds the string denoting the enable_simulate_order field in the database.
	FieldEnableSimulateOrder = "enable_simulate_order"
	// FieldMaxUnpaidOrders holds the string denoting the max_unpaid_orders field in the database.
	FieldMaxUnpaidOrders = "max_unpaid_orders"
	// Table holds the table name of the appconfig in the database.
	Table = "app_configs"
)

// Columns holds all SQL columns for appconfig fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldSimulateOrderCouponMode,
	FieldSimulateOrderCouponProbability,
	FieldSimulateOrderCashableProfitProbability,
	FieldEnableSimulateOrder,
	FieldMaxUnpaidOrders,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultSimulateOrderCouponMode holds the default value on creation for the "simulate_order_coupon_mode" field.
	DefaultSimulateOrderCouponMode string
	// DefaultSimulateOrderCouponProbability holds the default value on creation for the "simulate_order_coupon_probability" field.
	DefaultSimulateOrderCouponProbability decimal.Decimal
	// DefaultSimulateOrderCashableProfitProbability holds the default value on creation for the "simulate_order_cashable_profit_probability" field.
	DefaultSimulateOrderCashableProfitProbability decimal.Decimal
	// DefaultEnableSimulateOrder holds the default value on creation for the "enable_simulate_order" field.
	DefaultEnableSimulateOrder bool
	// DefaultMaxUnpaidOrders holds the default value on creation for the "max_unpaid_orders" field.
	DefaultMaxUnpaidOrders uint32
)
