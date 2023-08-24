// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderStateUpdate is the builder for updating OrderState entities.
type OrderStateUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderStateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderStateUpdate builder.
func (osu *OrderStateUpdate) Where(ps ...predicate.OrderState) *OrderStateUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetCreatedAt sets the "created_at" field.
func (osu *OrderStateUpdate) SetCreatedAt(u uint32) *OrderStateUpdate {
	osu.mutation.ResetCreatedAt()
	osu.mutation.SetCreatedAt(u)
	return osu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableCreatedAt(u *uint32) *OrderStateUpdate {
	if u != nil {
		osu.SetCreatedAt(*u)
	}
	return osu
}

// AddCreatedAt adds u to the "created_at" field.
func (osu *OrderStateUpdate) AddCreatedAt(u int32) *OrderStateUpdate {
	osu.mutation.AddCreatedAt(u)
	return osu
}

// SetUpdatedAt sets the "updated_at" field.
func (osu *OrderStateUpdate) SetUpdatedAt(u uint32) *OrderStateUpdate {
	osu.mutation.ResetUpdatedAt()
	osu.mutation.SetUpdatedAt(u)
	return osu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (osu *OrderStateUpdate) AddUpdatedAt(u int32) *OrderStateUpdate {
	osu.mutation.AddUpdatedAt(u)
	return osu
}

// SetDeletedAt sets the "deleted_at" field.
func (osu *OrderStateUpdate) SetDeletedAt(u uint32) *OrderStateUpdate {
	osu.mutation.ResetDeletedAt()
	osu.mutation.SetDeletedAt(u)
	return osu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableDeletedAt(u *uint32) *OrderStateUpdate {
	if u != nil {
		osu.SetDeletedAt(*u)
	}
	return osu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (osu *OrderStateUpdate) AddDeletedAt(u int32) *OrderStateUpdate {
	osu.mutation.AddDeletedAt(u)
	return osu
}

// SetOrderID sets the "order_id" field.
func (osu *OrderStateUpdate) SetOrderID(u uuid.UUID) *OrderStateUpdate {
	osu.mutation.SetOrderID(u)
	return osu
}

// SetOrderState sets the "order_state" field.
func (osu *OrderStateUpdate) SetOrderState(s string) *OrderStateUpdate {
	osu.mutation.SetOrderState(s)
	return osu
}

// SetNillableOrderState sets the "order_state" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableOrderState(s *string) *OrderStateUpdate {
	if s != nil {
		osu.SetOrderState(*s)
	}
	return osu
}

// ClearOrderState clears the value of the "order_state" field.
func (osu *OrderStateUpdate) ClearOrderState() *OrderStateUpdate {
	osu.mutation.ClearOrderState()
	return osu
}

// SetEndAt sets the "end_at" field.
func (osu *OrderStateUpdate) SetEndAt(u uint32) *OrderStateUpdate {
	osu.mutation.ResetEndAt()
	osu.mutation.SetEndAt(u)
	return osu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableEndAt(u *uint32) *OrderStateUpdate {
	if u != nil {
		osu.SetEndAt(*u)
	}
	return osu
}

// AddEndAt adds u to the "end_at" field.
func (osu *OrderStateUpdate) AddEndAt(u int32) *OrderStateUpdate {
	osu.mutation.AddEndAt(u)
	return osu
}

// ClearEndAt clears the value of the "end_at" field.
func (osu *OrderStateUpdate) ClearEndAt() *OrderStateUpdate {
	osu.mutation.ClearEndAt()
	return osu
}

// SetLastBenefitAt sets the "last_benefit_at" field.
func (osu *OrderStateUpdate) SetLastBenefitAt(u uint32) *OrderStateUpdate {
	osu.mutation.ResetLastBenefitAt()
	osu.mutation.SetLastBenefitAt(u)
	return osu
}

// SetNillableLastBenefitAt sets the "last_benefit_at" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableLastBenefitAt(u *uint32) *OrderStateUpdate {
	if u != nil {
		osu.SetLastBenefitAt(*u)
	}
	return osu
}

// AddLastBenefitAt adds u to the "last_benefit_at" field.
func (osu *OrderStateUpdate) AddLastBenefitAt(u int32) *OrderStateUpdate {
	osu.mutation.AddLastBenefitAt(u)
	return osu
}

// ClearLastBenefitAt clears the value of the "last_benefit_at" field.
func (osu *OrderStateUpdate) ClearLastBenefitAt() *OrderStateUpdate {
	osu.mutation.ClearLastBenefitAt()
	return osu
}

// SetBenefitState sets the "benefit_state" field.
func (osu *OrderStateUpdate) SetBenefitState(s string) *OrderStateUpdate {
	osu.mutation.SetBenefitState(s)
	return osu
}

// SetNillableBenefitState sets the "benefit_state" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableBenefitState(s *string) *OrderStateUpdate {
	if s != nil {
		osu.SetBenefitState(*s)
	}
	return osu
}

// ClearBenefitState clears the value of the "benefit_state" field.
func (osu *OrderStateUpdate) ClearBenefitState() *OrderStateUpdate {
	osu.mutation.ClearBenefitState()
	return osu
}

// SetUserSetPaid sets the "user_set_paid" field.
func (osu *OrderStateUpdate) SetUserSetPaid(b bool) *OrderStateUpdate {
	osu.mutation.SetUserSetPaid(b)
	return osu
}

// SetNillableUserSetPaid sets the "user_set_paid" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableUserSetPaid(b *bool) *OrderStateUpdate {
	if b != nil {
		osu.SetUserSetPaid(*b)
	}
	return osu
}

// ClearUserSetPaid clears the value of the "user_set_paid" field.
func (osu *OrderStateUpdate) ClearUserSetPaid() *OrderStateUpdate {
	osu.mutation.ClearUserSetPaid()
	return osu
}

// SetUserSetCancelled sets the "user_set_cancelled" field.
func (osu *OrderStateUpdate) SetUserSetCancelled(b bool) *OrderStateUpdate {
	osu.mutation.SetUserSetCancelled(b)
	return osu
}

// SetNillableUserSetCancelled sets the "user_set_cancelled" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillableUserSetCancelled(b *bool) *OrderStateUpdate {
	if b != nil {
		osu.SetUserSetCancelled(*b)
	}
	return osu
}

// ClearUserSetCancelled clears the value of the "user_set_cancelled" field.
func (osu *OrderStateUpdate) ClearUserSetCancelled() *OrderStateUpdate {
	osu.mutation.ClearUserSetCancelled()
	return osu
}

// SetPaymentTransactionID sets the "payment_transaction_id" field.
func (osu *OrderStateUpdate) SetPaymentTransactionID(s string) *OrderStateUpdate {
	osu.mutation.SetPaymentTransactionID(s)
	return osu
}

// SetNillablePaymentTransactionID sets the "payment_transaction_id" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillablePaymentTransactionID(s *string) *OrderStateUpdate {
	if s != nil {
		osu.SetPaymentTransactionID(*s)
	}
	return osu
}

// ClearPaymentTransactionID clears the value of the "payment_transaction_id" field.
func (osu *OrderStateUpdate) ClearPaymentTransactionID() *OrderStateUpdate {
	osu.mutation.ClearPaymentTransactionID()
	return osu
}

// SetPaymentFinishAmount sets the "payment_finish_amount" field.
func (osu *OrderStateUpdate) SetPaymentFinishAmount(d decimal.Decimal) *OrderStateUpdate {
	osu.mutation.SetPaymentFinishAmount(d)
	return osu
}

// SetNillablePaymentFinishAmount sets the "payment_finish_amount" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillablePaymentFinishAmount(d *decimal.Decimal) *OrderStateUpdate {
	if d != nil {
		osu.SetPaymentFinishAmount(*d)
	}
	return osu
}

// ClearPaymentFinishAmount clears the value of the "payment_finish_amount" field.
func (osu *OrderStateUpdate) ClearPaymentFinishAmount() *OrderStateUpdate {
	osu.mutation.ClearPaymentFinishAmount()
	return osu
}

// SetPaymentState sets the "payment_state" field.
func (osu *OrderStateUpdate) SetPaymentState(s string) *OrderStateUpdate {
	osu.mutation.SetPaymentState(s)
	return osu
}

// SetNillablePaymentState sets the "payment_state" field if the given value is not nil.
func (osu *OrderStateUpdate) SetNillablePaymentState(s *string) *OrderStateUpdate {
	if s != nil {
		osu.SetPaymentState(*s)
	}
	return osu
}

// ClearPaymentState clears the value of the "payment_state" field.
func (osu *OrderStateUpdate) ClearPaymentState() *OrderStateUpdate {
	osu.mutation.ClearPaymentState()
	return osu
}

// Mutation returns the OrderStateMutation object of the builder.
func (osu *OrderStateUpdate) Mutation() *OrderStateMutation {
	return osu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OrderStateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := osu.defaults(); err != nil {
		return 0, err
	}
	if len(osu.hooks) == 0 {
		affected, err = osu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osu.mutation = mutation
			affected, err = osu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(osu.hooks) - 1; i >= 0; i-- {
			if osu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = osu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, osu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OrderStateUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OrderStateUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OrderStateUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osu *OrderStateUpdate) defaults() error {
	if _, ok := osu.mutation.UpdatedAt(); !ok {
		if orderstate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderstate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderstate.UpdateDefaultUpdatedAt()
		osu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (osu *OrderStateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderStateUpdate {
	osu.modifiers = append(osu.modifiers, modifiers...)
	return osu
}

func (osu *OrderStateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstate.Table,
			Columns: orderstate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderstate.FieldID,
			},
		},
	}
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldCreatedAt,
		})
	}
	if value, ok := osu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldCreatedAt,
		})
	}
	if value, ok := osu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldUpdatedAt,
		})
	}
	if value, ok := osu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldUpdatedAt,
		})
	}
	if value, ok := osu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldDeletedAt,
		})
	}
	if value, ok := osu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldDeletedAt,
		})
	}
	if value, ok := osu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderstate.FieldOrderID,
		})
	}
	if value, ok := osu.mutation.OrderState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldOrderState,
		})
	}
	if osu.mutation.OrderStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldOrderState,
		})
	}
	if value, ok := osu.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldEndAt,
		})
	}
	if value, ok := osu.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldEndAt,
		})
	}
	if osu.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstate.FieldEndAt,
		})
	}
	if value, ok := osu.mutation.LastBenefitAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldLastBenefitAt,
		})
	}
	if value, ok := osu.mutation.AddedLastBenefitAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldLastBenefitAt,
		})
	}
	if osu.mutation.LastBenefitAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstate.FieldLastBenefitAt,
		})
	}
	if value, ok := osu.mutation.BenefitState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldBenefitState,
		})
	}
	if osu.mutation.BenefitStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldBenefitState,
		})
	}
	if value, ok := osu.mutation.UserSetPaid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: orderstate.FieldUserSetPaid,
		})
	}
	if osu.mutation.UserSetPaidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: orderstate.FieldUserSetPaid,
		})
	}
	if value, ok := osu.mutation.UserSetCancelled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: orderstate.FieldUserSetCancelled,
		})
	}
	if osu.mutation.UserSetCancelledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: orderstate.FieldUserSetCancelled,
		})
	}
	if value, ok := osu.mutation.PaymentTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldPaymentTransactionID,
		})
	}
	if osu.mutation.PaymentTransactionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldPaymentTransactionID,
		})
	}
	if value, ok := osu.mutation.PaymentFinishAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderstate.FieldPaymentFinishAmount,
		})
	}
	if osu.mutation.PaymentFinishAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderstate.FieldPaymentFinishAmount,
		})
	}
	if value, ok := osu.mutation.PaymentState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldPaymentState,
		})
	}
	if osu.mutation.PaymentStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldPaymentState,
		})
	}
	_spec.Modifiers = osu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderStateUpdateOne is the builder for updating a single OrderState entity.
type OrderStateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderStateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (osuo *OrderStateUpdateOne) SetCreatedAt(u uint32) *OrderStateUpdateOne {
	osuo.mutation.ResetCreatedAt()
	osuo.mutation.SetCreatedAt(u)
	return osuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableCreatedAt(u *uint32) *OrderStateUpdateOne {
	if u != nil {
		osuo.SetCreatedAt(*u)
	}
	return osuo
}

// AddCreatedAt adds u to the "created_at" field.
func (osuo *OrderStateUpdateOne) AddCreatedAt(u int32) *OrderStateUpdateOne {
	osuo.mutation.AddCreatedAt(u)
	return osuo
}

// SetUpdatedAt sets the "updated_at" field.
func (osuo *OrderStateUpdateOne) SetUpdatedAt(u uint32) *OrderStateUpdateOne {
	osuo.mutation.ResetUpdatedAt()
	osuo.mutation.SetUpdatedAt(u)
	return osuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (osuo *OrderStateUpdateOne) AddUpdatedAt(u int32) *OrderStateUpdateOne {
	osuo.mutation.AddUpdatedAt(u)
	return osuo
}

// SetDeletedAt sets the "deleted_at" field.
func (osuo *OrderStateUpdateOne) SetDeletedAt(u uint32) *OrderStateUpdateOne {
	osuo.mutation.ResetDeletedAt()
	osuo.mutation.SetDeletedAt(u)
	return osuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableDeletedAt(u *uint32) *OrderStateUpdateOne {
	if u != nil {
		osuo.SetDeletedAt(*u)
	}
	return osuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (osuo *OrderStateUpdateOne) AddDeletedAt(u int32) *OrderStateUpdateOne {
	osuo.mutation.AddDeletedAt(u)
	return osuo
}

// SetOrderID sets the "order_id" field.
func (osuo *OrderStateUpdateOne) SetOrderID(u uuid.UUID) *OrderStateUpdateOne {
	osuo.mutation.SetOrderID(u)
	return osuo
}

// SetOrderState sets the "order_state" field.
func (osuo *OrderStateUpdateOne) SetOrderState(s string) *OrderStateUpdateOne {
	osuo.mutation.SetOrderState(s)
	return osuo
}

// SetNillableOrderState sets the "order_state" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableOrderState(s *string) *OrderStateUpdateOne {
	if s != nil {
		osuo.SetOrderState(*s)
	}
	return osuo
}

// ClearOrderState clears the value of the "order_state" field.
func (osuo *OrderStateUpdateOne) ClearOrderState() *OrderStateUpdateOne {
	osuo.mutation.ClearOrderState()
	return osuo
}

// SetEndAt sets the "end_at" field.
func (osuo *OrderStateUpdateOne) SetEndAt(u uint32) *OrderStateUpdateOne {
	osuo.mutation.ResetEndAt()
	osuo.mutation.SetEndAt(u)
	return osuo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableEndAt(u *uint32) *OrderStateUpdateOne {
	if u != nil {
		osuo.SetEndAt(*u)
	}
	return osuo
}

// AddEndAt adds u to the "end_at" field.
func (osuo *OrderStateUpdateOne) AddEndAt(u int32) *OrderStateUpdateOne {
	osuo.mutation.AddEndAt(u)
	return osuo
}

// ClearEndAt clears the value of the "end_at" field.
func (osuo *OrderStateUpdateOne) ClearEndAt() *OrderStateUpdateOne {
	osuo.mutation.ClearEndAt()
	return osuo
}

// SetLastBenefitAt sets the "last_benefit_at" field.
func (osuo *OrderStateUpdateOne) SetLastBenefitAt(u uint32) *OrderStateUpdateOne {
	osuo.mutation.ResetLastBenefitAt()
	osuo.mutation.SetLastBenefitAt(u)
	return osuo
}

// SetNillableLastBenefitAt sets the "last_benefit_at" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableLastBenefitAt(u *uint32) *OrderStateUpdateOne {
	if u != nil {
		osuo.SetLastBenefitAt(*u)
	}
	return osuo
}

// AddLastBenefitAt adds u to the "last_benefit_at" field.
func (osuo *OrderStateUpdateOne) AddLastBenefitAt(u int32) *OrderStateUpdateOne {
	osuo.mutation.AddLastBenefitAt(u)
	return osuo
}

// ClearLastBenefitAt clears the value of the "last_benefit_at" field.
func (osuo *OrderStateUpdateOne) ClearLastBenefitAt() *OrderStateUpdateOne {
	osuo.mutation.ClearLastBenefitAt()
	return osuo
}

// SetBenefitState sets the "benefit_state" field.
func (osuo *OrderStateUpdateOne) SetBenefitState(s string) *OrderStateUpdateOne {
	osuo.mutation.SetBenefitState(s)
	return osuo
}

// SetNillableBenefitState sets the "benefit_state" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableBenefitState(s *string) *OrderStateUpdateOne {
	if s != nil {
		osuo.SetBenefitState(*s)
	}
	return osuo
}

// ClearBenefitState clears the value of the "benefit_state" field.
func (osuo *OrderStateUpdateOne) ClearBenefitState() *OrderStateUpdateOne {
	osuo.mutation.ClearBenefitState()
	return osuo
}

// SetUserSetPaid sets the "user_set_paid" field.
func (osuo *OrderStateUpdateOne) SetUserSetPaid(b bool) *OrderStateUpdateOne {
	osuo.mutation.SetUserSetPaid(b)
	return osuo
}

// SetNillableUserSetPaid sets the "user_set_paid" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableUserSetPaid(b *bool) *OrderStateUpdateOne {
	if b != nil {
		osuo.SetUserSetPaid(*b)
	}
	return osuo
}

// ClearUserSetPaid clears the value of the "user_set_paid" field.
func (osuo *OrderStateUpdateOne) ClearUserSetPaid() *OrderStateUpdateOne {
	osuo.mutation.ClearUserSetPaid()
	return osuo
}

// SetUserSetCancelled sets the "user_set_cancelled" field.
func (osuo *OrderStateUpdateOne) SetUserSetCancelled(b bool) *OrderStateUpdateOne {
	osuo.mutation.SetUserSetCancelled(b)
	return osuo
}

// SetNillableUserSetCancelled sets the "user_set_cancelled" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillableUserSetCancelled(b *bool) *OrderStateUpdateOne {
	if b != nil {
		osuo.SetUserSetCancelled(*b)
	}
	return osuo
}

// ClearUserSetCancelled clears the value of the "user_set_cancelled" field.
func (osuo *OrderStateUpdateOne) ClearUserSetCancelled() *OrderStateUpdateOne {
	osuo.mutation.ClearUserSetCancelled()
	return osuo
}

// SetPaymentTransactionID sets the "payment_transaction_id" field.
func (osuo *OrderStateUpdateOne) SetPaymentTransactionID(s string) *OrderStateUpdateOne {
	osuo.mutation.SetPaymentTransactionID(s)
	return osuo
}

// SetNillablePaymentTransactionID sets the "payment_transaction_id" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillablePaymentTransactionID(s *string) *OrderStateUpdateOne {
	if s != nil {
		osuo.SetPaymentTransactionID(*s)
	}
	return osuo
}

// ClearPaymentTransactionID clears the value of the "payment_transaction_id" field.
func (osuo *OrderStateUpdateOne) ClearPaymentTransactionID() *OrderStateUpdateOne {
	osuo.mutation.ClearPaymentTransactionID()
	return osuo
}

// SetPaymentFinishAmount sets the "payment_finish_amount" field.
func (osuo *OrderStateUpdateOne) SetPaymentFinishAmount(d decimal.Decimal) *OrderStateUpdateOne {
	osuo.mutation.SetPaymentFinishAmount(d)
	return osuo
}

// SetNillablePaymentFinishAmount sets the "payment_finish_amount" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillablePaymentFinishAmount(d *decimal.Decimal) *OrderStateUpdateOne {
	if d != nil {
		osuo.SetPaymentFinishAmount(*d)
	}
	return osuo
}

// ClearPaymentFinishAmount clears the value of the "payment_finish_amount" field.
func (osuo *OrderStateUpdateOne) ClearPaymentFinishAmount() *OrderStateUpdateOne {
	osuo.mutation.ClearPaymentFinishAmount()
	return osuo
}

// SetPaymentState sets the "payment_state" field.
func (osuo *OrderStateUpdateOne) SetPaymentState(s string) *OrderStateUpdateOne {
	osuo.mutation.SetPaymentState(s)
	return osuo
}

// SetNillablePaymentState sets the "payment_state" field if the given value is not nil.
func (osuo *OrderStateUpdateOne) SetNillablePaymentState(s *string) *OrderStateUpdateOne {
	if s != nil {
		osuo.SetPaymentState(*s)
	}
	return osuo
}

// ClearPaymentState clears the value of the "payment_state" field.
func (osuo *OrderStateUpdateOne) ClearPaymentState() *OrderStateUpdateOne {
	osuo.mutation.ClearPaymentState()
	return osuo
}

// Mutation returns the OrderStateMutation object of the builder.
func (osuo *OrderStateUpdateOne) Mutation() *OrderStateMutation {
	return osuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OrderStateUpdateOne) Select(field string, fields ...string) *OrderStateUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OrderState entity.
func (osuo *OrderStateUpdateOne) Save(ctx context.Context) (*OrderState, error) {
	var (
		err  error
		node *OrderState
	)
	if err := osuo.defaults(); err != nil {
		return nil, err
	}
	if len(osuo.hooks) == 0 {
		node, err = osuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderStateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			osuo.mutation = mutation
			node, err = osuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(osuo.hooks) - 1; i >= 0; i-- {
			if osuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = osuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, osuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*OrderState)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderStateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OrderStateUpdateOne) SaveX(ctx context.Context) *OrderState {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OrderStateUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OrderStateUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osuo *OrderStateUpdateOne) defaults() error {
	if _, ok := osuo.mutation.UpdatedAt(); !ok {
		if orderstate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderstate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderstate.UpdateDefaultUpdatedAt()
		osuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (osuo *OrderStateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderStateUpdateOne {
	osuo.modifiers = append(osuo.modifiers, modifiers...)
	return osuo
}

func (osuo *OrderStateUpdateOne) sqlSave(ctx context.Context) (_node *OrderState, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstate.Table,
			Columns: orderstate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderstate.FieldID,
			},
		},
	}
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderState.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstate.FieldID)
		for _, f := range fields {
			if !orderstate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldCreatedAt,
		})
	}
	if value, ok := osuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldCreatedAt,
		})
	}
	if value, ok := osuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldUpdatedAt,
		})
	}
	if value, ok := osuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldUpdatedAt,
		})
	}
	if value, ok := osuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldDeletedAt,
		})
	}
	if value, ok := osuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldDeletedAt,
		})
	}
	if value, ok := osuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: orderstate.FieldOrderID,
		})
	}
	if value, ok := osuo.mutation.OrderState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldOrderState,
		})
	}
	if osuo.mutation.OrderStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldOrderState,
		})
	}
	if value, ok := osuo.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldEndAt,
		})
	}
	if value, ok := osuo.mutation.AddedEndAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldEndAt,
		})
	}
	if osuo.mutation.EndAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstate.FieldEndAt,
		})
	}
	if value, ok := osuo.mutation.LastBenefitAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldLastBenefitAt,
		})
	}
	if value, ok := osuo.mutation.AddedLastBenefitAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: orderstate.FieldLastBenefitAt,
		})
	}
	if osuo.mutation.LastBenefitAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: orderstate.FieldLastBenefitAt,
		})
	}
	if value, ok := osuo.mutation.BenefitState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldBenefitState,
		})
	}
	if osuo.mutation.BenefitStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldBenefitState,
		})
	}
	if value, ok := osuo.mutation.UserSetPaid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: orderstate.FieldUserSetPaid,
		})
	}
	if osuo.mutation.UserSetPaidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: orderstate.FieldUserSetPaid,
		})
	}
	if value, ok := osuo.mutation.UserSetCancelled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: orderstate.FieldUserSetCancelled,
		})
	}
	if osuo.mutation.UserSetCancelledCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: orderstate.FieldUserSetCancelled,
		})
	}
	if value, ok := osuo.mutation.PaymentTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldPaymentTransactionID,
		})
	}
	if osuo.mutation.PaymentTransactionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldPaymentTransactionID,
		})
	}
	if value, ok := osuo.mutation.PaymentFinishAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: orderstate.FieldPaymentFinishAmount,
		})
	}
	if osuo.mutation.PaymentFinishAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: orderstate.FieldPaymentFinishAmount,
		})
	}
	if value, ok := osuo.mutation.PaymentState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: orderstate.FieldPaymentState,
		})
	}
	if osuo.mutation.PaymentStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: orderstate.FieldPaymentState,
		})
	}
	_spec.Modifiers = osuo.modifiers
	_node = &OrderState{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderstate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
