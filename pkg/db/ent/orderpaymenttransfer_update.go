// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderpaymenttransfer"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderPaymentTransferUpdate is the builder for updating OrderPaymentTransfer entities.
type OrderPaymentTransferUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderPaymentTransferMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderPaymentTransferUpdate builder.
func (optu *OrderPaymentTransferUpdate) Where(ps ...predicate.OrderPaymentTransfer) *OrderPaymentTransferUpdate {
	optu.mutation.Where(ps...)
	return optu
}

// SetCreatedAt sets the "created_at" field.
func (optu *OrderPaymentTransferUpdate) SetCreatedAt(u uint32) *OrderPaymentTransferUpdate {
	optu.mutation.ResetCreatedAt()
	optu.mutation.SetCreatedAt(u)
	return optu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (optu *OrderPaymentTransferUpdate) SetNillableCreatedAt(u *uint32) *OrderPaymentTransferUpdate {
	if u != nil {
		optu.SetCreatedAt(*u)
	}
	return optu
}

// AddCreatedAt adds u to the "created_at" field.
func (optu *OrderPaymentTransferUpdate) AddCreatedAt(u int32) *OrderPaymentTransferUpdate {
	optu.mutation.AddCreatedAt(u)
	return optu
}

// SetUpdatedAt sets the "updated_at" field.
func (optu *OrderPaymentTransferUpdate) SetUpdatedAt(u uint32) *OrderPaymentTransferUpdate {
	optu.mutation.ResetUpdatedAt()
	optu.mutation.SetUpdatedAt(u)
	return optu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (optu *OrderPaymentTransferUpdate) AddUpdatedAt(u int32) *OrderPaymentTransferUpdate {
	optu.mutation.AddUpdatedAt(u)
	return optu
}

// SetDeletedAt sets the "deleted_at" field.
func (optu *OrderPaymentTransferUpdate) SetDeletedAt(u uint32) *OrderPaymentTransferUpdate {
	optu.mutation.ResetDeletedAt()
	optu.mutation.SetDeletedAt(u)
	return optu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (optu *OrderPaymentTransferUpdate) SetNillableDeletedAt(u *uint32) *OrderPaymentTransferUpdate {
	if u != nil {
		optu.SetDeletedAt(*u)
	}
	return optu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (optu *OrderPaymentTransferUpdate) AddDeletedAt(u int32) *OrderPaymentTransferUpdate {
	optu.mutation.AddDeletedAt(u)
	return optu
}

// SetEntID sets the "ent_id" field.
func (optu *OrderPaymentTransferUpdate) SetEntID(u uuid.UUID) *OrderPaymentTransferUpdate {
	optu.mutation.SetEntID(u)
	return optu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (optu *OrderPaymentTransferUpdate) SetNillableEntID(u *uuid.UUID) *OrderPaymentTransferUpdate {
	if u != nil {
		optu.SetEntID(*u)
	}
	return optu
}

// SetOrderID sets the "order_id" field.
func (optu *OrderPaymentTransferUpdate) SetOrderID(u uuid.UUID) *OrderPaymentTransferUpdate {
	optu.mutation.SetOrderID(u)
	return optu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (optu *OrderPaymentTransferUpdate) SetNillableOrderID(u *uuid.UUID) *OrderPaymentTransferUpdate {
	if u != nil {
		optu.SetOrderID(*u)
	}
	return optu
}

// ClearOrderID clears the value of the "order_id" field.
func (optu *OrderPaymentTransferUpdate) ClearOrderID() *OrderPaymentTransferUpdate {
	optu.mutation.ClearOrderID()
	return optu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (optu *OrderPaymentTransferUpdate) SetCoinTypeID(u uuid.UUID) *OrderPaymentTransferUpdate {
	optu.mutation.SetCoinTypeID(u)
	return optu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (optu *OrderPaymentTransferUpdate) SetNillableCoinTypeID(u *uuid.UUID) *OrderPaymentTransferUpdate {
	if u != nil {
		optu.SetCoinTypeID(*u)
	}
	return optu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (optu *OrderPaymentTransferUpdate) ClearCoinTypeID() *OrderPaymentTransferUpdate {
	optu.mutation.ClearCoinTypeID()
	return optu
}

// SetStartAmount sets the "start_amount" field.
func (optu *OrderPaymentTransferUpdate) SetStartAmount(d decimal.Decimal) *OrderPaymentTransferUpdate {
	optu.mutation.SetStartAmount(d)
	return optu
}

// SetNillableStartAmount sets the "start_amount" field if the given value is not nil.
func (optu *OrderPaymentTransferUpdate) SetNillableStartAmount(d *decimal.Decimal) *OrderPaymentTransferUpdate {
	if d != nil {
		optu.SetStartAmount(*d)
	}
	return optu
}

// ClearStartAmount clears the value of the "start_amount" field.
func (optu *OrderPaymentTransferUpdate) ClearStartAmount() *OrderPaymentTransferUpdate {
	optu.mutation.ClearStartAmount()
	return optu
}

// Mutation returns the OrderPaymentTransferMutation object of the builder.
func (optu *OrderPaymentTransferUpdate) Mutation() *OrderPaymentTransferMutation {
	return optu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (optu *OrderPaymentTransferUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := optu.defaults(); err != nil {
		return 0, err
	}
	if len(optu.hooks) == 0 {
		affected, err = optu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderPaymentTransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			optu.mutation = mutation
			affected, err = optu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(optu.hooks) - 1; i >= 0; i-- {
			if optu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = optu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, optu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (optu *OrderPaymentTransferUpdate) SaveX(ctx context.Context) int {
	affected, err := optu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (optu *OrderPaymentTransferUpdate) Exec(ctx context.Context) error {
	_, err := optu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (optu *OrderPaymentTransferUpdate) ExecX(ctx context.Context) {
	if err := optu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (optu *OrderPaymentTransferUpdate) defaults() error {
	if _, ok := optu.mutation.UpdatedAt(); !ok {
		if orderpaymenttransfer.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpaymenttransfer.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderpaymenttransfer.UpdateDefaultUpdatedAt()
		optu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (optu *OrderPaymentTransferUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderPaymentTransferUpdate {
	optu.modifiers = append(optu.modifiers, modifiers...)
	return optu
}

func (optu *OrderPaymentTransferUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderpaymenttransfer.Table,
			Columns: orderpaymenttransfer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderpaymenttransfer.FieldID,
			},
		},
	}
	if ps := optu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := optu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldCreatedAt,
		})
	}
	if value, ok := optu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldCreatedAt,
		})
	}
	if value, ok := optu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldUpdatedAt,
		})
	}
	if value, ok := optu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldUpdatedAt,
		})
	}
	if value, ok := optu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldDeletedAt,
		})
	}
	if value, ok := optu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldDeletedAt,
		})
	}
	if value, ok := optu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymenttransfer.FieldEntID,
		})
	}
	if value, ok := optu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymenttransfer.FieldOrderID,
		})
	}
	if optu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymenttransfer.FieldOrderID,
		})
	}
	if value, ok := optu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymenttransfer.FieldCoinTypeID,
		})
	}
	if optu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymenttransfer.FieldCoinTypeID,
		})
	}
	if value, ok := optu.mutation.StartAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderpaymenttransfer.FieldStartAmount,
		})
	}
	if optu.mutation.StartAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderpaymenttransfer.FieldStartAmount,
		})
	}
	_spec.Modifiers = optu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, optu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpaymenttransfer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderPaymentTransferUpdateOne is the builder for updating a single OrderPaymentTransfer entity.
type OrderPaymentTransferUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderPaymentTransferMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (optuo *OrderPaymentTransferUpdateOne) SetCreatedAt(u uint32) *OrderPaymentTransferUpdateOne {
	optuo.mutation.ResetCreatedAt()
	optuo.mutation.SetCreatedAt(u)
	return optuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (optuo *OrderPaymentTransferUpdateOne) SetNillableCreatedAt(u *uint32) *OrderPaymentTransferUpdateOne {
	if u != nil {
		optuo.SetCreatedAt(*u)
	}
	return optuo
}

// AddCreatedAt adds u to the "created_at" field.
func (optuo *OrderPaymentTransferUpdateOne) AddCreatedAt(u int32) *OrderPaymentTransferUpdateOne {
	optuo.mutation.AddCreatedAt(u)
	return optuo
}

// SetUpdatedAt sets the "updated_at" field.
func (optuo *OrderPaymentTransferUpdateOne) SetUpdatedAt(u uint32) *OrderPaymentTransferUpdateOne {
	optuo.mutation.ResetUpdatedAt()
	optuo.mutation.SetUpdatedAt(u)
	return optuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (optuo *OrderPaymentTransferUpdateOne) AddUpdatedAt(u int32) *OrderPaymentTransferUpdateOne {
	optuo.mutation.AddUpdatedAt(u)
	return optuo
}

// SetDeletedAt sets the "deleted_at" field.
func (optuo *OrderPaymentTransferUpdateOne) SetDeletedAt(u uint32) *OrderPaymentTransferUpdateOne {
	optuo.mutation.ResetDeletedAt()
	optuo.mutation.SetDeletedAt(u)
	return optuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (optuo *OrderPaymentTransferUpdateOne) SetNillableDeletedAt(u *uint32) *OrderPaymentTransferUpdateOne {
	if u != nil {
		optuo.SetDeletedAt(*u)
	}
	return optuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (optuo *OrderPaymentTransferUpdateOne) AddDeletedAt(u int32) *OrderPaymentTransferUpdateOne {
	optuo.mutation.AddDeletedAt(u)
	return optuo
}

// SetEntID sets the "ent_id" field.
func (optuo *OrderPaymentTransferUpdateOne) SetEntID(u uuid.UUID) *OrderPaymentTransferUpdateOne {
	optuo.mutation.SetEntID(u)
	return optuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (optuo *OrderPaymentTransferUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderPaymentTransferUpdateOne {
	if u != nil {
		optuo.SetEntID(*u)
	}
	return optuo
}

// SetOrderID sets the "order_id" field.
func (optuo *OrderPaymentTransferUpdateOne) SetOrderID(u uuid.UUID) *OrderPaymentTransferUpdateOne {
	optuo.mutation.SetOrderID(u)
	return optuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (optuo *OrderPaymentTransferUpdateOne) SetNillableOrderID(u *uuid.UUID) *OrderPaymentTransferUpdateOne {
	if u != nil {
		optuo.SetOrderID(*u)
	}
	return optuo
}

// ClearOrderID clears the value of the "order_id" field.
func (optuo *OrderPaymentTransferUpdateOne) ClearOrderID() *OrderPaymentTransferUpdateOne {
	optuo.mutation.ClearOrderID()
	return optuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (optuo *OrderPaymentTransferUpdateOne) SetCoinTypeID(u uuid.UUID) *OrderPaymentTransferUpdateOne {
	optuo.mutation.SetCoinTypeID(u)
	return optuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (optuo *OrderPaymentTransferUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *OrderPaymentTransferUpdateOne {
	if u != nil {
		optuo.SetCoinTypeID(*u)
	}
	return optuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (optuo *OrderPaymentTransferUpdateOne) ClearCoinTypeID() *OrderPaymentTransferUpdateOne {
	optuo.mutation.ClearCoinTypeID()
	return optuo
}

// SetStartAmount sets the "start_amount" field.
func (optuo *OrderPaymentTransferUpdateOne) SetStartAmount(d decimal.Decimal) *OrderPaymentTransferUpdateOne {
	optuo.mutation.SetStartAmount(d)
	return optuo
}

// SetNillableStartAmount sets the "start_amount" field if the given value is not nil.
func (optuo *OrderPaymentTransferUpdateOne) SetNillableStartAmount(d *decimal.Decimal) *OrderPaymentTransferUpdateOne {
	if d != nil {
		optuo.SetStartAmount(*d)
	}
	return optuo
}

// ClearStartAmount clears the value of the "start_amount" field.
func (optuo *OrderPaymentTransferUpdateOne) ClearStartAmount() *OrderPaymentTransferUpdateOne {
	optuo.mutation.ClearStartAmount()
	return optuo
}

// Mutation returns the OrderPaymentTransferMutation object of the builder.
func (optuo *OrderPaymentTransferUpdateOne) Mutation() *OrderPaymentTransferMutation {
	return optuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (optuo *OrderPaymentTransferUpdateOne) Select(field string, fields ...string) *OrderPaymentTransferUpdateOne {
	optuo.fields = append([]string{field}, fields...)
	return optuo
}

// Save executes the query and returns the updated OrderPaymentTransfer entity.
func (optuo *OrderPaymentTransferUpdateOne) Save(ctx context.Context) (*OrderPaymentTransfer, error) {
	var (
		err  error
		node *OrderPaymentTransfer
	)
	if err := optuo.defaults(); err != nil {
		return nil, err
	}
	if len(optuo.hooks) == 0 {
		node, err = optuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderPaymentTransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			optuo.mutation = mutation
			node, err = optuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(optuo.hooks) - 1; i >= 0; i-- {
			if optuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = optuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, optuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderPaymentTransfer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderPaymentTransferMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (optuo *OrderPaymentTransferUpdateOne) SaveX(ctx context.Context) *OrderPaymentTransfer {
	node, err := optuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (optuo *OrderPaymentTransferUpdateOne) Exec(ctx context.Context) error {
	_, err := optuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (optuo *OrderPaymentTransferUpdateOne) ExecX(ctx context.Context) {
	if err := optuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (optuo *OrderPaymentTransferUpdateOne) defaults() error {
	if _, ok := optuo.mutation.UpdatedAt(); !ok {
		if orderpaymenttransfer.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderpaymenttransfer.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderpaymenttransfer.UpdateDefaultUpdatedAt()
		optuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (optuo *OrderPaymentTransferUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderPaymentTransferUpdateOne {
	optuo.modifiers = append(optuo.modifiers, modifiers...)
	return optuo
}

func (optuo *OrderPaymentTransferUpdateOne) sqlSave(ctx context.Context) (_node *OrderPaymentTransfer, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderpaymenttransfer.Table,
			Columns: orderpaymenttransfer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderpaymenttransfer.FieldID,
			},
		},
	}
	id, ok := optuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderPaymentTransfer.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := optuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpaymenttransfer.FieldID)
		for _, f := range fields {
			if !orderpaymenttransfer.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderpaymenttransfer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := optuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := optuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldCreatedAt,
		})
	}
	if value, ok := optuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldCreatedAt,
		})
	}
	if value, ok := optuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldUpdatedAt,
		})
	}
	if value, ok := optuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldUpdatedAt,
		})
	}
	if value, ok := optuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldDeletedAt,
		})
	}
	if value, ok := optuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderpaymenttransfer.FieldDeletedAt,
		})
	}
	if value, ok := optuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymenttransfer.FieldEntID,
		})
	}
	if value, ok := optuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymenttransfer.FieldOrderID,
		})
	}
	if optuo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymenttransfer.FieldOrderID,
		})
	}
	if value, ok := optuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderpaymenttransfer.FieldCoinTypeID,
		})
	}
	if optuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderpaymenttransfer.FieldCoinTypeID,
		})
	}
	if value, ok := optuo.mutation.StartAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderpaymenttransfer.FieldStartAmount,
		})
	}
	if optuo.mutation.StartAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderpaymenttransfer.FieldStartAmount,
		})
	}
	_spec.Modifiers = optuo.modifiers
	_node = &OrderPaymentTransfer{config: optuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, optuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpaymenttransfer.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
