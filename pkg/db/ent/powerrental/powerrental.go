// Code generated by ent, DO NOT EDIT.

package powerrental

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the powerrental type in the database.
	Label = "power_rental"
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
	// FieldAppGoodStockID holds the string denoting the app_good_stock_id field in the database.
	FieldAppGoodStockID = "app_good_stock_id"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// FieldGoodValueUsd holds the string denoting the good_value_usd field in the database.
	FieldGoodValueUsd = "good_value_usd"
	// FieldPaymentAmountUsd holds the string denoting the payment_amount_usd field in the database.
	FieldPaymentAmountUsd = "payment_amount_usd"
	// FieldDiscountAmountUsd holds the string denoting the discount_amount_usd field in the database.
	FieldDiscountAmountUsd = "discount_amount_usd"
	// FieldPromotionID holds the string denoting the promotion_id field in the database.
	FieldPromotionID = "promotion_id"
	// FieldInvestmentType holds the string denoting the investment_type field in the database.
	FieldInvestmentType = "investment_type"
	// FieldDurationSeconds holds the string denoting the duration_seconds field in the database.
	FieldDurationSeconds = "duration_seconds"
	// Table holds the table name of the powerrental in the database.
	Table = "power_rentals"
)

// Columns holds all SQL columns for powerrental fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldOrderID,
	FieldAppGoodStockID,
	FieldUnits,
	FieldGoodValueUsd,
	FieldPaymentAmountUsd,
	FieldDiscountAmountUsd,
	FieldPromotionID,
	FieldInvestmentType,
	FieldDurationSeconds,
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
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultAppGoodStockID holds the default value on creation for the "app_good_stock_id" field.
	DefaultAppGoodStockID func() uuid.UUID
	// DefaultUnits holds the default value on creation for the "units" field.
	DefaultUnits decimal.Decimal
	// DefaultGoodValueUsd holds the default value on creation for the "good_value_usd" field.
	DefaultGoodValueUsd decimal.Decimal
	// DefaultPaymentAmountUsd holds the default value on creation for the "payment_amount_usd" field.
	DefaultPaymentAmountUsd decimal.Decimal
	// DefaultDiscountAmountUsd holds the default value on creation for the "discount_amount_usd" field.
	DefaultDiscountAmountUsd decimal.Decimal
	// DefaultPromotionID holds the default value on creation for the "promotion_id" field.
	DefaultPromotionID func() uuid.UUID
	// DefaultInvestmentType holds the default value on creation for the "investment_type" field.
	DefaultInvestmentType string
	// DefaultDurationSeconds holds the default value on creation for the "duration_seconds" field.
	DefaultDurationSeconds uint32
)
