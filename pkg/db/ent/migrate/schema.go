// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompensatesColumns holds the columns for the "compensates" table.
	CompensatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "compensate_type", Type: field.TypeString, Nullable: true, Default: "DefaultCompensateType"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// CompensatesTable holds the schema information for the "compensates" table.
	CompensatesTable = &schema.Table{
		Name:       "compensates",
		Columns:    CompensatesColumns,
		PrimaryKey: []*schema.Column{CompensatesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "compensate_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CompensatesColumns[4]},
			},
			{
				Name:    "compensate_order_id",
				Unique:  false,
				Columns: []*schema.Column{CompensatesColumns[5]},
			},
		},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "app_good_id", Type: field.TypeUUID},
		{Name: "payment_id", Type: field.TypeUUID, Nullable: true},
		{Name: "parent_order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "units_v1", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "good_value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "good_value_usd", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "payment_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "discount_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "promotion_id", Type: field.TypeUUID, Nullable: true},
		{Name: "duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "order_type", Type: field.TypeString, Nullable: true, Default: "Normal"},
		{Name: "investment_type", Type: field.TypeString, Nullable: true, Default: "FullPayment"},
		{Name: "coupon_ids", Type: field.TypeJSON, Nullable: true},
		{Name: "payment_type", Type: field.TypeString, Nullable: true, Default: "PayWithBalanceOnly"},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "payment_coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "transfer_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "balance_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "local_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "live_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "create_method", Type: field.TypeString, Nullable: true, Default: "OrderCreatedByPurchase"},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "order_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrdersColumns[4]},
			},
			{
				Name:    "order_app_id_user_id_good_id_app_good_id",
				Unique:  false,
				Columns: []*schema.Column{OrdersColumns[5], OrdersColumns[6], OrdersColumns[7], OrdersColumns[8]},
			},
		},
	}
	// OrderLocksColumns holds the columns for the "order_locks" table.
	OrderLocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "lock_type", Type: field.TypeString, Nullable: true, Default: "DefaultOrderLockType"},
	}
	// OrderLocksTable holds the schema information for the "order_locks" table.
	OrderLocksTable = &schema.Table{
		Name:       "order_locks",
		Columns:    OrderLocksColumns,
		PrimaryKey: []*schema.Column{OrderLocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderlock_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderLocksColumns[4]},
			},
			{
				Name:    "orderlock_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderLocksColumns[7]},
			},
		},
	}
	// OrderStatesColumns holds the columns for the "order_states" table.
	OrderStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "order_state", Type: field.TypeString, Nullable: true, Default: "OrderStateCreated"},
		{Name: "cancel_state", Type: field.TypeString, Nullable: true, Default: "DefaultOrderState"},
		{Name: "start_mode", Type: field.TypeString, Nullable: true, Default: "OrderStartConfirmed"},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "paid_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "last_benefit_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "benefit_state", Type: field.TypeString, Nullable: true, Default: "BenefitWait"},
		{Name: "user_set_paid", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "user_set_canceled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "admin_set_canceled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_transaction_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "payment_finish_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "payment_state", Type: field.TypeString, Nullable: true, Default: "PaymentStateWait"},
		{Name: "outofgas_hours", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "compensate_hours", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "renew_state", Type: field.TypeString, Nullable: true, Default: "OrderRenewWait"},
		{Name: "renew_notify_at", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// OrderStatesTable holds the schema information for the "order_states" table.
	OrderStatesTable = &schema.Table{
		Name:       "order_states",
		Columns:    OrderStatesColumns,
		PrimaryKey: []*schema.Column{OrderStatesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderstate_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderStatesColumns[4]},
			},
			{
				Name:    "orderstate_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderStatesColumns[5]},
			},
		},
	}
	// OutOfGasColumns holds the columns for the "out_of_gas" table.
	OutOfGasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// OutOfGasTable holds the schema information for the "out_of_gas" table.
	OutOfGasTable = &schema.Table{
		Name:       "out_of_gas",
		Columns:    OutOfGasColumns,
		PrimaryKey: []*schema.Column{OutOfGasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "outofgas_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OutOfGasColumns[4]},
			},
		},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "order_id", Type: field.TypeUUID},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_info_id", Type: field.TypeUUID, Nullable: true},
		{Name: "start_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "multi_payment_coins", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_amounts", Type: field.TypeJSON, Nullable: true},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:       "payments",
		Columns:    PaymentsColumns,
		PrimaryKey: []*schema.Column{PaymentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "payment_ent_id",
				Unique:  true,
				Columns: []*schema.Column{PaymentsColumns[4]},
			},
			{
				Name:    "payment_order_id",
				Unique:  false,
				Columns: []*schema.Column{PaymentsColumns[8]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompensatesTable,
		OrdersTable,
		OrderLocksTable,
		OrderStatesTable,
		OutOfGasTable,
		PaymentsTable,
	}
)

func init() {
}
