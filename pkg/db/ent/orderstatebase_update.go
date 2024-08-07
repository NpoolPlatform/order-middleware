// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderStateBaseUpdate is the builder for updating OrderStateBase entities.
type OrderStateBaseUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderStateBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderStateBaseUpdate builder.
func (osbu *OrderStateBaseUpdate) Where(ps ...predicate.OrderStateBase) *OrderStateBaseUpdate {
	osbu.mutation.Where(ps...)
	return osbu
}

// SetCreatedAt sets the "created_at" field.
func (osbu *OrderStateBaseUpdate) SetCreatedAt(u uint32) *OrderStateBaseUpdate {
	osbu.mutation.ResetCreatedAt()
	osbu.mutation.SetCreatedAt(u)
	return osbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableCreatedAt(u *uint32) *OrderStateBaseUpdate {
	if u != nil {
		osbu.SetCreatedAt(*u)
	}
	return osbu
}

// AddCreatedAt adds u to the "created_at" field.
func (osbu *OrderStateBaseUpdate) AddCreatedAt(u int32) *OrderStateBaseUpdate {
	osbu.mutation.AddCreatedAt(u)
	return osbu
}

// SetUpdatedAt sets the "updated_at" field.
func (osbu *OrderStateBaseUpdate) SetUpdatedAt(u uint32) *OrderStateBaseUpdate {
	osbu.mutation.ResetUpdatedAt()
	osbu.mutation.SetUpdatedAt(u)
	return osbu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (osbu *OrderStateBaseUpdate) AddUpdatedAt(u int32) *OrderStateBaseUpdate {
	osbu.mutation.AddUpdatedAt(u)
	return osbu
}

// SetDeletedAt sets the "deleted_at" field.
func (osbu *OrderStateBaseUpdate) SetDeletedAt(u uint32) *OrderStateBaseUpdate {
	osbu.mutation.ResetDeletedAt()
	osbu.mutation.SetDeletedAt(u)
	return osbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableDeletedAt(u *uint32) *OrderStateBaseUpdate {
	if u != nil {
		osbu.SetDeletedAt(*u)
	}
	return osbu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (osbu *OrderStateBaseUpdate) AddDeletedAt(u int32) *OrderStateBaseUpdate {
	osbu.mutation.AddDeletedAt(u)
	return osbu
}

// SetEntID sets the "ent_id" field.
func (osbu *OrderStateBaseUpdate) SetEntID(u uuid.UUID) *OrderStateBaseUpdate {
	osbu.mutation.SetEntID(u)
	return osbu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableEntID(u *uuid.UUID) *OrderStateBaseUpdate {
	if u != nil {
		osbu.SetEntID(*u)
	}
	return osbu
}

// SetOrderID sets the "order_id" field.
func (osbu *OrderStateBaseUpdate) SetOrderID(u uuid.UUID) *OrderStateBaseUpdate {
	osbu.mutation.SetOrderID(u)
	return osbu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableOrderID(u *uuid.UUID) *OrderStateBaseUpdate {
	if u != nil {
		osbu.SetOrderID(*u)
	}
	return osbu
}

// ClearOrderID clears the value of the "order_id" field.
func (osbu *OrderStateBaseUpdate) ClearOrderID() *OrderStateBaseUpdate {
	osbu.mutation.ClearOrderID()
	return osbu
}

// SetOrderState sets the "order_state" field.
func (osbu *OrderStateBaseUpdate) SetOrderState(s string) *OrderStateBaseUpdate {
	osbu.mutation.SetOrderState(s)
	return osbu
}

// SetNillableOrderState sets the "order_state" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableOrderState(s *string) *OrderStateBaseUpdate {
	if s != nil {
		osbu.SetOrderState(*s)
	}
	return osbu
}

// ClearOrderState clears the value of the "order_state" field.
func (osbu *OrderStateBaseUpdate) ClearOrderState() *OrderStateBaseUpdate {
	osbu.mutation.ClearOrderState()
	return osbu
}

// SetStartMode sets the "start_mode" field.
func (osbu *OrderStateBaseUpdate) SetStartMode(s string) *OrderStateBaseUpdate {
	osbu.mutation.SetStartMode(s)
	return osbu
}

// SetNillableStartMode sets the "start_mode" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableStartMode(s *string) *OrderStateBaseUpdate {
	if s != nil {
		osbu.SetStartMode(*s)
	}
	return osbu
}

// ClearStartMode clears the value of the "start_mode" field.
func (osbu *OrderStateBaseUpdate) ClearStartMode() *OrderStateBaseUpdate {
	osbu.mutation.ClearStartMode()
	return osbu
}

// SetStartAt sets the "start_at" field.
func (osbu *OrderStateBaseUpdate) SetStartAt(u uint32) *OrderStateBaseUpdate {
	osbu.mutation.ResetStartAt()
	osbu.mutation.SetStartAt(u)
	return osbu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableStartAt(u *uint32) *OrderStateBaseUpdate {
	if u != nil {
		osbu.SetStartAt(*u)
	}
	return osbu
}

// AddStartAt adds u to the "start_at" field.
func (osbu *OrderStateBaseUpdate) AddStartAt(u int32) *OrderStateBaseUpdate {
	osbu.mutation.AddStartAt(u)
	return osbu
}

// ClearStartAt clears the value of the "start_at" field.
func (osbu *OrderStateBaseUpdate) ClearStartAt() *OrderStateBaseUpdate {
	osbu.mutation.ClearStartAt()
	return osbu
}

// SetLastBenefitAt sets the "last_benefit_at" field.
func (osbu *OrderStateBaseUpdate) SetLastBenefitAt(u uint32) *OrderStateBaseUpdate {
	osbu.mutation.ResetLastBenefitAt()
	osbu.mutation.SetLastBenefitAt(u)
	return osbu
}

// SetNillableLastBenefitAt sets the "last_benefit_at" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableLastBenefitAt(u *uint32) *OrderStateBaseUpdate {
	if u != nil {
		osbu.SetLastBenefitAt(*u)
	}
	return osbu
}

// AddLastBenefitAt adds u to the "last_benefit_at" field.
func (osbu *OrderStateBaseUpdate) AddLastBenefitAt(u int32) *OrderStateBaseUpdate {
	osbu.mutation.AddLastBenefitAt(u)
	return osbu
}

// ClearLastBenefitAt clears the value of the "last_benefit_at" field.
func (osbu *OrderStateBaseUpdate) ClearLastBenefitAt() *OrderStateBaseUpdate {
	osbu.mutation.ClearLastBenefitAt()
	return osbu
}

// SetBenefitState sets the "benefit_state" field.
func (osbu *OrderStateBaseUpdate) SetBenefitState(s string) *OrderStateBaseUpdate {
	osbu.mutation.SetBenefitState(s)
	return osbu
}

// SetNillableBenefitState sets the "benefit_state" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillableBenefitState(s *string) *OrderStateBaseUpdate {
	if s != nil {
		osbu.SetBenefitState(*s)
	}
	return osbu
}

// ClearBenefitState clears the value of the "benefit_state" field.
func (osbu *OrderStateBaseUpdate) ClearBenefitState() *OrderStateBaseUpdate {
	osbu.mutation.ClearBenefitState()
	return osbu
}

// SetPaymentType sets the "payment_type" field.
func (osbu *OrderStateBaseUpdate) SetPaymentType(s string) *OrderStateBaseUpdate {
	osbu.mutation.SetPaymentType(s)
	return osbu
}

// SetNillablePaymentType sets the "payment_type" field if the given value is not nil.
func (osbu *OrderStateBaseUpdate) SetNillablePaymentType(s *string) *OrderStateBaseUpdate {
	if s != nil {
		osbu.SetPaymentType(*s)
	}
	return osbu
}

// ClearPaymentType clears the value of the "payment_type" field.
func (osbu *OrderStateBaseUpdate) ClearPaymentType() *OrderStateBaseUpdate {
	osbu.mutation.ClearPaymentType()
	return osbu
}

// Mutation returns the OrderStateBaseMutation object of the builder.
func (osbu *OrderStateBaseUpdate) Mutation() *OrderStateBaseMutation {
	return osbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osbu *OrderStateBaseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := osbu.defaults(); err != nil {
		return 0, err
	}
	if len(osbu.hooks) == 0 {
		affected, err = osbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStateBaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osbu.mutation = mutation
			affected, err = osbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(osbu.hooks) - 1; i >= 0; i-- {
			if osbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = osbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (osbu *OrderStateBaseUpdate) SaveX(ctx context.Context) int {
	affected, err := osbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osbu *OrderStateBaseUpdate) Exec(ctx context.Context) error {
	_, err := osbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osbu *OrderStateBaseUpdate) ExecX(ctx context.Context) {
	if err := osbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osbu *OrderStateBaseUpdate) defaults() error {
	if _, ok := osbu.mutation.UpdatedAt(); !ok {
		if orderstatebase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderstatebase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderstatebase.UpdateDefaultUpdatedAt()
		osbu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (osbu *OrderStateBaseUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderStateBaseUpdate {
	osbu.modifiers = append(osbu.modifiers, modifiers...)
	return osbu
}

func (osbu *OrderStateBaseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstatebase.Table,
			Columns: orderstatebase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderstatebase.FieldID,
			},
		},
	}
	if ps := osbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldCreatedAt,
		})
	}
	if value, ok := osbu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldCreatedAt,
		})
	}
	if value, ok := osbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldUpdatedAt,
		})
	}
	if value, ok := osbu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldUpdatedAt,
		})
	}
	if value, ok := osbu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldDeletedAt,
		})
	}
	if value, ok := osbu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldDeletedAt,
		})
	}
	if value, ok := osbu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderstatebase.FieldEntID,
		})
	}
	if value, ok := osbu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderstatebase.FieldOrderID,
		})
	}
	if osbu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderstatebase.FieldOrderID,
		})
	}
	if value, ok := osbu.mutation.OrderState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldOrderState,
		})
	}
	if osbu.mutation.OrderStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldOrderState,
		})
	}
	if value, ok := osbu.mutation.StartMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldStartMode,
		})
	}
	if osbu.mutation.StartModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldStartMode,
		})
	}
	if value, ok := osbu.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldStartAt,
		})
	}
	if value, ok := osbu.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldStartAt,
		})
	}
	if osbu.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstatebase.FieldStartAt,
		})
	}
	if value, ok := osbu.mutation.LastBenefitAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldLastBenefitAt,
		})
	}
	if value, ok := osbu.mutation.AddedLastBenefitAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldLastBenefitAt,
		})
	}
	if osbu.mutation.LastBenefitAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstatebase.FieldLastBenefitAt,
		})
	}
	if value, ok := osbu.mutation.BenefitState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldBenefitState,
		})
	}
	if osbu.mutation.BenefitStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldBenefitState,
		})
	}
	if value, ok := osbu.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldPaymentType,
		})
	}
	if osbu.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldPaymentType,
		})
	}
	_spec.Modifiers = osbu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, osbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstatebase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderStateBaseUpdateOne is the builder for updating a single OrderStateBase entity.
type OrderStateBaseUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderStateBaseMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (osbuo *OrderStateBaseUpdateOne) SetCreatedAt(u uint32) *OrderStateBaseUpdateOne {
	osbuo.mutation.ResetCreatedAt()
	osbuo.mutation.SetCreatedAt(u)
	return osbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableCreatedAt(u *uint32) *OrderStateBaseUpdateOne {
	if u != nil {
		osbuo.SetCreatedAt(*u)
	}
	return osbuo
}

// AddCreatedAt adds u to the "created_at" field.
func (osbuo *OrderStateBaseUpdateOne) AddCreatedAt(u int32) *OrderStateBaseUpdateOne {
	osbuo.mutation.AddCreatedAt(u)
	return osbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (osbuo *OrderStateBaseUpdateOne) SetUpdatedAt(u uint32) *OrderStateBaseUpdateOne {
	osbuo.mutation.ResetUpdatedAt()
	osbuo.mutation.SetUpdatedAt(u)
	return osbuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (osbuo *OrderStateBaseUpdateOne) AddUpdatedAt(u int32) *OrderStateBaseUpdateOne {
	osbuo.mutation.AddUpdatedAt(u)
	return osbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (osbuo *OrderStateBaseUpdateOne) SetDeletedAt(u uint32) *OrderStateBaseUpdateOne {
	osbuo.mutation.ResetDeletedAt()
	osbuo.mutation.SetDeletedAt(u)
	return osbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableDeletedAt(u *uint32) *OrderStateBaseUpdateOne {
	if u != nil {
		osbuo.SetDeletedAt(*u)
	}
	return osbuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (osbuo *OrderStateBaseUpdateOne) AddDeletedAt(u int32) *OrderStateBaseUpdateOne {
	osbuo.mutation.AddDeletedAt(u)
	return osbuo
}

// SetEntID sets the "ent_id" field.
func (osbuo *OrderStateBaseUpdateOne) SetEntID(u uuid.UUID) *OrderStateBaseUpdateOne {
	osbuo.mutation.SetEntID(u)
	return osbuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableEntID(u *uuid.UUID) *OrderStateBaseUpdateOne {
	if u != nil {
		osbuo.SetEntID(*u)
	}
	return osbuo
}

// SetOrderID sets the "order_id" field.
func (osbuo *OrderStateBaseUpdateOne) SetOrderID(u uuid.UUID) *OrderStateBaseUpdateOne {
	osbuo.mutation.SetOrderID(u)
	return osbuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableOrderID(u *uuid.UUID) *OrderStateBaseUpdateOne {
	if u != nil {
		osbuo.SetOrderID(*u)
	}
	return osbuo
}

// ClearOrderID clears the value of the "order_id" field.
func (osbuo *OrderStateBaseUpdateOne) ClearOrderID() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearOrderID()
	return osbuo
}

// SetOrderState sets the "order_state" field.
func (osbuo *OrderStateBaseUpdateOne) SetOrderState(s string) *OrderStateBaseUpdateOne {
	osbuo.mutation.SetOrderState(s)
	return osbuo
}

// SetNillableOrderState sets the "order_state" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableOrderState(s *string) *OrderStateBaseUpdateOne {
	if s != nil {
		osbuo.SetOrderState(*s)
	}
	return osbuo
}

// ClearOrderState clears the value of the "order_state" field.
func (osbuo *OrderStateBaseUpdateOne) ClearOrderState() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearOrderState()
	return osbuo
}

// SetStartMode sets the "start_mode" field.
func (osbuo *OrderStateBaseUpdateOne) SetStartMode(s string) *OrderStateBaseUpdateOne {
	osbuo.mutation.SetStartMode(s)
	return osbuo
}

// SetNillableStartMode sets the "start_mode" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableStartMode(s *string) *OrderStateBaseUpdateOne {
	if s != nil {
		osbuo.SetStartMode(*s)
	}
	return osbuo
}

// ClearStartMode clears the value of the "start_mode" field.
func (osbuo *OrderStateBaseUpdateOne) ClearStartMode() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearStartMode()
	return osbuo
}

// SetStartAt sets the "start_at" field.
func (osbuo *OrderStateBaseUpdateOne) SetStartAt(u uint32) *OrderStateBaseUpdateOne {
	osbuo.mutation.ResetStartAt()
	osbuo.mutation.SetStartAt(u)
	return osbuo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableStartAt(u *uint32) *OrderStateBaseUpdateOne {
	if u != nil {
		osbuo.SetStartAt(*u)
	}
	return osbuo
}

// AddStartAt adds u to the "start_at" field.
func (osbuo *OrderStateBaseUpdateOne) AddStartAt(u int32) *OrderStateBaseUpdateOne {
	osbuo.mutation.AddStartAt(u)
	return osbuo
}

// ClearStartAt clears the value of the "start_at" field.
func (osbuo *OrderStateBaseUpdateOne) ClearStartAt() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearStartAt()
	return osbuo
}

// SetLastBenefitAt sets the "last_benefit_at" field.
func (osbuo *OrderStateBaseUpdateOne) SetLastBenefitAt(u uint32) *OrderStateBaseUpdateOne {
	osbuo.mutation.ResetLastBenefitAt()
	osbuo.mutation.SetLastBenefitAt(u)
	return osbuo
}

// SetNillableLastBenefitAt sets the "last_benefit_at" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableLastBenefitAt(u *uint32) *OrderStateBaseUpdateOne {
	if u != nil {
		osbuo.SetLastBenefitAt(*u)
	}
	return osbuo
}

// AddLastBenefitAt adds u to the "last_benefit_at" field.
func (osbuo *OrderStateBaseUpdateOne) AddLastBenefitAt(u int32) *OrderStateBaseUpdateOne {
	osbuo.mutation.AddLastBenefitAt(u)
	return osbuo
}

// ClearLastBenefitAt clears the value of the "last_benefit_at" field.
func (osbuo *OrderStateBaseUpdateOne) ClearLastBenefitAt() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearLastBenefitAt()
	return osbuo
}

// SetBenefitState sets the "benefit_state" field.
func (osbuo *OrderStateBaseUpdateOne) SetBenefitState(s string) *OrderStateBaseUpdateOne {
	osbuo.mutation.SetBenefitState(s)
	return osbuo
}

// SetNillableBenefitState sets the "benefit_state" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillableBenefitState(s *string) *OrderStateBaseUpdateOne {
	if s != nil {
		osbuo.SetBenefitState(*s)
	}
	return osbuo
}

// ClearBenefitState clears the value of the "benefit_state" field.
func (osbuo *OrderStateBaseUpdateOne) ClearBenefitState() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearBenefitState()
	return osbuo
}

// SetPaymentType sets the "payment_type" field.
func (osbuo *OrderStateBaseUpdateOne) SetPaymentType(s string) *OrderStateBaseUpdateOne {
	osbuo.mutation.SetPaymentType(s)
	return osbuo
}

// SetNillablePaymentType sets the "payment_type" field if the given value is not nil.
func (osbuo *OrderStateBaseUpdateOne) SetNillablePaymentType(s *string) *OrderStateBaseUpdateOne {
	if s != nil {
		osbuo.SetPaymentType(*s)
	}
	return osbuo
}

// ClearPaymentType clears the value of the "payment_type" field.
func (osbuo *OrderStateBaseUpdateOne) ClearPaymentType() *OrderStateBaseUpdateOne {
	osbuo.mutation.ClearPaymentType()
	return osbuo
}

// Mutation returns the OrderStateBaseMutation object of the builder.
func (osbuo *OrderStateBaseUpdateOne) Mutation() *OrderStateBaseMutation {
	return osbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osbuo *OrderStateBaseUpdateOne) Select(field string, fields ...string) *OrderStateBaseUpdateOne {
	osbuo.fields = append([]string{field}, fields...)
	return osbuo
}

// Save executes the query and returns the updated OrderStateBase entity.
func (osbuo *OrderStateBaseUpdateOne) Save(ctx context.Context) (*OrderStateBase, error) {
	var (
		err  error
		node *OrderStateBase
	)
	if err := osbuo.defaults(); err != nil {
		return nil, err
	}
	if len(osbuo.hooks) == 0 {
		node, err = osbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStateBaseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osbuo.mutation = mutation
			node, err = osbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osbuo.hooks) - 1; i >= 0; i-- {
			if osbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = osbuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, osbuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderStateBase)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderStateBaseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (osbuo *OrderStateBaseUpdateOne) SaveX(ctx context.Context) *OrderStateBase {
	node, err := osbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osbuo *OrderStateBaseUpdateOne) Exec(ctx context.Context) error {
	_, err := osbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osbuo *OrderStateBaseUpdateOne) ExecX(ctx context.Context) {
	if err := osbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osbuo *OrderStateBaseUpdateOne) defaults() error {
	if _, ok := osbuo.mutation.UpdatedAt(); !ok {
		if orderstatebase.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderstatebase.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderstatebase.UpdateDefaultUpdatedAt()
		osbuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (osbuo *OrderStateBaseUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderStateBaseUpdateOne {
	osbuo.modifiers = append(osbuo.modifiers, modifiers...)
	return osbuo
}

func (osbuo *OrderStateBaseUpdateOne) sqlSave(ctx context.Context) (_node *OrderStateBase, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstatebase.Table,
			Columns: orderstatebase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderstatebase.FieldID,
			},
		},
	}
	id, ok := osbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderStateBase.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatebase.FieldID)
		for _, f := range fields {
			if !orderstatebase.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderstatebase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldCreatedAt,
		})
	}
	if value, ok := osbuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldCreatedAt,
		})
	}
	if value, ok := osbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldUpdatedAt,
		})
	}
	if value, ok := osbuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldUpdatedAt,
		})
	}
	if value, ok := osbuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldDeletedAt,
		})
	}
	if value, ok := osbuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldDeletedAt,
		})
	}
	if value, ok := osbuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderstatebase.FieldEntID,
		})
	}
	if value, ok := osbuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderstatebase.FieldOrderID,
		})
	}
	if osbuo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: orderstatebase.FieldOrderID,
		})
	}
	if value, ok := osbuo.mutation.OrderState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldOrderState,
		})
	}
	if osbuo.mutation.OrderStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldOrderState,
		})
	}
	if value, ok := osbuo.mutation.StartMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldStartMode,
		})
	}
	if osbuo.mutation.StartModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldStartMode,
		})
	}
	if value, ok := osbuo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldStartAt,
		})
	}
	if value, ok := osbuo.mutation.AddedStartAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldStartAt,
		})
	}
	if osbuo.mutation.StartAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstatebase.FieldStartAt,
		})
	}
	if value, ok := osbuo.mutation.LastBenefitAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldLastBenefitAt,
		})
	}
	if value, ok := osbuo.mutation.AddedLastBenefitAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstatebase.FieldLastBenefitAt,
		})
	}
	if osbuo.mutation.LastBenefitAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstatebase.FieldLastBenefitAt,
		})
	}
	if value, ok := osbuo.mutation.BenefitState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldBenefitState,
		})
	}
	if osbuo.mutation.BenefitStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldBenefitState,
		})
	}
	if value, ok := osbuo.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstatebase.FieldPaymentType,
		})
	}
	if osbuo.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstatebase.FieldPaymentType,
		})
	}
	_spec.Modifiers = osbuo.modifiers
	_node = &OrderStateBase{config: osbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstatebase.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
