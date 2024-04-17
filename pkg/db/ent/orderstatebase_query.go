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
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstatebase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// OrderStateBaseQuery is the builder for querying OrderStateBase entities.
type OrderStateBaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderStateBase
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderStateBaseQuery builder.
func (osbq *OrderStateBaseQuery) Where(ps ...predicate.OrderStateBase) *OrderStateBaseQuery {
	osbq.predicates = append(osbq.predicates, ps...)
	return osbq
}

// Limit adds a limit step to the query.
func (osbq *OrderStateBaseQuery) Limit(limit int) *OrderStateBaseQuery {
	osbq.limit = &limit
	return osbq
}

// Offset adds an offset step to the query.
func (osbq *OrderStateBaseQuery) Offset(offset int) *OrderStateBaseQuery {
	osbq.offset = &offset
	return osbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osbq *OrderStateBaseQuery) Unique(unique bool) *OrderStateBaseQuery {
	osbq.unique = &unique
	return osbq
}

// Order adds an order step to the query.
func (osbq *OrderStateBaseQuery) Order(o ...OrderFunc) *OrderStateBaseQuery {
	osbq.order = append(osbq.order, o...)
	return osbq
}

// First returns the first OrderStateBase entity from the query.
// Returns a *NotFoundError when no OrderStateBase was found.
func (osbq *OrderStateBaseQuery) First(ctx context.Context) (*OrderStateBase, error) {
	nodes, err := osbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderstatebase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) FirstX(ctx context.Context) *OrderStateBase {
	node, err := osbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderStateBase ID from the query.
// Returns a *NotFoundError when no OrderStateBase ID was found.
func (osbq *OrderStateBaseQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = osbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderstatebase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := osbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderStateBase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderStateBase entity is found.
// Returns a *NotFoundError when no OrderStateBase entities are found.
func (osbq *OrderStateBaseQuery) Only(ctx context.Context) (*OrderStateBase, error) {
	nodes, err := osbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderstatebase.Label}
	default:
		return nil, &NotSingularError{orderstatebase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) OnlyX(ctx context.Context) *OrderStateBase {
	node, err := osbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderStateBase ID in the query.
// Returns a *NotSingularError when more than one OrderStateBase ID is found.
// Returns a *NotFoundError when no entities are found.
func (osbq *OrderStateBaseQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = osbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderstatebase.Label}
	default:
		err = &NotSingularError{orderstatebase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := osbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderStateBases.
func (osbq *OrderStateBaseQuery) All(ctx context.Context) ([]*OrderStateBase, error) {
	if err := osbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return osbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) AllX(ctx context.Context) []*OrderStateBase {
	nodes, err := osbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderStateBase IDs.
func (osbq *OrderStateBaseQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := osbq.Select(orderstatebase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := osbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osbq *OrderStateBaseQuery) Count(ctx context.Context) (int, error) {
	if err := osbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return osbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) CountX(ctx context.Context) int {
	count, err := osbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osbq *OrderStateBaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := osbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return osbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (osbq *OrderStateBaseQuery) ExistX(ctx context.Context) bool {
	exist, err := osbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderStateBaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osbq *OrderStateBaseQuery) Clone() *OrderStateBaseQuery {
	if osbq == nil {
		return nil
	}
	return &OrderStateBaseQuery{
		config:     osbq.config,
		limit:      osbq.limit,
		offset:     osbq.offset,
		order:      append([]OrderFunc{}, osbq.order...),
		predicates: append([]predicate.OrderStateBase{}, osbq.predicates...),
		// clone intermediate query.
		sql:    osbq.sql.Clone(),
		path:   osbq.path,
		unique: osbq.unique,
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
//	client.OrderStateBase.Query().
//		GroupBy(orderstatebase.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (osbq *OrderStateBaseQuery) GroupBy(field string, fields ...string) *OrderStateBaseGroupBy {
	grbuild := &OrderStateBaseGroupBy{config: osbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := osbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return osbq.sqlQuery(ctx), nil
	}
	grbuild.label = orderstatebase.Label
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
//	client.OrderStateBase.Query().
//		Select(orderstatebase.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (osbq *OrderStateBaseQuery) Select(fields ...string) *OrderStateBaseSelect {
	osbq.fields = append(osbq.fields, fields...)
	selbuild := &OrderStateBaseSelect{OrderStateBaseQuery: osbq}
	selbuild.label = orderstatebase.Label
	selbuild.flds, selbuild.scan = &osbq.fields, selbuild.Scan
	return selbuild
}

func (osbq *OrderStateBaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range osbq.fields {
		if !orderstatebase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if osbq.path != nil {
		prev, err := osbq.path(ctx)
		if err != nil {
			return err
		}
		osbq.sql = prev
	}
	if orderstatebase.Policy == nil {
		return errors.New("ent: uninitialized orderstatebase.Policy (forgotten import ent/runtime?)")
	}
	if err := orderstatebase.Policy.EvalQuery(ctx, osbq); err != nil {
		return err
	}
	return nil
}

func (osbq *OrderStateBaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderStateBase, error) {
	var (
		nodes = []*OrderStateBase{}
		_spec = osbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrderStateBase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrderStateBase{config: osbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(osbq.modifiers) > 0 {
		_spec.Modifiers = osbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (osbq *OrderStateBaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osbq.querySpec()
	if len(osbq.modifiers) > 0 {
		_spec.Modifiers = osbq.modifiers
	}
	_spec.Node.Columns = osbq.fields
	if len(osbq.fields) > 0 {
		_spec.Unique = osbq.unique != nil && *osbq.unique
	}
	return sqlgraph.CountNodes(ctx, osbq.driver, _spec)
}

func (osbq *OrderStateBaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := osbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (osbq *OrderStateBaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstatebase.Table,
			Columns: orderstatebase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderstatebase.FieldID,
			},
		},
		From:   osbq.sql,
		Unique: true,
	}
	if unique := osbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := osbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatebase.FieldID)
		for i := range fields {
			if fields[i] != orderstatebase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osbq *OrderStateBaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osbq.driver.Dialect())
	t1 := builder.Table(orderstatebase.Table)
	columns := osbq.fields
	if len(columns) == 0 {
		columns = orderstatebase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osbq.sql != nil {
		selector = osbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osbq.unique != nil && *osbq.unique {
		selector.Distinct()
	}
	for _, m := range osbq.modifiers {
		m(selector)
	}
	for _, p := range osbq.predicates {
		p(selector)
	}
	for _, p := range osbq.order {
		p(selector)
	}
	if offset := osbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (osbq *OrderStateBaseQuery) ForUpdate(opts ...sql.LockOption) *OrderStateBaseQuery {
	if osbq.driver.Dialect() == dialect.Postgres {
		osbq.Unique(false)
	}
	osbq.modifiers = append(osbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return osbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (osbq *OrderStateBaseQuery) ForShare(opts ...sql.LockOption) *OrderStateBaseQuery {
	if osbq.driver.Dialect() == dialect.Postgres {
		osbq.Unique(false)
	}
	osbq.modifiers = append(osbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return osbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osbq *OrderStateBaseQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderStateBaseSelect {
	osbq.modifiers = append(osbq.modifiers, modifiers...)
	return osbq.Select()
}

// OrderStateBaseGroupBy is the group-by builder for OrderStateBase entities.
type OrderStateBaseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osbgb *OrderStateBaseGroupBy) Aggregate(fns ...AggregateFunc) *OrderStateBaseGroupBy {
	osbgb.fns = append(osbgb.fns, fns...)
	return osbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (osbgb *OrderStateBaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := osbgb.path(ctx)
	if err != nil {
		return err
	}
	osbgb.sql = query
	return osbgb.sqlScan(ctx, v)
}

func (osbgb *OrderStateBaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range osbgb.fields {
		if !orderstatebase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := osbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (osbgb *OrderStateBaseGroupBy) sqlQuery() *sql.Selector {
	selector := osbgb.sql.Select()
	aggregation := make([]string, 0, len(osbgb.fns))
	for _, fn := range osbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(osbgb.fields)+len(osbgb.fns))
		for _, f := range osbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(osbgb.fields...)...)
}

// OrderStateBaseSelect is the builder for selecting fields of OrderStateBase entities.
type OrderStateBaseSelect struct {
	*OrderStateBaseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (osbs *OrderStateBaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := osbs.prepareQuery(ctx); err != nil {
		return err
	}
	osbs.sql = osbs.OrderStateBaseQuery.sqlQuery(ctx)
	return osbs.sqlScan(ctx, v)
}

func (osbs *OrderStateBaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := osbs.sql.Query()
	if err := osbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osbs *OrderStateBaseSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderStateBaseSelect {
	osbs.modifiers = append(osbs.modifiers, modifiers...)
	return osbs
}