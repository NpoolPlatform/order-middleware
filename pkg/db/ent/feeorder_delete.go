// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorder"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// FeeOrderDelete is the builder for deleting a FeeOrder entity.
type FeeOrderDelete struct {
	config
	hooks    []Hook
	mutation *FeeOrderMutation
}

// Where appends a list predicates to the FeeOrderDelete builder.
func (fod *FeeOrderDelete) Where(ps ...predicate.FeeOrder) *FeeOrderDelete {
	fod.mutation.Where(ps...)
	return fod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fod *FeeOrderDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fod.hooks) == 0 {
		affected, err = fod.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeeOrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fod.mutation = mutation
			affected, err = fod.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fod.hooks) - 1; i >= 0; i-- {
			if fod.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fod.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fod.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fod *FeeOrderDelete) ExecX(ctx context.Context) int {
	n, err := fod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fod *FeeOrderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: feeorder.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: feeorder.FieldID,
			},
		},
	}
	if ps := fod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FeeOrderDeleteOne is the builder for deleting a single FeeOrder entity.
type FeeOrderDeleteOne struct {
	fod *FeeOrderDelete
}

// Exec executes the deletion query.
func (fodo *FeeOrderDeleteOne) Exec(ctx context.Context) error {
	n, err := fodo.fod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feeorder.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fodo *FeeOrderDeleteOne) ExecX(ctx context.Context) {
	fodo.fod.ExecX(ctx)
}