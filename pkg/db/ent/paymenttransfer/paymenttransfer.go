// Code generated by ent, DO NOT EDIT.

package paymenttransfer

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the paymenttransfer type in the database.
	Label = "payment_transfer"
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
	// FieldCoinTypeID holds the string denoting the coin_type_id field in the database.
	FieldCoinTypeID = "coin_type_id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldStartAmount holds the string denoting the start_amount field in the database.
	FieldStartAmount = "start_amount"
	// FieldTransactionID holds the string denoting the transaction_id field in the database.
	FieldTransactionID = "transaction_id"
	// FieldFinishAmount holds the string denoting the finish_amount field in the database.
	FieldFinishAmount = "finish_amount"
	// FieldCoinUsdCurrency holds the string denoting the coin_usd_currency field in the database.
	FieldCoinUsdCurrency = "coin_usd_currency"
	// FieldLocalCoinUsdCurrency holds the string denoting the local_coin_usd_currency field in the database.
	FieldLocalCoinUsdCurrency = "local_coin_usd_currency"
	// FieldLiveCoinUsdCurrency holds the string denoting the live_coin_usd_currency field in the database.
	FieldLiveCoinUsdCurrency = "live_coin_usd_currency"
	// Table holds the table name of the paymenttransfer in the database.
	Table = "payment_transfers"
)

// Columns holds all SQL columns for paymenttransfer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldOrderID,
	FieldCoinTypeID,
	FieldAccountID,
	FieldAmount,
	FieldStartAmount,
	FieldTransactionID,
	FieldFinishAmount,
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
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultOrderID holds the default value on creation for the "order_id" field.
	DefaultOrderID func() uuid.UUID
	// DefaultCoinTypeID holds the default value on creation for the "coin_type_id" field.
	DefaultCoinTypeID func() uuid.UUID
	// DefaultAccountID holds the default value on creation for the "account_id" field.
	DefaultAccountID func() uuid.UUID
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount decimal.Decimal
	// DefaultStartAmount holds the default value on creation for the "start_amount" field.
	DefaultStartAmount decimal.Decimal
	// DefaultTransactionID holds the default value on creation for the "transaction_id" field.
	DefaultTransactionID string
	// DefaultFinishAmount holds the default value on creation for the "finish_amount" field.
	DefaultFinishAmount decimal.Decimal
	// DefaultCoinUsdCurrency holds the default value on creation for the "coin_usd_currency" field.
	DefaultCoinUsdCurrency decimal.Decimal
	// DefaultLocalCoinUsdCurrency holds the default value on creation for the "local_coin_usd_currency" field.
	DefaultLocalCoinUsdCurrency decimal.Decimal
	// DefaultLiveCoinUsdCurrency holds the default value on creation for the "live_coin_usd_currency" field.
	DefaultLiveCoinUsdCurrency decimal.Decimal
)
