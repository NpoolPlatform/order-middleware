// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/appconfig"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppConfig is the model entity for the AppConfig schema.
type AppConfig struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// EnableSimulateOrder holds the value of the "enable_simulate_order" field.
	EnableSimulateOrder bool `json:"enable_simulate_order,omitempty"`
	// SimulateOrderUnits holds the value of the "simulate_order_units" field.
	SimulateOrderUnits decimal.Decimal `json:"simulate_order_units,omitempty"`
	// SimulateOrderCouponMode holds the value of the "simulate_order_coupon_mode" field.
	SimulateOrderCouponMode string `json:"simulate_order_coupon_mode,omitempty"`
	// SimulateOrderCouponProbability holds the value of the "simulate_order_coupon_probability" field.
	SimulateOrderCouponProbability decimal.Decimal `json:"simulate_order_coupon_probability,omitempty"`
	// SimulateOrderCashableProfitProbability holds the value of the "simulate_order_cashable_profit_probability" field.
	SimulateOrderCashableProfitProbability decimal.Decimal `json:"simulate_order_cashable_profit_probability,omitempty"`
	// MaxUnpaidOrders holds the value of the "max_unpaid_orders" field.
	MaxUnpaidOrders uint32 `json:"max_unpaid_orders,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appconfig.FieldSimulateOrderUnits, appconfig.FieldSimulateOrderCouponProbability, appconfig.FieldSimulateOrderCashableProfitProbability:
			values[i] = new(decimal.Decimal)
		case appconfig.FieldEnableSimulateOrder:
			values[i] = new(sql.NullBool)
		case appconfig.FieldID, appconfig.FieldCreatedAt, appconfig.FieldUpdatedAt, appconfig.FieldDeletedAt, appconfig.FieldMaxUnpaidOrders:
			values[i] = new(sql.NullInt64)
		case appconfig.FieldSimulateOrderCouponMode:
			values[i] = new(sql.NullString)
		case appconfig.FieldEntID, appconfig.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppConfig fields.
func (ac *AppConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ac.ID = uint32(value.Int64)
		case appconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ac.CreatedAt = uint32(value.Int64)
			}
		case appconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ac.UpdatedAt = uint32(value.Int64)
			}
		case appconfig.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ac.DeletedAt = uint32(value.Int64)
			}
		case appconfig.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ac.EntID = *value
			}
		case appconfig.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ac.AppID = *value
			}
		case appconfig.FieldEnableSimulateOrder:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enable_simulate_order", values[i])
			} else if value.Valid {
				ac.EnableSimulateOrder = value.Bool
			}
		case appconfig.FieldSimulateOrderUnits:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field simulate_order_units", values[i])
			} else if value != nil {
				ac.SimulateOrderUnits = *value
			}
		case appconfig.FieldSimulateOrderCouponMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field simulate_order_coupon_mode", values[i])
			} else if value.Valid {
				ac.SimulateOrderCouponMode = value.String
			}
		case appconfig.FieldSimulateOrderCouponProbability:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field simulate_order_coupon_probability", values[i])
			} else if value != nil {
				ac.SimulateOrderCouponProbability = *value
			}
		case appconfig.FieldSimulateOrderCashableProfitProbability:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field simulate_order_cashable_profit_probability", values[i])
			} else if value != nil {
				ac.SimulateOrderCashableProfitProbability = *value
			}
		case appconfig.FieldMaxUnpaidOrders:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_unpaid_orders", values[i])
			} else if value.Valid {
				ac.MaxUnpaidOrders = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppConfig.
// Note that you need to call AppConfig.Unwrap() before calling this method if this AppConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AppConfig) Update() *AppConfigUpdateOne {
	return (&AppConfigClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the AppConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ac *AppConfig) Unwrap() *AppConfig {
	_tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppConfig is not a transactional entity")
	}
	ac.config.driver = _tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AppConfig) String() string {
	var builder strings.Builder
	builder.WriteString("AppConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ac.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ac.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.AppID))
	builder.WriteString(", ")
	builder.WriteString("enable_simulate_order=")
	builder.WriteString(fmt.Sprintf("%v", ac.EnableSimulateOrder))
	builder.WriteString(", ")
	builder.WriteString("simulate_order_units=")
	builder.WriteString(fmt.Sprintf("%v", ac.SimulateOrderUnits))
	builder.WriteString(", ")
	builder.WriteString("simulate_order_coupon_mode=")
	builder.WriteString(ac.SimulateOrderCouponMode)
	builder.WriteString(", ")
	builder.WriteString("simulate_order_coupon_probability=")
	builder.WriteString(fmt.Sprintf("%v", ac.SimulateOrderCouponProbability))
	builder.WriteString(", ")
	builder.WriteString("simulate_order_cashable_profit_probability=")
	builder.WriteString(fmt.Sprintf("%v", ac.SimulateOrderCashableProfitProbability))
	builder.WriteString(", ")
	builder.WriteString("max_unpaid_orders=")
	builder.WriteString(fmt.Sprintf("%v", ac.MaxUnpaidOrders))
	builder.WriteByte(')')
	return builder.String()
}

// AppConfigs is a parsable slice of AppConfig.
type AppConfigs []*AppConfig

func (ac AppConfigs) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}
