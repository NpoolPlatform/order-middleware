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
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderpaymentbalance"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// OrderPaymentBalanceQuery is the builder for querying OrderPaymentBalance entities.
type OrderPaymentBalanceQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderPaymentBalance
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderPaymentBalanceQuery builder.
func (opbq *OrderPaymentBalanceQuery) Where(ps ...predicate.OrderPaymentBalance) *OrderPaymentBalanceQuery {
	opbq.predicates = append(opbq.predicates, ps...)
	return opbq
}

// Limit adds a limit step to the query.
func (opbq *OrderPaymentBalanceQuery) Limit(limit int) *OrderPaymentBalanceQuery {
	opbq.limit = &limit
	return opbq
}

// Offset adds an offset step to the query.
func (opbq *OrderPaymentBalanceQuery) Offset(offset int) *OrderPaymentBalanceQuery {
	opbq.offset = &offset
	return opbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opbq *OrderPaymentBalanceQuery) Unique(unique bool) *OrderPaymentBalanceQuery {
	opbq.unique = &unique
	return opbq
}

// Order adds an order step to the query.
func (opbq *OrderPaymentBalanceQuery) Order(o ...OrderFunc) *OrderPaymentBalanceQuery {
	opbq.order = append(opbq.order, o...)
	return opbq
}

// First returns the first OrderPaymentBalance entity from the query.
// Returns a *NotFoundError when no OrderPaymentBalance was found.
func (opbq *OrderPaymentBalanceQuery) First(ctx context.Context) (*OrderPaymentBalance, error) {
	nodes, err := opbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderpaymentbalance.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) FirstX(ctx context.Context) *OrderPaymentBalance {
	node, err := opbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderPaymentBalance ID from the query.
// Returns a *NotFoundError when no OrderPaymentBalance ID was found.
func (opbq *OrderPaymentBalanceQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = opbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderpaymentbalance.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := opbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderPaymentBalance entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderPaymentBalance entity is found.
// Returns a *NotFoundError when no OrderPaymentBalance entities are found.
func (opbq *OrderPaymentBalanceQuery) Only(ctx context.Context) (*OrderPaymentBalance, error) {
	nodes, err := opbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderpaymentbalance.Label}
	default:
		return nil, &NotSingularError{orderpaymentbalance.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) OnlyX(ctx context.Context) *OrderPaymentBalance {
	node, err := opbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderPaymentBalance ID in the query.
// Returns a *NotSingularError when more than one OrderPaymentBalance ID is found.
// Returns a *NotFoundError when no entities are found.
func (opbq *OrderPaymentBalanceQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = opbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderpaymentbalance.Label}
	default:
		err = &NotSingularError{orderpaymentbalance.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := opbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderPaymentBalances.
func (opbq *OrderPaymentBalanceQuery) All(ctx context.Context) ([]*OrderPaymentBalance, error) {
	if err := opbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return opbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) AllX(ctx context.Context) []*OrderPaymentBalance {
	nodes, err := opbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderPaymentBalance IDs.
func (opbq *OrderPaymentBalanceQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := opbq.Select(orderpaymentbalance.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := opbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opbq *OrderPaymentBalanceQuery) Count(ctx context.Context) (int, error) {
	if err := opbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return opbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) CountX(ctx context.Context) int {
	count, err := opbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opbq *OrderPaymentBalanceQuery) Exist(ctx context.Context) (bool, error) {
	if err := opbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return opbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (opbq *OrderPaymentBalanceQuery) ExistX(ctx context.Context) bool {
	exist, err := opbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderPaymentBalanceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opbq *OrderPaymentBalanceQuery) Clone() *OrderPaymentBalanceQuery {
	if opbq == nil {
		return nil
	}
	return &OrderPaymentBalanceQuery{
		config:     opbq.config,
		limit:      opbq.limit,
		offset:     opbq.offset,
		order:      append([]OrderFunc{}, opbq.order...),
		predicates: append([]predicate.OrderPaymentBalance{}, opbq.predicates...),
		// clone intermediate query.
		sql:    opbq.sql.Clone(),
		path:   opbq.path,
		unique: opbq.unique,
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
//	client.OrderPaymentBalance.Query().
//		GroupBy(orderpaymentbalance.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (opbq *OrderPaymentBalanceQuery) GroupBy(field string, fields ...string) *OrderPaymentBalanceGroupBy {
	grbuild := &OrderPaymentBalanceGroupBy{config: opbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := opbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return opbq.sqlQuery(ctx), nil
	}
	grbuild.label = orderpaymentbalance.Label
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
//	client.OrderPaymentBalance.Query().
//		Select(orderpaymentbalance.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (opbq *OrderPaymentBalanceQuery) Select(fields ...string) *OrderPaymentBalanceSelect {
	opbq.fields = append(opbq.fields, fields...)
	selbuild := &OrderPaymentBalanceSelect{OrderPaymentBalanceQuery: opbq}
	selbuild.label = orderpaymentbalance.Label
	selbuild.flds, selbuild.scan = &opbq.fields, selbuild.Scan
	return selbuild
}

func (opbq *OrderPaymentBalanceQuery) prepareQuery(ctx context.Context) error {
	for _, f := range opbq.fields {
		if !orderpaymentbalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opbq.path != nil {
		prev, err := opbq.path(ctx)
		if err != nil {
			return err
		}
		opbq.sql = prev
	}
	if orderpaymentbalance.Policy == nil {
		return errors.New("ent: uninitialized orderpaymentbalance.Policy (forgotten import ent/runtime?)")
	}
	if err := orderpaymentbalance.Policy.EvalQuery(ctx, opbq); err != nil {
		return err
	}
	return nil
}

func (opbq *OrderPaymentBalanceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderPaymentBalance, error) {
	var (
		nodes = []*OrderPaymentBalance{}
		_spec = opbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrderPaymentBalance).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrderPaymentBalance{config: opbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(opbq.modifiers) > 0 {
		_spec.Modifiers = opbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (opbq *OrderPaymentBalanceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opbq.querySpec()
	if len(opbq.modifiers) > 0 {
		_spec.Modifiers = opbq.modifiers
	}
	_spec.Node.Columns = opbq.fields
	if len(opbq.fields) > 0 {
		_spec.Unique = opbq.unique != nil && *opbq.unique
	}
	return sqlgraph.CountNodes(ctx, opbq.driver, _spec)
}

func (opbq *OrderPaymentBalanceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := opbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (opbq *OrderPaymentBalanceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderpaymentbalance.Table,
			Columns: orderpaymentbalance.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderpaymentbalance.FieldID,
			},
		},
		From:   opbq.sql,
		Unique: true,
	}
	if unique := opbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := opbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpaymentbalance.FieldID)
		for i := range fields {
			if fields[i] != orderpaymentbalance.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := opbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opbq *OrderPaymentBalanceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opbq.driver.Dialect())
	t1 := builder.Table(orderpaymentbalance.Table)
	columns := opbq.fields
	if len(columns) == 0 {
		columns = orderpaymentbalance.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opbq.sql != nil {
		selector = opbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opbq.unique != nil && *opbq.unique {
		selector.Distinct()
	}
	for _, m := range opbq.modifiers {
		m(selector)
	}
	for _, p := range opbq.predicates {
		p(selector)
	}
	for _, p := range opbq.order {
		p(selector)
	}
	if offset := opbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (opbq *OrderPaymentBalanceQuery) ForUpdate(opts ...sql.LockOption) *OrderPaymentBalanceQuery {
	if opbq.driver.Dialect() == dialect.Postgres {
		opbq.Unique(false)
	}
	opbq.modifiers = append(opbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return opbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (opbq *OrderPaymentBalanceQuery) ForShare(opts ...sql.LockOption) *OrderPaymentBalanceQuery {
	if opbq.driver.Dialect() == dialect.Postgres {
		opbq.Unique(false)
	}
	opbq.modifiers = append(opbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return opbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opbq *OrderPaymentBalanceQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderPaymentBalanceSelect {
	opbq.modifiers = append(opbq.modifiers, modifiers...)
	return opbq.Select()
}

// OrderPaymentBalanceGroupBy is the group-by builder for OrderPaymentBalance entities.
type OrderPaymentBalanceGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opbgb *OrderPaymentBalanceGroupBy) Aggregate(fns ...AggregateFunc) *OrderPaymentBalanceGroupBy {
	opbgb.fns = append(opbgb.fns, fns...)
	return opbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (opbgb *OrderPaymentBalanceGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := opbgb.path(ctx)
	if err != nil {
		return err
	}
	opbgb.sql = query
	return opbgb.sqlScan(ctx, v)
}

func (opbgb *OrderPaymentBalanceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range opbgb.fields {
		if !orderpaymentbalance.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := opbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (opbgb *OrderPaymentBalanceGroupBy) sqlQuery() *sql.Selector {
	selector := opbgb.sql.Select()
	aggregation := make([]string, 0, len(opbgb.fns))
	for _, fn := range opbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(opbgb.fields)+len(opbgb.fns))
		for _, f := range opbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(opbgb.fields...)...)
}

// OrderPaymentBalanceSelect is the builder for selecting fields of OrderPaymentBalance entities.
type OrderPaymentBalanceSelect struct {
	*OrderPaymentBalanceQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (opbs *OrderPaymentBalanceSelect) Scan(ctx context.Context, v interface{}) error {
	if err := opbs.prepareQuery(ctx); err != nil {
		return err
	}
	opbs.sql = opbs.OrderPaymentBalanceQuery.sqlQuery(ctx)
	return opbs.sqlScan(ctx, v)
}

func (opbs *OrderPaymentBalanceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := opbs.sql.Query()
	if err := opbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opbs *OrderPaymentBalanceSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderPaymentBalanceSelect {
	opbs.modifiers = append(opbs.modifiers, modifiers...)
	return opbs
}
