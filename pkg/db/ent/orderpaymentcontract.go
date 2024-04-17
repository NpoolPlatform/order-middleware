// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderpaymentcontract"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderPaymentContract is the model entity for the OrderPaymentContract schema.
type OrderPaymentContract struct {
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
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderPaymentContract) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderpaymentcontract.FieldAmount:
			values[i] = new(decimal.Decimal)
		case orderpaymentcontract.FieldID, orderpaymentcontract.FieldCreatedAt, orderpaymentcontract.FieldUpdatedAt, orderpaymentcontract.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case orderpaymentcontract.FieldEntID, orderpaymentcontract.FieldOrderID, orderpaymentcontract.FieldCoinTypeID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrderPaymentContract", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderPaymentContract fields.
func (opc *OrderPaymentContract) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderpaymentcontract.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			opc.ID = uint32(value.Int64)
		case orderpaymentcontract.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				opc.CreatedAt = uint32(value.Int64)
			}
		case orderpaymentcontract.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				opc.UpdatedAt = uint32(value.Int64)
			}
		case orderpaymentcontract.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				opc.DeletedAt = uint32(value.Int64)
			}
		case orderpaymentcontract.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				opc.EntID = *value
			}
		case orderpaymentcontract.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				opc.OrderID = *value
			}
		case orderpaymentcontract.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				opc.CoinTypeID = *value
			}
		case orderpaymentcontract.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				opc.Amount = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this OrderPaymentContract.
// Note that you need to call OrderPaymentContract.Unwrap() before calling this method if this OrderPaymentContract
// was returned from a transaction, and the transaction was committed or rolled back.
func (opc *OrderPaymentContract) Update() *OrderPaymentContractUpdateOne {
	return (&OrderPaymentContractClient{config: opc.config}).UpdateOne(opc)
}

// Unwrap unwraps the OrderPaymentContract entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (opc *OrderPaymentContract) Unwrap() *OrderPaymentContract {
	_tx, ok := opc.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderPaymentContract is not a transactional entity")
	}
	opc.config.driver = _tx.drv
	return opc
}

// String implements the fmt.Stringer.
func (opc *OrderPaymentContract) String() string {
	var builder strings.Builder
	builder.WriteString("OrderPaymentContract(")
	builder.WriteString(fmt.Sprintf("id=%v, ", opc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", opc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", opc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", opc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", opc.EntID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", opc.OrderID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", opc.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", opc.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// OrderPaymentContracts is a parsable slice of OrderPaymentContract.
type OrderPaymentContracts []*OrderPaymentContract

func (opc OrderPaymentContracts) config(cfg config) {
	for _i := range opc {
		opc[_i].config = cfg
	}
}
