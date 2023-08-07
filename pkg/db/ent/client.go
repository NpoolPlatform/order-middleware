// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/compensate"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/order"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/outofgas"
	"github.com/NpoolPlatform/order-middleware/pkg/db/ent/payment"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Compensate is the client for interacting with the Compensate builders.
	Compensate *CompensateClient
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// OutOfGas is the client for interacting with the OutOfGas builders.
	OutOfGas *OutOfGasClient
	// Payment is the client for interacting with the Payment builders.
	Payment *PaymentClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Compensate = NewCompensateClient(c.config)
	c.Order = NewOrderClient(c.config)
	c.OutOfGas = NewOutOfGasClient(c.config)
	c.Payment = NewPaymentClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Compensate: NewCompensateClient(cfg),
		Order:      NewOrderClient(cfg),
		OutOfGas:   NewOutOfGasClient(cfg),
		Payment:    NewPaymentClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Compensate: NewCompensateClient(cfg),
		Order:      NewOrderClient(cfg),
		OutOfGas:   NewOutOfGasClient(cfg),
		Payment:    NewPaymentClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Compensate.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Compensate.Use(hooks...)
	c.Order.Use(hooks...)
	c.OutOfGas.Use(hooks...)
	c.Payment.Use(hooks...)
}

// CompensateClient is a client for the Compensate schema.
type CompensateClient struct {
	config
}

// NewCompensateClient returns a client for the Compensate from the given config.
func NewCompensateClient(c config) *CompensateClient {
	return &CompensateClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `compensate.Hooks(f(g(h())))`.
func (c *CompensateClient) Use(hooks ...Hook) {
	c.hooks.Compensate = append(c.hooks.Compensate, hooks...)
}

// Create returns a builder for creating a Compensate entity.
func (c *CompensateClient) Create() *CompensateCreate {
	mutation := newCompensateMutation(c.config, OpCreate)
	return &CompensateCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Compensate entities.
func (c *CompensateClient) CreateBulk(builders ...*CompensateCreate) *CompensateCreateBulk {
	return &CompensateCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Compensate.
func (c *CompensateClient) Update() *CompensateUpdate {
	mutation := newCompensateMutation(c.config, OpUpdate)
	return &CompensateUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CompensateClient) UpdateOne(co *Compensate) *CompensateUpdateOne {
	mutation := newCompensateMutation(c.config, OpUpdateOne, withCompensate(co))
	return &CompensateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CompensateClient) UpdateOneID(id uuid.UUID) *CompensateUpdateOne {
	mutation := newCompensateMutation(c.config, OpUpdateOne, withCompensateID(id))
	return &CompensateUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Compensate.
func (c *CompensateClient) Delete() *CompensateDelete {
	mutation := newCompensateMutation(c.config, OpDelete)
	return &CompensateDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CompensateClient) DeleteOne(co *Compensate) *CompensateDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CompensateClient) DeleteOneID(id uuid.UUID) *CompensateDeleteOne {
	builder := c.Delete().Where(compensate.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CompensateDeleteOne{builder}
}

// Query returns a query builder for Compensate.
func (c *CompensateClient) Query() *CompensateQuery {
	return &CompensateQuery{
		config: c.config,
	}
}

// Get returns a Compensate entity by its id.
func (c *CompensateClient) Get(ctx context.Context, id uuid.UUID) (*Compensate, error) {
	return c.Query().Where(compensate.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CompensateClient) GetX(ctx context.Context, id uuid.UUID) *Compensate {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CompensateClient) Hooks() []Hook {
	hooks := c.hooks.Compensate
	return append(hooks[:len(hooks):len(hooks)], compensate.Hooks[:]...)
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Create returns a builder for creating a Order entity.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id uuid.UUID) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *OrderClient) DeleteOneID(id uuid.UUID) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{
		config: c.config,
	}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id uuid.UUID) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id uuid.UUID) *Order {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	hooks := c.hooks.Order
	return append(hooks[:len(hooks):len(hooks)], order.Hooks[:]...)
}

// OutOfGasClient is a client for the OutOfGas schema.
type OutOfGasClient struct {
	config
}

// NewOutOfGasClient returns a client for the OutOfGas from the given config.
func NewOutOfGasClient(c config) *OutOfGasClient {
	return &OutOfGasClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `outofgas.Hooks(f(g(h())))`.
func (c *OutOfGasClient) Use(hooks ...Hook) {
	c.hooks.OutOfGas = append(c.hooks.OutOfGas, hooks...)
}

// Create returns a builder for creating a OutOfGas entity.
func (c *OutOfGasClient) Create() *OutOfGasCreate {
	mutation := newOutOfGasMutation(c.config, OpCreate)
	return &OutOfGasCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OutOfGas entities.
func (c *OutOfGasClient) CreateBulk(builders ...*OutOfGasCreate) *OutOfGasCreateBulk {
	return &OutOfGasCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OutOfGas.
func (c *OutOfGasClient) Update() *OutOfGasUpdate {
	mutation := newOutOfGasMutation(c.config, OpUpdate)
	return &OutOfGasUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OutOfGasClient) UpdateOne(oog *OutOfGas) *OutOfGasUpdateOne {
	mutation := newOutOfGasMutation(c.config, OpUpdateOne, withOutOfGas(oog))
	return &OutOfGasUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OutOfGasClient) UpdateOneID(id uuid.UUID) *OutOfGasUpdateOne {
	mutation := newOutOfGasMutation(c.config, OpUpdateOne, withOutOfGasID(id))
	return &OutOfGasUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OutOfGas.
func (c *OutOfGasClient) Delete() *OutOfGasDelete {
	mutation := newOutOfGasMutation(c.config, OpDelete)
	return &OutOfGasDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OutOfGasClient) DeleteOne(oog *OutOfGas) *OutOfGasDeleteOne {
	return c.DeleteOneID(oog.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *OutOfGasClient) DeleteOneID(id uuid.UUID) *OutOfGasDeleteOne {
	builder := c.Delete().Where(outofgas.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OutOfGasDeleteOne{builder}
}

// Query returns a query builder for OutOfGas.
func (c *OutOfGasClient) Query() *OutOfGasQuery {
	return &OutOfGasQuery{
		config: c.config,
	}
}

// Get returns a OutOfGas entity by its id.
func (c *OutOfGasClient) Get(ctx context.Context, id uuid.UUID) (*OutOfGas, error) {
	return c.Query().Where(outofgas.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OutOfGasClient) GetX(ctx context.Context, id uuid.UUID) *OutOfGas {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *OutOfGasClient) Hooks() []Hook {
	hooks := c.hooks.OutOfGas
	return append(hooks[:len(hooks):len(hooks)], outofgas.Hooks[:]...)
}

// PaymentClient is a client for the Payment schema.
type PaymentClient struct {
	config
}

// NewPaymentClient returns a client for the Payment from the given config.
func NewPaymentClient(c config) *PaymentClient {
	return &PaymentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `payment.Hooks(f(g(h())))`.
func (c *PaymentClient) Use(hooks ...Hook) {
	c.hooks.Payment = append(c.hooks.Payment, hooks...)
}

// Create returns a builder for creating a Payment entity.
func (c *PaymentClient) Create() *PaymentCreate {
	mutation := newPaymentMutation(c.config, OpCreate)
	return &PaymentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Payment entities.
func (c *PaymentClient) CreateBulk(builders ...*PaymentCreate) *PaymentCreateBulk {
	return &PaymentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Payment.
func (c *PaymentClient) Update() *PaymentUpdate {
	mutation := newPaymentMutation(c.config, OpUpdate)
	return &PaymentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PaymentClient) UpdateOne(pa *Payment) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPayment(pa))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PaymentClient) UpdateOneID(id uuid.UUID) *PaymentUpdateOne {
	mutation := newPaymentMutation(c.config, OpUpdateOne, withPaymentID(id))
	return &PaymentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Payment.
func (c *PaymentClient) Delete() *PaymentDelete {
	mutation := newPaymentMutation(c.config, OpDelete)
	return &PaymentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PaymentClient) DeleteOne(pa *Payment) *PaymentDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PaymentClient) DeleteOneID(id uuid.UUID) *PaymentDeleteOne {
	builder := c.Delete().Where(payment.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PaymentDeleteOne{builder}
}

// Query returns a query builder for Payment.
func (c *PaymentClient) Query() *PaymentQuery {
	return &PaymentQuery{
		config: c.config,
	}
}

// Get returns a Payment entity by its id.
func (c *PaymentClient) Get(ctx context.Context, id uuid.UUID) (*Payment, error) {
	return c.Query().Where(payment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PaymentClient) GetX(ctx context.Context, id uuid.UUID) *Payment {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PaymentClient) Hooks() []Hook {
	hooks := c.hooks.Payment
	return append(hooks[:len(hooks):len(hooks)], payment.Hooks[:]...)
}
