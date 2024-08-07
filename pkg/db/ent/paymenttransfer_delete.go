// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymenttransfer"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// PaymentTransferDelete is the builder for deleting a PaymentTransfer entity.
type PaymentTransferDelete struct {
	config
	hooks    []Hook
	mutation *PaymentTransferMutation
}

// Where appends a list predicates to the PaymentTransferDelete builder.
func (ptd *PaymentTransferDelete) Where(ps ...predicate.PaymentTransfer) *PaymentTransferDelete {
	ptd.mutation.Where(ps...)
	return ptd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ptd *PaymentTransferDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptd.hooks) == 0 {
		affected, err = ptd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PaymentTransferMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ptd.mutation = mutation
			affected, err = ptd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptd.hooks) - 1; i >= 0; i-- {
			if ptd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptd *PaymentTransferDelete) ExecX(ctx context.Context) int {
	n, err := ptd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ptd *PaymentTransferDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: paymenttransfer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymenttransfer.FieldID,
			},
		},
	}
	if ps := ptd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ptd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PaymentTransferDeleteOne is the builder for deleting a single PaymentTransfer entity.
type PaymentTransferDeleteOne struct {
	ptd *PaymentTransferDelete
}

// Exec executes the deletion query.
func (ptdo *PaymentTransferDeleteOne) Exec(ctx context.Context) error {
	n, err := ptdo.ptd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{paymenttransfer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdo *PaymentTransferDeleteOne) ExecX(ctx context.Context) {
	ptdo.ptd.ExecX(ctx)
}
