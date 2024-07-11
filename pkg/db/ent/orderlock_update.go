// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderlock"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderLockUpdate is the builder for updating OrderLock entities.
type OrderLockUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderLockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderLockUpdate builder.
func (olu *OrderLockUpdate) Where(ps ...predicate.OrderLock) *OrderLockUpdate {
	olu.mutation.Where(ps...)
	return olu
}

// SetCreatedAt sets the "created_at" field.
func (olu *OrderLockUpdate) SetCreatedAt(u uint32) *OrderLockUpdate {
	olu.mutation.ResetCreatedAt()
	olu.mutation.SetCreatedAt(u)
	return olu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (olu *OrderLockUpdate) SetNillableCreatedAt(u *uint32) *OrderLockUpdate {
	if u != nil {
		olu.SetCreatedAt(*u)
	}
	return olu
}

// AddCreatedAt adds u to the "created_at" field.
func (olu *OrderLockUpdate) AddCreatedAt(u int32) *OrderLockUpdate {
	olu.mutation.AddCreatedAt(u)
	return olu
}

// SetUpdatedAt sets the "updated_at" field.
func (olu *OrderLockUpdate) SetUpdatedAt(u uint32) *OrderLockUpdate {
	olu.mutation.ResetUpdatedAt()
	olu.mutation.SetUpdatedAt(u)
	return olu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (olu *OrderLockUpdate) AddUpdatedAt(u int32) *OrderLockUpdate {
	olu.mutation.AddUpdatedAt(u)
	return olu
}

// SetDeletedAt sets the "deleted_at" field.
func (olu *OrderLockUpdate) SetDeletedAt(u uint32) *OrderLockUpdate {
	olu.mutation.ResetDeletedAt()
	olu.mutation.SetDeletedAt(u)
	return olu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (olu *OrderLockUpdate) SetNillableDeletedAt(u *uint32) *OrderLockUpdate {
	if u != nil {
		olu.SetDeletedAt(*u)
	}
	return olu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (olu *OrderLockUpdate) AddDeletedAt(u int32) *OrderLockUpdate {
	olu.mutation.AddDeletedAt(u)
	return olu
}

// SetEntID sets the "ent_id" field.
func (olu *OrderLockUpdate) SetEntID(u uuid.UUID) *OrderLockUpdate {
	olu.mutation.SetEntID(u)
	return olu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (olu *OrderLockUpdate) SetNillableEntID(u *uuid.UUID) *OrderLockUpdate {
	if u != nil {
		olu.SetEntID(*u)
	}
	return olu
}

// SetOrderID sets the "order_id" field.
func (olu *OrderLockUpdate) SetOrderID(u uuid.UUID) *OrderLockUpdate {
	olu.mutation.SetOrderID(u)
	return olu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (olu *OrderLockUpdate) SetNillableOrderID(u *uuid.UUID) *OrderLockUpdate {
	if u != nil {
		olu.SetOrderID(*u)
	}
	return olu
}

// ClearOrderID clears the value of the "order_id" field.
func (olu *OrderLockUpdate) ClearOrderID() *OrderLockUpdate {
	olu.mutation.ClearOrderID()
	return olu
}

// SetUserID sets the "user_id" field.
func (olu *OrderLockUpdate) SetUserID(u uuid.UUID) *OrderLockUpdate {
	olu.mutation.SetUserID(u)
	return olu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (olu *OrderLockUpdate) SetNillableUserID(u *uuid.UUID) *OrderLockUpdate {
	if u != nil {
		olu.SetUserID(*u)
	}
	return olu
}

// ClearUserID clears the value of the "user_id" field.
func (olu *OrderLockUpdate) ClearUserID() *OrderLockUpdate {
	olu.mutation.ClearUserID()
	return olu
}

// SetLockType sets the "lock_type" field.
func (olu *OrderLockUpdate) SetLockType(s string) *OrderLockUpdate {
	olu.mutation.SetLockType(s)
	return olu
}

// SetNillableLockType sets the "lock_type" field if the given value is not nil.
func (olu *OrderLockUpdate) SetNillableLockType(s *string) *OrderLockUpdate {
	if s != nil {
		olu.SetLockType(*s)
	}
	return olu
}

// ClearLockType clears the value of the "lock_type" field.
func (olu *OrderLockUpdate) ClearLockType() *OrderLockUpdate {
	olu.mutation.ClearLockType()
	return olu
}

// Mutation returns the OrderLockMutation object of the builder.
func (olu *OrderLockUpdate) Mutation() *OrderLockMutation {
	return olu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (olu *OrderLockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := olu.defaults(); err != nil {
		return 0, err
	}
	if len(olu.hooks) == 0 {
		affected, err = olu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			olu.mutation = mutation
			affected, err = olu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(olu.hooks) - 1; i >= 0; i-- {
			if olu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = olu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, olu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (olu *OrderLockUpdate) SaveX(ctx context.Context) int {
	affected, err := olu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (olu *OrderLockUpdate) Exec(ctx context.Context) error {
	_, err := olu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olu *OrderLockUpdate) ExecX(ctx context.Context) {
	if err := olu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olu *OrderLockUpdate) defaults() error {
	if _, ok := olu.mutation.UpdatedAt(); !ok {
		if orderlock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderlock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderlock.UpdateDefaultUpdatedAt()
		olu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (olu *OrderLockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderLockUpdate {
	olu.modifiers = append(olu.modifiers, modifiers...)
	return olu
}

func (olu *OrderLockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderlock.Table,
			Columns: orderlock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderlock.FieldID,
			},
		},
	}
	if ps := olu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := olu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldCreatedAt,
		})
	}
	if value, ok := olu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldCreatedAt,
		})
	}
	if value, ok := olu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldUpdatedAt,
		})
	}
	if value, ok := olu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldUpdatedAt,
		})
	}
	if value, ok := olu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldDeletedAt,
		})
	}
	if value, ok := olu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldDeletedAt,
		})
	}
	if value, ok := olu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldEntID,
		})
	}
	if value, ok := olu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldOrderID,
		})
	}
	if olu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderlock.FieldOrderID,
		})
	}
	if value, ok := olu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldUserID,
		})
	}
	if olu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderlock.FieldUserID,
		})
	}
	if value, ok := olu.mutation.LockType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderlock.FieldLockType,
		})
	}
	if olu.mutation.LockTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderlock.FieldLockType,
		})
	}
	_spec.Modifiers = olu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, olu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderlock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderLockUpdateOne is the builder for updating a single OrderLock entity.
type OrderLockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderLockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (oluo *OrderLockUpdateOne) SetCreatedAt(u uint32) *OrderLockUpdateOne {
	oluo.mutation.ResetCreatedAt()
	oluo.mutation.SetCreatedAt(u)
	return oluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oluo *OrderLockUpdateOne) SetNillableCreatedAt(u *uint32) *OrderLockUpdateOne {
	if u != nil {
		oluo.SetCreatedAt(*u)
	}
	return oluo
}

// AddCreatedAt adds u to the "created_at" field.
func (oluo *OrderLockUpdateOne) AddCreatedAt(u int32) *OrderLockUpdateOne {
	oluo.mutation.AddCreatedAt(u)
	return oluo
}

// SetUpdatedAt sets the "updated_at" field.
func (oluo *OrderLockUpdateOne) SetUpdatedAt(u uint32) *OrderLockUpdateOne {
	oluo.mutation.ResetUpdatedAt()
	oluo.mutation.SetUpdatedAt(u)
	return oluo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (oluo *OrderLockUpdateOne) AddUpdatedAt(u int32) *OrderLockUpdateOne {
	oluo.mutation.AddUpdatedAt(u)
	return oluo
}

// SetDeletedAt sets the "deleted_at" field.
func (oluo *OrderLockUpdateOne) SetDeletedAt(u uint32) *OrderLockUpdateOne {
	oluo.mutation.ResetDeletedAt()
	oluo.mutation.SetDeletedAt(u)
	return oluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oluo *OrderLockUpdateOne) SetNillableDeletedAt(u *uint32) *OrderLockUpdateOne {
	if u != nil {
		oluo.SetDeletedAt(*u)
	}
	return oluo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (oluo *OrderLockUpdateOne) AddDeletedAt(u int32) *OrderLockUpdateOne {
	oluo.mutation.AddDeletedAt(u)
	return oluo
}

// SetEntID sets the "ent_id" field.
func (oluo *OrderLockUpdateOne) SetEntID(u uuid.UUID) *OrderLockUpdateOne {
	oluo.mutation.SetEntID(u)
	return oluo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (oluo *OrderLockUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderLockUpdateOne {
	if u != nil {
		oluo.SetEntID(*u)
	}
	return oluo
}

// SetOrderID sets the "order_id" field.
func (oluo *OrderLockUpdateOne) SetOrderID(u uuid.UUID) *OrderLockUpdateOne {
	oluo.mutation.SetOrderID(u)
	return oluo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oluo *OrderLockUpdateOne) SetNillableOrderID(u *uuid.UUID) *OrderLockUpdateOne {
	if u != nil {
		oluo.SetOrderID(*u)
	}
	return oluo
}

// ClearOrderID clears the value of the "order_id" field.
func (oluo *OrderLockUpdateOne) ClearOrderID() *OrderLockUpdateOne {
	oluo.mutation.ClearOrderID()
	return oluo
}

// SetUserID sets the "user_id" field.
func (oluo *OrderLockUpdateOne) SetUserID(u uuid.UUID) *OrderLockUpdateOne {
	oluo.mutation.SetUserID(u)
	return oluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (oluo *OrderLockUpdateOne) SetNillableUserID(u *uuid.UUID) *OrderLockUpdateOne {
	if u != nil {
		oluo.SetUserID(*u)
	}
	return oluo
}

// ClearUserID clears the value of the "user_id" field.
func (oluo *OrderLockUpdateOne) ClearUserID() *OrderLockUpdateOne {
	oluo.mutation.ClearUserID()
	return oluo
}

// SetLockType sets the "lock_type" field.
func (oluo *OrderLockUpdateOne) SetLockType(s string) *OrderLockUpdateOne {
	oluo.mutation.SetLockType(s)
	return oluo
}

// SetNillableLockType sets the "lock_type" field if the given value is not nil.
func (oluo *OrderLockUpdateOne) SetNillableLockType(s *string) *OrderLockUpdateOne {
	if s != nil {
		oluo.SetLockType(*s)
	}
	return oluo
}

// ClearLockType clears the value of the "lock_type" field.
func (oluo *OrderLockUpdateOne) ClearLockType() *OrderLockUpdateOne {
	oluo.mutation.ClearLockType()
	return oluo
}

// Mutation returns the OrderLockMutation object of the builder.
func (oluo *OrderLockUpdateOne) Mutation() *OrderLockMutation {
	return oluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oluo *OrderLockUpdateOne) Select(field string, fields ...string) *OrderLockUpdateOne {
	oluo.fields = append([]string{field}, fields...)
	return oluo
}

// Save executes the query and returns the updated OrderLock entity.
func (oluo *OrderLockUpdateOne) Save(ctx context.Context) (*OrderLock, error) {
	var (
		err  error
		node *OrderLock
	)
	if err := oluo.defaults(); err != nil {
		return nil, err
	}
	if len(oluo.hooks) == 0 {
		node, err = oluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oluo.mutation = mutation
			node, err = oluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oluo.hooks) - 1; i >= 0; i-- {
			if oluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, oluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderLock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderLockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (oluo *OrderLockUpdateOne) SaveX(ctx context.Context) *OrderLock {
	node, err := oluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oluo *OrderLockUpdateOne) Exec(ctx context.Context) error {
	_, err := oluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oluo *OrderLockUpdateOne) ExecX(ctx context.Context) {
	if err := oluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oluo *OrderLockUpdateOne) defaults() error {
	if _, ok := oluo.mutation.UpdatedAt(); !ok {
		if orderlock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderlock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderlock.UpdateDefaultUpdatedAt()
		oluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (oluo *OrderLockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderLockUpdateOne {
	oluo.modifiers = append(oluo.modifiers, modifiers...)
	return oluo
}

func (oluo *OrderLockUpdateOne) sqlSave(ctx context.Context) (_node *OrderLock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderlock.Table,
			Columns: orderlock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderlock.FieldID,
			},
		},
	}
	id, ok := oluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderLock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderlock.FieldID)
		for _, f := range fields {
			if !orderlock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderlock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldCreatedAt,
		})
	}
	if value, ok := oluo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldCreatedAt,
		})
	}
	if value, ok := oluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldUpdatedAt,
		})
	}
	if value, ok := oluo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldUpdatedAt,
		})
	}
	if value, ok := oluo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldDeletedAt,
		})
	}
	if value, ok := oluo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldDeletedAt,
		})
	}
	if value, ok := oluo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldEntID,
		})
	}
	if value, ok := oluo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldOrderID,
		})
	}
	if oluo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderlock.FieldOrderID,
		})
	}
	if value, ok := oluo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldUserID,
		})
	}
	if oluo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderlock.FieldUserID,
		})
	}
	if value, ok := oluo.mutation.LockType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderlock.FieldLockType,
		})
	}
	if oluo.mutation.LockTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderlock.FieldLockType,
		})
	}
	_spec.Modifiers = oluo.modifiers
	_node = &OrderLock{config: oluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderlock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
