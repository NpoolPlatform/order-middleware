// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentcontract"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// PaymentContractCreate is the builder for creating a PaymentContract entity.
type PaymentContractCreate struct {
	config
	mutation *PaymentContractMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (pcc *PaymentContractCreate) SetCreatedAt(u uint32) *PaymentContractCreate {
	pcc.mutation.SetCreatedAt(u)
	return pcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableCreatedAt(u *uint32) *PaymentContractCreate {
	if u != nil {
		pcc.SetCreatedAt(*u)
	}
	return pcc
}

// SetUpdatedAt sets the "updated_at" field.
func (pcc *PaymentContractCreate) SetUpdatedAt(u uint32) *PaymentContractCreate {
	pcc.mutation.SetUpdatedAt(u)
	return pcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableUpdatedAt(u *uint32) *PaymentContractCreate {
	if u != nil {
		pcc.SetUpdatedAt(*u)
	}
	return pcc
}

// SetDeletedAt sets the "deleted_at" field.
func (pcc *PaymentContractCreate) SetDeletedAt(u uint32) *PaymentContractCreate {
	pcc.mutation.SetDeletedAt(u)
	return pcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableDeletedAt(u *uint32) *PaymentContractCreate {
	if u != nil {
		pcc.SetDeletedAt(*u)
	}
	return pcc
}

// SetEntID sets the "ent_id" field.
func (pcc *PaymentContractCreate) SetEntID(u uuid.UUID) *PaymentContractCreate {
	pcc.mutation.SetEntID(u)
	return pcc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableEntID(u *uuid.UUID) *PaymentContractCreate {
	if u != nil {
		pcc.SetEntID(*u)
	}
	return pcc
}

// SetOrderID sets the "order_id" field.
func (pcc *PaymentContractCreate) SetOrderID(u uuid.UUID) *PaymentContractCreate {
	pcc.mutation.SetOrderID(u)
	return pcc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableOrderID(u *uuid.UUID) *PaymentContractCreate {
	if u != nil {
		pcc.SetOrderID(*u)
	}
	return pcc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (pcc *PaymentContractCreate) SetCoinTypeID(u uuid.UUID) *PaymentContractCreate {
	pcc.mutation.SetCoinTypeID(u)
	return pcc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableCoinTypeID(u *uuid.UUID) *PaymentContractCreate {
	if u != nil {
		pcc.SetCoinTypeID(*u)
	}
	return pcc
}

// SetAmount sets the "amount" field.
func (pcc *PaymentContractCreate) SetAmount(d decimal.Decimal) *PaymentContractCreate {
	pcc.mutation.SetAmount(d)
	return pcc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pcc *PaymentContractCreate) SetNillableAmount(d *decimal.Decimal) *PaymentContractCreate {
	if d != nil {
		pcc.SetAmount(*d)
	}
	return pcc
}

// SetID sets the "id" field.
func (pcc *PaymentContractCreate) SetID(u uint32) *PaymentContractCreate {
	pcc.mutation.SetID(u)
	return pcc
}

// Mutation returns the PaymentContractMutation object of the builder.
func (pcc *PaymentContractCreate) Mutation() *PaymentContractMutation {
	return pcc.mutation
}

// Save creates the PaymentContract in the database.
func (pcc *PaymentContractCreate) Save(ctx context.Context) (*PaymentContract, error) {
	var (
		err  error
		node *PaymentContract
	)
	if err := pcc.defaults(); err != nil {
		return nil, err
	}
	if len(pcc.hooks) == 0 {
		if err = pcc.check(); err != nil {
			return nil, err
		}
		node, err = pcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcc.check(); err != nil {
				return nil, err
			}
			pcc.mutation = mutation
			if node, err = pcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pcc.hooks) - 1; i >= 0; i-- {
			if pcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PaymentContract)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PaymentContractMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *PaymentContractCreate) SaveX(ctx context.Context) *PaymentContract {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *PaymentContractCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *PaymentContractCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *PaymentContractCreate) defaults() error {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		if paymentcontract.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentcontract.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := paymentcontract.DefaultCreatedAt()
		pcc.mutation.SetCreatedAt(v)
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		if paymentcontract.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentcontract.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := paymentcontract.DefaultUpdatedAt()
		pcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pcc.mutation.DeletedAt(); !ok {
		if paymentcontract.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized paymentcontract.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := paymentcontract.DefaultDeletedAt()
		pcc.mutation.SetDeletedAt(v)
	}
	if _, ok := pcc.mutation.EntID(); !ok {
		if paymentcontract.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized paymentcontract.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := paymentcontract.DefaultEntID()
		pcc.mutation.SetEntID(v)
	}
	if _, ok := pcc.mutation.OrderID(); !ok {
		if paymentcontract.DefaultOrderID == nil {
			return fmt.Errorf("ent: uninitialized paymentcontract.DefaultOrderID (forgotten import ent/runtime?)")
		}
		v := paymentcontract.DefaultOrderID()
		pcc.mutation.SetOrderID(v)
	}
	if _, ok := pcc.mutation.CoinTypeID(); !ok {
		if paymentcontract.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized paymentcontract.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := paymentcontract.DefaultCoinTypeID()
		pcc.mutation.SetCoinTypeID(v)
	}
	if _, ok := pcc.mutation.Amount(); !ok {
		v := paymentcontract.DefaultAmount
		pcc.mutation.SetAmount(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pcc *PaymentContractCreate) check() error {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "PaymentContract.created_at"`)}
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "PaymentContract.updated_at"`)}
	}
	if _, ok := pcc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "PaymentContract.deleted_at"`)}
	}
	if _, ok := pcc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "PaymentContract.ent_id"`)}
	}
	return nil
}

func (pcc *PaymentContractCreate) sqlSave(ctx context.Context) (*PaymentContract, error) {
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
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

func (pcc *PaymentContractCreate) createSpec() (*PaymentContract, *sqlgraph.CreateSpec) {
	var (
		_node = &PaymentContract{config: pcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: paymentcontract.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentcontract.FieldID,
			},
		}
	)
	_spec.OnConflict = pcc.conflict
	if id, ok := pcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentcontract.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentcontract.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := pcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: paymentcontract.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := pcc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentcontract.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := pcc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentcontract.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := pcc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: paymentcontract.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := pcc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: paymentcontract.FieldAmount,
		})
		_node.Amount = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PaymentContract.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PaymentContractUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pcc *PaymentContractCreate) OnConflict(opts ...sql.ConflictOption) *PaymentContractUpsertOne {
	pcc.conflict = opts
	return &PaymentContractUpsertOne{
		create: pcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PaymentContract.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcc *PaymentContractCreate) OnConflictColumns(columns ...string) *PaymentContractUpsertOne {
	pcc.conflict = append(pcc.conflict, sql.ConflictColumns(columns...))
	return &PaymentContractUpsertOne{
		create: pcc,
	}
}

type (
	// PaymentContractUpsertOne is the builder for "upsert"-ing
	//  one PaymentContract node.
	PaymentContractUpsertOne struct {
		create *PaymentContractCreate
	}

	// PaymentContractUpsert is the "OnConflict" setter.
	PaymentContractUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *PaymentContractUpsert) SetCreatedAt(v uint32) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateCreatedAt() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PaymentContractUpsert) AddCreatedAt(v uint32) *PaymentContractUpsert {
	u.Add(paymentcontract.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PaymentContractUpsert) SetUpdatedAt(v uint32) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateUpdatedAt() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PaymentContractUpsert) AddUpdatedAt(v uint32) *PaymentContractUpsert {
	u.Add(paymentcontract.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PaymentContractUpsert) SetDeletedAt(v uint32) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateDeletedAt() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PaymentContractUpsert) AddDeletedAt(v uint32) *PaymentContractUpsert {
	u.Add(paymentcontract.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *PaymentContractUpsert) SetEntID(v uuid.UUID) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateEntID() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldEntID)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *PaymentContractUpsert) SetOrderID(v uuid.UUID) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateOrderID() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *PaymentContractUpsert) ClearOrderID() *PaymentContractUpsert {
	u.SetNull(paymentcontract.FieldOrderID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *PaymentContractUpsert) SetCoinTypeID(v uuid.UUID) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateCoinTypeID() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *PaymentContractUpsert) ClearCoinTypeID() *PaymentContractUpsert {
	u.SetNull(paymentcontract.FieldCoinTypeID)
	return u
}

// SetAmount sets the "amount" field.
func (u *PaymentContractUpsert) SetAmount(v decimal.Decimal) *PaymentContractUpsert {
	u.Set(paymentcontract.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PaymentContractUpsert) UpdateAmount() *PaymentContractUpsert {
	u.SetExcluded(paymentcontract.FieldAmount)
	return u
}

// ClearAmount clears the value of the "amount" field.
func (u *PaymentContractUpsert) ClearAmount() *PaymentContractUpsert {
	u.SetNull(paymentcontract.FieldAmount)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PaymentContract.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(paymentcontract.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PaymentContractUpsertOne) UpdateNewValues() *PaymentContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(paymentcontract.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PaymentContract.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PaymentContractUpsertOne) Ignore() *PaymentContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PaymentContractUpsertOne) DoNothing() *PaymentContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PaymentContractCreate.OnConflict
// documentation for more info.
func (u *PaymentContractUpsertOne) Update(set func(*PaymentContractUpsert)) *PaymentContractUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PaymentContractUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PaymentContractUpsertOne) SetCreatedAt(v uint32) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PaymentContractUpsertOne) AddCreatedAt(v uint32) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateCreatedAt() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PaymentContractUpsertOne) SetUpdatedAt(v uint32) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PaymentContractUpsertOne) AddUpdatedAt(v uint32) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateUpdatedAt() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PaymentContractUpsertOne) SetDeletedAt(v uint32) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PaymentContractUpsertOne) AddDeletedAt(v uint32) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateDeletedAt() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *PaymentContractUpsertOne) SetEntID(v uuid.UUID) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateEntID() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateEntID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *PaymentContractUpsertOne) SetOrderID(v uuid.UUID) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateOrderID() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *PaymentContractUpsertOne) ClearOrderID() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.ClearOrderID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *PaymentContractUpsertOne) SetCoinTypeID(v uuid.UUID) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateCoinTypeID() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *PaymentContractUpsertOne) ClearCoinTypeID() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *PaymentContractUpsertOne) SetAmount(v decimal.Decimal) *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PaymentContractUpsertOne) UpdateAmount() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *PaymentContractUpsertOne) ClearAmount() *PaymentContractUpsertOne {
	return u.Update(func(s *PaymentContractUpsert) {
		s.ClearAmount()
	})
}

// Exec executes the query.
func (u *PaymentContractUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PaymentContractCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PaymentContractUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PaymentContractUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PaymentContractUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PaymentContractCreateBulk is the builder for creating many PaymentContract entities in bulk.
type PaymentContractCreateBulk struct {
	config
	builders []*PaymentContractCreate
	conflict []sql.ConflictOption
}

// Save creates the PaymentContract entities in the database.
func (pccb *PaymentContractCreateBulk) Save(ctx context.Context) ([]*PaymentContract, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*PaymentContract, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PaymentContractMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *PaymentContractCreateBulk) SaveX(ctx context.Context) []*PaymentContract {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *PaymentContractCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *PaymentContractCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PaymentContract.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PaymentContractUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (pccb *PaymentContractCreateBulk) OnConflict(opts ...sql.ConflictOption) *PaymentContractUpsertBulk {
	pccb.conflict = opts
	return &PaymentContractUpsertBulk{
		create: pccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PaymentContract.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pccb *PaymentContractCreateBulk) OnConflictColumns(columns ...string) *PaymentContractUpsertBulk {
	pccb.conflict = append(pccb.conflict, sql.ConflictColumns(columns...))
	return &PaymentContractUpsertBulk{
		create: pccb,
	}
}

// PaymentContractUpsertBulk is the builder for "upsert"-ing
// a bulk of PaymentContract nodes.
type PaymentContractUpsertBulk struct {
	create *PaymentContractCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PaymentContract.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(paymentcontract.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PaymentContractUpsertBulk) UpdateNewValues() *PaymentContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(paymentcontract.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PaymentContract.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PaymentContractUpsertBulk) Ignore() *PaymentContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PaymentContractUpsertBulk) DoNothing() *PaymentContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PaymentContractCreateBulk.OnConflict
// documentation for more info.
func (u *PaymentContractUpsertBulk) Update(set func(*PaymentContractUpsert)) *PaymentContractUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PaymentContractUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *PaymentContractUpsertBulk) SetCreatedAt(v uint32) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *PaymentContractUpsertBulk) AddCreatedAt(v uint32) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateCreatedAt() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *PaymentContractUpsertBulk) SetUpdatedAt(v uint32) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *PaymentContractUpsertBulk) AddUpdatedAt(v uint32) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateUpdatedAt() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *PaymentContractUpsertBulk) SetDeletedAt(v uint32) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *PaymentContractUpsertBulk) AddDeletedAt(v uint32) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateDeletedAt() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *PaymentContractUpsertBulk) SetEntID(v uuid.UUID) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateEntID() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateEntID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *PaymentContractUpsertBulk) SetOrderID(v uuid.UUID) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateOrderID() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *PaymentContractUpsertBulk) ClearOrderID() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.ClearOrderID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *PaymentContractUpsertBulk) SetCoinTypeID(v uuid.UUID) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateCoinTypeID() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *PaymentContractUpsertBulk) ClearCoinTypeID() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetAmount sets the "amount" field.
func (u *PaymentContractUpsertBulk) SetAmount(v decimal.Decimal) *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *PaymentContractUpsertBulk) UpdateAmount() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.UpdateAmount()
	})
}

// ClearAmount clears the value of the "amount" field.
func (u *PaymentContractUpsertBulk) ClearAmount() *PaymentContractUpsertBulk {
	return u.Update(func(s *PaymentContractUpsert) {
		s.ClearAmount()
	})
}

// Exec executes the query.
func (u *PaymentContractUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PaymentContractCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PaymentContractCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PaymentContractUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
