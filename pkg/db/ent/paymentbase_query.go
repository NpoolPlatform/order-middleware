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
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentbase"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// PaymentBaseQuery is the builder for querying PaymentBase entities.
type PaymentBaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PaymentBase
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaymentBaseQuery builder.
func (pbq *PaymentBaseQuery) Where(ps ...predicate.PaymentBase) *PaymentBaseQuery {
	pbq.predicates = append(pbq.predicates, ps...)
	return pbq
}

// Limit adds a limit step to the query.
func (pbq *PaymentBaseQuery) Limit(limit int) *PaymentBaseQuery {
	pbq.limit = &limit
	return pbq
}

// Offset adds an offset step to the query.
func (pbq *PaymentBaseQuery) Offset(offset int) *PaymentBaseQuery {
	pbq.offset = &offset
	return pbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pbq *PaymentBaseQuery) Unique(unique bool) *PaymentBaseQuery {
	pbq.unique = &unique
	return pbq
}

// Order adds an order step to the query.
func (pbq *PaymentBaseQuery) Order(o ...OrderFunc) *PaymentBaseQuery {
	pbq.order = append(pbq.order, o...)
	return pbq
}

// First returns the first PaymentBase entity from the query.
// Returns a *NotFoundError when no PaymentBase was found.
func (pbq *PaymentBaseQuery) First(ctx context.Context) (*PaymentBase, error) {
	nodes, err := pbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{paymentbase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pbq *PaymentBaseQuery) FirstX(ctx context.Context) *PaymentBase {
	node, err := pbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaymentBase ID from the query.
// Returns a *NotFoundError when no PaymentBase ID was found.
func (pbq *PaymentBaseQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{paymentbase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pbq *PaymentBaseQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := pbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaymentBase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PaymentBase entity is found.
// Returns a *NotFoundError when no PaymentBase entities are found.
func (pbq *PaymentBaseQuery) Only(ctx context.Context) (*PaymentBase, error) {
	nodes, err := pbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{paymentbase.Label}
	default:
		return nil, &NotSingularError{paymentbase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pbq *PaymentBaseQuery) OnlyX(ctx context.Context) *PaymentBase {
	node, err := pbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaymentBase ID in the query.
// Returns a *NotSingularError when more than one PaymentBase ID is found.
// Returns a *NotFoundError when no entities are found.
func (pbq *PaymentBaseQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{paymentbase.Label}
	default:
		err = &NotSingularError{paymentbase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pbq *PaymentBaseQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := pbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaymentBases.
func (pbq *PaymentBaseQuery) All(ctx context.Context) ([]*PaymentBase, error) {
	if err := pbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pbq *PaymentBaseQuery) AllX(ctx context.Context) []*PaymentBase {
	nodes, err := pbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaymentBase IDs.
func (pbq *PaymentBaseQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := pbq.Select(paymentbase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pbq *PaymentBaseQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := pbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pbq *PaymentBaseQuery) Count(ctx context.Context) (int, error) {
	if err := pbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pbq *PaymentBaseQuery) CountX(ctx context.Context) int {
	count, err := pbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pbq *PaymentBaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := pbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pbq *PaymentBaseQuery) ExistX(ctx context.Context) bool {
	exist, err := pbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaymentBaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pbq *PaymentBaseQuery) Clone() *PaymentBaseQuery {
	if pbq == nil {
		return nil
	}
	return &PaymentBaseQuery{
		config:     pbq.config,
		limit:      pbq.limit,
		offset:     pbq.offset,
		order:      append([]OrderFunc{}, pbq.order...),
		predicates: append([]predicate.PaymentBase{}, pbq.predicates...),
		// clone intermediate query.
		sql:    pbq.sql.Clone(),
		path:   pbq.path,
		unique: pbq.unique,
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
//	client.PaymentBase.Query().
//		GroupBy(paymentbase.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pbq *PaymentBaseQuery) GroupBy(field string, fields ...string) *PaymentBaseGroupBy {
	grbuild := &PaymentBaseGroupBy{config: pbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pbq.sqlQuery(ctx), nil
	}
	grbuild.label = paymentbase.Label
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
//	client.PaymentBase.Query().
//		Select(paymentbase.FieldCreatedAt).
//		Scan(ctx, &v)
func (pbq *PaymentBaseQuery) Select(fields ...string) *PaymentBaseSelect {
	pbq.fields = append(pbq.fields, fields...)
	selbuild := &PaymentBaseSelect{PaymentBaseQuery: pbq}
	selbuild.label = paymentbase.Label
	selbuild.flds, selbuild.scan = &pbq.fields, selbuild.Scan
	return selbuild
}

func (pbq *PaymentBaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pbq.fields {
		if !paymentbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pbq.path != nil {
		prev, err := pbq.path(ctx)
		if err != nil {
			return err
		}
		pbq.sql = prev
	}
	if paymentbase.Policy == nil {
		return errors.New("ent: uninitialized paymentbase.Policy (forgotten import ent/runtime?)")
	}
	if err := paymentbase.Policy.EvalQuery(ctx, pbq); err != nil {
		return err
	}
	return nil
}

func (pbq *PaymentBaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PaymentBase, error) {
	var (
		nodes = []*PaymentBase{}
		_spec = pbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PaymentBase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PaymentBase{config: pbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(pbq.modifiers) > 0 {
		_spec.Modifiers = pbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pbq *PaymentBaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pbq.querySpec()
	if len(pbq.modifiers) > 0 {
		_spec.Modifiers = pbq.modifiers
	}
	_spec.Node.Columns = pbq.fields
	if len(pbq.fields) > 0 {
		_spec.Unique = pbq.unique != nil && *pbq.unique
	}
	return sqlgraph.CountNodes(ctx, pbq.driver, _spec)
}

func (pbq *PaymentBaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pbq *PaymentBaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentbase.Table,
			Columns: paymentbase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentbase.FieldID,
			},
		},
		From:   pbq.sql,
		Unique: true,
	}
	if unique := pbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentbase.FieldID)
		for i := range fields {
			if fields[i] != paymentbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pbq *PaymentBaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pbq.driver.Dialect())
	t1 := builder.Table(paymentbase.Table)
	columns := pbq.fields
	if len(columns) == 0 {
		columns = paymentbase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pbq.sql != nil {
		selector = pbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pbq.unique != nil && *pbq.unique {
		selector.Distinct()
	}
	for _, m := range pbq.modifiers {
		m(selector)
	}
	for _, p := range pbq.predicates {
		p(selector)
	}
	for _, p := range pbq.order {
		p(selector)
	}
	if offset := pbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (pbq *PaymentBaseQuery) ForUpdate(opts ...sql.LockOption) *PaymentBaseQuery {
	if pbq.driver.Dialect() == dialect.Postgres {
		pbq.Unique(false)
	}
	pbq.modifiers = append(pbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return pbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (pbq *PaymentBaseQuery) ForShare(opts ...sql.LockOption) *PaymentBaseQuery {
	if pbq.driver.Dialect() == dialect.Postgres {
		pbq.Unique(false)
	}
	pbq.modifiers = append(pbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return pbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pbq *PaymentBaseQuery) Modify(modifiers ...func(s *sql.Selector)) *PaymentBaseSelect {
	pbq.modifiers = append(pbq.modifiers, modifiers...)
	return pbq.Select()
}

// PaymentBaseGroupBy is the group-by builder for PaymentBase entities.
type PaymentBaseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pbgb *PaymentBaseGroupBy) Aggregate(fns ...AggregateFunc) *PaymentBaseGroupBy {
	pbgb.fns = append(pbgb.fns, fns...)
	return pbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pbgb *PaymentBaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pbgb.path(ctx)
	if err != nil {
		return err
	}
	pbgb.sql = query
	return pbgb.sqlScan(ctx, v)
}

func (pbgb *PaymentBaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pbgb.fields {
		if !paymentbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pbgb *PaymentBaseGroupBy) sqlQuery() *sql.Selector {
	selector := pbgb.sql.Select()
	aggregation := make([]string, 0, len(pbgb.fns))
	for _, fn := range pbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pbgb.fields)+len(pbgb.fns))
		for _, f := range pbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pbgb.fields...)...)
}

// PaymentBaseSelect is the builder for selecting fields of PaymentBase entities.
type PaymentBaseSelect struct {
	*PaymentBaseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pbs *PaymentBaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pbs.prepareQuery(ctx); err != nil {
		return err
	}
	pbs.sql = pbs.PaymentBaseQuery.sqlQuery(ctx)
	return pbs.sqlScan(ctx, v)
}

func (pbs *PaymentBaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pbs.sql.Query()
	if err := pbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pbs *PaymentBaseSelect) Modify(modifiers ...func(s *sql.Selector)) *PaymentBaseSelect {
	pbs.modifiers = append(pbs.modifiers, modifiers...)
	return pbs
}
