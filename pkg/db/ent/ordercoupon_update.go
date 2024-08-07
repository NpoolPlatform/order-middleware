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
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderCouponUpdate is the builder for updating OrderCoupon entities.
type OrderCouponUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderCouponMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderCouponUpdate builder.
func (ocu *OrderCouponUpdate) Where(ps ...predicate.OrderCoupon) *OrderCouponUpdate {
	ocu.mutation.Where(ps...)
	return ocu
}

// SetCreatedAt sets the "created_at" field.
func (ocu *OrderCouponUpdate) SetCreatedAt(u uint32) *OrderCouponUpdate {
	ocu.mutation.ResetCreatedAt()
	ocu.mutation.SetCreatedAt(u)
	return ocu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ocu *OrderCouponUpdate) SetNillableCreatedAt(u *uint32) *OrderCouponUpdate {
	if u != nil {
		ocu.SetCreatedAt(*u)
	}
	return ocu
}

// AddCreatedAt adds u to the "created_at" field.
func (ocu *OrderCouponUpdate) AddCreatedAt(u int32) *OrderCouponUpdate {
	ocu.mutation.AddCreatedAt(u)
	return ocu
}

// SetUpdatedAt sets the "updated_at" field.
func (ocu *OrderCouponUpdate) SetUpdatedAt(u uint32) *OrderCouponUpdate {
	ocu.mutation.ResetUpdatedAt()
	ocu.mutation.SetUpdatedAt(u)
	return ocu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ocu *OrderCouponUpdate) AddUpdatedAt(u int32) *OrderCouponUpdate {
	ocu.mutation.AddUpdatedAt(u)
	return ocu
}

// SetDeletedAt sets the "deleted_at" field.
func (ocu *OrderCouponUpdate) SetDeletedAt(u uint32) *OrderCouponUpdate {
	ocu.mutation.ResetDeletedAt()
	ocu.mutation.SetDeletedAt(u)
	return ocu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ocu *OrderCouponUpdate) SetNillableDeletedAt(u *uint32) *OrderCouponUpdate {
	if u != nil {
		ocu.SetDeletedAt(*u)
	}
	return ocu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ocu *OrderCouponUpdate) AddDeletedAt(u int32) *OrderCouponUpdate {
	ocu.mutation.AddDeletedAt(u)
	return ocu
}

// SetEntID sets the "ent_id" field.
func (ocu *OrderCouponUpdate) SetEntID(u uuid.UUID) *OrderCouponUpdate {
	ocu.mutation.SetEntID(u)
	return ocu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ocu *OrderCouponUpdate) SetNillableEntID(u *uuid.UUID) *OrderCouponUpdate {
	if u != nil {
		ocu.SetEntID(*u)
	}
	return ocu
}

// SetOrderID sets the "order_id" field.
func (ocu *OrderCouponUpdate) SetOrderID(u uuid.UUID) *OrderCouponUpdate {
	ocu.mutation.SetOrderID(u)
	return ocu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (ocu *OrderCouponUpdate) SetNillableOrderID(u *uuid.UUID) *OrderCouponUpdate {
	if u != nil {
		ocu.SetOrderID(*u)
	}
	return ocu
}

// ClearOrderID clears the value of the "order_id" field.
func (ocu *OrderCouponUpdate) ClearOrderID() *OrderCouponUpdate {
	ocu.mutation.ClearOrderID()
	return ocu
}

// SetCouponID sets the "coupon_id" field.
func (ocu *OrderCouponUpdate) SetCouponID(u uuid.UUID) *OrderCouponUpdate {
	ocu.mutation.SetCouponID(u)
	return ocu
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ocu *OrderCouponUpdate) SetNillableCouponID(u *uuid.UUID) *OrderCouponUpdate {
	if u != nil {
		ocu.SetCouponID(*u)
	}
	return ocu
}

// ClearCouponID clears the value of the "coupon_id" field.
func (ocu *OrderCouponUpdate) ClearCouponID() *OrderCouponUpdate {
	ocu.mutation.ClearCouponID()
	return ocu
}

// Mutation returns the OrderCouponMutation object of the builder.
func (ocu *OrderCouponUpdate) Mutation() *OrderCouponMutation {
	return ocu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ocu *OrderCouponUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ocu.defaults(); err != nil {
		return 0, err
	}
	if len(ocu.hooks) == 0 {
		affected, err = ocu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocu.mutation = mutation
			affected, err = ocu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ocu.hooks) - 1; i >= 0; i-- {
			if ocu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ocu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ocu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ocu *OrderCouponUpdate) SaveX(ctx context.Context) int {
	affected, err := ocu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ocu *OrderCouponUpdate) Exec(ctx context.Context) error {
	_, err := ocu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocu *OrderCouponUpdate) ExecX(ctx context.Context) {
	if err := ocu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ocu *OrderCouponUpdate) defaults() error {
	if _, ok := ocu.mutation.UpdatedAt(); !ok {
		if ordercoupon.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ordercoupon.UpdateDefaultUpdatedAt()
		ocu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ocu *OrderCouponUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderCouponUpdate {
	ocu.modifiers = append(ocu.modifiers, modifiers...)
	return ocu
}

func (ocu *OrderCouponUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordercoupon.Table,
			Columns: ordercoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: ordercoupon.FieldID,
			},
		},
	}
	if ps := ocu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldCreatedAt,
		})
	}
	if value, ok := ocu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldCreatedAt,
		})
	}
	if value, ok := ocu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ocu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ocu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldDeletedAt,
		})
	}
	if value, ok := ocu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldDeletedAt,
		})
	}
	if value, ok := ocu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldEntID,
		})
	}
	if value, ok := ocu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldOrderID,
		})
	}
	if ocu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: ordercoupon.FieldOrderID,
		})
	}
	if value, ok := ocu.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldCouponID,
		})
	}
	if ocu.mutation.CouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: ordercoupon.FieldCouponID,
		})
	}
	_spec.Modifiers = ocu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ocu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordercoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderCouponUpdateOne is the builder for updating a single OrderCoupon entity.
type OrderCouponUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderCouponMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ocuo *OrderCouponUpdateOne) SetCreatedAt(u uint32) *OrderCouponUpdateOne {
	ocuo.mutation.ResetCreatedAt()
	ocuo.mutation.SetCreatedAt(u)
	return ocuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ocuo *OrderCouponUpdateOne) SetNillableCreatedAt(u *uint32) *OrderCouponUpdateOne {
	if u != nil {
		ocuo.SetCreatedAt(*u)
	}
	return ocuo
}

// AddCreatedAt adds u to the "created_at" field.
func (ocuo *OrderCouponUpdateOne) AddCreatedAt(u int32) *OrderCouponUpdateOne {
	ocuo.mutation.AddCreatedAt(u)
	return ocuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ocuo *OrderCouponUpdateOne) SetUpdatedAt(u uint32) *OrderCouponUpdateOne {
	ocuo.mutation.ResetUpdatedAt()
	ocuo.mutation.SetUpdatedAt(u)
	return ocuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ocuo *OrderCouponUpdateOne) AddUpdatedAt(u int32) *OrderCouponUpdateOne {
	ocuo.mutation.AddUpdatedAt(u)
	return ocuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ocuo *OrderCouponUpdateOne) SetDeletedAt(u uint32) *OrderCouponUpdateOne {
	ocuo.mutation.ResetDeletedAt()
	ocuo.mutation.SetDeletedAt(u)
	return ocuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ocuo *OrderCouponUpdateOne) SetNillableDeletedAt(u *uint32) *OrderCouponUpdateOne {
	if u != nil {
		ocuo.SetDeletedAt(*u)
	}
	return ocuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ocuo *OrderCouponUpdateOne) AddDeletedAt(u int32) *OrderCouponUpdateOne {
	ocuo.mutation.AddDeletedAt(u)
	return ocuo
}

// SetEntID sets the "ent_id" field.
func (ocuo *OrderCouponUpdateOne) SetEntID(u uuid.UUID) *OrderCouponUpdateOne {
	ocuo.mutation.SetEntID(u)
	return ocuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ocuo *OrderCouponUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderCouponUpdateOne {
	if u != nil {
		ocuo.SetEntID(*u)
	}
	return ocuo
}

// SetOrderID sets the "order_id" field.
func (ocuo *OrderCouponUpdateOne) SetOrderID(u uuid.UUID) *OrderCouponUpdateOne {
	ocuo.mutation.SetOrderID(u)
	return ocuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (ocuo *OrderCouponUpdateOne) SetNillableOrderID(u *uuid.UUID) *OrderCouponUpdateOne {
	if u != nil {
		ocuo.SetOrderID(*u)
	}
	return ocuo
}

// ClearOrderID clears the value of the "order_id" field.
func (ocuo *OrderCouponUpdateOne) ClearOrderID() *OrderCouponUpdateOne {
	ocuo.mutation.ClearOrderID()
	return ocuo
}

// SetCouponID sets the "coupon_id" field.
func (ocuo *OrderCouponUpdateOne) SetCouponID(u uuid.UUID) *OrderCouponUpdateOne {
	ocuo.mutation.SetCouponID(u)
	return ocuo
}

// SetNillableCouponID sets the "coupon_id" field if the given value is not nil.
func (ocuo *OrderCouponUpdateOne) SetNillableCouponID(u *uuid.UUID) *OrderCouponUpdateOne {
	if u != nil {
		ocuo.SetCouponID(*u)
	}
	return ocuo
}

// ClearCouponID clears the value of the "coupon_id" field.
func (ocuo *OrderCouponUpdateOne) ClearCouponID() *OrderCouponUpdateOne {
	ocuo.mutation.ClearCouponID()
	return ocuo
}

// Mutation returns the OrderCouponMutation object of the builder.
func (ocuo *OrderCouponUpdateOne) Mutation() *OrderCouponMutation {
	return ocuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ocuo *OrderCouponUpdateOne) Select(field string, fields ...string) *OrderCouponUpdateOne {
	ocuo.fields = append([]string{field}, fields...)
	return ocuo
}

// Save executes the query and returns the updated OrderCoupon entity.
func (ocuo *OrderCouponUpdateOne) Save(ctx context.Context) (*OrderCoupon, error) {
	var (
		err  error
		node *OrderCoupon
	)
	if err := ocuo.defaults(); err != nil {
		return nil, err
	}
	if len(ocuo.hooks) == 0 {
		node, err = ocuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderCouponMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ocuo.mutation = mutation
			node, err = ocuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ocuo.hooks) - 1; i >= 0; i-- {
			if ocuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ocuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ocuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (ocuo *OrderCouponUpdateOne) SaveX(ctx context.Context) *OrderCoupon {
	node, err := ocuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ocuo *OrderCouponUpdateOne) Exec(ctx context.Context) error {
	_, err := ocuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocuo *OrderCouponUpdateOne) ExecX(ctx context.Context) {
	if err := ocuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ocuo *OrderCouponUpdateOne) defaults() error {
	if _, ok := ocuo.mutation.UpdatedAt(); !ok {
		if ordercoupon.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized ordercoupon.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := ordercoupon.UpdateDefaultUpdatedAt()
		ocuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ocuo *OrderCouponUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderCouponUpdateOne {
	ocuo.modifiers = append(ocuo.modifiers, modifiers...)
	return ocuo
}

func (ocuo *OrderCouponUpdateOne) sqlSave(ctx context.Context) (_node *OrderCoupon, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ordercoupon.Table,
			Columns: ordercoupon.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: ordercoupon.FieldID,
			},
		},
	}
	id, ok := ocuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderCoupon.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ocuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordercoupon.FieldID)
		for _, f := range fields {
			if !ordercoupon.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ordercoupon.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ocuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ocuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldCreatedAt,
		})
	}
	if value, ok := ocuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldCreatedAt,
		})
	}
	if value, ok := ocuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ocuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldUpdatedAt,
		})
	}
	if value, ok := ocuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldDeletedAt,
		})
	}
	if value, ok := ocuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: ordercoupon.FieldDeletedAt,
		})
	}
	if value, ok := ocuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldEntID,
		})
	}
	if value, ok := ocuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldOrderID,
		})
	}
	if ocuo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: ordercoupon.FieldOrderID,
		})
	}
	if value, ok := ocuo.mutation.CouponID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ordercoupon.FieldCouponID,
		})
	}
	if ocuo.mutation.CouponIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: ordercoupon.FieldCouponID,
		})
	}
	_spec.Modifiers = ocuo.modifiers
	_node = &OrderCoupon{config: ocuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ocuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordercoupon.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
