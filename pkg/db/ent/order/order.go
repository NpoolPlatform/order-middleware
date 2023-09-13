// Code generated by ent, DO NOT EDIT.

package order

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldPaymentID holds the string denoting the payment_id field in the database.
	FieldPaymentID = "payment_id"
	// FieldParentOrderID holds the string denoting the parent_order_id field in the database.
	FieldParentOrderID = "parent_order_id"
	// FieldUnitsV1 holds the string denoting the units_v1 field in the database.
	FieldUnitsV1 = "units_v1"
	// FieldGoodValue holds the string denoting the good_value field in the database.
	FieldGoodValue = "good_value"
	// FieldGoodValueUsd holds the string denoting the good_value_usd field in the database.
	FieldGoodValueUsd = "good_value_usd"
	// FieldPaymentAmount holds the string denoting the payment_amount field in the database.
	FieldPaymentAmount = "payment_amount"
	// FieldDiscountAmount holds the string denoting the discount_amount field in the database.
	FieldDiscountAmount = "discount_amount"
	// FieldPromotionID holds the string denoting the promotion_id field in the database.
	FieldPromotionID = "promotion_id"
	// FieldDurationDays holds the string denoting the duration_days field in the database.
	FieldDurationDays = "duration_days"
	// FieldOrderType holds the string denoting the order_type field in the database.
	FieldOrderType = "order_type"
	// FieldInvestmentType holds the string denoting the investment_type field in the database.
	FieldInvestmentType = "investment_type"
	// FieldCouponIds holds the string denoting the coupon_ids field in the database.
	FieldCouponIds = "coupon_ids"
	// FieldPaymentType holds the string denoting the payment_type field in the database.
	FieldPaymentType = "payment_type"
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldPaymentCoinTypeID holds the string denoting the payment_coin_type_id field in the database.
	FieldPaymentCoinTypeID = "payment_coin_type_id"
	// FieldTransferAmount holds the string denoting the transfer_amount field in the database.
	FieldTransferAmount = "transfer_amount"
	// FieldBalanceAmount holds the string denoting the balance_amount field in the database.
	FieldBalanceAmount = "balance_amount"
	// FieldCoinUsdCurrency holds the string denoting the coin_usd_currency field in the database.
	FieldCoinUsdCurrency = "coin_usd_currency"
	// FieldLocalCoinUsdCurrency holds the string denoting the local_coin_usd_currency field in the database.
	FieldLocalCoinUsdCurrency = "local_coin_usd_currency"
	// FieldLiveCoinUsdCurrency holds the string denoting the live_coin_usd_currency field in the database.
	FieldLiveCoinUsdCurrency = "live_coin_usd_currency"
	// Table holds the table name of the order in the database.
	Table = "orders"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldAppGoodID,
	FieldPaymentID,
	FieldParentOrderID,
	FieldUnitsV1,
	FieldGoodValue,
	FieldGoodValueUsd,
	FieldPaymentAmount,
	FieldDiscountAmount,
	FieldPromotionID,
	FieldDurationDays,
	FieldOrderType,
	FieldInvestmentType,
	FieldCouponIds,
	FieldPaymentType,
	FieldCoinTypeID,
	FieldPaymentCoinTypeID,
	FieldTransferAmount,
	FieldBalanceAmount,
	FieldCoinUsdCurrency,
	FieldLocalCoinUsdCurrency,
	FieldLiveCoinUsdCurrency,
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
	// DefaultPaymentID holds the default value on creation for the "payment_id" field.
	DefaultPaymentID func() uuid.UUID
	// DefaultParentOrderID holds the default value on creation for the "parent_order_id" field.
	DefaultParentOrderID func() uuid.UUID
	// DefaultUnitsV1 holds the default value on creation for the "units_v1" field.
	DefaultUnitsV1 decimal.Decimal
	// DefaultGoodValue holds the default value on creation for the "good_value" field.
	DefaultGoodValue decimal.Decimal
	// DefaultGoodValueUsd holds the default value on creation for the "good_value_usd" field.
	DefaultGoodValueUsd decimal.Decimal
	// DefaultPaymentAmount holds the default value on creation for the "payment_amount" field.
	DefaultPaymentAmount decimal.Decimal
	// DefaultDiscountAmount holds the default value on creation for the "discount_amount" field.
	DefaultDiscountAmount decimal.Decimal
	// DefaultPromotionID holds the default value on creation for the "promotion_id" field.
	DefaultPromotionID func() uuid.UUID
	// DefaultDurationDays holds the default value on creation for the "duration_days" field.
	DefaultDurationDays uint32
	// DefaultOrderType holds the default value on creation for the "order_type" field.
	DefaultOrderType string
	// DefaultInvestmentType holds the default value on creation for the "investment_type" field.
	DefaultInvestmentType string
	// DefaultCouponIds holds the default value on creation for the "coupon_ids" field.
	DefaultCouponIds func() []uuid.UUID
	// DefaultPaymentType holds the default value on creation for the "payment_type" field.
	DefaultPaymentType string
	// DefaultTransferAmount holds the default value on creation for the "transfer_amount" field.
	DefaultTransferAmount decimal.Decimal
	// DefaultBalanceAmount holds the default value on creation for the "balance_amount" field.
	DefaultBalanceAmount decimal.Decimal
	// DefaultCoinUsdCurrency holds the default value on creation for the "coin_usd_currency" field.
	DefaultCoinUsdCurrency decimal.Decimal
	// DefaultLocalCoinUsdCurrency holds the default value on creation for the "local_coin_usd_currency" field.
	DefaultLocalCoinUsdCurrency decimal.Decimal
	// DefaultLiveCoinUsdCurrency holds the default value on creation for the "live_coin_usd_currency" field.
	DefaultLiveCoinUsdCurrency decimal.Decimal
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
