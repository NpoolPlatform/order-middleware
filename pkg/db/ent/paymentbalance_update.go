// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalance"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PaymentBalanceUpdate is the builder for updating PaymentBalance entities.
type PaymentBalanceUpdate struct {
	config
	hooks     []Hook
	mutation  *PaymentBalanceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PaymentBalanceUpdate builder.
func (pbu *PaymentBalanceUpdate) Where(ps ...predicate.PaymentBalance) *PaymentBalanceUpdate {
	pbu.mutation.Where(ps...)
	return pbu
}

// SetCreatedAt sets the "created_at" field.
func (pbu *PaymentBalanceUpdate) SetCreatedAt(u uint32) *PaymentBalanceUpdate {
	pbu.mutation.ResetCreatedAt()
	pbu.mutation.SetCreatedAt(u)
	return pbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableCreatedAt(u *uint32) *PaymentBalanceUpdate {
	if u != nil {
		pbu.SetCreatedAt(*u)
	}
	return pbu
}

// AddCreatedAt adds u to the "created_at" field.
func (pbu *PaymentBalanceUpdate) AddCreatedAt(u int32) *PaymentBalanceUpdate {
	pbu.mutation.AddCreatedAt(u)
	return pbu
}

// SetUpdatedAt sets the "updated_at" field.
func (pbu *PaymentBalanceUpdate) SetUpdatedAt(u uint32) *PaymentBalanceUpdate {
	pbu.mutation.ResetUpdatedAt()
	pbu.mutation.SetUpdatedAt(u)
	return pbu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pbu *PaymentBalanceUpdate) AddUpdatedAt(u int32) *PaymentBalanceUpdate {
	pbu.mutation.AddUpdatedAt(u)
	return pbu
}

// SetDeletedAt sets the "deleted_at" field.
func (pbu *PaymentBalanceUpdate) SetDeletedAt(u uint32) *PaymentBalanceUpdate {
	pbu.mutation.ResetDeletedAt()
	pbu.mutation.SetDeletedAt(u)
	return pbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableDeletedAt(u *uint32) *PaymentBalanceUpdate {
	if u != nil {
		pbu.SetDeletedAt(*u)
	}
	return pbu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pbu *PaymentBalanceUpdate) AddDeletedAt(u int32) *PaymentBalanceUpdate {
	pbu.mutation.AddDeletedAt(u)
	return pbu
}

// SetEntID sets the "ent_id" field.
func (pbu *PaymentBalanceUpdate) SetEntID(u uuid.UUID) *PaymentBalanceUpdate {
	pbu.mutation.SetEntID(u)
	return pbu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableEntID(u *uuid.UUID) *PaymentBalanceUpdate {
	if u != nil {
		pbu.SetEntID(*u)
	}
	return pbu
}

// SetPaymentID sets the "payment_id" field.
func (pbu *PaymentBalanceUpdate) SetPaymentID(u uuid.UUID) *PaymentBalanceUpdate {
	pbu.mutation.SetPaymentID(u)
	return pbu
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillablePaymentID(u *uuid.UUID) *PaymentBalanceUpdate {
	if u != nil {
		pbu.SetPaymentID(*u)
	}
	return pbu
}

// ClearPaymentID clears the value of the "payment_id" field.
func (pbu *PaymentBalanceUpdate) ClearPaymentID() *PaymentBalanceUpdate {
	pbu.mutation.ClearPaymentID()
	return pbu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pbu *PaymentBalanceUpdate) SetCoinTypeID(u uuid.UUID) *PaymentBalanceUpdate {
	pbu.mutation.SetCoinTypeID(u)
	return pbu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableCoinTypeID(u *uuid.UUID) *PaymentBalanceUpdate {
	if u != nil {
		pbu.SetCoinTypeID(*u)
	}
	return pbu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (pbu *PaymentBalanceUpdate) ClearCoinTypeID() *PaymentBalanceUpdate {
	pbu.mutation.ClearCoinTypeID()
	return pbu
}

// SetAmount sets the "amount" field.
func (pbu *PaymentBalanceUpdate) SetAmount(d decimal.Decimal) *PaymentBalanceUpdate {
	pbu.mutation.SetAmount(d)
	return pbu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableAmount(d *decimal.Decimal) *PaymentBalanceUpdate {
	if d != nil {
		pbu.SetAmount(*d)
	}
	return pbu
}

// ClearAmount clears the value of the "amount" field.
func (pbu *PaymentBalanceUpdate) ClearAmount() *PaymentBalanceUpdate {
	pbu.mutation.ClearAmount()
	return pbu
}

// SetCoinUsdCurrency sets the "coin_usd_currency" field.
func (pbu *PaymentBalanceUpdate) SetCoinUsdCurrency(d decimal.Decimal) *PaymentBalanceUpdate {
	pbu.mutation.SetCoinUsdCurrency(d)
	return pbu
}

// SetNillableCoinUsdCurrency sets the "coin_usd_currency" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableCoinUsdCurrency(d *decimal.Decimal) *PaymentBalanceUpdate {
	if d != nil {
		pbu.SetCoinUsdCurrency(*d)
	}
	return pbu
}

// ClearCoinUsdCurrency clears the value of the "coin_usd_currency" field.
func (pbu *PaymentBalanceUpdate) ClearCoinUsdCurrency() *PaymentBalanceUpdate {
	pbu.mutation.ClearCoinUsdCurrency()
	return pbu
}

// SetLocalCoinUsdCurrency sets the "local_coin_usd_currency" field.
func (pbu *PaymentBalanceUpdate) SetLocalCoinUsdCurrency(d decimal.Decimal) *PaymentBalanceUpdate {
	pbu.mutation.SetLocalCoinUsdCurrency(d)
	return pbu
}

// SetNillableLocalCoinUsdCurrency sets the "local_coin_usd_currency" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableLocalCoinUsdCurrency(d *decimal.Decimal) *PaymentBalanceUpdate {
	if d != nil {
		pbu.SetLocalCoinUsdCurrency(*d)
	}
	return pbu
}

// ClearLocalCoinUsdCurrency clears the value of the "local_coin_usd_currency" field.
func (pbu *PaymentBalanceUpdate) ClearLocalCoinUsdCurrency() *PaymentBalanceUpdate {
	pbu.mutation.ClearLocalCoinUsdCurrency()
	return pbu
}

// SetLiveCoinUsdCurrency sets the "live_coin_usd_currency" field.
func (pbu *PaymentBalanceUpdate) SetLiveCoinUsdCurrency(d decimal.Decimal) *PaymentBalanceUpdate {
	pbu.mutation.SetLiveCoinUsdCurrency(d)
	return pbu
}

// SetNillableLiveCoinUsdCurrency sets the "live_coin_usd_currency" field if the given value is not nil.
func (pbu *PaymentBalanceUpdate) SetNillableLiveCoinUsdCurrency(d *decimal.Decimal) *PaymentBalanceUpdate {
	if d != nil {
		pbu.SetLiveCoinUsdCurrency(*d)
	}
	return pbu
}

// ClearLiveCoinUsdCurrency clears the value of the "live_coin_usd_currency" field.
func (pbu *PaymentBalanceUpdate) ClearLiveCoinUsdCurrency() *PaymentBalanceUpdate {
	pbu.mutation.ClearLiveCoinUsdCurrency()
	return pbu
}

// Mutation returns the PaymentBalanceMutation object of the builder.
func (pbu *PaymentBalanceUpdate) Mutation() *PaymentBalanceMutation {
	return pbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pbu *PaymentBalanceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pbu.defaults(); err != nil {
		return 0, err
	}
	if len(pbu.hooks) == 0 {
		affected, err = pbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentBalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pbu.mutation = mutation
			affected, err = pbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pbu.hooks) - 1; i >= 0; i-- {
			if pbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pbu *PaymentBalanceUpdate) SaveX(ctx context.Context) int {
	affected, err := pbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pbu *PaymentBalanceUpdate) Exec(ctx context.Context) error {
	_, err := pbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbu *PaymentBalanceUpdate) ExecX(ctx context.Context) {
	if err := pbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbu *PaymentBalanceUpdate) defaults() error {
	if _, ok := pbu.mutation.UpdatedAt(); !ok {
		if paymentbalance.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentbalance.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := paymentbalance.UpdateDefaultUpdatedAt()
		pbu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pbu *PaymentBalanceUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PaymentBalanceUpdate {
	pbu.modifiers = append(pbu.modifiers, modifiers...)
	return pbu
}

func (pbu *PaymentBalanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentbalance.Table,
			Columns: paymentbalance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentbalance.FieldID,
			},
		},
	}
	if ps := pbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldCreatedAt,
		})
	}
	if value, ok := pbu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldCreatedAt,
		})
	}
	if value, ok := pbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldUpdatedAt,
		})
	}
	if value, ok := pbu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldUpdatedAt,
		})
	}
	if value, ok := pbu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldDeletedAt,
		})
	}
	if value, ok := pbu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldDeletedAt,
		})
	}
	if value, ok := pbu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalance.FieldEntID,
		})
	}
	if value, ok := pbu.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalance.FieldPaymentID,
		})
	}
	if pbu.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalance.FieldPaymentID,
		})
	}
	if value, ok := pbu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalance.FieldCoinTypeID,
		})
	}
	if pbu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalance.FieldCoinTypeID,
		})
	}
	if value, ok := pbu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldAmount,
		})
	}
	if pbu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldAmount,
		})
	}
	if value, ok := pbu.mutation.CoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldCoinUsdCurrency,
		})
	}
	if pbu.mutation.CoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldCoinUsdCurrency,
		})
	}
	if value, ok := pbu.mutation.LocalCoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldLocalCoinUsdCurrency,
		})
	}
	if pbu.mutation.LocalCoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldLocalCoinUsdCurrency,
		})
	}
	if value, ok := pbu.mutation.LiveCoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldLiveCoinUsdCurrency,
		})
	}
	if pbu.mutation.LiveCoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldLiveCoinUsdCurrency,
		})
	}
	_spec.Modifiers = pbu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, pbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentbalance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PaymentBalanceUpdateOne is the builder for updating a single PaymentBalance entity.
type PaymentBalanceUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PaymentBalanceMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (pbuo *PaymentBalanceUpdateOne) SetCreatedAt(u uint32) *PaymentBalanceUpdateOne {
	pbuo.mutation.ResetCreatedAt()
	pbuo.mutation.SetCreatedAt(u)
	return pbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableCreatedAt(u *uint32) *PaymentBalanceUpdateOne {
	if u != nil {
		pbuo.SetCreatedAt(*u)
	}
	return pbuo
}

// AddCreatedAt adds u to the "created_at" field.
func (pbuo *PaymentBalanceUpdateOne) AddCreatedAt(u int32) *PaymentBalanceUpdateOne {
	pbuo.mutation.AddCreatedAt(u)
	return pbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (pbuo *PaymentBalanceUpdateOne) SetUpdatedAt(u uint32) *PaymentBalanceUpdateOne {
	pbuo.mutation.ResetUpdatedAt()
	pbuo.mutation.SetUpdatedAt(u)
	return pbuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pbuo *PaymentBalanceUpdateOne) AddUpdatedAt(u int32) *PaymentBalanceUpdateOne {
	pbuo.mutation.AddUpdatedAt(u)
	return pbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (pbuo *PaymentBalanceUpdateOne) SetDeletedAt(u uint32) *PaymentBalanceUpdateOne {
	pbuo.mutation.ResetDeletedAt()
	pbuo.mutation.SetDeletedAt(u)
	return pbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableDeletedAt(u *uint32) *PaymentBalanceUpdateOne {
	if u != nil {
		pbuo.SetDeletedAt(*u)
	}
	return pbuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pbuo *PaymentBalanceUpdateOne) AddDeletedAt(u int32) *PaymentBalanceUpdateOne {
	pbuo.mutation.AddDeletedAt(u)
	return pbuo
}

// SetEntID sets the "ent_id" field.
func (pbuo *PaymentBalanceUpdateOne) SetEntID(u uuid.UUID) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetEntID(u)
	return pbuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableEntID(u *uuid.UUID) *PaymentBalanceUpdateOne {
	if u != nil {
		pbuo.SetEntID(*u)
	}
	return pbuo
}

// SetPaymentID sets the "payment_id" field.
func (pbuo *PaymentBalanceUpdateOne) SetPaymentID(u uuid.UUID) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetPaymentID(u)
	return pbuo
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillablePaymentID(u *uuid.UUID) *PaymentBalanceUpdateOne {
	if u != nil {
		pbuo.SetPaymentID(*u)
	}
	return pbuo
}

// ClearPaymentID clears the value of the "payment_id" field.
func (pbuo *PaymentBalanceUpdateOne) ClearPaymentID() *PaymentBalanceUpdateOne {
	pbuo.mutation.ClearPaymentID()
	return pbuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pbuo *PaymentBalanceUpdateOne) SetCoinTypeID(u uuid.UUID) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetCoinTypeID(u)
	return pbuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *PaymentBalanceUpdateOne {
	if u != nil {
		pbuo.SetCoinTypeID(*u)
	}
	return pbuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (pbuo *PaymentBalanceUpdateOne) ClearCoinTypeID() *PaymentBalanceUpdateOne {
	pbuo.mutation.ClearCoinTypeID()
	return pbuo
}

// SetAmount sets the "amount" field.
func (pbuo *PaymentBalanceUpdateOne) SetAmount(d decimal.Decimal) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetAmount(d)
	return pbuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableAmount(d *decimal.Decimal) *PaymentBalanceUpdateOne {
	if d != nil {
		pbuo.SetAmount(*d)
	}
	return pbuo
}

// ClearAmount clears the value of the "amount" field.
func (pbuo *PaymentBalanceUpdateOne) ClearAmount() *PaymentBalanceUpdateOne {
	pbuo.mutation.ClearAmount()
	return pbuo
}

// SetCoinUsdCurrency sets the "coin_usd_currency" field.
func (pbuo *PaymentBalanceUpdateOne) SetCoinUsdCurrency(d decimal.Decimal) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetCoinUsdCurrency(d)
	return pbuo
}

// SetNillableCoinUsdCurrency sets the "coin_usd_currency" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableCoinUsdCurrency(d *decimal.Decimal) *PaymentBalanceUpdateOne {
	if d != nil {
		pbuo.SetCoinUsdCurrency(*d)
	}
	return pbuo
}

// ClearCoinUsdCurrency clears the value of the "coin_usd_currency" field.
func (pbuo *PaymentBalanceUpdateOne) ClearCoinUsdCurrency() *PaymentBalanceUpdateOne {
	pbuo.mutation.ClearCoinUsdCurrency()
	return pbuo
}

// SetLocalCoinUsdCurrency sets the "local_coin_usd_currency" field.
func (pbuo *PaymentBalanceUpdateOne) SetLocalCoinUsdCurrency(d decimal.Decimal) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetLocalCoinUsdCurrency(d)
	return pbuo
}

// SetNillableLocalCoinUsdCurrency sets the "local_coin_usd_currency" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableLocalCoinUsdCurrency(d *decimal.Decimal) *PaymentBalanceUpdateOne {
	if d != nil {
		pbuo.SetLocalCoinUsdCurrency(*d)
	}
	return pbuo
}

// ClearLocalCoinUsdCurrency clears the value of the "local_coin_usd_currency" field.
func (pbuo *PaymentBalanceUpdateOne) ClearLocalCoinUsdCurrency() *PaymentBalanceUpdateOne {
	pbuo.mutation.ClearLocalCoinUsdCurrency()
	return pbuo
}

// SetLiveCoinUsdCurrency sets the "live_coin_usd_currency" field.
func (pbuo *PaymentBalanceUpdateOne) SetLiveCoinUsdCurrency(d decimal.Decimal) *PaymentBalanceUpdateOne {
	pbuo.mutation.SetLiveCoinUsdCurrency(d)
	return pbuo
}

// SetNillableLiveCoinUsdCurrency sets the "live_coin_usd_currency" field if the given value is not nil.
func (pbuo *PaymentBalanceUpdateOne) SetNillableLiveCoinUsdCurrency(d *decimal.Decimal) *PaymentBalanceUpdateOne {
	if d != nil {
		pbuo.SetLiveCoinUsdCurrency(*d)
	}
	return pbuo
}

// ClearLiveCoinUsdCurrency clears the value of the "live_coin_usd_currency" field.
func (pbuo *PaymentBalanceUpdateOne) ClearLiveCoinUsdCurrency() *PaymentBalanceUpdateOne {
	pbuo.mutation.ClearLiveCoinUsdCurrency()
	return pbuo
}

// Mutation returns the PaymentBalanceMutation object of the builder.
func (pbuo *PaymentBalanceUpdateOne) Mutation() *PaymentBalanceMutation {
	return pbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pbuo *PaymentBalanceUpdateOne) Select(field string, fields ...string) *PaymentBalanceUpdateOne {
	pbuo.fields = append([]string{field}, fields...)
	return pbuo
}

// Save executes the query and returns the updated PaymentBalance entity.
func (pbuo *PaymentBalanceUpdateOne) Save(ctx context.Context) (*PaymentBalance, error) {
	var (
		err  error
		node *PaymentBalance
	)
	if err := pbuo.defaults(); err != nil {
		return nil, err
	}
	if len(pbuo.hooks) == 0 {
		node, err = pbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentBalanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pbuo.mutation = mutation
			node, err = pbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pbuo.hooks) - 1; i >= 0; i-- {
			if pbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pbuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pbuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PaymentBalance)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PaymentBalanceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pbuo *PaymentBalanceUpdateOne) SaveX(ctx context.Context) *PaymentBalance {
	node, err := pbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pbuo *PaymentBalanceUpdateOne) Exec(ctx context.Context) error {
	_, err := pbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbuo *PaymentBalanceUpdateOne) ExecX(ctx context.Context) {
	if err := pbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbuo *PaymentBalanceUpdateOne) defaults() error {
	if _, ok := pbuo.mutation.UpdatedAt(); !ok {
		if paymentbalance.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentbalance.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := paymentbalance.UpdateDefaultUpdatedAt()
		pbuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pbuo *PaymentBalanceUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PaymentBalanceUpdateOne {
	pbuo.modifiers = append(pbuo.modifiers, modifiers...)
	return pbuo
}

func (pbuo *PaymentBalanceUpdateOne) sqlSave(ctx context.Context) (_node *PaymentBalance, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentbalance.Table,
			Columns: paymentbalance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentbalance.FieldID,
			},
		},
	}
	id, ok := pbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PaymentBalance.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentbalance.FieldID)
		for _, f := range fields {
			if !paymentbalance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != paymentbalance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldCreatedAt,
		})
	}
	if value, ok := pbuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldCreatedAt,
		})
	}
	if value, ok := pbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldUpdatedAt,
		})
	}
	if value, ok := pbuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldUpdatedAt,
		})
	}
	if value, ok := pbuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldDeletedAt,
		})
	}
	if value, ok := pbuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalance.FieldDeletedAt,
		})
	}
	if value, ok := pbuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalance.FieldEntID,
		})
	}
	if value, ok := pbuo.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalance.FieldPaymentID,
		})
	}
	if pbuo.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalance.FieldPaymentID,
		})
	}
	if value, ok := pbuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalance.FieldCoinTypeID,
		})
	}
	if pbuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalance.FieldCoinTypeID,
		})
	}
	if value, ok := pbuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldAmount,
		})
	}
	if pbuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldAmount,
		})
	}
	if value, ok := pbuo.mutation.CoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldCoinUsdCurrency,
		})
	}
	if pbuo.mutation.CoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldCoinUsdCurrency,
		})
	}
	if value, ok := pbuo.mutation.LocalCoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldLocalCoinUsdCurrency,
		})
	}
	if pbuo.mutation.LocalCoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldLocalCoinUsdCurrency,
		})
	}
	if value, ok := pbuo.mutation.LiveCoinUsdCurrency(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentbalance.FieldLiveCoinUsdCurrency,
		})
	}
	if pbuo.mutation.LiveCoinUsdCurrencyCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: paymentbalance.FieldLiveCoinUsdCurrency,
		})
	}
	_spec.Modifiers = pbuo.modifiers
	_node = &PaymentBalance{config: pbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentbalance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}