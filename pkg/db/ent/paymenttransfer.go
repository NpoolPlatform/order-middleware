// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymenttransfer"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PaymentTransfer is the model entity for the PaymentTransfer schema.
type PaymentTransfer struct {
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
	// AccountID holds the value of the "account_id" field.
	AccountID uuid.UUID `json:"account_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount decimal.Decimal `json:"amount,omitempty"`
	// StartAmount holds the value of the "start_amount" field.
	StartAmount decimal.Decimal `json:"start_amount,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID string `json:"transaction_id,omitempty"`
	// FinishAmount holds the value of the "finish_amount" field.
	FinishAmount decimal.Decimal `json:"finish_amount,omitempty"`
	// CoinUsdCurrency holds the value of the "coin_usd_currency" field.
	CoinUsdCurrency decimal.Decimal `json:"coin_usd_currency,omitempty"`
	// LocalCoinUsdCurrency holds the value of the "local_coin_usd_currency" field.
	LocalCoinUsdCurrency decimal.Decimal `json:"local_coin_usd_currency,omitempty"`
	// LiveCoinUsdCurrency holds the value of the "live_coin_usd_currency" field.
	LiveCoinUsdCurrency decimal.Decimal `json:"live_coin_usd_currency,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PaymentTransfer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case paymenttransfer.FieldAmount, paymenttransfer.FieldStartAmount, paymenttransfer.FieldFinishAmount, paymenttransfer.FieldCoinUsdCurrency, paymenttransfer.FieldLocalCoinUsdCurrency, paymenttransfer.FieldLiveCoinUsdCurrency:
			values[i] = new(decimal.Decimal)
		case paymenttransfer.FieldID, paymenttransfer.FieldCreatedAt, paymenttransfer.FieldUpdatedAt, paymenttransfer.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case paymenttransfer.FieldTransactionID:
			values[i] = new(sql.NullString)
		case paymenttransfer.FieldEntID, paymenttransfer.FieldOrderID, paymenttransfer.FieldCoinTypeID, paymenttransfer.FieldAccountID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PaymentTransfer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PaymentTransfer fields.
func (pt *PaymentTransfer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case paymenttransfer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = uint32(value.Int64)
		case paymenttransfer.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = uint32(value.Int64)
			}
		case paymenttransfer.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pt.UpdatedAt = uint32(value.Int64)
			}
		case paymenttransfer.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pt.DeletedAt = uint32(value.Int64)
			}
		case paymenttransfer.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				pt.EntID = *value
			}
		case paymenttransfer.FieldOrderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value != nil {
				pt.OrderID = *value
			}
		case paymenttransfer.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				pt.CoinTypeID = *value
			}
		case paymenttransfer.FieldAccountID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value != nil {
				pt.AccountID = *value
			}
		case paymenttransfer.FieldAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value != nil {
				pt.Amount = *value
			}
		case paymenttransfer.FieldStartAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field start_amount", values[i])
			} else if value != nil {
				pt.StartAmount = *value
			}
		case paymenttransfer.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				pt.TransactionID = value.String
			}
		case paymenttransfer.FieldFinishAmount:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field finish_amount", values[i])
			} else if value != nil {
				pt.FinishAmount = *value
			}
		case paymenttransfer.FieldCoinUsdCurrency:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field coin_usd_currency", values[i])
			} else if value != nil {
				pt.CoinUsdCurrency = *value
			}
		case paymenttransfer.FieldLocalCoinUsdCurrency:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field local_coin_usd_currency", values[i])
			} else if value != nil {
				pt.LocalCoinUsdCurrency = *value
			}
		case paymenttransfer.FieldLiveCoinUsdCurrency:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field live_coin_usd_currency", values[i])
			} else if value != nil {
				pt.LiveCoinUsdCurrency = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this PaymentTransfer.
// Note that you need to call PaymentTransfer.Unwrap() before calling this method if this PaymentTransfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PaymentTransfer) Update() *PaymentTransferUpdateOne {
	return (&PaymentTransferClient{config: pt.config}).UpdateOne(pt)
}

// Unwrap unwraps the PaymentTransfer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PaymentTransfer) Unwrap() *PaymentTransfer {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PaymentTransfer is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PaymentTransfer) String() string {
	var builder strings.Builder
	builder.WriteString("PaymentTransfer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pt.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pt.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pt.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.EntID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.OrderID))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.AccountID))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", pt.Amount))
	builder.WriteString(", ")
	builder.WriteString("start_amount=")
	builder.WriteString(fmt.Sprintf("%v", pt.StartAmount))
	builder.WriteString(", ")
	builder.WriteString("transaction_id=")
	builder.WriteString(pt.TransactionID)
	builder.WriteString(", ")
	builder.WriteString("finish_amount=")
	builder.WriteString(fmt.Sprintf("%v", pt.FinishAmount))
	builder.WriteString(", ")
	builder.WriteString("coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", pt.CoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("local_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", pt.LocalCoinUsdCurrency))
	builder.WriteString(", ")
	builder.WriteString("live_coin_usd_currency=")
	builder.WriteString(fmt.Sprintf("%v", pt.LiveCoinUsdCurrency))
	builder.WriteByte(')')
	return builder.String()
}

// PaymentTransfers is a parsable slice of PaymentTransfer.
type PaymentTransfers []*PaymentTransfer

func (pt PaymentTransfers) config(cfg config) {
	for _i := range pt {
		pt[_i].config = cfg
	}
}
