// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/powerrentalstate"
	"github.com/google/uuid"
)

// PowerRentalState is the model entity for the PowerRentalState schema.
type PowerRentalState struct {
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
	// CancelState holds the value of the "cancel_state" field.
	CancelState string `json:"cancel_state,omitempty"`
	// EndAt holds the value of the "end_at" field.
	EndAt uint32 `json:"end_at,omitempty"`
	// PaidAt holds the value of the "paid_at" field.
	PaidAt uint32 `json:"paid_at,omitempty"`
	// UserSetPaid holds the value of the "user_set_paid" field.
	UserSetPaid bool `json:"user_set_paid,omitempty"`
	// UserSetCanceled holds the value of the "user_set_canceled" field.
	UserSetCanceled bool `json:"user_set_canceled,omitempty"`
	// AdminSetCanceled holds the value of the "admin_set_canceled" field.
	AdminSetCanceled bool `json:"admin_set_canceled,omitempty"`
	// PaymentState holds the value of the "payment_state" field.
	PaymentState string `json:"payment_state,omitempty"`
	// OutofgasHours holds the value of the "outofgas_hours" field.
	OutofgasHours uint32 `json:"outofgas_hours,omitempty"`
	// CompensateHours holds the value of the "compensate_hours" field.
	CompensateHours uint32 `json:"compensate_hours,omitempty"`
	// RenewState holds the value of the "renew_state" field.
	RenewState string `json:"renew_state,omitempty"`
	// RenewNotifyAt holds the value of the "renew_notify_at" field.
	RenewNotifyAt uint32 `json:"renew_notify_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PowerRentalState) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case powerrentalstate.FieldUserSetPaid, powerrentalstate.FieldUserSetCanceled, powerrentalstate.FieldAdminSetCanceled:
			values[i] = new(sql.NullBool)
		case powerrentalstate.FieldID, powerrentalstate.FieldCreatedAt, powerrentalstate.FieldUpdatedAt, powerrentalstate.FieldDeletedAt, powerrentalstate.FieldEndAt, powerrentalstate.FieldPaidAt, powerrentalstate.FieldOutofgasHours, powerrentalstate.FieldCompensateHours, powerrentalstate.FieldRenewNotifyAt:
			values[i] = new(sql.NullInt64)
		case powerrentalstate.FieldCancelState, powerrentalstate.FieldPaymentState, powerrentalstate.FieldRenewState:
			values[i] = new(sql.NullString)
		case powerrentalstate.FieldEntID, powerrentalstate.FieldOrderID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PowerRentalState", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PowerRentalState fields.
func (prs *PowerRentalState) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case powerrentalstate.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			prs.ID = uint32(value.Int64)
		case powerrentalstate.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				prs.CreatedAt = uint32(value.Int64)
			}
		case powerrentalstate.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				prs.UpdatedAt = uint32(value.Int64)
			}
		case powerrentalstate.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				prs.DeletedAt = uint32(value.Int64)
			}
		case powerrentalstate.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				prs.EntID = *value
			}
		case powerrentalstate.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				prs.OrderID = *value
			}
		case powerrentalstate.FieldCancelState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cancel_state", values[i])
			} else if value.Valid {
				prs.CancelState = value.String
			}
		case powerrentalstate.FieldEndAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				prs.EndAt = uint32(value.Int64)
			}
		case powerrentalstate.FieldPaidAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field paid_at", values[i])
			} else if value.Valid {
				prs.PaidAt = uint32(value.Int64)
			}
		case powerrentalstate.FieldUserSetPaid:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field user_set_paid", values[i])
			} else if value.Valid {
				prs.UserSetPaid = value.Bool
			}
		case powerrentalstate.FieldUserSetCanceled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field user_set_canceled", values[i])
			} else if value.Valid {
				prs.UserSetCanceled = value.Bool
			}
		case powerrentalstate.FieldAdminSetCanceled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field admin_set_canceled", values[i])
			} else if value.Valid {
				prs.AdminSetCanceled = value.Bool
			}
		case powerrentalstate.FieldPaymentState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_state", values[i])
			} else if value.Valid {
				prs.PaymentState = value.String
			}
		case powerrentalstate.FieldOutofgasHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field outofgas_hours", values[i])
			} else if value.Valid {
				prs.OutofgasHours = uint32(value.Int64)
			}
		case powerrentalstate.FieldCompensateHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field compensate_hours", values[i])
			} else if value.Valid {
				prs.CompensateHours = uint32(value.Int64)
			}
		case powerrentalstate.FieldRenewState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field renew_state", values[i])
			} else if value.Valid {
				prs.RenewState = value.String
			}
		case powerrentalstate.FieldRenewNotifyAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field renew_notify_at", values[i])
			} else if value.Valid {
				prs.RenewNotifyAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PowerRentalState.
// Note that you need to call PowerRentalState.Unwrap() before calling this method if this PowerRentalState
// was returned from a transaction, and the transaction was committed or rolled back.
func (prs *PowerRentalState) Update() *PowerRentalStateUpdateOne {
	return (&PowerRentalStateClient{config: prs.config}).UpdateOne(prs)
}

// Unwrap unwraps the PowerRentalState entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (prs *PowerRentalState) Unwrap() *PowerRentalState {
	_tx, ok := prs.config.driver.(*txDriver)
	if !ok {
		panic("ent: PowerRentalState is not a transactional entity")
	}
	prs.config.driver = _tx.drv
	return prs
}

// String implements the fmt.Stringer.
func (prs *PowerRentalState) String() string {
	var builder strings.Builder
	builder.WriteString("PowerRentalState(")
	builder.WriteString(fmt.Sprintf("id=%v, ", prs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", prs.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", prs.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", prs.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", prs.EntID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", prs.OrderID))
	builder.WriteString(", ")
	builder.WriteString("cancel_state=")
	builder.WriteString(prs.CancelState)
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(fmt.Sprintf("%v", prs.EndAt))
	builder.WriteString(", ")
	builder.WriteString("paid_at=")
	builder.WriteString(fmt.Sprintf("%v", prs.PaidAt))
	builder.WriteString(", ")
	builder.WriteString("user_set_paid=")
	builder.WriteString(fmt.Sprintf("%v", prs.UserSetPaid))
	builder.WriteString(", ")
	builder.WriteString("user_set_canceled=")
	builder.WriteString(fmt.Sprintf("%v", prs.UserSetCanceled))
	builder.WriteString(", ")
	builder.WriteString("admin_set_canceled=")
	builder.WriteString(fmt.Sprintf("%v", prs.AdminSetCanceled))
	builder.WriteString(", ")
	builder.WriteString("payment_state=")
	builder.WriteString(prs.PaymentState)
	builder.WriteString(", ")
	builder.WriteString("outofgas_hours=")
	builder.WriteString(fmt.Sprintf("%v", prs.OutofgasHours))
	builder.WriteString(", ")
	builder.WriteString("compensate_hours=")
	builder.WriteString(fmt.Sprintf("%v", prs.CompensateHours))
	builder.WriteString(", ")
	builder.WriteString("renew_state=")
	builder.WriteString(prs.RenewState)
	builder.WriteString(", ")
	builder.WriteString("renew_notify_at=")
	builder.WriteString(fmt.Sprintf("%v", prs.RenewNotifyAt))
	builder.WriteByte(')')
	return builder.String()
}

// PowerRentalStates is a parsable slice of PowerRentalState.
type PowerRentalStates []*PowerRentalState

func (prs PowerRentalStates) config(cfg config) {
	for _i := range prs {
		prs[_i].config = cfg
	}
}
