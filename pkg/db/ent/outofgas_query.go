// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OutOfGasQuery is the builder for querying OutOfGas entities.
type OutOfGasQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OutOfGas
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OutOfGasQuery builder.
func (oogq *OutOfGasQuery) Where(ps ...predicate.OutOfGas) *OutOfGasQuery {
	oogq.predicates = append(oogq.predicates, ps...)
	return oogq
}

// Limit adds a limit step to the query.
func (oogq *OutOfGasQuery) Limit(limit int) *OutOfGasQuery {
	oogq.limit = &limit
	return oogq
}

// Offset adds an offset step to the query.
func (oogq *OutOfGasQuery) Offset(offset int) *OutOfGasQuery {
	oogq.offset = &offset
	return oogq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oogq *OutOfGasQuery) Unique(unique bool) *OutOfGasQuery {
	oogq.unique = &unique
	return oogq
}

// Order adds an order step to the query.
func (oogq *OutOfGasQuery) Order(o ...OrderFunc) *OutOfGasQuery {
	oogq.order = append(oogq.order, o...)
	return oogq
}

// First returns the first OutOfGas entity from the query.
// Returns a *NotFoundError when no OutOfGas was found.
func (oogq *OutOfGasQuery) First(ctx context.Context) (*OutOfGas, error) {
	nodes, err := oogq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{outofgas.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oogq *OutOfGasQuery) FirstX(ctx context.Context) *OutOfGas {
	node, err := oogq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OutOfGas ID from the query.
// Returns a *NotFoundError when no OutOfGas ID was found.
func (oogq *OutOfGasQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oogq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{outofgas.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oogq *OutOfGasQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := oogq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OutOfGas entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OutOfGas entity is found.
// Returns a *NotFoundError when no OutOfGas entities are found.
func (oogq *OutOfGasQuery) Only(ctx context.Context) (*OutOfGas, error) {
	nodes, err := oogq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{outofgas.Label}
	default:
		return nil, &NotSingularError{outofgas.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oogq *OutOfGasQuery) OnlyX(ctx context.Context) *OutOfGas {
	node, err := oogq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OutOfGas ID in the query.
// Returns a *NotSingularError when more than one OutOfGas ID is found.
// Returns a *NotFoundError when no entities are found.
func (oogq *OutOfGasQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = oogq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{outofgas.Label}
	default:
		err = &NotSingularError{outofgas.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oogq *OutOfGasQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := oogq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OutOfGasSlice.
func (oogq *OutOfGasQuery) All(ctx context.Context) ([]*OutOfGas, error) {
	if err := oogq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oogq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oogq *OutOfGasQuery) AllX(ctx context.Context) []*OutOfGas {
	nodes, err := oogq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OutOfGas IDs.
func (oogq *OutOfGasQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := oogq.Select(outofgas.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oogq *OutOfGasQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := oogq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oogq *OutOfGasQuery) Count(ctx context.Context) (int, error) {
	if err := oogq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oogq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oogq *OutOfGasQuery) CountX(ctx context.Context) int {
	count, err := oogq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oogq *OutOfGasQuery) Exist(ctx context.Context) (bool, error) {
	if err := oogq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oogq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oogq *OutOfGasQuery) ExistX(ctx context.Context) bool {
	exist, err := oogq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OutOfGasQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oogq *OutOfGasQuery) Clone() *OutOfGasQuery {
	if oogq == nil {
		return nil
	}
	return &OutOfGasQuery{
		config:     oogq.config,
		limit:      oogq.limit,
		offset:     oogq.offset,
		order:      append([]OrderFunc{}, oogq.order...),
		predicates: append([]predicate.OutOfGas{}, oogq.predicates...),
		// clone intermediate query.
		sql:    oogq.sql.Clone(),
		path:   oogq.path,
		unique: oogq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OutOfGas.Query().
//		GroupBy(outofgas.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (oogq *OutOfGasQuery) GroupBy(field string, fields ...string) *OutOfGasGroupBy {
	grbuild := &OutOfGasGroupBy{config: oogq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oogq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oogq.sqlQuery(ctx), nil
	}
	grbuild.label = outofgas.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.OutOfGas.Query().
//		Select(outofgas.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (oogq *OutOfGasQuery) Select(fields ...string) *OutOfGasSelect {
	oogq.fields = append(oogq.fields, fields...)
	selbuild := &OutOfGasSelect{OutOfGasQuery: oogq}
	selbuild.label = outofgas.Label
	selbuild.flds, selbuild.scan = &oogq.fields, selbuild.Scan
	return selbuild
}

func (oogq *OutOfGasQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oogq.fields {
		if !outofgas.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oogq.path != nil {
		prev, err := oogq.path(ctx)
		if err != nil {
			return err
		}
		oogq.sql = prev
	}
	if outofgas.Policy == nil {
		return errors.New("ent: uninitialized outofgas.Policy (forgotten import ent/runtime?)")
	}
	if err := outofgas.Policy.EvalQuery(ctx, oogq); err != nil {
		return err
	}
	return nil
}

func (oogq *OutOfGasQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OutOfGas, error) {
	var (
		nodes = []*OutOfGas{}
		_spec = oogq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OutOfGas).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OutOfGas{config: oogq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(oogq.modifiers) > 0 {
		_spec.Modifiers = oogq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oogq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oogq *OutOfGasQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oogq.querySpec()
	if len(oogq.modifiers) > 0 {
		_spec.Modifiers = oogq.modifiers
	}
	_spec.Node.Columns = oogq.fields
	if len(oogq.fields) > 0 {
		_spec.Unique = oogq.unique != nil && *oogq.unique
	}
	return sqlgraph.CountNodes(ctx, oogq.driver, _spec)
}

func (oogq *OutOfGasQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := oogq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (oogq *OutOfGasQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   outofgas.Table,
			Columns: outofgas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: outofgas.FieldID,
			},
		},
		From:   oogq.sql,
		Unique: true,
	}
	if unique := oogq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oogq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, outofgas.FieldID)
		for i := range fields {
			if fields[i] != outofgas.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oogq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oogq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oogq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oogq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oogq *OutOfGasQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oogq.driver.Dialect())
	t1 := builder.Table(outofgas.Table)
	columns := oogq.fields
	if len(columns) == 0 {
		columns = outofgas.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oogq.sql != nil {
		selector = oogq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oogq.unique != nil && *oogq.unique {
		selector.Distinct()
	}
	for _, m := range oogq.modifiers {
		m(selector)
	}
	for _, p := range oogq.predicates {
		p(selector)
	}
	for _, p := range oogq.order {
		p(selector)
	}
	if offset := oogq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oogq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (oogq *OutOfGasQuery) ForUpdate(opts ...sql.LockOption) *OutOfGasQuery {
	if oogq.driver.Dialect() == dialect.Postgres {
		oogq.Unique(false)
	}
	oogq.modifiers = append(oogq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return oogq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (oogq *OutOfGasQuery) ForShare(opts ...sql.LockOption) *OutOfGasQuery {
	if oogq.driver.Dialect() == dialect.Postgres {
		oogq.Unique(false)
	}
	oogq.modifiers = append(oogq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return oogq
}

// OutOfGasGroupBy is the group-by builder for OutOfGas entities.
type OutOfGasGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ooggb *OutOfGasGroupBy) Aggregate(fns ...AggregateFunc) *OutOfGasGroupBy {
	ooggb.fns = append(ooggb.fns, fns...)
	return ooggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ooggb *OutOfGasGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ooggb.path(ctx)
	if err != nil {
		return err
	}
	ooggb.sql = query
	return ooggb.sqlScan(ctx, v)
}

func (ooggb *OutOfGasGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ooggb.fields {
		if !outofgas.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ooggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ooggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ooggb *OutOfGasGroupBy) sqlQuery() *sql.Selector {
	selector := ooggb.sql.Select()
	aggregation := make([]string, 0, len(ooggb.fns))
	for _, fn := range ooggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ooggb.fields)+len(ooggb.fns))
		for _, f := range ooggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ooggb.fields...)...)
}

// OutOfGasSelect is the builder for selecting fields of OutOfGas entities.
type OutOfGasSelect struct {
	*OutOfGasQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oogs *OutOfGasSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oogs.prepareQuery(ctx); err != nil {
		return err
	}
	oogs.sql = oogs.OutOfGasQuery.sqlQuery(ctx)
	return oogs.sqlScan(ctx, v)
}

func (oogs *OutOfGasSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oogs.sql.Query()
	if err := oogs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
