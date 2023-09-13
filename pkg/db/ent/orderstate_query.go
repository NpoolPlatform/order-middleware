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
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/orderstate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// OrderStateQuery is the builder for querying OrderState entities.
type OrderStateQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.OrderState
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderStateQuery builder.
func (osq *OrderStateQuery) Where(ps ...predicate.OrderState) *OrderStateQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit adds a limit step to the query.
func (osq *OrderStateQuery) Limit(limit int) *OrderStateQuery {
	osq.limit = &limit
	return osq
}

// Offset adds an offset step to the query.
func (osq *OrderStateQuery) Offset(offset int) *OrderStateQuery {
	osq.offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OrderStateQuery) Unique(unique bool) *OrderStateQuery {
	osq.unique = &unique
	return osq
}

// Order adds an order step to the query.
func (osq *OrderStateQuery) Order(o ...OrderFunc) *OrderStateQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// First returns the first OrderState entity from the query.
// Returns a *NotFoundError when no OrderState was found.
func (osq *OrderStateQuery) First(ctx context.Context) (*OrderState, error) {
	nodes, err := osq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderstate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OrderStateQuery) FirstX(ctx context.Context) *OrderState {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderState ID from the query.
// Returns a *NotFoundError when no OrderState ID was found.
func (osq *OrderStateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = osq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderstate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OrderStateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderState entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderState entity is found.
// Returns a *NotFoundError when no OrderState entities are found.
func (osq *OrderStateQuery) Only(ctx context.Context) (*OrderState, error) {
	nodes, err := osq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderstate.Label}
	default:
		return nil, &NotSingularError{orderstate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OrderStateQuery) OnlyX(ctx context.Context) *OrderState {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderState ID in the query.
// Returns a *NotSingularError when more than one OrderState ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OrderStateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = osq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderstate.Label}
	default:
		err = &NotSingularError{orderstate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OrderStateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderStates.
func (osq *OrderStateQuery) All(ctx context.Context) ([]*OrderState, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return osq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (osq *OrderStateQuery) AllX(ctx context.Context) []*OrderState {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderState IDs.
func (osq *OrderStateQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := osq.Select(orderstate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OrderStateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OrderStateQuery) Count(ctx context.Context) (int, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return osq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OrderStateQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OrderStateQuery) Exist(ctx context.Context) (bool, error) {
	if err := osq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return osq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OrderStateQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderStateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OrderStateQuery) Clone() *OrderStateQuery {
	if osq == nil {
		return nil
	}
	return &OrderStateQuery{
		config:     osq.config,
		limit:      osq.limit,
		offset:     osq.offset,
		order:      append([]OrderFunc{}, osq.order...),
		predicates: append([]predicate.OrderState{}, osq.predicates...),
		// clone intermediate query.
		sql:    osq.sql.Clone(),
		path:   osq.path,
		unique: osq.unique,
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
//	client.OrderState.Query().
//		GroupBy(orderstate.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (osq *OrderStateQuery) GroupBy(field string, fields ...string) *OrderStateGroupBy {
	grbuild := &OrderStateGroupBy{config: osq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return osq.sqlQuery(ctx), nil
	}
	grbuild.label = orderstate.Label
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
//	client.OrderState.Query().
//		Select(orderstate.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (osq *OrderStateQuery) Select(fields ...string) *OrderStateSelect {
	osq.fields = append(osq.fields, fields...)
	selbuild := &OrderStateSelect{OrderStateQuery: osq}
	selbuild.label = orderstate.Label
	selbuild.flds, selbuild.scan = &osq.fields, selbuild.Scan
	return selbuild
}

func (osq *OrderStateQuery) prepareQuery(ctx context.Context) error {
	for _, f := range osq.fields {
		if !orderstate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	if orderstate.Policy == nil {
		return errors.New("ent: uninitialized orderstate.Policy (forgotten import ent/runtime?)")
	}
	if err := orderstate.Policy.EvalQuery(ctx, osq); err != nil {
		return err
	}
	return nil
}

func (osq *OrderStateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderState, error) {
	var (
		nodes = []*OrderState{}
		_spec = osq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*OrderState).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &OrderState{config: osq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (osq *OrderStateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	_spec.Node.Columns = osq.fields
	if len(osq.fields) > 0 {
		_spec.Unique = osq.unique != nil && *osq.unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OrderStateQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := osq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (osq *OrderStateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderstate.Table,
			Columns: orderstate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderstate.FieldID,
			},
		},
		From:   osq.sql,
		Unique: true,
	}
	if unique := osq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := osq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstate.FieldID)
		for i := range fields {
			if fields[i] != orderstate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OrderStateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(orderstate.Table)
	columns := osq.fields
	if len(columns) == 0 {
		columns = orderstate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.unique != nil && *osq.unique {
		selector.Distinct()
	}
	for _, m := range osq.modifiers {
		m(selector)
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (osq *OrderStateQuery) ForUpdate(opts ...sql.LockOption) *OrderStateQuery {
	if osq.driver.Dialect() == dialect.Postgres {
		osq.Unique(false)
	}
	osq.modifiers = append(osq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return osq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (osq *OrderStateQuery) ForShare(opts ...sql.LockOption) *OrderStateQuery {
	if osq.driver.Dialect() == dialect.Postgres {
		osq.Unique(false)
	}
	osq.modifiers = append(osq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return osq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (osq *OrderStateQuery) Modify(modifiers ...func(s *sql.Selector)) *OrderStateSelect {
	osq.modifiers = append(osq.modifiers, modifiers...)
	return osq.Select()
}

// OrderStateGroupBy is the group-by builder for OrderState entities.
type OrderStateGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OrderStateGroupBy) Aggregate(fns ...AggregateFunc) *OrderStateGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the group-by query and scans the result into the given value.
func (osgb *OrderStateGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := osgb.path(ctx)
	if err != nil {
		return err
	}
	osgb.sql = query
	return osgb.sqlScan(ctx, v)
}

func (osgb *OrderStateGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range osgb.fields {
		if !orderstate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := osgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (osgb *OrderStateGroupBy) sqlQuery() *sql.Selector {
	selector := osgb.sql.Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(osgb.fields)+len(osgb.fns))
		for _, f := range osgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(osgb.fields...)...)
}

// OrderStateSelect is the builder for selecting fields of OrderState entities.
type OrderStateSelect struct {
	*OrderStateQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OrderStateSelect) Scan(ctx context.Context, v interface{}) error {
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	oss.sql = oss.OrderStateQuery.sqlQuery(ctx)
	return oss.sqlScan(ctx, v)
}

func (oss *OrderStateSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := oss.sql.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (oss *OrderStateSelect) Modify(modifiers ...func(s *sql.Selector)) *OrderStateSelect {
	oss.modifiers = append(oss.modifiers, modifiers...)
	return oss
}
