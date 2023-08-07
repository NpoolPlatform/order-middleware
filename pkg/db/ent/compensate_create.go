// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	"github.com/google/uuid"
)

// CompensateCreate is the builder for creating a Compensate entity.
type CompensateCreate struct {
	config
	mutation *CompensateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CompensateCreate) SetCreatedAt(u uint32) *CompensateCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableCreatedAt(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CompensateCreate) SetUpdatedAt(u uint32) *CompensateCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableUpdatedAt(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CompensateCreate) SetDeletedAt(u uint32) *CompensateCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableDeletedAt(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetOrderID sets the "order_id" field.
func (cc *CompensateCreate) SetOrderID(u uuid.UUID) *CompensateCreate {
	cc.mutation.SetOrderID(u)
	return cc
}

// SetStart sets the "start" field.
func (cc *CompensateCreate) SetStart(u uint32) *CompensateCreate {
	cc.mutation.SetStart(u)
	return cc
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableStart(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetStart(*u)
	}
	return cc
}

// SetEnd sets the "end" field.
func (cc *CompensateCreate) SetEnd(u uint32) *CompensateCreate {
	cc.mutation.SetEnd(u)
	return cc
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableEnd(u *uint32) *CompensateCreate {
	if u != nil {
		cc.SetEnd(*u)
	}
	return cc
}

// SetMessage sets the "message" field.
func (cc *CompensateCreate) SetMessage(s string) *CompensateCreate {
	cc.mutation.SetMessage(s)
	return cc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableMessage(s *string) *CompensateCreate {
	if s != nil {
		cc.SetMessage(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CompensateCreate) SetID(u uuid.UUID) *CompensateCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CompensateCreate) SetNillableID(u *uuid.UUID) *CompensateCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CompensateMutation object of the builder.
func (cc *CompensateCreate) Mutation() *CompensateMutation {
	return cc.mutation
}

// Save creates the Compensate in the database.
func (cc *CompensateCreate) Save(ctx context.Context) (*Compensate, error) {
	var (
		err  error
		node *Compensate
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompensateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Compensate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CompensateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CompensateCreate) SaveX(ctx context.Context) *Compensate {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CompensateCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CompensateCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CompensateCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if compensate.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized compensate.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := compensate.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if compensate.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized compensate.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := compensate.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if compensate.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized compensate.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := compensate.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.Start(); !ok {
		v := compensate.DefaultStart
		cc.mutation.SetStart(v)
	}
	if _, ok := cc.mutation.End(); !ok {
		v := compensate.DefaultEnd
		cc.mutation.SetEnd(v)
	}
	if _, ok := cc.mutation.Message(); !ok {
		v := compensate.DefaultMessage
		cc.mutation.SetMessage(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if compensate.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized compensate.DefaultID (forgotten import ent/runtime?)")
		}
		v := compensate.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CompensateCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Compensate.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Compensate.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Compensate.deleted_at"`)}
	}
	if _, ok := cc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Compensate.order_id"`)}
	}
	return nil
}

func (cc *CompensateCreate) sqlSave(ctx context.Context) (*Compensate, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cc *CompensateCreate) createSpec() (*Compensate, *sqlgraph.CreateSpec) {
	var (
		_node = &Compensate{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: compensate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: compensate.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: compensate.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := cc.mutation.Start(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldStart,
		})
		_node.Start = value
	}
	if value, ok := cc.mutation.End(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: compensate.FieldEnd,
		})
		_node.End = value
	}
	if value, ok := cc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: compensate.FieldMessage,
		})
		_node.Message = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Compensate.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CompensateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CompensateCreate) OnConflict(opts ...sql.ConflictOption) *CompensateUpsertOne {
	cc.conflict = opts
	return &CompensateUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CompensateCreate) OnConflictColumns(columns ...string) *CompensateUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CompensateUpsertOne{
		create: cc,
	}
}

type (
	// CompensateUpsertOne is the builder for "upsert"-ing
	//  one Compensate node.
	CompensateUpsertOne struct {
		create *CompensateCreate
	}

	// CompensateUpsert is the "OnConflict" setter.
	CompensateUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CompensateUpsert) SetCreatedAt(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateCreatedAt() *CompensateUpsert {
	u.SetExcluded(compensate.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CompensateUpsert) AddCreatedAt(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CompensateUpsert) SetUpdatedAt(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateUpdatedAt() *CompensateUpsert {
	u.SetExcluded(compensate.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CompensateUpsert) AddUpdatedAt(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CompensateUpsert) SetDeletedAt(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateDeletedAt() *CompensateUpsert {
	u.SetExcluded(compensate.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CompensateUpsert) AddDeletedAt(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldDeletedAt, v)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *CompensateUpsert) SetOrderID(v uuid.UUID) *CompensateUpsert {
	u.Set(compensate.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateOrderID() *CompensateUpsert {
	u.SetExcluded(compensate.FieldOrderID)
	return u
}

// SetStart sets the "start" field.
func (u *CompensateUpsert) SetStart(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldStart, v)
	return u
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateStart() *CompensateUpsert {
	u.SetExcluded(compensate.FieldStart)
	return u
}

// AddStart adds v to the "start" field.
func (u *CompensateUpsert) AddStart(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldStart, v)
	return u
}

// ClearStart clears the value of the "start" field.
func (u *CompensateUpsert) ClearStart() *CompensateUpsert {
	u.SetNull(compensate.FieldStart)
	return u
}

// SetEnd sets the "end" field.
func (u *CompensateUpsert) SetEnd(v uint32) *CompensateUpsert {
	u.Set(compensate.FieldEnd, v)
	return u
}

// UpdateEnd sets the "end" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateEnd() *CompensateUpsert {
	u.SetExcluded(compensate.FieldEnd)
	return u
}

// AddEnd adds v to the "end" field.
func (u *CompensateUpsert) AddEnd(v uint32) *CompensateUpsert {
	u.Add(compensate.FieldEnd, v)
	return u
}

// ClearEnd clears the value of the "end" field.
func (u *CompensateUpsert) ClearEnd() *CompensateUpsert {
	u.SetNull(compensate.FieldEnd)
	return u
}

// SetMessage sets the "message" field.
func (u *CompensateUpsert) SetMessage(v string) *CompensateUpsert {
	u.Set(compensate.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CompensateUpsert) UpdateMessage() *CompensateUpsert {
	u.SetExcluded(compensate.FieldMessage)
	return u
}

// ClearMessage clears the value of the "message" field.
func (u *CompensateUpsert) ClearMessage() *CompensateUpsert {
	u.SetNull(compensate.FieldMessage)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(compensate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CompensateUpsertOne) UpdateNewValues() *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(compensate.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Compensate.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CompensateUpsertOne) Ignore() *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CompensateUpsertOne) DoNothing() *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CompensateCreate.OnConflict
// documentation for more info.
func (u *CompensateUpsertOne) Update(set func(*CompensateUpsert)) *CompensateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CompensateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CompensateUpsertOne) SetCreatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CompensateUpsertOne) AddCreatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateCreatedAt() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CompensateUpsertOne) SetUpdatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CompensateUpsertOne) AddUpdatedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateUpdatedAt() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CompensateUpsertOne) SetDeletedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CompensateUpsertOne) AddDeletedAt(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateDeletedAt() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *CompensateUpsertOne) SetOrderID(v uuid.UUID) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateOrderID() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateOrderID()
	})
}

// SetStart sets the "start" field.
func (u *CompensateUpsertOne) SetStart(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetStart(v)
	})
}

// AddStart adds v to the "start" field.
func (u *CompensateUpsertOne) AddStart(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateStart() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateStart()
	})
}

// ClearStart clears the value of the "start" field.
func (u *CompensateUpsertOne) ClearStart() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearStart()
	})
}

// SetEnd sets the "end" field.
func (u *CompensateUpsertOne) SetEnd(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetEnd(v)
	})
}

// AddEnd adds v to the "end" field.
func (u *CompensateUpsertOne) AddEnd(v uint32) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.AddEnd(v)
	})
}

// UpdateEnd sets the "end" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateEnd() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateEnd()
	})
}

// ClearEnd clears the value of the "end" field.
func (u *CompensateUpsertOne) ClearEnd() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearEnd()
	})
}

// SetMessage sets the "message" field.
func (u *CompensateUpsertOne) SetMessage(v string) *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CompensateUpsertOne) UpdateMessage() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *CompensateUpsertOne) ClearMessage() *CompensateUpsertOne {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearMessage()
	})
}

// Exec executes the query.
func (u *CompensateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CompensateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CompensateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CompensateUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CompensateUpsertOne.ID is not supported by MySQL driver. Use CompensateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CompensateUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CompensateCreateBulk is the builder for creating many Compensate entities in bulk.
type CompensateCreateBulk struct {
	config
	builders []*CompensateCreate
	conflict []sql.ConflictOption
}

// Save creates the Compensate entities in the database.
func (ccb *CompensateCreateBulk) Save(ctx context.Context) ([]*Compensate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Compensate, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompensateMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CompensateCreateBulk) SaveX(ctx context.Context) []*Compensate {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CompensateCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CompensateCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Compensate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CompensateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CompensateCreateBulk) OnConflict(opts ...sql.ConflictOption) *CompensateUpsertBulk {
	ccb.conflict = opts
	return &CompensateUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CompensateCreateBulk) OnConflictColumns(columns ...string) *CompensateUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CompensateUpsertBulk{
		create: ccb,
	}
}

// CompensateUpsertBulk is the builder for "upsert"-ing
// a bulk of Compensate nodes.
type CompensateUpsertBulk struct {
	create *CompensateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(compensate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CompensateUpsertBulk) UpdateNewValues() *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(compensate.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Compensate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CompensateUpsertBulk) Ignore() *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CompensateUpsertBulk) DoNothing() *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CompensateCreateBulk.OnConflict
// documentation for more info.
func (u *CompensateUpsertBulk) Update(set func(*CompensateUpsert)) *CompensateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CompensateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CompensateUpsertBulk) SetCreatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CompensateUpsertBulk) AddCreatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateCreatedAt() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CompensateUpsertBulk) SetUpdatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CompensateUpsertBulk) AddUpdatedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateUpdatedAt() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CompensateUpsertBulk) SetDeletedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CompensateUpsertBulk) AddDeletedAt(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateDeletedAt() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetOrderID sets the "order_id" field.
func (u *CompensateUpsertBulk) SetOrderID(v uuid.UUID) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateOrderID() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateOrderID()
	})
}

// SetStart sets the "start" field.
func (u *CompensateUpsertBulk) SetStart(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetStart(v)
	})
}

// AddStart adds v to the "start" field.
func (u *CompensateUpsertBulk) AddStart(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddStart(v)
	})
}

// UpdateStart sets the "start" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateStart() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateStart()
	})
}

// ClearStart clears the value of the "start" field.
func (u *CompensateUpsertBulk) ClearStart() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearStart()
	})
}

// SetEnd sets the "end" field.
func (u *CompensateUpsertBulk) SetEnd(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetEnd(v)
	})
}

// AddEnd adds v to the "end" field.
func (u *CompensateUpsertBulk) AddEnd(v uint32) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.AddEnd(v)
	})
}

// UpdateEnd sets the "end" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateEnd() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateEnd()
	})
}

// ClearEnd clears the value of the "end" field.
func (u *CompensateUpsertBulk) ClearEnd() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearEnd()
	})
}

// SetMessage sets the "message" field.
func (u *CompensateUpsertBulk) SetMessage(v string) *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *CompensateUpsertBulk) UpdateMessage() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *CompensateUpsertBulk) ClearMessage() *CompensateUpsertBulk {
	return u.Update(func(s *CompensateUpsert) {
		s.ClearMessage()
	})
}

// Exec executes the query.
func (u *CompensateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CompensateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CompensateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CompensateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
