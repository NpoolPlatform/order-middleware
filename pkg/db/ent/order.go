// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// ParentOrderID holds the value of the "parent_order_id" field.
	ParentOrderID uuid.UUID `json:"parent_order_id,omitempty"`
	// PayWithParent holds the value of the "pay_with_parent" field.
	PayWithParent bool `json:"pay_with_parent,omitempty"`
	// Units holds the value of the "units" field.
	Units uint32 `json:"units,omitempty"`
	// UnitsV1 holds the value of the "units_v1" field.
	UnitsV1 decimal.Decimal `json:"units_v1,omitempty"`
	// PromotionID holds the value of the "promotion_id" field.
	PromotionID uuid.UUID `json:"promotion_id,omitempty"`
	// DiscountCouponID holds the value of the "discount_coupon_id" field.
	DiscountCouponID uuid.UUID `json:"discount_coupon_id,omitempty"`
	// UserSpecialReductionID holds the value of the "user_special_reduction_id" field.
	UserSpecialReductionID uuid.UUID `json:"user_special_reduction_id,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt uint32 `json:"end_at,omitempty"`
	// FixAmountCouponID holds the value of the "fix_amount_coupon_id" field.
	FixAmountCouponID uuid.UUID `json:"fix_amount_coupon_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// InvestmentType holds the value of the "investment_type" field.
	InvestmentType string `json:"investment_type,omitempty"`
	// CouponIds holds the value of the "coupon_ids" field.
	CouponIds []uuid.UUID `json:"coupon_ids,omitempty"`
	// LastBenefitAt holds the value of the "last_benefit_at" field.
	LastBenefitAt uint32 `json:"last_benefit_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldCouponIds:
			values[i] = new([]byte)
		case order.FieldUnitsV1:
			values[i] = new(decimal.Decimal)
		case order.FieldPayWithParent:
			values[i] = new(sql.NullBool)
		case order.FieldCreatedAt, order.FieldUpdatedAt, order.FieldDeletedAt, order.FieldUnits, order.FieldStartAt, order.FieldEndAt, order.FieldLastBenefitAt:
			values[i] = new(sql.NullInt64)
		case order.FieldType, order.FieldState, order.FieldInvestmentType:
			values[i] = new(sql.NullString)
		case order.FieldID, order.FieldGoodID, order.FieldAppID, order.FieldUserID, order.FieldParentOrderID, order.FieldPromotionID, order.FieldDiscountCouponID, order.FieldUserSpecialReductionID, order.FieldFixAmountCouponID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				o.ID = *value
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = uint32(value.Int64)
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = uint32(value.Int64)
			}
		case order.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = uint32(value.Int64)
			}
		case order.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				o.GoodID = *value
			}
		case order.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				o.AppID = *value
			}
		case order.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				o.UserID = *value
			}
		case order.FieldParentOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field parent_order_id", values[i])
			} else if value != nil {
				o.ParentOrderID = *value
			}
		case order.FieldPayWithParent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field pay_with_parent", values[i])
			} else if value.Valid {
				o.PayWithParent = value.Bool
			}
		case order.FieldUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value.Valid {
				o.Units = uint32(value.Int64)
			}
		case order.FieldUnitsV1:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field units_v1", values[i])
			} else if value != nil {
				o.UnitsV1 = *value
			}
		case order.FieldPromotionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field promotion_id", values[i])
			} else if value != nil {
				o.PromotionID = *value
			}
		case order.FieldDiscountCouponID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field discount_coupon_id", values[i])
			} else if value != nil {
				o.DiscountCouponID = *value
			}
		case order.FieldUserSpecialReductionID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_special_reduction_id", values[i])
			} else if value != nil {
				o.UserSpecialReductionID = *value
			}
		case order.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				o.StartAt = uint32(value.Int64)
			}
		case order.FieldEndAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				o.EndAt = uint32(value.Int64)
			}
		case order.FieldFixAmountCouponID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field fix_amount_coupon_id", values[i])
			} else if value != nil {
				o.FixAmountCouponID = *value
			}
		case order.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = value.String
			}
		case order.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				o.State = value.String
			}
		case order.FieldInvestmentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field investment_type", values[i])
			} else if value.Valid {
				o.InvestmentType = value.String
			}
		case order.FieldCouponIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field coupon_ids", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.CouponIds); err != nil {
					return fmt.Errorf("unmarshal field coupon_ids: %w", err)
				}
			}
		case order.FieldLastBenefitAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_benefit_at", values[i])
			} else if value.Valid {
				o.LastBenefitAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", o.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", o.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", o.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", o.GoodID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", o.AppID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", o.UserID))
	builder.WriteString(", ")
	builder.WriteString("parent_order_id=")
	builder.WriteString(fmt.Sprintf("%v", o.ParentOrderID))
	builder.WriteString(", ")
	builder.WriteString("pay_with_parent=")
	builder.WriteString(fmt.Sprintf("%v", o.PayWithParent))
	builder.WriteString(", ")
	builder.WriteString("units=")
	builder.WriteString(fmt.Sprintf("%v", o.Units))
	builder.WriteString(", ")
	builder.WriteString("units_v1=")
	builder.WriteString(fmt.Sprintf("%v", o.UnitsV1))
	builder.WriteString(", ")
	builder.WriteString("promotion_id=")
	builder.WriteString(fmt.Sprintf("%v", o.PromotionID))
	builder.WriteString(", ")
	builder.WriteString("discount_coupon_id=")
	builder.WriteString(fmt.Sprintf("%v", o.DiscountCouponID))
	builder.WriteString(", ")
	builder.WriteString("user_special_reduction_id=")
	builder.WriteString(fmt.Sprintf("%v", o.UserSpecialReductionID))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", o.StartAt))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(fmt.Sprintf("%v", o.EndAt))
	builder.WriteString(", ")
	builder.WriteString("fix_amount_coupon_id=")
	builder.WriteString(fmt.Sprintf("%v", o.FixAmountCouponID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(o.Type)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(o.State)
	builder.WriteString(", ")
	builder.WriteString("investment_type=")
	builder.WriteString(o.InvestmentType)
	builder.WriteString(", ")
	builder.WriteString("coupon_ids=")
	builder.WriteString(fmt.Sprintf("%v", o.CouponIds))
	builder.WriteString(", ")
	builder.WriteString("last_benefit_at=")
	builder.WriteString(fmt.Sprintf("%v", o.LastBenefitAt))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
