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
	"github.com/google/uuid"
)

// OrderLockCreate is the builder for creating a OrderLock entity.
type OrderLockCreate struct {
	config
	mutation *OrderLockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (olc *OrderLockCreate) SetCreatedAt(u uint32) *OrderLockCreate {
	olc.mutation.SetCreatedAt(u)
	return olc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (olc *OrderLockCreate) SetNillableCreatedAt(u *uint32) *OrderLockCreate {
	if u != nil {
		olc.SetCreatedAt(*u)
	}
	return olc
}

// SetUpdatedAt sets the "updated_at" field.
func (olc *OrderLockCreate) SetUpdatedAt(u uint32) *OrderLockCreate {
	olc.mutation.SetUpdatedAt(u)
	return olc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (olc *OrderLockCreate) SetNillableUpdatedAt(u *uint32) *OrderLockCreate {
	if u != nil {
		olc.SetUpdatedAt(*u)
	}
	return olc
}

// SetDeletedAt sets the "deleted_at" field.
func (olc *OrderLockCreate) SetDeletedAt(u uint32) *OrderLockCreate {
	olc.mutation.SetDeletedAt(u)
	return olc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (olc *OrderLockCreate) SetNillableDeletedAt(u *uint32) *OrderLockCreate {
	if u != nil {
		olc.SetDeletedAt(*u)
	}
	return olc
}

// SetEntID sets the "ent_id" field.
func (olc *OrderLockCreate) SetEntID(u uuid.UUID) *OrderLockCreate {
	olc.mutation.SetEntID(u)
	return olc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (olc *OrderLockCreate) SetNillableEntID(u *uuid.UUID) *OrderLockCreate {
	if u != nil {
		olc.SetEntID(*u)
	}
	return olc
}

// SetAppID sets the "app_id" field.
func (olc *OrderLockCreate) SetAppID(u uuid.UUID) *OrderLockCreate {
	olc.mutation.SetAppID(u)
	return olc
}

// SetUserID sets the "user_id" field.
func (olc *OrderLockCreate) SetUserID(u uuid.UUID) *OrderLockCreate {
	olc.mutation.SetUserID(u)
	return olc
}

// SetOrderID sets the "order_id" field.
func (olc *OrderLockCreate) SetOrderID(u uuid.UUID) *OrderLockCreate {
	olc.mutation.SetOrderID(u)
	return olc
}

// SetLockType sets the "lock_type" field.
func (olc *OrderLockCreate) SetLockType(s string) *OrderLockCreate {
	olc.mutation.SetLockType(s)
	return olc
}

// SetNillableLockType sets the "lock_type" field if the given value is not nil.
func (olc *OrderLockCreate) SetNillableLockType(s *string) *OrderLockCreate {
	if s != nil {
		olc.SetLockType(*s)
	}
	return olc
}

// SetID sets the "id" field.
func (olc *OrderLockCreate) SetID(u uint32) *OrderLockCreate {
	olc.mutation.SetID(u)
	return olc
}

// Mutation returns the OrderLockMutation object of the builder.
func (olc *OrderLockCreate) Mutation() *OrderLockMutation {
	return olc.mutation
}

// Save creates the OrderLock in the database.
func (olc *OrderLockCreate) Save(ctx context.Context) (*OrderLock, error) {
	var (
		err  error
		node *OrderLock
	)
	if err := olc.defaults(); err != nil {
		return nil, err
	}
	if len(olc.hooks) == 0 {
		if err = olc.check(); err != nil {
			return nil, err
		}
		node, err = olc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = olc.check(); err != nil {
				return nil, err
			}
			olc.mutation = mutation
			if node, err = olc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(olc.hooks) - 1; i >= 0; i-- {
			if olc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = olc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, olc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (olc *OrderLockCreate) SaveX(ctx context.Context) *OrderLock {
	v, err := olc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olc *OrderLockCreate) Exec(ctx context.Context) error {
	_, err := olc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olc *OrderLockCreate) ExecX(ctx context.Context) {
	if err := olc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (olc *OrderLockCreate) defaults() error {
	if _, ok := olc.mutation.CreatedAt(); !ok {
		if orderlock.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderlock.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orderlock.DefaultCreatedAt()
		olc.mutation.SetCreatedAt(v)
	}
	if _, ok := olc.mutation.UpdatedAt(); !ok {
		if orderlock.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderlock.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderlock.DefaultUpdatedAt()
		olc.mutation.SetUpdatedAt(v)
	}
	if _, ok := olc.mutation.DeletedAt(); !ok {
		if orderlock.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized orderlock.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := orderlock.DefaultDeletedAt()
		olc.mutation.SetDeletedAt(v)
	}
	if _, ok := olc.mutation.EntID(); !ok {
		if orderlock.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized orderlock.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := orderlock.DefaultEntID()
		olc.mutation.SetEntID(v)
	}
	if _, ok := olc.mutation.LockType(); !ok {
		v := orderlock.DefaultLockType
		olc.mutation.SetLockType(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (olc *OrderLockCreate) check() error {
	if _, ok := olc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderLock.created_at"`)}
	}
	if _, ok := olc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderLock.updated_at"`)}
	}
	if _, ok := olc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "OrderLock.deleted_at"`)}
	}
	if _, ok := olc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "OrderLock.ent_id"`)}
	}
	if _, ok := olc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "OrderLock.app_id"`)}
	}
	if _, ok := olc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "OrderLock.user_id"`)}
	}
	if _, ok := olc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderLock.order_id"`)}
	}
	return nil
}

func (olc *OrderLockCreate) sqlSave(ctx context.Context) (*OrderLock, error) {
	_node, _spec := olc.createSpec()
	if err := sqlgraph.CreateNode(ctx, olc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (olc *OrderLockCreate) createSpec() (*OrderLock, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderLock{config: olc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: orderlock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderlock.FieldID,
			},
		}
	)
	_spec.OnConflict = olc.conflict
	if id, ok := olc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := olc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := olc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := olc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderlock.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := olc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := olc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := olc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := olc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderlock.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := olc.mutation.LockType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderlock.FieldLockType,
		})
		_node.LockType = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderLock.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (olc *OrderLockCreate) OnConflict(opts ...sql.ConflictOption) *OrderLockUpsertOne {
	olc.conflict = opts
	return &OrderLockUpsertOne{
		create: olc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (olc *OrderLockCreate) OnConflictColumns(columns ...string) *OrderLockUpsertOne {
	olc.conflict = append(olc.conflict, sql.ConflictColumns(columns...))
	return &OrderLockUpsertOne{
		create: olc,
	}
}

type (
	// OrderLockUpsertOne is the builder for "upsert"-ing
	//  one OrderLock node.
	OrderLockUpsertOne struct {
		create *OrderLockCreate
	}

	// OrderLockUpsert is the "OnConflict" setter.
	OrderLockUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *OrderLockUpsert) SetCreatedAt(v uint32) *OrderLockUpsert {
	u.Set(orderlock.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateCreatedAt() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderLockUpsert) AddCreatedAt(v uint32) *OrderLockUpsert {
	u.Add(orderlock.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderLockUpsert) SetUpdatedAt(v uint32) *OrderLockUpsert {
	u.Set(orderlock.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateUpdatedAt() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderLockUpsert) AddUpdatedAt(v uint32) *OrderLockUpsert {
	u.Add(orderlock.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderLockUpsert) SetDeletedAt(v uint32) *OrderLockUpsert {
	u.Set(orderlock.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateDeletedAt() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderLockUpsert) AddDeletedAt(v uint32) *OrderLockUpsert {
	u.Add(orderlock.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *OrderLockUpsert) SetEntID(v uuid.UUID) *OrderLockUpsert {
	u.Set(orderlock.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateEntID() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *OrderLockUpsert) SetAppID(v uuid.UUID) *OrderLockUpsert {
	u.Set(orderlock.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateAppID() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *OrderLockUpsert) SetUserID(v uuid.UUID) *OrderLockUpsert {
	u.Set(orderlock.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateUserID() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldUserID)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderLockUpsert) SetOrderID(v uuid.UUID) *OrderLockUpsert {
	u.Set(orderlock.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateOrderID() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldOrderID)
	return u
}

// SetLockType sets the "lock_type" field.
func (u *OrderLockUpsert) SetLockType(v string) *OrderLockUpsert {
	u.Set(orderlock.FieldLockType, v)
	return u
}

// UpdateLockType sets the "lock_type" field to the value that was provided on create.
func (u *OrderLockUpsert) UpdateLockType() *OrderLockUpsert {
	u.SetExcluded(orderlock.FieldLockType)
	return u
}

// ClearLockType clears the value of the "lock_type" field.
func (u *OrderLockUpsert) ClearLockType() *OrderLockUpsert {
	u.SetNull(orderlock.FieldLockType)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderlock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *OrderLockUpsertOne) UpdateNewValues() *OrderLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(orderlock.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.OrderLock.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OrderLockUpsertOne) Ignore() *OrderLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderLockUpsertOne) DoNothing() *OrderLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderLockCreate.OnConflict
// documentation for more info.
func (u *OrderLockUpsertOne) Update(set func(*OrderLockUpsert)) *OrderLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderLockUpsertOne) SetCreatedAt(v uint32) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderLockUpsertOne) AddCreatedAt(v uint32) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateCreatedAt() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderLockUpsertOne) SetUpdatedAt(v uint32) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderLockUpsertOne) AddUpdatedAt(v uint32) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateUpdatedAt() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderLockUpsertOne) SetDeletedAt(v uint32) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderLockUpsertOne) AddDeletedAt(v uint32) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateDeletedAt() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *OrderLockUpsertOne) SetEntID(v uuid.UUID) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateEntID() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *OrderLockUpsertOne) SetAppID(v uuid.UUID) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateAppID() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrderLockUpsertOne) SetUserID(v uuid.UUID) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateUserID() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateUserID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *OrderLockUpsertOne) SetOrderID(v uuid.UUID) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateOrderID() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateOrderID()
	})
}

// SetLockType sets the "lock_type" field.
func (u *OrderLockUpsertOne) SetLockType(v string) *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetLockType(v)
	})
}

// UpdateLockType sets the "lock_type" field to the value that was provided on create.
func (u *OrderLockUpsertOne) UpdateLockType() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateLockType()
	})
}

// ClearLockType clears the value of the "lock_type" field.
func (u *OrderLockUpsertOne) ClearLockType() *OrderLockUpsertOne {
	return u.Update(func(s *OrderLockUpsert) {
		s.ClearLockType()
	})
}

// Exec executes the query.
func (u *OrderLockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderLockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderLockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderLockUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderLockUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderLockCreateBulk is the builder for creating many OrderLock entities in bulk.
type OrderLockCreateBulk struct {
	config
	builders []*OrderLockCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderLock entities in the database.
func (olcb *OrderLockCreateBulk) Save(ctx context.Context) ([]*OrderLock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(olcb.builders))
	nodes := make([]*OrderLock, len(olcb.builders))
	mutators := make([]Mutator, len(olcb.builders))
	for i := range olcb.builders {
		func(i int, root context.Context) {
			builder := olcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderLockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, olcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = olcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, olcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, olcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (olcb *OrderLockCreateBulk) SaveX(ctx context.Context) []*OrderLock {
	v, err := olcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olcb *OrderLockCreateBulk) Exec(ctx context.Context) error {
	_, err := olcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olcb *OrderLockCreateBulk) ExecX(ctx context.Context) {
	if err := olcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderLock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (olcb *OrderLockCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderLockUpsertBulk {
	olcb.conflict = opts
	return &OrderLockUpsertBulk{
		create: olcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (olcb *OrderLockCreateBulk) OnConflictColumns(columns ...string) *OrderLockUpsertBulk {
	olcb.conflict = append(olcb.conflict, sql.ConflictColumns(columns...))
	return &OrderLockUpsertBulk{
		create: olcb,
	}
}

// OrderLockUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderLock nodes.
type OrderLockUpsertBulk struct {
	create *OrderLockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(orderlock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *OrderLockUpsertBulk) UpdateNewValues() *OrderLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(orderlock.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderLock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OrderLockUpsertBulk) Ignore() *OrderLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderLockUpsertBulk) DoNothing() *OrderLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderLockCreateBulk.OnConflict
// documentation for more info.
func (u *OrderLockUpsertBulk) Update(set func(*OrderLockUpsert)) *OrderLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderLockUpsertBulk) SetCreatedAt(v uint32) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderLockUpsertBulk) AddCreatedAt(v uint32) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateCreatedAt() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderLockUpsertBulk) SetUpdatedAt(v uint32) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderLockUpsertBulk) AddUpdatedAt(v uint32) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateUpdatedAt() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderLockUpsertBulk) SetDeletedAt(v uint32) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderLockUpsertBulk) AddDeletedAt(v uint32) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateDeletedAt() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *OrderLockUpsertBulk) SetEntID(v uuid.UUID) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateEntID() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *OrderLockUpsertBulk) SetAppID(v uuid.UUID) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateAppID() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *OrderLockUpsertBulk) SetUserID(v uuid.UUID) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateUserID() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateUserID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *OrderLockUpsertBulk) SetOrderID(v uuid.UUID) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateOrderID() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateOrderID()
	})
}

// SetLockType sets the "lock_type" field.
func (u *OrderLockUpsertBulk) SetLockType(v string) *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.SetLockType(v)
	})
}

// UpdateLockType sets the "lock_type" field to the value that was provided on create.
func (u *OrderLockUpsertBulk) UpdateLockType() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.UpdateLockType()
	})
}

// ClearLockType clears the value of the "lock_type" field.
func (u *OrderLockUpsertBulk) ClearLockType() *OrderLockUpsertBulk {
	return u.Update(func(s *OrderLockUpsert) {
		s.ClearLockType()
	})
}

// Exec executes the query.
func (u *OrderLockUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderLockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderLockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderLockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
