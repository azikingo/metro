// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"team-manager/ent/migrate"

	"team-manager/ent/direction"
	"team-manager/ent/schedule"
	"team-manager/ent/station"
	"team-manager/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Direction is the client for interacting with the Direction builders.
	Direction *DirectionClient
	// Schedule is the client for interacting with the Schedule builders.
	Schedule *ScheduleClient
	// Station is the client for interacting with the Station builders.
	Station *StationClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Direction = NewDirectionClient(c.config)
	c.Schedule = NewScheduleClient(c.config)
	c.Station = NewStationClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Direction: NewDirectionClient(cfg),
		Schedule:  NewScheduleClient(cfg),
		Station:   NewStationClient(cfg),
		User:      NewUserClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Direction: NewDirectionClient(cfg),
		Schedule:  NewScheduleClient(cfg),
		Station:   NewStationClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Direction.
//		Query().
//		Count(ctx)
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
	c.Direction.Use(hooks...)
	c.Schedule.Use(hooks...)
	c.Station.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Direction.Intercept(interceptors...)
	c.Schedule.Intercept(interceptors...)
	c.Station.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DirectionMutation:
		return c.Direction.mutate(ctx, m)
	case *ScheduleMutation:
		return c.Schedule.mutate(ctx, m)
	case *StationMutation:
		return c.Station.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DirectionClient is a client for the Direction schema.
type DirectionClient struct {
	config
}

// NewDirectionClient returns a client for the Direction from the given config.
func NewDirectionClient(c config) *DirectionClient {
	return &DirectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `direction.Hooks(f(g(h())))`.
func (c *DirectionClient) Use(hooks ...Hook) {
	c.hooks.Direction = append(c.hooks.Direction, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `direction.Intercept(f(g(h())))`.
func (c *DirectionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Direction = append(c.inters.Direction, interceptors...)
}

// Create returns a builder for creating a Direction entity.
func (c *DirectionClient) Create() *DirectionCreate {
	mutation := newDirectionMutation(c.config, OpCreate)
	return &DirectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Direction entities.
func (c *DirectionClient) CreateBulk(builders ...*DirectionCreate) *DirectionCreateBulk {
	return &DirectionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DirectionClient) MapCreateBulk(slice any, setFunc func(*DirectionCreate, int)) *DirectionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DirectionCreateBulk{err: fmt.Errorf("calling to DirectionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DirectionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DirectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Direction.
func (c *DirectionClient) Update() *DirectionUpdate {
	mutation := newDirectionMutation(c.config, OpUpdate)
	return &DirectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DirectionClient) UpdateOne(d *Direction) *DirectionUpdateOne {
	mutation := newDirectionMutation(c.config, OpUpdateOne, withDirection(d))
	return &DirectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DirectionClient) UpdateOneID(id int64) *DirectionUpdateOne {
	mutation := newDirectionMutation(c.config, OpUpdateOne, withDirectionID(id))
	return &DirectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Direction.
func (c *DirectionClient) Delete() *DirectionDelete {
	mutation := newDirectionMutation(c.config, OpDelete)
	return &DirectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DirectionClient) DeleteOne(d *Direction) *DirectionDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DirectionClient) DeleteOneID(id int64) *DirectionDeleteOne {
	builder := c.Delete().Where(direction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DirectionDeleteOne{builder}
}

// Query returns a query builder for Direction.
func (c *DirectionClient) Query() *DirectionQuery {
	return &DirectionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDirection},
		inters: c.Interceptors(),
	}
}

// Get returns a Direction entity by its id.
func (c *DirectionClient) Get(ctx context.Context, id int64) (*Direction, error) {
	return c.Query().Where(direction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DirectionClient) GetX(ctx context.Context, id int64) *Direction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDirectionSchedules queries the direction_schedules edge of a Direction.
func (c *DirectionClient) QueryDirectionSchedules(d *Direction) *ScheduleQuery {
	query := (&ScheduleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(direction.Table, direction.FieldID, id),
			sqlgraph.To(schedule.Table, schedule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, direction.DirectionSchedulesTable, direction.DirectionSchedulesColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DirectionClient) Hooks() []Hook {
	return c.hooks.Direction
}

// Interceptors returns the client interceptors.
func (c *DirectionClient) Interceptors() []Interceptor {
	return c.inters.Direction
}

func (c *DirectionClient) mutate(ctx context.Context, m *DirectionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DirectionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DirectionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DirectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DirectionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Direction mutation op: %q", m.Op())
	}
}

// ScheduleClient is a client for the Schedule schema.
type ScheduleClient struct {
	config
}

// NewScheduleClient returns a client for the Schedule from the given config.
func NewScheduleClient(c config) *ScheduleClient {
	return &ScheduleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schedule.Hooks(f(g(h())))`.
func (c *ScheduleClient) Use(hooks ...Hook) {
	c.hooks.Schedule = append(c.hooks.Schedule, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `schedule.Intercept(f(g(h())))`.
func (c *ScheduleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Schedule = append(c.inters.Schedule, interceptors...)
}

// Create returns a builder for creating a Schedule entity.
func (c *ScheduleClient) Create() *ScheduleCreate {
	mutation := newScheduleMutation(c.config, OpCreate)
	return &ScheduleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Schedule entities.
func (c *ScheduleClient) CreateBulk(builders ...*ScheduleCreate) *ScheduleCreateBulk {
	return &ScheduleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ScheduleClient) MapCreateBulk(slice any, setFunc func(*ScheduleCreate, int)) *ScheduleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ScheduleCreateBulk{err: fmt.Errorf("calling to ScheduleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ScheduleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ScheduleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Schedule.
func (c *ScheduleClient) Update() *ScheduleUpdate {
	mutation := newScheduleMutation(c.config, OpUpdate)
	return &ScheduleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScheduleClient) UpdateOne(s *Schedule) *ScheduleUpdateOne {
	mutation := newScheduleMutation(c.config, OpUpdateOne, withSchedule(s))
	return &ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScheduleClient) UpdateOneID(id int64) *ScheduleUpdateOne {
	mutation := newScheduleMutation(c.config, OpUpdateOne, withScheduleID(id))
	return &ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Schedule.
func (c *ScheduleClient) Delete() *ScheduleDelete {
	mutation := newScheduleMutation(c.config, OpDelete)
	return &ScheduleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ScheduleClient) DeleteOne(s *Schedule) *ScheduleDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ScheduleClient) DeleteOneID(id int64) *ScheduleDeleteOne {
	builder := c.Delete().Where(schedule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScheduleDeleteOne{builder}
}

// Query returns a query builder for Schedule.
func (c *ScheduleClient) Query() *ScheduleQuery {
	return &ScheduleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchedule},
		inters: c.Interceptors(),
	}
}

// Get returns a Schedule entity by its id.
func (c *ScheduleClient) Get(ctx context.Context, id int64) (*Schedule, error) {
	return c.Query().Where(schedule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScheduleClient) GetX(ctx context.Context, id int64) *Schedule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStation queries the station edge of a Schedule.
func (c *ScheduleClient) QueryStation(s *Schedule) *StationQuery {
	query := (&StationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schedule.Table, schedule.FieldID, id),
			sqlgraph.To(station.Table, station.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, schedule.StationTable, schedule.StationColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDirection queries the direction edge of a Schedule.
func (c *ScheduleClient) QueryDirection(s *Schedule) *DirectionQuery {
	query := (&DirectionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schedule.Table, schedule.FieldID, id),
			sqlgraph.To(direction.Table, direction.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, schedule.DirectionTable, schedule.DirectionColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ScheduleClient) Hooks() []Hook {
	return c.hooks.Schedule
}

// Interceptors returns the client interceptors.
func (c *ScheduleClient) Interceptors() []Interceptor {
	return c.inters.Schedule
}

func (c *ScheduleClient) mutate(ctx context.Context, m *ScheduleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ScheduleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ScheduleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ScheduleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Schedule mutation op: %q", m.Op())
	}
}

// StationClient is a client for the Station schema.
type StationClient struct {
	config
}

// NewStationClient returns a client for the Station from the given config.
func NewStationClient(c config) *StationClient {
	return &StationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `station.Hooks(f(g(h())))`.
func (c *StationClient) Use(hooks ...Hook) {
	c.hooks.Station = append(c.hooks.Station, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `station.Intercept(f(g(h())))`.
func (c *StationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Station = append(c.inters.Station, interceptors...)
}

// Create returns a builder for creating a Station entity.
func (c *StationClient) Create() *StationCreate {
	mutation := newStationMutation(c.config, OpCreate)
	return &StationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Station entities.
func (c *StationClient) CreateBulk(builders ...*StationCreate) *StationCreateBulk {
	return &StationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *StationClient) MapCreateBulk(slice any, setFunc func(*StationCreate, int)) *StationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &StationCreateBulk{err: fmt.Errorf("calling to StationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*StationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &StationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Station.
func (c *StationClient) Update() *StationUpdate {
	mutation := newStationMutation(c.config, OpUpdate)
	return &StationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StationClient) UpdateOne(s *Station) *StationUpdateOne {
	mutation := newStationMutation(c.config, OpUpdateOne, withStation(s))
	return &StationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StationClient) UpdateOneID(id int64) *StationUpdateOne {
	mutation := newStationMutation(c.config, OpUpdateOne, withStationID(id))
	return &StationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Station.
func (c *StationClient) Delete() *StationDelete {
	mutation := newStationMutation(c.config, OpDelete)
	return &StationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StationClient) DeleteOne(s *Station) *StationDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StationClient) DeleteOneID(id int64) *StationDeleteOne {
	builder := c.Delete().Where(station.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StationDeleteOne{builder}
}

// Query returns a query builder for Station.
func (c *StationClient) Query() *StationQuery {
	return &StationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStation},
		inters: c.Interceptors(),
	}
}

// Get returns a Station entity by its id.
func (c *StationClient) Get(ctx context.Context, id int64) (*Station, error) {
	return c.Query().Where(station.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StationClient) GetX(ctx context.Context, id int64) *Station {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStationSchedules queries the station_schedules edge of a Station.
func (c *StationClient) QueryStationSchedules(s *Station) *ScheduleQuery {
	query := (&ScheduleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(station.Table, station.FieldID, id),
			sqlgraph.To(schedule.Table, schedule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, station.StationSchedulesTable, station.StationSchedulesColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StationClient) Hooks() []Hook {
	return c.hooks.Station
}

// Interceptors returns the client interceptors.
func (c *StationClient) Interceptors() []Interceptor {
	return c.inters.Station
}

func (c *StationClient) mutate(ctx context.Context, m *StationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Station mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Direction, Schedule, Station, User []ent.Hook
	}
	inters struct {
		Direction, Schedule, Station, User []ent.Interceptor
	}
)
