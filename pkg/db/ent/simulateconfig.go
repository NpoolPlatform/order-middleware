// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/simulateconfig"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// SimulateConfig is the model entity for the SimulateConfig schema.
type SimulateConfig struct {
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
	// SendCouponMode holds the value of the "send_coupon_mode" field.
	SendCouponMode string `json:"send_coupon_mode,omitempty"`
	// SendCouponProbability holds the value of the "send_coupon_probability" field.
	SendCouponProbability decimal.Decimal `json:"send_coupon_probability,omitempty"`
	// EnabledProfitTx holds the value of the "enabled_profit_tx" field.
	EnabledProfitTx bool `json:"enabled_profit_tx,omitempty"`
	// ProfitTxProbability holds the value of the "profit_tx_probability" field.
	ProfitTxProbability decimal.Decimal `json:"profit_tx_probability,omitempty"`
	// Enabled holds the value of the "enabled" field.
	Enabled bool `json:"enabled,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SimulateConfig) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case simulateconfig.FieldSendCouponProbability, simulateconfig.FieldProfitTxProbability:
			values[i] = new(decimal.Decimal)
		case simulateconfig.FieldEnabledProfitTx, simulateconfig.FieldEnabled:
			values[i] = new(sql.NullBool)
		case simulateconfig.FieldID, simulateconfig.FieldCreatedAt, simulateconfig.FieldUpdatedAt, simulateconfig.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case simulateconfig.FieldSendCouponMode:
			values[i] = new(sql.NullString)
		case simulateconfig.FieldEntID, simulateconfig.FieldAppID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SimulateConfig", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SimulateConfig fields.
func (sc *SimulateConfig) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case simulateconfig.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sc.ID = uint32(value.Int64)
		case simulateconfig.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sc.CreatedAt = uint32(value.Int64)
			}
		case simulateconfig.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sc.UpdatedAt = uint32(value.Int64)
			}
		case simulateconfig.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				sc.DeletedAt = uint32(value.Int64)
			}
		case simulateconfig.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				sc.EntID = *value
			}
		case simulateconfig.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				sc.AppID = *value
			}
		case simulateconfig.FieldSendCouponMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field send_coupon_mode", values[i])
			} else if value.Valid {
				sc.SendCouponMode = value.String
			}
		case simulateconfig.FieldSendCouponProbability:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field send_coupon_probability", values[i])
			} else if value != nil {
				sc.SendCouponProbability = *value
			}
		case simulateconfig.FieldEnabledProfitTx:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled_profit_tx", values[i])
			} else if value.Valid {
				sc.EnabledProfitTx = value.Bool
			}
		case simulateconfig.FieldProfitTxProbability:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field profit_tx_probability", values[i])
			} else if value != nil {
				sc.ProfitTxProbability = *value
			}
		case simulateconfig.FieldEnabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field enabled", values[i])
			} else if value.Valid {
				sc.Enabled = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this SimulateConfig.
// Note that you need to call SimulateConfig.Unwrap() before calling this method if this SimulateConfig
// was returned from a transaction, and the transaction was committed or rolled back.
func (sc *SimulateConfig) Update() *SimulateConfigUpdateOne {
	return (&SimulateConfigClient{config: sc.config}).UpdateOne(sc)
}

// Unwrap unwraps the SimulateConfig entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sc *SimulateConfig) Unwrap() *SimulateConfig {
	_tx, ok := sc.config.driver.(*txDriver)
	if !ok {
		panic("ent: SimulateConfig is not a transactional entity")
	}
	sc.config.driver = _tx.drv
	return sc
}

// String implements the fmt.Stringer.
func (sc *SimulateConfig) String() string {
	var builder strings.Builder
	builder.WriteString("SimulateConfig(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", sc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", sc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", sc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", sc.AppID))
	builder.WriteString(", ")
	builder.WriteString("send_coupon_mode=")
	builder.WriteString(sc.SendCouponMode)
	builder.WriteString(", ")
	builder.WriteString("send_coupon_probability=")
	builder.WriteString(fmt.Sprintf("%v", sc.SendCouponProbability))
	builder.WriteString(", ")
	builder.WriteString("enabled_profit_tx=")
	builder.WriteString(fmt.Sprintf("%v", sc.EnabledProfitTx))
	builder.WriteString(", ")
	builder.WriteString("profit_tx_probability=")
	builder.WriteString(fmt.Sprintf("%v", sc.ProfitTxProbability))
	builder.WriteString(", ")
	builder.WriteString("enabled=")
	builder.WriteString(fmt.Sprintf("%v", sc.Enabled))
	builder.WriteByte(')')
	return builder.String()
}

// SimulateConfigs is a parsable slice of SimulateConfig.
type SimulateConfigs []*SimulateConfig

func (sc SimulateConfigs) config(cfg config) {
	for _i := range sc {
		sc[_i].config = cfg
	}
}
