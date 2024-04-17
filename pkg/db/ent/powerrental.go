// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrental"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PowerRental is the model entity for the PowerRental schema.
type PowerRental struct {
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
	// OrderID holds the value of the "order_id" field.
	OrderID uuid.UUID `json:"order_id,omitempty"`
	// Units holds the value of the "units" field.
	Units decimal.Decimal `json:"units,omitempty"`
	// GoodValue holds the value of the "good_value" field.
	GoodValue decimal.Decimal `json:"good_value,omitempty"`
	// GoodValueUsd holds the value of the "good_value_usd" field.
	GoodValueUsd decimal.Decimal `json:"good_value_usd,omitempty"`
	// PaymentAmount holds the value of the "payment_amount" field.
	PaymentAmount decimal.Decimal `json:"payment_amount,omitempty"`
	// DiscountAmount holds the value of the "discount_amount" field.
	DiscountAmount decimal.Decimal `json:"discount_amount,omitempty"`
	// PromotionID holds the value of the "promotion_id" field.
	PromotionID uuid.UUID `json:"promotion_id,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration uint32 `json:"duration,omitempty"`
	// InvestmentType holds the value of the "investment_type" field.
	InvestmentType string `json:"investment_type,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PowerRental) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case powerrental.FieldUnits, powerrental.FieldGoodValue, powerrental.FieldGoodValueUsd, powerrental.FieldPaymentAmount, powerrental.FieldDiscountAmount:
			values[i] = new(decimal.Decimal)
		case powerrental.FieldID, powerrental.FieldCreatedAt, powerrental.FieldUpdatedAt, powerrental.FieldDeletedAt, powerrental.FieldDuration:
			values[i] = new(sql.NullInt64)
		case powerrental.FieldInvestmentType:
			values[i] = new(sql.NullString)
		case powerrental.FieldEntID, powerrental.FieldOrderID, powerrental.FieldPromotionID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PowerRental", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PowerRental fields.
func (pr *PowerRental) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case powerrental.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = uint32(value.Int64)
		case powerrental.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = uint32(value.Int64)
			}
		case powerrental.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pr.UpdatedAt = uint32(value.Int64)
			}
		case powerrental.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pr.DeletedAt = uint32(value.Int64)
			}
		case powerrental.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				pr.EntID = *value
			}
		case powerrental.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				pr.OrderID = *value
			}
		case powerrental.FieldUnits:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value != nil {
				pr.Units = *value
			}
		case powerrental.FieldGoodValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field good_value", values[i])
			} else if value != nil {
				pr.GoodValue = *value
			}
		case powerrental.FieldGoodValueUsd:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field good_value_usd", values[i])
			} else if value != nil {
				pr.GoodValueUsd = *value
			}
		case powerrental.FieldPaymentAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field payment_amount", values[i])
			} else if value != nil {
				pr.PaymentAmount = *value
			}
		case powerrental.FieldDiscountAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field discount_amount", values[i])
			} else if value != nil {
				pr.DiscountAmount = *value
			}
		case powerrental.FieldPromotionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field promotion_id", values[i])
			} else if value != nil {
				pr.PromotionID = *value
			}
		case powerrental.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				pr.Duration = uint32(value.Int64)
			}
		case powerrental.FieldInvestmentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field investment_type", values[i])
			} else if value.Valid {
				pr.InvestmentType = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PowerRental.
// Note that you need to call PowerRental.Unwrap() before calling this method if this PowerRental
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PowerRental) Update() *PowerRentalUpdateOne {
	return (&PowerRentalClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the PowerRental entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PowerRental) Unwrap() *PowerRental {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PowerRental is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PowerRental) String() string {
	var builder strings.Builder
	builder.WriteString("PowerRental(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pr.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.EntID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.OrderID))
	builder.WriteString(", ")
	builder.WriteString("units=")
	builder.WriteString(fmt.Sprintf("%v", pr.Units))
	builder.WriteString(", ")
	builder.WriteString("good_value=")
	builder.WriteString(fmt.Sprintf("%v", pr.GoodValue))
	builder.WriteString(", ")
	builder.WriteString("good_value_usd=")
	builder.WriteString(fmt.Sprintf("%v", pr.GoodValueUsd))
	builder.WriteString(", ")
	builder.WriteString("payment_amount=")
	builder.WriteString(fmt.Sprintf("%v", pr.PaymentAmount))
	builder.WriteString(", ")
	builder.WriteString("discount_amount=")
	builder.WriteString(fmt.Sprintf("%v", pr.DiscountAmount))
	builder.WriteString(", ")
	builder.WriteString("promotion_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.PromotionID))
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", pr.Duration))
	builder.WriteString(", ")
	builder.WriteString("investment_type=")
	builder.WriteString(pr.InvestmentType)
	builder.WriteByte(')')
	return builder.String()
}

// PowerRentals is a parsable slice of PowerRental.
type PowerRentals []*PowerRental

func (pr PowerRentals) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
