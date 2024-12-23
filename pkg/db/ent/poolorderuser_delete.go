// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/poolorderuser"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// PoolOrderUserDelete is the builder for deleting a PoolOrderUser entity.
type PoolOrderUserDelete struct {
	config
	hooks    []Hook
	mutation *PoolOrderUserMutation
}

// Where appends a list predicates to the PoolOrderUserDelete builder.
func (poud *PoolOrderUserDelete) Where(ps ...predicate.PoolOrderUser) *PoolOrderUserDelete {
	poud.mutation.Where(ps...)
	return poud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (poud *PoolOrderUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(poud.hooks) == 0 {
		affected, err = poud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolOrderUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			poud.mutation = mutation
			affected, err = poud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(poud.hooks) - 1; i >= 0; i-- {
			if poud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = poud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, poud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (poud *PoolOrderUserDelete) ExecX(ctx context.Context) int {
	n, err := poud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (poud *PoolOrderUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: poolorderuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: poolorderuser.FieldID,
			},
		},
	}
	if ps := poud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, poud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PoolOrderUserDeleteOne is the builder for deleting a single PoolOrderUser entity.
type PoolOrderUserDeleteOne struct {
	poud *PoolOrderUserDelete
}

// Exec executes the deletion query.
func (poudo *PoolOrderUserDeleteOne) Exec(ctx context.Context) error {
	n, err := poudo.poud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{poolorderuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (poudo *PoolOrderUserDeleteOne) ExecX(ctx context.Context) {
	poudo.poud.ExecX(ctx)
}
