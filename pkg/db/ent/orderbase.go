// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderbase"
	"github.com/google/uuid"
)

// OrderBase is the model entity for the OrderBase schema.
type OrderBase struct {
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
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// ParentOrderID holds the value of the "parent_order_id" field.
	ParentOrderID uuid.UUID `json:"parent_order_id,omitempty"`
	// OrderType holds the value of the "order_type" field.
	OrderType string `json:"order_type,omitempty"`
	// PaymentType holds the value of the "payment_type" field.
	PaymentType string `json:"payment_type,omitempty"`
	// CreateMethod holds the value of the "create_method" field.
	CreateMethod string `json:"create_method,omitempty"`
	// Simulate holds the value of the "simulate" field.
	Simulate bool `json:"simulate,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderBase) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderbase.FieldSimulate:
			values[i] = new(sql.NullBool)
		case orderbase.FieldID, orderbase.FieldCreatedAt, orderbase.FieldUpdatedAt, orderbase.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case orderbase.FieldOrderType, orderbase.FieldPaymentType, orderbase.FieldCreateMethod:
			values[i] = new(sql.NullString)
		case orderbase.FieldEntID, orderbase.FieldAppID, orderbase.FieldUserID, orderbase.FieldGoodID, orderbase.FieldAppGoodID, orderbase.FieldParentOrderID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderBase", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderBase fields.
func (ob *OrderBase) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderbase.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ob.ID = uint32(value.Int64)
		case orderbase.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ob.CreatedAt = uint32(value.Int64)
			}
		case orderbase.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ob.UpdatedAt = uint32(value.Int64)
			}
		case orderbase.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ob.DeletedAt = uint32(value.Int64)
			}
		case orderbase.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				ob.EntID = *value
			}
		case orderbase.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ob.AppID = *value
			}
		case orderbase.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ob.UserID = *value
			}
		case orderbase.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				ob.GoodID = *value
			}
		case orderbase.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				ob.AppGoodID = *value
			}
		case orderbase.FieldParentOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_order_id", values[i])
			} else if value != nil {
				ob.ParentOrderID = *value
			}
		case orderbase.FieldOrderType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_type", values[i])
			} else if value.Valid {
				ob.OrderType = value.String
			}
		case orderbase.FieldPaymentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_type", values[i])
			} else if value.Valid {
				ob.PaymentType = value.String
			}
		case orderbase.FieldCreateMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field create_method", values[i])
			} else if value.Valid {
				ob.CreateMethod = value.String
			}
		case orderbase.FieldSimulate:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field simulate", values[i])
			} else if value.Valid {
				ob.Simulate = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OrderBase.
// Note that you need to call OrderBase.Unwrap() before calling this method if this OrderBase
// was returned from a transaction, and the transaction was committed or rolled back.
func (ob *OrderBase) Update() *OrderBaseUpdateOne {
	return (&OrderBaseClient{config: ob.config}).UpdateOne(ob)
}

// Unwrap unwraps the OrderBase entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ob *OrderBase) Unwrap() *OrderBase {
	_tx, ok := ob.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderBase is not a transactional entity")
	}
	ob.config.driver = _tx.drv
	return ob
}

// String implements the fmt.Stringer.
func (ob *OrderBase) String() string {
	var builder strings.Builder
	builder.WriteString("OrderBase(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ob.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", ob.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", ob.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ob.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", ob.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", ob.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ob.UserID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", ob.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", ob.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("parent_order_id=")
	builder.WriteString(fmt.Sprintf("%v", ob.ParentOrderID))
	builder.WriteString(", ")
	builder.WriteString("order_type=")
	builder.WriteString(ob.OrderType)
	builder.WriteString(", ")
	builder.WriteString("payment_type=")
	builder.WriteString(ob.PaymentType)
	builder.WriteString(", ")
	builder.WriteString("create_method=")
	builder.WriteString(ob.CreateMethod)
	builder.WriteString(", ")
	builder.WriteString("simulate=")
	builder.WriteString(fmt.Sprintf("%v", ob.Simulate))
	builder.WriteByte(')')
	return builder.String()
}

// OrderBases is a parsable slice of OrderBase.
type OrderBases []*OrderBase

func (ob OrderBases) config(cfg config) {
	for _i := range ob {
		ob[_i].config = cfg
	}
}