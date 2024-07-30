// Code generated by ent, DO NOT EDIT.

package orderbase

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the orderbase type in the database.
	Label = "order_base"
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
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldGoodType holds the string denoting the good_type field in the database.
	FieldGoodType = "good_type"
	// FieldParentOrderID holds the string denoting the parent_order_id field in the database.
	FieldParentOrderID = "parent_order_id"
	// FieldOrderType holds the string denoting the order_type field in the database.
	FieldOrderType = "order_type"
	// FieldCreateMethod holds the string denoting the create_method field in the database.
	FieldCreateMethod = "create_method"
	// FieldSimulate holds the string denoting the simulate field in the database.
	FieldSimulate = "simulate"
	// Table holds the table name of the orderbase in the database.
	Table = "order_bases"
)

// Columns holds all SQL columns for orderbase fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldUserID,
	FieldGoodID,
	FieldAppGoodID,
	FieldGoodType,
	FieldParentOrderID,
	FieldOrderType,
	FieldCreateMethod,
	FieldSimulate,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultGoodType holds the default value on creation for the "good_type" field.
	DefaultGoodType string
	// DefaultParentOrderID holds the default value on creation for the "parent_order_id" field.
	DefaultParentOrderID func() uuid.UUID
	// DefaultOrderType holds the default value on creation for the "order_type" field.
	DefaultOrderType string
	// DefaultCreateMethod holds the default value on creation for the "create_method" field.
	DefaultCreateMethod string
	// DefaultSimulate holds the default value on creation for the "simulate" field.
	DefaultSimulate bool
)
