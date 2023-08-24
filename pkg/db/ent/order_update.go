// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks     []Hook
	mutation  *OrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetCreatedAt sets the "created_at" field.
func (ou *OrderUpdate) SetCreatedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetCreatedAt()
	ou.mutation.SetCreatedAt(u)
	return ou
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableCreatedAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetCreatedAt(*u)
	}
	return ou
}

// AddCreatedAt adds u to the "created_at" field.
func (ou *OrderUpdate) AddCreatedAt(u int32) *OrderUpdate {
	ou.mutation.AddCreatedAt(u)
	return ou
}

// SetUpdatedAt sets the "updated_at" field.
func (ou *OrderUpdate) SetUpdatedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetUpdatedAt()
	ou.mutation.SetUpdatedAt(u)
	return ou
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ou *OrderUpdate) AddUpdatedAt(u int32) *OrderUpdate {
	ou.mutation.AddUpdatedAt(u)
	return ou
}

// SetDeletedAt sets the "deleted_at" field.
func (ou *OrderUpdate) SetDeletedAt(u uint32) *OrderUpdate {
	ou.mutation.ResetDeletedAt()
	ou.mutation.SetDeletedAt(u)
	return ou
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDeletedAt(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetDeletedAt(*u)
	}
	return ou
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ou *OrderUpdate) AddDeletedAt(u int32) *OrderUpdate {
	ou.mutation.AddDeletedAt(u)
	return ou
}

// SetAppID sets the "app_id" field.
func (ou *OrderUpdate) SetAppID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetAppID(u)
	return ou
}

// SetUserID sets the "user_id" field.
func (ou *OrderUpdate) SetUserID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetUserID(u)
	return ou
}

// SetGoodID sets the "good_id" field.
func (ou *OrderUpdate) SetGoodID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetGoodID(u)
	return ou
}

// SetAppGoodID sets the "app_good_id" field.
func (ou *OrderUpdate) SetAppGoodID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetAppGoodID(u)
	return ou
}

// SetPaymentID sets the "payment_id" field.
func (ou *OrderUpdate) SetPaymentID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetPaymentID(u)
	return ou
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePaymentID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetPaymentID(*u)
	}
	return ou
}

// ClearPaymentID clears the value of the "payment_id" field.
func (ou *OrderUpdate) ClearPaymentID() *OrderUpdate {
	ou.mutation.ClearPaymentID()
	return ou
}

// SetParentOrderID sets the "parent_order_id" field.
func (ou *OrderUpdate) SetParentOrderID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetParentOrderID(u)
	return ou
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableParentOrderID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetParentOrderID(*u)
	}
	return ou
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (ou *OrderUpdate) ClearParentOrderID() *OrderUpdate {
	ou.mutation.ClearParentOrderID()
	return ou
}

// SetUnitsV1 sets the "units_v1" field.
func (ou *OrderUpdate) SetUnitsV1(d decimal.Decimal) *OrderUpdate {
	ou.mutation.SetUnitsV1(d)
	return ou
}

// SetNillableUnitsV1 sets the "units_v1" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableUnitsV1(d *decimal.Decimal) *OrderUpdate {
	if d != nil {
		ou.SetUnitsV1(*d)
	}
	return ou
}

// ClearUnitsV1 clears the value of the "units_v1" field.
func (ou *OrderUpdate) ClearUnitsV1() *OrderUpdate {
	ou.mutation.ClearUnitsV1()
	return ou
}

// SetGoodValue sets the "good_value" field.
func (ou *OrderUpdate) SetGoodValue(d decimal.Decimal) *OrderUpdate {
	ou.mutation.SetGoodValue(d)
	return ou
}

// SetNillableGoodValue sets the "good_value" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableGoodValue(d *decimal.Decimal) *OrderUpdate {
	if d != nil {
		ou.SetGoodValue(*d)
	}
	return ou
}

// ClearGoodValue clears the value of the "good_value" field.
func (ou *OrderUpdate) ClearGoodValue() *OrderUpdate {
	ou.mutation.ClearGoodValue()
	return ou
}

// SetPaymentAmount sets the "payment_amount" field.
func (ou *OrderUpdate) SetPaymentAmount(d decimal.Decimal) *OrderUpdate {
	ou.mutation.SetPaymentAmount(d)
	return ou
}

// SetNillablePaymentAmount sets the "payment_amount" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePaymentAmount(d *decimal.Decimal) *OrderUpdate {
	if d != nil {
		ou.SetPaymentAmount(*d)
	}
	return ou
}

// ClearPaymentAmount clears the value of the "payment_amount" field.
func (ou *OrderUpdate) ClearPaymentAmount() *OrderUpdate {
	ou.mutation.ClearPaymentAmount()
	return ou
}

// SetDiscountAmount sets the "discount_amount" field.
func (ou *OrderUpdate) SetDiscountAmount(d decimal.Decimal) *OrderUpdate {
	ou.mutation.SetDiscountAmount(d)
	return ou
}

// SetNillableDiscountAmount sets the "discount_amount" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDiscountAmount(d *decimal.Decimal) *OrderUpdate {
	if d != nil {
		ou.SetDiscountAmount(*d)
	}
	return ou
}

// ClearDiscountAmount clears the value of the "discount_amount" field.
func (ou *OrderUpdate) ClearDiscountAmount() *OrderUpdate {
	ou.mutation.ClearDiscountAmount()
	return ou
}

// SetPromotionID sets the "promotion_id" field.
func (ou *OrderUpdate) SetPromotionID(u uuid.UUID) *OrderUpdate {
	ou.mutation.SetPromotionID(u)
	return ou
}

// SetNillablePromotionID sets the "promotion_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePromotionID(u *uuid.UUID) *OrderUpdate {
	if u != nil {
		ou.SetPromotionID(*u)
	}
	return ou
}

// ClearPromotionID clears the value of the "promotion_id" field.
func (ou *OrderUpdate) ClearPromotionID() *OrderUpdate {
	ou.mutation.ClearPromotionID()
	return ou
}

// SetDurationDays sets the "duration_days" field.
func (ou *OrderUpdate) SetDurationDays(u uint32) *OrderUpdate {
	ou.mutation.ResetDurationDays()
	ou.mutation.SetDurationDays(u)
	return ou
}

// SetNillableDurationDays sets the "duration_days" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDurationDays(u *uint32) *OrderUpdate {
	if u != nil {
		ou.SetDurationDays(*u)
	}
	return ou
}

// AddDurationDays adds u to the "duration_days" field.
func (ou *OrderUpdate) AddDurationDays(u int32) *OrderUpdate {
	ou.mutation.AddDurationDays(u)
	return ou
}

// ClearDurationDays clears the value of the "duration_days" field.
func (ou *OrderUpdate) ClearDurationDays() *OrderUpdate {
	ou.mutation.ClearDurationDays()
	return ou
}

// SetOrderType sets the "order_type" field.
func (ou *OrderUpdate) SetOrderType(s string) *OrderUpdate {
	ou.mutation.SetOrderType(s)
	return ou
}

// SetNillableOrderType sets the "order_type" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableOrderType(s *string) *OrderUpdate {
	if s != nil {
		ou.SetOrderType(*s)
	}
	return ou
}

// ClearOrderType clears the value of the "order_type" field.
func (ou *OrderUpdate) ClearOrderType() *OrderUpdate {
	ou.mutation.ClearOrderType()
	return ou
}

// SetInvestmentType sets the "investment_type" field.
func (ou *OrderUpdate) SetInvestmentType(s string) *OrderUpdate {
	ou.mutation.SetInvestmentType(s)
	return ou
}

// SetNillableInvestmentType sets the "investment_type" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableInvestmentType(s *string) *OrderUpdate {
	if s != nil {
		ou.SetInvestmentType(*s)
	}
	return ou
}

// ClearInvestmentType clears the value of the "investment_type" field.
func (ou *OrderUpdate) ClearInvestmentType() *OrderUpdate {
	ou.mutation.ClearInvestmentType()
	return ou
}

// SetCouponIds sets the "coupon_ids" field.
func (ou *OrderUpdate) SetCouponIds(u []uuid.UUID) *OrderUpdate {
	ou.mutation.SetCouponIds(u)
	return ou
}

// ClearCouponIds clears the value of the "coupon_ids" field.
func (ou *OrderUpdate) ClearCouponIds() *OrderUpdate {
	ou.mutation.ClearCouponIds()
	return ou
}

// SetPaymentType sets the "payment_type" field.
func (ou *OrderUpdate) SetPaymentType(s string) *OrderUpdate {
	ou.mutation.SetPaymentType(s)
	return ou
}

// SetNillablePaymentType sets the "payment_type" field if the given value is not nil.
func (ou *OrderUpdate) SetNillablePaymentType(s *string) *OrderUpdate {
	if s != nil {
		ou.SetPaymentType(*s)
	}
	return ou
}

// ClearPaymentType clears the value of the "payment_type" field.
func (ou *OrderUpdate) ClearPaymentType() *OrderUpdate {
	ou.mutation.ClearPaymentType()
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ou.defaults(); err != nil {
		return 0, err
	}
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ou *OrderUpdate) defaults() error {
	if _, ok := ou.mutation.UpdatedAt(); !ok {
		if order.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized order.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := order.UpdateDefaultUpdatedAt()
		ou.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ou *OrderUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUpdate {
	ou.modifiers = append(ou.modifiers, modifiers...)
	return ou
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ou.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ou.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ou.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppID,
		})
	}
	if value, ok := ou.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ou.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldGoodID,
		})
	}
	if value, ok := ou.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppGoodID,
		})
	}
	if value, ok := ou.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldPaymentID,
		})
	}
	if ou.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldPaymentID,
		})
	}
	if value, ok := ou.mutation.ParentOrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldParentOrderID,
		})
	}
	if ou.mutation.ParentOrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldParentOrderID,
		})
	}
	if value, ok := ou.mutation.UnitsV1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldUnitsV1,
		})
	}
	if ou.mutation.UnitsV1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldUnitsV1,
		})
	}
	if value, ok := ou.mutation.GoodValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldGoodValue,
		})
	}
	if ou.mutation.GoodValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldGoodValue,
		})
	}
	if value, ok := ou.mutation.PaymentAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldPaymentAmount,
		})
	}
	if ou.mutation.PaymentAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldPaymentAmount,
		})
	}
	if value, ok := ou.mutation.DiscountAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldDiscountAmount,
		})
	}
	if ou.mutation.DiscountAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldDiscountAmount,
		})
	}
	if value, ok := ou.mutation.PromotionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldPromotionID,
		})
	}
	if ou.mutation.PromotionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldPromotionID,
		})
	}
	if value, ok := ou.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDurationDays,
		})
	}
	if value, ok := ou.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDurationDays,
		})
	}
	if ou.mutation.DurationDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: order.FieldDurationDays,
		})
	}
	if value, ok := ou.mutation.OrderType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldOrderType,
		})
	}
	if ou.mutation.OrderTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldOrderType,
		})
	}
	if value, ok := ou.mutation.InvestmentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldInvestmentType,
		})
	}
	if ou.mutation.InvestmentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldInvestmentType,
		})
	}
	if value, ok := ou.mutation.CouponIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: order.FieldCouponIds,
		})
	}
	if ou.mutation.CouponIdsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: order.FieldCouponIds,
		})
	}
	if value, ok := ou.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldPaymentType,
		})
	}
	if ou.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldPaymentType,
		})
	}
	_spec.Modifiers = ou.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrderMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (ouo *OrderUpdateOne) SetCreatedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetCreatedAt()
	ouo.mutation.SetCreatedAt(u)
	return ouo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableCreatedAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetCreatedAt(*u)
	}
	return ouo
}

// AddCreatedAt adds u to the "created_at" field.
func (ouo *OrderUpdateOne) AddCreatedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddCreatedAt(u)
	return ouo
}

// SetUpdatedAt sets the "updated_at" field.
func (ouo *OrderUpdateOne) SetUpdatedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetUpdatedAt()
	ouo.mutation.SetUpdatedAt(u)
	return ouo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ouo *OrderUpdateOne) AddUpdatedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddUpdatedAt(u)
	return ouo
}

// SetDeletedAt sets the "deleted_at" field.
func (ouo *OrderUpdateOne) SetDeletedAt(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetDeletedAt()
	ouo.mutation.SetDeletedAt(u)
	return ouo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDeletedAt(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetDeletedAt(*u)
	}
	return ouo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ouo *OrderUpdateOne) AddDeletedAt(u int32) *OrderUpdateOne {
	ouo.mutation.AddDeletedAt(u)
	return ouo
}

// SetAppID sets the "app_id" field.
func (ouo *OrderUpdateOne) SetAppID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetAppID(u)
	return ouo
}

// SetUserID sets the "user_id" field.
func (ouo *OrderUpdateOne) SetUserID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetUserID(u)
	return ouo
}

// SetGoodID sets the "good_id" field.
func (ouo *OrderUpdateOne) SetGoodID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetGoodID(u)
	return ouo
}

// SetAppGoodID sets the "app_good_id" field.
func (ouo *OrderUpdateOne) SetAppGoodID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetAppGoodID(u)
	return ouo
}

// SetPaymentID sets the "payment_id" field.
func (ouo *OrderUpdateOne) SetPaymentID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetPaymentID(u)
	return ouo
}

// SetNillablePaymentID sets the "payment_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePaymentID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetPaymentID(*u)
	}
	return ouo
}

// ClearPaymentID clears the value of the "payment_id" field.
func (ouo *OrderUpdateOne) ClearPaymentID() *OrderUpdateOne {
	ouo.mutation.ClearPaymentID()
	return ouo
}

// SetParentOrderID sets the "parent_order_id" field.
func (ouo *OrderUpdateOne) SetParentOrderID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetParentOrderID(u)
	return ouo
}

// SetNillableParentOrderID sets the "parent_order_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableParentOrderID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetParentOrderID(*u)
	}
	return ouo
}

// ClearParentOrderID clears the value of the "parent_order_id" field.
func (ouo *OrderUpdateOne) ClearParentOrderID() *OrderUpdateOne {
	ouo.mutation.ClearParentOrderID()
	return ouo
}

// SetUnitsV1 sets the "units_v1" field.
func (ouo *OrderUpdateOne) SetUnitsV1(d decimal.Decimal) *OrderUpdateOne {
	ouo.mutation.SetUnitsV1(d)
	return ouo
}

// SetNillableUnitsV1 sets the "units_v1" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUnitsV1(d *decimal.Decimal) *OrderUpdateOne {
	if d != nil {
		ouo.SetUnitsV1(*d)
	}
	return ouo
}

// ClearUnitsV1 clears the value of the "units_v1" field.
func (ouo *OrderUpdateOne) ClearUnitsV1() *OrderUpdateOne {
	ouo.mutation.ClearUnitsV1()
	return ouo
}

// SetGoodValue sets the "good_value" field.
func (ouo *OrderUpdateOne) SetGoodValue(d decimal.Decimal) *OrderUpdateOne {
	ouo.mutation.SetGoodValue(d)
	return ouo
}

// SetNillableGoodValue sets the "good_value" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableGoodValue(d *decimal.Decimal) *OrderUpdateOne {
	if d != nil {
		ouo.SetGoodValue(*d)
	}
	return ouo
}

// ClearGoodValue clears the value of the "good_value" field.
func (ouo *OrderUpdateOne) ClearGoodValue() *OrderUpdateOne {
	ouo.mutation.ClearGoodValue()
	return ouo
}

// SetPaymentAmount sets the "payment_amount" field.
func (ouo *OrderUpdateOne) SetPaymentAmount(d decimal.Decimal) *OrderUpdateOne {
	ouo.mutation.SetPaymentAmount(d)
	return ouo
}

// SetNillablePaymentAmount sets the "payment_amount" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePaymentAmount(d *decimal.Decimal) *OrderUpdateOne {
	if d != nil {
		ouo.SetPaymentAmount(*d)
	}
	return ouo
}

// ClearPaymentAmount clears the value of the "payment_amount" field.
func (ouo *OrderUpdateOne) ClearPaymentAmount() *OrderUpdateOne {
	ouo.mutation.ClearPaymentAmount()
	return ouo
}

// SetDiscountAmount sets the "discount_amount" field.
func (ouo *OrderUpdateOne) SetDiscountAmount(d decimal.Decimal) *OrderUpdateOne {
	ouo.mutation.SetDiscountAmount(d)
	return ouo
}

// SetNillableDiscountAmount sets the "discount_amount" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDiscountAmount(d *decimal.Decimal) *OrderUpdateOne {
	if d != nil {
		ouo.SetDiscountAmount(*d)
	}
	return ouo
}

// ClearDiscountAmount clears the value of the "discount_amount" field.
func (ouo *OrderUpdateOne) ClearDiscountAmount() *OrderUpdateOne {
	ouo.mutation.ClearDiscountAmount()
	return ouo
}

// SetPromotionID sets the "promotion_id" field.
func (ouo *OrderUpdateOne) SetPromotionID(u uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetPromotionID(u)
	return ouo
}

// SetNillablePromotionID sets the "promotion_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePromotionID(u *uuid.UUID) *OrderUpdateOne {
	if u != nil {
		ouo.SetPromotionID(*u)
	}
	return ouo
}

// ClearPromotionID clears the value of the "promotion_id" field.
func (ouo *OrderUpdateOne) ClearPromotionID() *OrderUpdateOne {
	ouo.mutation.ClearPromotionID()
	return ouo
}

// SetDurationDays sets the "duration_days" field.
func (ouo *OrderUpdateOne) SetDurationDays(u uint32) *OrderUpdateOne {
	ouo.mutation.ResetDurationDays()
	ouo.mutation.SetDurationDays(u)
	return ouo
}

// SetNillableDurationDays sets the "duration_days" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDurationDays(u *uint32) *OrderUpdateOne {
	if u != nil {
		ouo.SetDurationDays(*u)
	}
	return ouo
}

// AddDurationDays adds u to the "duration_days" field.
func (ouo *OrderUpdateOne) AddDurationDays(u int32) *OrderUpdateOne {
	ouo.mutation.AddDurationDays(u)
	return ouo
}

// ClearDurationDays clears the value of the "duration_days" field.
func (ouo *OrderUpdateOne) ClearDurationDays() *OrderUpdateOne {
	ouo.mutation.ClearDurationDays()
	return ouo
}

// SetOrderType sets the "order_type" field.
func (ouo *OrderUpdateOne) SetOrderType(s string) *OrderUpdateOne {
	ouo.mutation.SetOrderType(s)
	return ouo
}

// SetNillableOrderType sets the "order_type" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableOrderType(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetOrderType(*s)
	}
	return ouo
}

// ClearOrderType clears the value of the "order_type" field.
func (ouo *OrderUpdateOne) ClearOrderType() *OrderUpdateOne {
	ouo.mutation.ClearOrderType()
	return ouo
}

// SetInvestmentType sets the "investment_type" field.
func (ouo *OrderUpdateOne) SetInvestmentType(s string) *OrderUpdateOne {
	ouo.mutation.SetInvestmentType(s)
	return ouo
}

// SetNillableInvestmentType sets the "investment_type" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableInvestmentType(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetInvestmentType(*s)
	}
	return ouo
}

// ClearInvestmentType clears the value of the "investment_type" field.
func (ouo *OrderUpdateOne) ClearInvestmentType() *OrderUpdateOne {
	ouo.mutation.ClearInvestmentType()
	return ouo
}

// SetCouponIds sets the "coupon_ids" field.
func (ouo *OrderUpdateOne) SetCouponIds(u []uuid.UUID) *OrderUpdateOne {
	ouo.mutation.SetCouponIds(u)
	return ouo
}

// ClearCouponIds clears the value of the "coupon_ids" field.
func (ouo *OrderUpdateOne) ClearCouponIds() *OrderUpdateOne {
	ouo.mutation.ClearCouponIds()
	return ouo
}

// SetPaymentType sets the "payment_type" field.
func (ouo *OrderUpdateOne) SetPaymentType(s string) *OrderUpdateOne {
	ouo.mutation.SetPaymentType(s)
	return ouo
}

// SetNillablePaymentType sets the "payment_type" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillablePaymentType(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetPaymentType(*s)
	}
	return ouo
}

// ClearPaymentType clears the value of the "payment_type" field.
func (ouo *OrderUpdateOne) ClearPaymentType() *OrderUpdateOne {
	ouo.mutation.ClearPaymentType()
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if err := ouo.defaults(); err != nil {
		return nil, err
	}
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Order)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ouo *OrderUpdateOne) defaults() error {
	if _, ok := ouo.mutation.UpdatedAt(); !ok {
		if order.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized order.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := order.UpdateDefaultUpdatedAt()
		ouo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ouo *OrderUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrderUpdateOne {
	ouo.modifiers = append(ouo.modifiers, modifiers...)
	return ouo
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldCreatedAt,
		})
	}
	if value, ok := ouo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldUpdatedAt,
		})
	}
	if value, ok := ouo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDeletedAt,
		})
	}
	if value, ok := ouo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppID,
		})
	}
	if value, ok := ouo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldUserID,
		})
	}
	if value, ok := ouo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldGoodID,
		})
	}
	if value, ok := ouo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldAppGoodID,
		})
	}
	if value, ok := ouo.mutation.PaymentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldPaymentID,
		})
	}
	if ouo.mutation.PaymentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldPaymentID,
		})
	}
	if value, ok := ouo.mutation.ParentOrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldParentOrderID,
		})
	}
	if ouo.mutation.ParentOrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldParentOrderID,
		})
	}
	if value, ok := ouo.mutation.UnitsV1(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldUnitsV1,
		})
	}
	if ouo.mutation.UnitsV1Cleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldUnitsV1,
		})
	}
	if value, ok := ouo.mutation.GoodValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldGoodValue,
		})
	}
	if ouo.mutation.GoodValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldGoodValue,
		})
	}
	if value, ok := ouo.mutation.PaymentAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldPaymentAmount,
		})
	}
	if ouo.mutation.PaymentAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldPaymentAmount,
		})
	}
	if value, ok := ouo.mutation.DiscountAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: order.FieldDiscountAmount,
		})
	}
	if ouo.mutation.DiscountAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: order.FieldDiscountAmount,
		})
	}
	if value, ok := ouo.mutation.PromotionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: order.FieldPromotionID,
		})
	}
	if ouo.mutation.PromotionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: order.FieldPromotionID,
		})
	}
	if value, ok := ouo.mutation.DurationDays(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDurationDays,
		})
	}
	if value, ok := ouo.mutation.AddedDurationDays(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: order.FieldDurationDays,
		})
	}
	if ouo.mutation.DurationDaysCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: order.FieldDurationDays,
		})
	}
	if value, ok := ouo.mutation.OrderType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldOrderType,
		})
	}
	if ouo.mutation.OrderTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldOrderType,
		})
	}
	if value, ok := ouo.mutation.InvestmentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldInvestmentType,
		})
	}
	if ouo.mutation.InvestmentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldInvestmentType,
		})
	}
	if value, ok := ouo.mutation.CouponIds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: order.FieldCouponIds,
		})
	}
	if ouo.mutation.CouponIdsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: order.FieldCouponIds,
		})
	}
	if value, ok := ouo.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: order.FieldPaymentType,
		})
	}
	if ouo.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: order.FieldPaymentType,
		})
	}
	_spec.Modifiers = ouo.modifiers
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
