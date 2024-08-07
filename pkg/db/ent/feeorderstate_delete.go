// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/feeorderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// FeeOrderStateDelete is the builder for deleting a FeeOrderState entity.
type FeeOrderStateDelete struct {
	config
	hooks    []Hook
	mutation *FeeOrderStateMutation
}

// Where appends a list predicates to the FeeOrderStateDelete builder.
func (fosd *FeeOrderStateDelete) Where(ps ...predicate.FeeOrderState) *FeeOrderStateDelete {
	fosd.mutation.Where(ps...)
	return fosd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fosd *FeeOrderStateDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fosd.hooks) == 0 {
		affected, err = fosd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeeOrderStateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fosd.mutation = mutation
			affected, err = fosd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fosd.hooks) - 1; i >= 0; i-- {
			if fosd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fosd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fosd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fosd *FeeOrderStateDelete) ExecX(ctx context.Context) int {
	n, err := fosd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fosd *FeeOrderStateDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: feeorderstate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: feeorderstate.FieldID,
			},
		},
	}
	if ps := fosd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fosd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FeeOrderStateDeleteOne is the builder for deleting a single FeeOrderState entity.
type FeeOrderStateDeleteOne struct {
	fosd *FeeOrderStateDelete
}

// Exec executes the deletion query.
func (fosdo *FeeOrderStateDeleteOne) Exec(ctx context.Context) error {
	n, err := fosdo.fosd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feeorderstate.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fosdo *FeeOrderStateDeleteOne) ExecX(ctx context.Context) {
	fosdo.fosd.ExecX(ctx)
}
