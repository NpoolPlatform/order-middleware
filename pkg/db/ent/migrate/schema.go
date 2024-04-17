// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppConfigsColumns holds the columns for the "app_configs" table.
	AppConfigsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "simulate_order_coupon_mode", Type: field.TypeString, Nullable: true, Default: "WithoutCoupon"},
		{Name: "simulate_order_coupon_probability", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "simulate_order_cashable_profit_probability", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "enable_simulate_order", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "max_unpaid_orders", Type: field.TypeUint32, Nullable: true, Default: 5},
	}
	// AppConfigsTable holds the schema information for the "app_configs" table.
	AppConfigsTable = &schema.Table{
		Name:       "app_configs",
		Columns:    AppConfigsColumns,
		PrimaryKey: []*schema.Column{AppConfigsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appconfig_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppConfigsColumns[4]},
			},
		},
	}
	// CompensatesColumns holds the columns for the "compensates" table.
	CompensatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
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
		{Name: "simulate", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "create_method", Type: field.TypeString, Nullable: true, Default: "OrderCreatedByPurchase"},
		{Name: "multi_payment_coins", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_amounts", Type: field.TypeJSON, Nullable: true},
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
	// OrderBasesColumns holds the columns for the "order_bases" table.
	OrderBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "parent_order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_type", Type: field.TypeString, Nullable: true, Default: "Normal"},
		{Name: "payment_type", Type: field.TypeString, Nullable: true, Default: "PayWithBalanceOnly"},
		{Name: "create_method", Type: field.TypeString, Nullable: true, Default: "OrderCreatedByPurchase"},
	}
	// OrderBasesTable holds the schema information for the "order_bases" table.
	OrderBasesTable = &schema.Table{
		Name:       "order_bases",
		Columns:    OrderBasesColumns,
		PrimaryKey: []*schema.Column{OrderBasesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderbase_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderBasesColumns[4]},
			},
			{
				Name:    "orderbase_app_id_user_id_good_id_app_good_id",
				Unique:  false,
				Columns: []*schema.Column{OrderBasesColumns[5], OrderBasesColumns[6], OrderBasesColumns[7], OrderBasesColumns[8]},
			},
		},
	}
	// OrderCouponsColumns holds the columns for the "order_coupons" table.
	OrderCouponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coupon_id", Type: field.TypeUUID, Nullable: true},
	}
	// OrderCouponsTable holds the schema information for the "order_coupons" table.
	OrderCouponsTable = &schema.Table{
		Name:       "order_coupons",
		Columns:    OrderCouponsColumns,
		PrimaryKey: []*schema.Column{OrderCouponsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "ordercoupon_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderCouponsColumns[4]},
			},
			{
				Name:    "ordercoupon_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderCouponsColumns[5]},
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
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
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
				Columns: []*schema.Column{OrderLocksColumns[5]},
			},
		},
	}
	// OrderPaymentBalancesColumns holds the columns for the "order_payment_balances" table.
	OrderPaymentBalancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "local_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
		{Name: "live_coin_usd_currency", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37, 18)"}},
	}
	// OrderPaymentBalancesTable holds the schema information for the "order_payment_balances" table.
	OrderPaymentBalancesTable = &schema.Table{
		Name:       "order_payment_balances",
		Columns:    OrderPaymentBalancesColumns,
		PrimaryKey: []*schema.Column{OrderPaymentBalancesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderpaymentbalance_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderPaymentBalancesColumns[4]},
			},
			{
				Name:    "orderpaymentbalance_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderPaymentBalancesColumns[5]},
			},
		},
	}
	// OrderPaymentContractsColumns holds the columns for the "order_payment_contracts" table.
	OrderPaymentContractsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// OrderPaymentContractsTable holds the schema information for the "order_payment_contracts" table.
	OrderPaymentContractsTable = &schema.Table{
		Name:       "order_payment_contracts",
		Columns:    OrderPaymentContractsColumns,
		PrimaryKey: []*schema.Column{OrderPaymentContractsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderpaymentcontract_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderPaymentContractsColumns[4]},
			},
			{
				Name:    "orderpaymentcontract_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderPaymentContractsColumns[5]},
			},
		},
	}
	// OrderPaymentTransfersColumns holds the columns for the "order_payment_transfers" table.
	OrderPaymentTransfersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
		{Name: "start_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "payment_transaction_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "payment_finish_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// OrderPaymentTransfersTable holds the schema information for the "order_payment_transfers" table.
	OrderPaymentTransfersTable = &schema.Table{
		Name:       "order_payment_transfers",
		Columns:    OrderPaymentTransfersColumns,
		PrimaryKey: []*schema.Column{OrderPaymentTransfersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderpaymenttransfer_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderPaymentTransfersColumns[4]},
			},
			{
				Name:    "orderpaymenttransfer_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderPaymentTransfersColumns[5]},
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
	// OrderStateBasesColumns holds the columns for the "order_state_bases" table.
	OrderStateBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_state", Type: field.TypeString, Nullable: true, Default: "OrderStateCreated"},
		{Name: "start_mode", Type: field.TypeString, Nullable: true, Default: "OrderStartConfirmed"},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "last_benefit_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "benefit_state", Type: field.TypeString, Nullable: true, Default: "BenefitWait"},
	}
	// OrderStateBasesTable holds the schema information for the "order_state_bases" table.
	OrderStateBasesTable = &schema.Table{
		Name:       "order_state_bases",
		Columns:    OrderStateBasesColumns,
		PrimaryKey: []*schema.Column{OrderStateBasesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "orderstatebase_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OrderStateBasesColumns[4]},
			},
			{
				Name:    "orderstatebase_order_id",
				Unique:  false,
				Columns: []*schema.Column{OrderStateBasesColumns[5]},
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
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
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
	// PowerRentalsColumns holds the columns for the "power_rentals" table.
	PowerRentalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "units", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "good_value", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "good_value_usd", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "payment_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "discount_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "promotion_id", Type: field.TypeUUID, Nullable: true},
		{Name: "duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "investment_type", Type: field.TypeString, Nullable: true, Default: "FullPayment"},
		{Name: "simulate", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// PowerRentalsTable holds the schema information for the "power_rentals" table.
	PowerRentalsTable = &schema.Table{
		Name:       "power_rentals",
		Columns:    PowerRentalsColumns,
		PrimaryKey: []*schema.Column{PowerRentalsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "powerrental_ent_id",
				Unique:  true,
				Columns: []*schema.Column{PowerRentalsColumns[4]},
			},
			{
				Name:    "powerrental_order_id",
				Unique:  false,
				Columns: []*schema.Column{PowerRentalsColumns[5]},
			},
		},
	}
	// PowerRentalStatesColumns holds the columns for the "power_rental_states" table.
	PowerRentalStatesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "cancel_state", Type: field.TypeString, Nullable: true, Default: "DefaultOrderState"},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "paid_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "user_set_paid", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "user_set_canceled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "admin_set_canceled", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "payment_state", Type: field.TypeString, Nullable: true, Default: "PaymentStateWait"},
		{Name: "outofgas_hours", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "compensate_hours", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "renew_state", Type: field.TypeString, Nullable: true, Default: "OrderRenewWait"},
		{Name: "renew_notify_at", Type: field.TypeUint32, Nullable: true, Default: 0},
	}
	// PowerRentalStatesTable holds the schema information for the "power_rental_states" table.
	PowerRentalStatesTable = &schema.Table{
		Name:       "power_rental_states",
		Columns:    PowerRentalStatesColumns,
		PrimaryKey: []*schema.Column{PowerRentalStatesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "powerrentalstate_ent_id",
				Unique:  true,
				Columns: []*schema.Column{PowerRentalStatesColumns[4]},
			},
			{
				Name:    "powerrentalstate_order_id",
				Unique:  false,
				Columns: []*schema.Column{PowerRentalStatesColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppConfigsTable,
		CompensatesTable,
		OrdersTable,
		OrderBasesTable,
		OrderCouponsTable,
		OrderLocksTable,
		OrderPaymentBalancesTable,
		OrderPaymentContractsTable,
		OrderPaymentTransfersTable,
		OrderStatesTable,
		OrderStateBasesTable,
		OutOfGasTable,
		PaymentsTable,
		PowerRentalsTable,
		PowerRentalStatesTable,
	}
)

func init() {
}
