// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/ordercoupon"
	"github.com/google/uuid"
)

// OrderCouponCreate is the builder for creating a OrderCoupon entity.
type OrderCouponCreate struct {
	config
	mutation *OrderCouponMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (occ *OrderCouponCreate) SetCreatedAt(u uint32) *OrderCouponCreate {
	occ.mutation.SetCreatedAt(u)
	return occ
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (occ *OrderCouponCreate) SetNillableCreatedAt(u *uint32) *OrderCouponCreate {
	if u != nil {
		occ.SetCreatedAt(*u)
	}
	return occ
}

// SetUpdatedAt sets the "updated_at" field.
func (occ *OrderCouponCreate) SetUpdatedAt(u uint32) *OrderCouponCreate {
	occ.mutation.SetUpdatedAt(u)
	return occ
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (occ *OrderCouponCreate) SetNillableUpdatedAt(u *uint32) *OrderCouponCreate {
	if u != nil {
		occ.SetUpdatedAt(*u)
	}
	return occ
}

// SetDeletedAt sets the "deleted_at" field.
func (occ *OrderCouponCreate) SetDeletedAt(u uint32) *OrderCouponCreate {
	occ.mutation.SetDeletedAt(u)
	return occ
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (occ *OrderCouponCreate) SetNillableDeletedAt(u *uint32) *OrderCouponCreate {
	if u != nil {
		occ.SetDeletedAt(*u)
	}
	return occ
}

// SetEntID sets the "ent_id" field.
func (occ *OrderCouponCreate) SetEntID(u uuid.UUID) *OrderCouponCreate {
	occ.mutation.SetEntID(u)
	return occ
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (occ *OrderCouponCreate) SetNillableEntID(u *uuid.UUID) *OrderCouponCreate {
	if u != nil {
		occ.SetEntID(*u)
	}
	return occ
}

// SetOrderID sets the "order_id" field.
func (occ *OrderCouponCreate) SetOrderID(u uuid.UUID) *OrderCouponCreate {
	occ.mutation.SetOrderID(u)
	return occ
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (occ *OrderCouponCreate) SetNillableOrderID(u *uuid.UUID) *OrderCouponCreate {
	if u != nil {
		occ.SetOrderID(*u)
	}
	return occ
}

// SetCouponID sets the "coupon_id" field.
func (occ *OrderCouponCreate) SetCouponID(u uuid.UUID) *OrderCouponCreate {
	occ.mutation.SetCouponID(u)
	return occ
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (occ *OrderCouponCreate) SetNillableCouponID(u *uuid.UUID) *OrderCouponCreate {
	if u != nil {
		occ.SetCouponID(*u)
	}
	return occ
}

// SetID sets the "id" field.
func (occ *OrderCouponCreate) SetID(u uint32) *OrderCouponCreate {
	occ.mutation.SetID(u)
	return occ
}

// Mutation returns the OrderCouponMutation object of the builder.
func (occ *OrderCouponCreate) Mutation() *OrderCouponMutation {
	return occ.mutation
}

// Save creates the OrderCoupon in the database.
func (occ *OrderCouponCreate) Save(ctx context.Context) (*OrderCoupon, error) {
	var (
		err  error
		node *OrderCoupon
	)
	if err := occ.defaults(); err != nil {
		return nil, err
	}
	if len(occ.hooks) == 0 {
		if err = occ.check(); err != nil {
			return nil, err
		}
		node, err = occ.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = occ.check(); err != nil {
				return nil, err
			}
			occ.mutation = mutation
			if node, err = occ.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(occ.hooks) - 1; i >= 0; i-- {
			if occ.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = occ.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, occ.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderCoupon)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderCouponMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (occ *OrderCouponCreate) SaveX(ctx context.Context) *OrderCoupon {
	v, err := occ.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occ *OrderCouponCreate) Exec(ctx context.Context) error {
	_, err := occ.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occ *OrderCouponCreate) ExecX(ctx context.Context) {
	if err := occ.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (occ *OrderCouponCreate) defaults() error {
	if _, ok := occ.mutation.CreatedAt(); !ok {
		if ordercoupon.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := ordercoupon.DefaultCreatedAt()
		occ.mutation.SetCreatedAt(v)
	}
	if _, ok := occ.mutation.UpdatedAt(); !ok {
		if ordercoupon.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ordercoupon.DefaultUpdatedAt()
		occ.mutation.SetUpdatedAt(v)
	}
	if _, ok := occ.mutation.DeletedAt(); !ok {
		if ordercoupon.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := ordercoupon.DefaultDeletedAt()
		occ.mutation.SetDeletedAt(v)
	}
	if _, ok := occ.mutation.EntID(); !ok {
		if ordercoupon.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := ordercoupon.DefaultEntID()
		occ.mutation.SetEntID(v)
	}
	if _, ok := occ.mutation.OrderID(); !ok {
		if ordercoupon.DefaultOrderID == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.DefaultOrderID (forgotten import ent/runtime?)")
		}
		v := ordercoupon.DefaultOrderID()
		occ.mutation.SetOrderID(v)
	}
	if _, ok := occ.mutation.CouponID(); !ok {
		if ordercoupon.DefaultCouponID == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.DefaultCouponID (forgotten import ent/runtime?)")
		}
		v := ordercoupon.DefaultCouponID()
		occ.mutation.SetCouponID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (occ *OrderCouponCreate) check() error {
	if _, ok := occ.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderCoupon.created_at"`)}
	}
	if _, ok := occ.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderCoupon.updated_at"`)}
	}
	if _, ok := occ.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "OrderCoupon.deleted_at"`)}
	}
	if _, ok := occ.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "OrderCoupon.ent_id"`)}
	}
	return nil
}

func (occ *OrderCouponCreate) sqlSave(ctx context.Context) (*OrderCoupon, error) {
	_node, _spec := occ.createSpec()
	if err := sqlgraph.CreateNode(ctx, occ.driver, _spec); err != nil {
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

func (occ *OrderCouponCreate) createSpec() (*OrderCoupon, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderCoupon{config: occ.config}
		_spec = &sqlgraph.CreateSpec{
			Table: ordercoupon.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: ordercoupon.FieldID,
			},
		}
	)
	_spec.OnConflict = occ.conflict
	if id, ok := occ.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := occ.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := occ.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := occ.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := occ.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := occ.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := occ.mutation.CouponID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldCouponID,
		})
		_node.CouponID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderCoupon.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderCouponUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (occ *OrderCouponCreate) OnConflict(opts ...sql.ConflictOption) *OrderCouponUpsertOne {
	occ.conflict = opts
	return &OrderCouponUpsertOne{
		create: occ,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderCoupon.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (occ *OrderCouponCreate) OnConflictColumns(columns ...string) *OrderCouponUpsertOne {
	occ.conflict = append(occ.conflict, sql.ConflictColumns(columns...))
	return &OrderCouponUpsertOne{
		create: occ,
	}
}

type (
	// OrderCouponUpsertOne is the builder for "upsert"-ing
	//  one OrderCoupon node.
	OrderCouponUpsertOne struct {
		create *OrderCouponCreate
	}

	// OrderCouponUpsert is the "OnConflict" setter.
	OrderCouponUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *OrderCouponUpsert) SetCreatedAt(v uint32) *OrderCouponUpsert {
	u.Set(ordercoupon.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderCouponUpsert) UpdateCreatedAt() *OrderCouponUpsert {
	u.SetExcluded(ordercoupon.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderCouponUpsert) AddCreatedAt(v uint32) *OrderCouponUpsert {
	u.Add(ordercoupon.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderCouponUpsert) SetUpdatedAt(v uint32) *OrderCouponUpsert {
	u.Set(ordercoupon.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderCouponUpsert) UpdateUpdatedAt() *OrderCouponUpsert {
	u.SetExcluded(ordercoupon.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderCouponUpsert) AddUpdatedAt(v uint32) *OrderCouponUpsert {
	u.Add(ordercoupon.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderCouponUpsert) SetDeletedAt(v uint32) *OrderCouponUpsert {
	u.Set(ordercoupon.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderCouponUpsert) UpdateDeletedAt() *OrderCouponUpsert {
	u.SetExcluded(ordercoupon.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderCouponUpsert) AddDeletedAt(v uint32) *OrderCouponUpsert {
	u.Add(ordercoupon.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *OrderCouponUpsert) SetEntID(v uuid.UUID) *OrderCouponUpsert {
	u.Set(ordercoupon.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderCouponUpsert) UpdateEntID() *OrderCouponUpsert {
	u.SetExcluded(ordercoupon.FieldEntID)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *OrderCouponUpsert) SetOrderID(v uuid.UUID) *OrderCouponUpsert {
	u.Set(ordercoupon.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderCouponUpsert) UpdateOrderID() *OrderCouponUpsert {
	u.SetExcluded(ordercoupon.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *OrderCouponUpsert) ClearOrderID() *OrderCouponUpsert {
	u.SetNull(ordercoupon.FieldOrderID)
	return u
}

// SetCouponID sets the "coupon_id" field.
func (u *OrderCouponUpsert) SetCouponID(v uuid.UUID) *OrderCouponUpsert {
	u.Set(ordercoupon.FieldCouponID, v)
	return u
}

// UpdateCouponID sets the "coupon_id" field to the value that was provided on create.
func (u *OrderCouponUpsert) UpdateCouponID() *OrderCouponUpsert {
	u.SetExcluded(ordercoupon.FieldCouponID)
	return u
}

// ClearCouponID clears the value of the "coupon_id" field.
func (u *OrderCouponUpsert) ClearCouponID() *OrderCouponUpsert {
	u.SetNull(ordercoupon.FieldCouponID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.OrderCoupon.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ordercoupon.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *OrderCouponUpsertOne) UpdateNewValues() *OrderCouponUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(ordercoupon.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.OrderCoupon.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *OrderCouponUpsertOne) Ignore() *OrderCouponUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderCouponUpsertOne) DoNothing() *OrderCouponUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCouponCreate.OnConflict
// documentation for more info.
func (u *OrderCouponUpsertOne) Update(set func(*OrderCouponUpsert)) *OrderCouponUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderCouponUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderCouponUpsertOne) SetCreatedAt(v uint32) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderCouponUpsertOne) AddCreatedAt(v uint32) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderCouponUpsertOne) UpdateCreatedAt() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderCouponUpsertOne) SetUpdatedAt(v uint32) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderCouponUpsertOne) AddUpdatedAt(v uint32) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderCouponUpsertOne) UpdateUpdatedAt() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderCouponUpsertOne) SetDeletedAt(v uint32) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderCouponUpsertOne) AddDeletedAt(v uint32) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderCouponUpsertOne) UpdateDeletedAt() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *OrderCouponUpsertOne) SetEntID(v uuid.UUID) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderCouponUpsertOne) UpdateEntID() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateEntID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *OrderCouponUpsertOne) SetOrderID(v uuid.UUID) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderCouponUpsertOne) UpdateOrderID() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *OrderCouponUpsertOne) ClearOrderID() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.ClearOrderID()
	})
}

// SetCouponID sets the "coupon_id" field.
func (u *OrderCouponUpsertOne) SetCouponID(v uuid.UUID) *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetCouponID(v)
	})
}

// UpdateCouponID sets the "coupon_id" field to the value that was provided on create.
func (u *OrderCouponUpsertOne) UpdateCouponID() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateCouponID()
	})
}

// ClearCouponID clears the value of the "coupon_id" field.
func (u *OrderCouponUpsertOne) ClearCouponID() *OrderCouponUpsertOne {
	return u.Update(func(s *OrderCouponUpsert) {
		s.ClearCouponID()
	})
}

// Exec executes the query.
func (u *OrderCouponUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderCouponCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderCouponUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *OrderCouponUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *OrderCouponUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// OrderCouponCreateBulk is the builder for creating many OrderCoupon entities in bulk.
type OrderCouponCreateBulk struct {
	config
	builders []*OrderCouponCreate
	conflict []sql.ConflictOption
}

// Save creates the OrderCoupon entities in the database.
func (occb *OrderCouponCreateBulk) Save(ctx context.Context) ([]*OrderCoupon, error) {
	specs := make([]*sqlgraph.CreateSpec, len(occb.builders))
	nodes := make([]*OrderCoupon, len(occb.builders))
	mutators := make([]Mutator, len(occb.builders))
	for i := range occb.builders {
		func(i int, root context.Context) {
			builder := occb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderCouponMutation)
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
					_, err = mutators[i+1].Mutate(root, occb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = occb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, occb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, occb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (occb *OrderCouponCreateBulk) SaveX(ctx context.Context) []*OrderCoupon {
	v, err := occb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (occb *OrderCouponCreateBulk) Exec(ctx context.Context) error {
	_, err := occb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (occb *OrderCouponCreateBulk) ExecX(ctx context.Context) {
	if err := occb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.OrderCoupon.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.OrderCouponUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (occb *OrderCouponCreateBulk) OnConflict(opts ...sql.ConflictOption) *OrderCouponUpsertBulk {
	occb.conflict = opts
	return &OrderCouponUpsertBulk{
		create: occb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.OrderCoupon.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (occb *OrderCouponCreateBulk) OnConflictColumns(columns ...string) *OrderCouponUpsertBulk {
	occb.conflict = append(occb.conflict, sql.ConflictColumns(columns...))
	return &OrderCouponUpsertBulk{
		create: occb,
	}
}

// OrderCouponUpsertBulk is the builder for "upsert"-ing
// a bulk of OrderCoupon nodes.
type OrderCouponUpsertBulk struct {
	create *OrderCouponCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.OrderCoupon.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ordercoupon.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *OrderCouponUpsertBulk) UpdateNewValues() *OrderCouponUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(ordercoupon.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.OrderCoupon.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *OrderCouponUpsertBulk) Ignore() *OrderCouponUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *OrderCouponUpsertBulk) DoNothing() *OrderCouponUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the OrderCouponCreateBulk.OnConflict
// documentation for more info.
func (u *OrderCouponUpsertBulk) Update(set func(*OrderCouponUpsert)) *OrderCouponUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&OrderCouponUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *OrderCouponUpsertBulk) SetCreatedAt(v uint32) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *OrderCouponUpsertBulk) AddCreatedAt(v uint32) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *OrderCouponUpsertBulk) UpdateCreatedAt() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *OrderCouponUpsertBulk) SetUpdatedAt(v uint32) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *OrderCouponUpsertBulk) AddUpdatedAt(v uint32) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *OrderCouponUpsertBulk) UpdateUpdatedAt() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *OrderCouponUpsertBulk) SetDeletedAt(v uint32) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *OrderCouponUpsertBulk) AddDeletedAt(v uint32) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *OrderCouponUpsertBulk) UpdateDeletedAt() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *OrderCouponUpsertBulk) SetEntID(v uuid.UUID) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *OrderCouponUpsertBulk) UpdateEntID() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateEntID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *OrderCouponUpsertBulk) SetOrderID(v uuid.UUID) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *OrderCouponUpsertBulk) UpdateOrderID() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *OrderCouponUpsertBulk) ClearOrderID() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.ClearOrderID()
	})
}

// SetCouponID sets the "coupon_id" field.
func (u *OrderCouponUpsertBulk) SetCouponID(v uuid.UUID) *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.SetCouponID(v)
	})
}

// UpdateCouponID sets the "coupon_id" field to the value that was provided on create.
func (u *OrderCouponUpsertBulk) UpdateCouponID() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.UpdateCouponID()
	})
}

// ClearCouponID clears the value of the "coupon_id" field.
func (u *OrderCouponUpsertBulk) ClearCouponID() *OrderCouponUpsertBulk {
	return u.Update(func(s *OrderCouponUpsert) {
		s.ClearCouponID()
	})
}

// Exec executes the query.
func (u *OrderCouponUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the OrderCouponCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for OrderCouponCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *OrderCouponUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
