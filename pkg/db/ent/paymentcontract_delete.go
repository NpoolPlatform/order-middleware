// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentcontract"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// PaymentContractDelete is the builder for deleting a PaymentContract entity.
type PaymentContractDelete struct {
	config
	hooks    []Hook
	mutation *PaymentContractMutation
}

// Where appends a list predicates to the PaymentContractDelete builder.
func (pcd *PaymentContractDelete) Where(ps ...predicate.PaymentContract) *PaymentContractDelete {
	pcd.mutation.Where(ps...)
	return pcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pcd *PaymentContractDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcd.hooks) == 0 {
		affected, err = pcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentContractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pcd.mutation = mutation
			affected, err = pcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pcd.hooks) - 1; i >= 0; i-- {
			if pcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcd *PaymentContractDelete) ExecX(ctx context.Context) int {
	n, err := pcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pcd *PaymentContractDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: paymentcontract.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentcontract.FieldID,
			},
		},
	}
	if ps := pcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PaymentContractDeleteOne is the builder for deleting a single PaymentContract entity.
type PaymentContractDeleteOne struct {
	pcd *PaymentContractDelete
}

// Exec executes the deletion query.
func (pcdo *PaymentContractDeleteOne) Exec(ctx context.Context) error {
	n, err := pcdo.pcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{paymentcontract.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pcdo *PaymentContractDeleteOne) ExecX(ctx context.Context) {
	pcdo.pcd.ExecX(ctx)
}
