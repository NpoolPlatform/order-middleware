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
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/paymentcontract"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/predicate"
)

// PaymentContractQuery is the builder for querying PaymentContract entities.
type PaymentContractQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PaymentContract
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PaymentContractQuery builder.
func (pcq *PaymentContractQuery) Where(ps ...predicate.PaymentContract) *PaymentContractQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit adds a limit step to the query.
func (pcq *PaymentContractQuery) Limit(limit int) *PaymentContractQuery {
	pcq.limit = &limit
	return pcq
}

// Offset adds an offset step to the query.
func (pcq *PaymentContractQuery) Offset(offset int) *PaymentContractQuery {
	pcq.offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PaymentContractQuery) Unique(unique bool) *PaymentContractQuery {
	pcq.unique = &unique
	return pcq
}

// Order adds an order step to the query.
func (pcq *PaymentContractQuery) Order(o ...OrderFunc) *PaymentContractQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// First returns the first PaymentContract entity from the query.
// Returns a *NotFoundError when no PaymentContract was found.
func (pcq *PaymentContractQuery) First(ctx context.Context) (*PaymentContract, error) {
	nodes, err := pcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{paymentcontract.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PaymentContractQuery) FirstX(ctx context.Context) *PaymentContract {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PaymentContract ID from the query.
// Returns a *NotFoundError when no PaymentContract ID was found.
func (pcq *PaymentContractQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{paymentcontract.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PaymentContractQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PaymentContract entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PaymentContract entity is found.
// Returns a *NotFoundError when no PaymentContract entities are found.
func (pcq *PaymentContractQuery) Only(ctx context.Context) (*PaymentContract, error) {
	nodes, err := pcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{paymentcontract.Label}
	default:
		return nil, &NotSingularError{paymentcontract.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PaymentContractQuery) OnlyX(ctx context.Context) *PaymentContract {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PaymentContract ID in the query.
// Returns a *NotSingularError when more than one PaymentContract ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PaymentContractQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{paymentcontract.Label}
	default:
		err = &NotSingularError{paymentcontract.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PaymentContractQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PaymentContracts.
func (pcq *PaymentContractQuery) All(ctx context.Context) ([]*PaymentContract, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PaymentContractQuery) AllX(ctx context.Context) []*PaymentContract {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PaymentContract IDs.
func (pcq *PaymentContractQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := pcq.Select(paymentcontract.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PaymentContractQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PaymentContractQuery) Count(ctx context.Context) (int, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PaymentContractQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PaymentContractQuery) Exist(ctx context.Context) (bool, error) {
	if err := pcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PaymentContractQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PaymentContractQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PaymentContractQuery) Clone() *PaymentContractQuery {
	if pcq == nil {
		return nil
	}
	return &PaymentContractQuery{
		config:     pcq.config,
		limit:      pcq.limit,
		offset:     pcq.offset,
		order:      append([]OrderFunc{}, pcq.order...),
		predicates: append([]predicate.PaymentContract{}, pcq.predicates...),
		// clone intermediate query.
		sql:    pcq.sql.Clone(),
		path:   pcq.path,
		unique: pcq.unique,
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
//	client.PaymentContract.Query().
//		GroupBy(paymentcontract.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PaymentContractQuery) GroupBy(field string, fields ...string) *PaymentContractGroupBy {
	grbuild := &PaymentContractGroupBy{config: pcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pcq.sqlQuery(ctx), nil
	}
	grbuild.label = paymentcontract.Label
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
//	client.PaymentContract.Query().
//		Select(paymentcontract.FieldCreatedAt).
//		Scan(ctx, &v)
func (pcq *PaymentContractQuery) Select(fields ...string) *PaymentContractSelect {
	pcq.fields = append(pcq.fields, fields...)
	selbuild := &PaymentContractSelect{PaymentContractQuery: pcq}
	selbuild.label = paymentcontract.Label
	selbuild.flds, selbuild.scan = &pcq.fields, selbuild.Scan
	return selbuild
}

func (pcq *PaymentContractQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pcq.fields {
		if !paymentcontract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	if paymentcontract.Policy == nil {
		return errors.New("ent: uninitialized paymentcontract.Policy (forgotten import ent/runtime?)")
	}
	if err := paymentcontract.Policy.EvalQuery(ctx, pcq); err != nil {
		return err
	}
	return nil
}

func (pcq *PaymentContractQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PaymentContract, error) {
	var (
		nodes = []*PaymentContract{}
		_spec = pcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PaymentContract).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PaymentContract{config: pcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(pcq.modifiers) > 0 {
		_spec.Modifiers = pcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pcq *PaymentContractQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	if len(pcq.modifiers) > 0 {
		_spec.Modifiers = pcq.modifiers
	}
	_spec.Node.Columns = pcq.fields
	if len(pcq.fields) > 0 {
		_spec.Unique = pcq.unique != nil && *pcq.unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PaymentContractQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pcq *PaymentContractQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   paymentcontract.Table,
			Columns: paymentcontract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: paymentcontract.FieldID,
			},
		},
		From:   pcq.sql,
		Unique: true,
	}
	if unique := pcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, paymentcontract.FieldID)
		for i := range fields {
			if fields[i] != paymentcontract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PaymentContractQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(paymentcontract.Table)
	columns := pcq.fields
	if len(columns) == 0 {
		columns = paymentcontract.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.unique != nil && *pcq.unique {
		selector.Distinct()
	}
	for _, m := range pcq.modifiers {
		m(selector)
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (pcq *PaymentContractQuery) ForUpdate(opts ...sql.LockOption) *PaymentContractQuery {
	if pcq.driver.Dialect() == dialect.Postgres {
		pcq.Unique(false)
	}
	pcq.modifiers = append(pcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return pcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (pcq *PaymentContractQuery) ForShare(opts ...sql.LockOption) *PaymentContractQuery {
	if pcq.driver.Dialect() == dialect.Postgres {
		pcq.Unique(false)
	}
	pcq.modifiers = append(pcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return pcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pcq *PaymentContractQuery) Modify(modifiers ...func(s *sql.Selector)) *PaymentContractSelect {
	pcq.modifiers = append(pcq.modifiers, modifiers...)
	return pcq.Select()
}

// PaymentContractGroupBy is the group-by builder for PaymentContract entities.
type PaymentContractGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PaymentContractGroupBy) Aggregate(fns ...AggregateFunc) *PaymentContractGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pcgb *PaymentContractGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pcgb.path(ctx)
	if err != nil {
		return err
	}
	pcgb.sql = query
	return pcgb.sqlScan(ctx, v)
}

func (pcgb *PaymentContractGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pcgb.fields {
		if !paymentcontract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pcgb *PaymentContractGroupBy) sqlQuery() *sql.Selector {
	selector := pcgb.sql.Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pcgb.fields)+len(pcgb.fns))
		for _, f := range pcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pcgb.fields...)...)
}

// PaymentContractSelect is the builder for selecting fields of PaymentContract entities.
type PaymentContractSelect struct {
	*PaymentContractQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PaymentContractSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	pcs.sql = pcs.PaymentContractQuery.sqlQuery(ctx)
	return pcs.sqlScan(ctx, v)
}

func (pcs *PaymentContractSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pcs.sql.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pcs *PaymentContractSelect) Modify(modifiers ...func(s *sql.Selector)) *PaymentContractSelect {
	pcs.modifiers = append(pcs.modifiers, modifiers...)
	return pcs
}
