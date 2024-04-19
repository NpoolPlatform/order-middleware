// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbalancelock"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PaymentBalanceLockUpdate is the builder for updating PaymentBalanceLock entities.
type PaymentBalanceLockUpdate struct {
	config
	hooks     []Hook
	mutation  *PaymentBalanceLockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PaymentBalanceLockUpdate builder.
func (pblu *PaymentBalanceLockUpdate) Where(ps ...predicate.PaymentBalanceLock) *PaymentBalanceLockUpdate {
	pblu.mutation.Where(ps...)
	return pblu
}

// SetCreatedAt sets the "created_at" field.
func (pblu *PaymentBalanceLockUpdate) SetCreatedAt(u uint32) *PaymentBalanceLockUpdate {
	pblu.mutation.ResetCreatedAt()
	pblu.mutation.SetCreatedAt(u)
	return pblu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pblu *PaymentBalanceLockUpdate) SetNillableCreatedAt(u *uint32) *PaymentBalanceLockUpdate {
	if u != nil {
		pblu.SetCreatedAt(*u)
	}
	return pblu
}

// AddCreatedAt adds u to the "created_at" field.
func (pblu *PaymentBalanceLockUpdate) AddCreatedAt(u int32) *PaymentBalanceLockUpdate {
	pblu.mutation.AddCreatedAt(u)
	return pblu
}

// SetUpdatedAt sets the "updated_at" field.
func (pblu *PaymentBalanceLockUpdate) SetUpdatedAt(u uint32) *PaymentBalanceLockUpdate {
	pblu.mutation.ResetUpdatedAt()
	pblu.mutation.SetUpdatedAt(u)
	return pblu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pblu *PaymentBalanceLockUpdate) AddUpdatedAt(u int32) *PaymentBalanceLockUpdate {
	pblu.mutation.AddUpdatedAt(u)
	return pblu
}

// SetDeletedAt sets the "deleted_at" field.
func (pblu *PaymentBalanceLockUpdate) SetDeletedAt(u uint32) *PaymentBalanceLockUpdate {
	pblu.mutation.ResetDeletedAt()
	pblu.mutation.SetDeletedAt(u)
	return pblu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pblu *PaymentBalanceLockUpdate) SetNillableDeletedAt(u *uint32) *PaymentBalanceLockUpdate {
	if u != nil {
		pblu.SetDeletedAt(*u)
	}
	return pblu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pblu *PaymentBalanceLockUpdate) AddDeletedAt(u int32) *PaymentBalanceLockUpdate {
	pblu.mutation.AddDeletedAt(u)
	return pblu
}

// SetEntID sets the "ent_id" field.
func (pblu *PaymentBalanceLockUpdate) SetEntID(u uuid.UUID) *PaymentBalanceLockUpdate {
	pblu.mutation.SetEntID(u)
	return pblu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pblu *PaymentBalanceLockUpdate) SetNillableEntID(u *uuid.UUID) *PaymentBalanceLockUpdate {
	if u != nil {
		pblu.SetEntID(*u)
	}
	return pblu
}

// SetPaymentID sets the "payment_id" field.
func (pblu *PaymentBalanceLockUpdate) SetPaymentID(u uuid.UUID) *PaymentBalanceLockUpdate {
	pblu.mutation.SetPaymentID(u)
	return pblu
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (pblu *PaymentBalanceLockUpdate) SetNillablePaymentID(u *uuid.UUID) *PaymentBalanceLockUpdate {
	if u != nil {
		pblu.SetPaymentID(*u)
	}
	return pblu
}

// ClearPaymentID clears the value of the "payment_id" field.
func (pblu *PaymentBalanceLockUpdate) ClearPaymentID() *PaymentBalanceLockUpdate {
	pblu.mutation.ClearPaymentID()
	return pblu
}

// SetLedgerLockID sets the "ledger_lock_id" field.
func (pblu *PaymentBalanceLockUpdate) SetLedgerLockID(u uuid.UUID) *PaymentBalanceLockUpdate {
	pblu.mutation.SetLedgerLockID(u)
	return pblu
}

// SetNillableLedgerLockID sets the "ledger_lock_id" field if the given value is not nil.
func (pblu *PaymentBalanceLockUpdate) SetNillableLedgerLockID(u *uuid.UUID) *PaymentBalanceLockUpdate {
	if u != nil {
		pblu.SetLedgerLockID(*u)
	}
	return pblu
}

// ClearLedgerLockID clears the value of the "ledger_lock_id" field.
func (pblu *PaymentBalanceLockUpdate) ClearLedgerLockID() *PaymentBalanceLockUpdate {
	pblu.mutation.ClearLedgerLockID()
	return pblu
}

// Mutation returns the PaymentBalanceLockMutation object of the builder.
func (pblu *PaymentBalanceLockUpdate) Mutation() *PaymentBalanceLockMutation {
	return pblu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pblu *PaymentBalanceLockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pblu.defaults(); err != nil {
		return 0, err
	}
	if len(pblu.hooks) == 0 {
		affected, err = pblu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentBalanceLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pblu.mutation = mutation
			affected, err = pblu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pblu.hooks) - 1; i >= 0; i-- {
			if pblu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pblu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pblu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pblu *PaymentBalanceLockUpdate) SaveX(ctx context.Context) int {
	affected, err := pblu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pblu *PaymentBalanceLockUpdate) Exec(ctx context.Context) error {
	_, err := pblu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pblu *PaymentBalanceLockUpdate) ExecX(ctx context.Context) {
	if err := pblu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pblu *PaymentBalanceLockUpdate) defaults() error {
	if _, ok := pblu.mutation.UpdatedAt(); !ok {
		if paymentbalancelock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentbalancelock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := paymentbalancelock.UpdateDefaultUpdatedAt()
		pblu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pblu *PaymentBalanceLockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PaymentBalanceLockUpdate {
	pblu.modifiers = append(pblu.modifiers, modifiers...)
	return pblu
}

func (pblu *PaymentBalanceLockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentbalancelock.Table,
			Columns: paymentbalancelock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentbalancelock.FieldID,
			},
		},
	}
	if ps := pblu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pblu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldCreatedAt,
		})
	}
	if value, ok := pblu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldCreatedAt,
		})
	}
	if value, ok := pblu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldUpdatedAt,
		})
	}
	if value, ok := pblu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldUpdatedAt,
		})
	}
	if value, ok := pblu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldDeletedAt,
		})
	}
	if value, ok := pblu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldDeletedAt,
		})
	}
	if value, ok := pblu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalancelock.FieldEntID,
		})
	}
	if value, ok := pblu.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalancelock.FieldPaymentID,
		})
	}
	if pblu.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalancelock.FieldPaymentID,
		})
	}
	if value, ok := pblu.mutation.LedgerLockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalancelock.FieldLedgerLockID,
		})
	}
	if pblu.mutation.LedgerLockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalancelock.FieldLedgerLockID,
		})
	}
	_spec.Modifiers = pblu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, pblu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentbalancelock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PaymentBalanceLockUpdateOne is the builder for updating a single PaymentBalanceLock entity.
type PaymentBalanceLockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PaymentBalanceLockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (pbluo *PaymentBalanceLockUpdateOne) SetCreatedAt(u uint32) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.ResetCreatedAt()
	pbluo.mutation.SetCreatedAt(u)
	return pbluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pbluo *PaymentBalanceLockUpdateOne) SetNillableCreatedAt(u *uint32) *PaymentBalanceLockUpdateOne {
	if u != nil {
		pbluo.SetCreatedAt(*u)
	}
	return pbluo
}

// AddCreatedAt adds u to the "created_at" field.
func (pbluo *PaymentBalanceLockUpdateOne) AddCreatedAt(u int32) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.AddCreatedAt(u)
	return pbluo
}

// SetUpdatedAt sets the "updated_at" field.
func (pbluo *PaymentBalanceLockUpdateOne) SetUpdatedAt(u uint32) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.ResetUpdatedAt()
	pbluo.mutation.SetUpdatedAt(u)
	return pbluo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pbluo *PaymentBalanceLockUpdateOne) AddUpdatedAt(u int32) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.AddUpdatedAt(u)
	return pbluo
}

// SetDeletedAt sets the "deleted_at" field.
func (pbluo *PaymentBalanceLockUpdateOne) SetDeletedAt(u uint32) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.ResetDeletedAt()
	pbluo.mutation.SetDeletedAt(u)
	return pbluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pbluo *PaymentBalanceLockUpdateOne) SetNillableDeletedAt(u *uint32) *PaymentBalanceLockUpdateOne {
	if u != nil {
		pbluo.SetDeletedAt(*u)
	}
	return pbluo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pbluo *PaymentBalanceLockUpdateOne) AddDeletedAt(u int32) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.AddDeletedAt(u)
	return pbluo
}

// SetEntID sets the "ent_id" field.
func (pbluo *PaymentBalanceLockUpdateOne) SetEntID(u uuid.UUID) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.SetEntID(u)
	return pbluo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pbluo *PaymentBalanceLockUpdateOne) SetNillableEntID(u *uuid.UUID) *PaymentBalanceLockUpdateOne {
	if u != nil {
		pbluo.SetEntID(*u)
	}
	return pbluo
}

// SetPaymentID sets the "payment_id" field.
func (pbluo *PaymentBalanceLockUpdateOne) SetPaymentID(u uuid.UUID) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.SetPaymentID(u)
	return pbluo
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (pbluo *PaymentBalanceLockUpdateOne) SetNillablePaymentID(u *uuid.UUID) *PaymentBalanceLockUpdateOne {
	if u != nil {
		pbluo.SetPaymentID(*u)
	}
	return pbluo
}

// ClearPaymentID clears the value of the "payment_id" field.
func (pbluo *PaymentBalanceLockUpdateOne) ClearPaymentID() *PaymentBalanceLockUpdateOne {
	pbluo.mutation.ClearPaymentID()
	return pbluo
}

// SetLedgerLockID sets the "ledger_lock_id" field.
func (pbluo *PaymentBalanceLockUpdateOne) SetLedgerLockID(u uuid.UUID) *PaymentBalanceLockUpdateOne {
	pbluo.mutation.SetLedgerLockID(u)
	return pbluo
}

// SetNillableLedgerLockID sets the "ledger_lock_id" field if the given value is not nil.
func (pbluo *PaymentBalanceLockUpdateOne) SetNillableLedgerLockID(u *uuid.UUID) *PaymentBalanceLockUpdateOne {
	if u != nil {
		pbluo.SetLedgerLockID(*u)
	}
	return pbluo
}

// ClearLedgerLockID clears the value of the "ledger_lock_id" field.
func (pbluo *PaymentBalanceLockUpdateOne) ClearLedgerLockID() *PaymentBalanceLockUpdateOne {
	pbluo.mutation.ClearLedgerLockID()
	return pbluo
}

// Mutation returns the PaymentBalanceLockMutation object of the builder.
func (pbluo *PaymentBalanceLockUpdateOne) Mutation() *PaymentBalanceLockMutation {
	return pbluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pbluo *PaymentBalanceLockUpdateOne) Select(field string, fields ...string) *PaymentBalanceLockUpdateOne {
	pbluo.fields = append([]string{field}, fields...)
	return pbluo
}

// Save executes the query and returns the updated PaymentBalanceLock entity.
func (pbluo *PaymentBalanceLockUpdateOne) Save(ctx context.Context) (*PaymentBalanceLock, error) {
	var (
		err  error
		node *PaymentBalanceLock
	)
	if err := pbluo.defaults(); err != nil {
		return nil, err
	}
	if len(pbluo.hooks) == 0 {
		node, err = pbluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentBalanceLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pbluo.mutation = mutation
			node, err = pbluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pbluo.hooks) - 1; i >= 0; i-- {
			if pbluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pbluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pbluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PaymentBalanceLock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PaymentBalanceLockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pbluo *PaymentBalanceLockUpdateOne) SaveX(ctx context.Context) *PaymentBalanceLock {
	node, err := pbluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pbluo *PaymentBalanceLockUpdateOne) Exec(ctx context.Context) error {
	_, err := pbluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pbluo *PaymentBalanceLockUpdateOne) ExecX(ctx context.Context) {
	if err := pbluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pbluo *PaymentBalanceLockUpdateOne) defaults() error {
	if _, ok := pbluo.mutation.UpdatedAt(); !ok {
		if paymentbalancelock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentbalancelock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := paymentbalancelock.UpdateDefaultUpdatedAt()
		pbluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pbluo *PaymentBalanceLockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PaymentBalanceLockUpdateOne {
	pbluo.modifiers = append(pbluo.modifiers, modifiers...)
	return pbluo
}

func (pbluo *PaymentBalanceLockUpdateOne) sqlSave(ctx context.Context) (_node *PaymentBalanceLock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentbalancelock.Table,
			Columns: paymentbalancelock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentbalancelock.FieldID,
			},
		},
	}
	id, ok := pbluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PaymentBalanceLock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pbluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentbalancelock.FieldID)
		for _, f := range fields {
			if !paymentbalancelock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != paymentbalancelock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pbluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pbluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldCreatedAt,
		})
	}
	if value, ok := pbluo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldCreatedAt,
		})
	}
	if value, ok := pbluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldUpdatedAt,
		})
	}
	if value, ok := pbluo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldUpdatedAt,
		})
	}
	if value, ok := pbluo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldDeletedAt,
		})
	}
	if value, ok := pbluo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentbalancelock.FieldDeletedAt,
		})
	}
	if value, ok := pbluo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalancelock.FieldEntID,
		})
	}
	if value, ok := pbluo.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalancelock.FieldPaymentID,
		})
	}
	if pbluo.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalancelock.FieldPaymentID,
		})
	}
	if value, ok := pbluo.mutation.LedgerLockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentbalancelock.FieldLedgerLockID,
		})
	}
	if pbluo.mutation.LedgerLockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: paymentbalancelock.FieldLedgerLockID,
		})
	}
	_spec.Modifiers = pbluo.modifiers
	_node = &PaymentBalanceLock{config: pbluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pbluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{paymentbalancelock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}