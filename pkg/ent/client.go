// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/pkg/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/morning-night-dream/platform-app/pkg/ent/article"
	"github.com/morning-night-dream/platform-app/pkg/ent/articletag"
	"github.com/morning-night-dream/platform-app/pkg/ent/readarticle"
	"github.com/morning-night-dream/platform-app/pkg/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Article is the client for interacting with the Article builders.
	Article *ArticleClient
	// ArticleTag is the client for interacting with the ArticleTag builders.
	ArticleTag *ArticleTagClient
	// ReadArticle is the client for interacting with the ReadArticle builders.
	ReadArticle *ReadArticleClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Article = NewArticleClient(c.config)
	c.ArticleTag = NewArticleTagClient(c.config)
	c.ReadArticle = NewReadArticleClient(c.config)
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
		ctx:         ctx,
		config:      cfg,
		Article:     NewArticleClient(cfg),
		ArticleTag:  NewArticleTagClient(cfg),
		ReadArticle: NewReadArticleClient(cfg),
		User:        NewUserClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		Article:     NewArticleClient(cfg),
		ArticleTag:  NewArticleTagClient(cfg),
		ReadArticle: NewReadArticleClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Article.
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
	c.Article.Use(hooks...)
	c.ArticleTag.Use(hooks...)
	c.ReadArticle.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Article.Intercept(interceptors...)
	c.ArticleTag.Intercept(interceptors...)
	c.ReadArticle.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ArticleMutation:
		return c.Article.mutate(ctx, m)
	case *ArticleTagMutation:
		return c.ArticleTag.mutate(ctx, m)
	case *ReadArticleMutation:
		return c.ReadArticle.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ArticleClient is a client for the Article schema.
type ArticleClient struct {
	config
}

// NewArticleClient returns a client for the Article from the given config.
func NewArticleClient(c config) *ArticleClient {
	return &ArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `article.Hooks(f(g(h())))`.
func (c *ArticleClient) Use(hooks ...Hook) {
	c.hooks.Article = append(c.hooks.Article, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `article.Intercept(f(g(h())))`.
func (c *ArticleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Article = append(c.inters.Article, interceptors...)
}

// Create returns a builder for creating a Article entity.
func (c *ArticleClient) Create() *ArticleCreate {
	mutation := newArticleMutation(c.config, OpCreate)
	return &ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Article entities.
func (c *ArticleClient) CreateBulk(builders ...*ArticleCreate) *ArticleCreateBulk {
	return &ArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Article.
func (c *ArticleClient) Update() *ArticleUpdate {
	mutation := newArticleMutation(c.config, OpUpdate)
	return &ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleClient) UpdateOne(a *Article) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticle(a))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleClient) UpdateOneID(id uuid.UUID) *ArticleUpdateOne {
	mutation := newArticleMutation(c.config, OpUpdateOne, withArticleID(id))
	return &ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Article.
func (c *ArticleClient) Delete() *ArticleDelete {
	mutation := newArticleMutation(c.config, OpDelete)
	return &ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleClient) DeleteOne(a *Article) *ArticleDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleClient) DeleteOneID(id uuid.UUID) *ArticleDeleteOne {
	builder := c.Delete().Where(article.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleDeleteOne{builder}
}

// Query returns a query builder for Article.
func (c *ArticleClient) Query() *ArticleQuery {
	return &ArticleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeArticle},
		inters: c.Interceptors(),
	}
}

// Get returns a Article entity by its id.
func (c *ArticleClient) Get(ctx context.Context, id uuid.UUID) (*Article, error) {
	return c.Query().Where(article.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleClient) GetX(ctx context.Context, id uuid.UUID) *Article {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTags queries the tags edge of a Article.
func (c *ArticleClient) QueryTags(a *Article) *ArticleTagQuery {
	query := (&ArticleTagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(articletag.Table, articletag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, article.TagsTable, article.TagsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReadArticles queries the read_articles edge of a Article.
func (c *ArticleClient) QueryReadArticles(a *Article) *ReadArticleQuery {
	query := (&ReadArticleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(article.Table, article.FieldID, id),
			sqlgraph.To(readarticle.Table, readarticle.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, article.ReadArticlesTable, article.ReadArticlesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleClient) Hooks() []Hook {
	return c.hooks.Article
}

// Interceptors returns the client interceptors.
func (c *ArticleClient) Interceptors() []Interceptor {
	return c.inters.Article
}

func (c *ArticleClient) mutate(ctx context.Context, m *ArticleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ArticleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ArticleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Article mutation op: %q", m.Op())
	}
}

// ArticleTagClient is a client for the ArticleTag schema.
type ArticleTagClient struct {
	config
}

// NewArticleTagClient returns a client for the ArticleTag from the given config.
func NewArticleTagClient(c config) *ArticleTagClient {
	return &ArticleTagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `articletag.Hooks(f(g(h())))`.
func (c *ArticleTagClient) Use(hooks ...Hook) {
	c.hooks.ArticleTag = append(c.hooks.ArticleTag, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `articletag.Intercept(f(g(h())))`.
func (c *ArticleTagClient) Intercept(interceptors ...Interceptor) {
	c.inters.ArticleTag = append(c.inters.ArticleTag, interceptors...)
}

// Create returns a builder for creating a ArticleTag entity.
func (c *ArticleTagClient) Create() *ArticleTagCreate {
	mutation := newArticleTagMutation(c.config, OpCreate)
	return &ArticleTagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ArticleTag entities.
func (c *ArticleTagClient) CreateBulk(builders ...*ArticleTagCreate) *ArticleTagCreateBulk {
	return &ArticleTagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ArticleTag.
func (c *ArticleTagClient) Update() *ArticleTagUpdate {
	mutation := newArticleTagMutation(c.config, OpUpdate)
	return &ArticleTagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArticleTagClient) UpdateOne(at *ArticleTag) *ArticleTagUpdateOne {
	mutation := newArticleTagMutation(c.config, OpUpdateOne, withArticleTag(at))
	return &ArticleTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArticleTagClient) UpdateOneID(id uuid.UUID) *ArticleTagUpdateOne {
	mutation := newArticleTagMutation(c.config, OpUpdateOne, withArticleTagID(id))
	return &ArticleTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ArticleTag.
func (c *ArticleTagClient) Delete() *ArticleTagDelete {
	mutation := newArticleTagMutation(c.config, OpDelete)
	return &ArticleTagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ArticleTagClient) DeleteOne(at *ArticleTag) *ArticleTagDeleteOne {
	return c.DeleteOneID(at.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ArticleTagClient) DeleteOneID(id uuid.UUID) *ArticleTagDeleteOne {
	builder := c.Delete().Where(articletag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArticleTagDeleteOne{builder}
}

// Query returns a query builder for ArticleTag.
func (c *ArticleTagClient) Query() *ArticleTagQuery {
	return &ArticleTagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeArticleTag},
		inters: c.Interceptors(),
	}
}

// Get returns a ArticleTag entity by its id.
func (c *ArticleTagClient) Get(ctx context.Context, id uuid.UUID) (*ArticleTag, error) {
	return c.Query().Where(articletag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArticleTagClient) GetX(ctx context.Context, id uuid.UUID) *ArticleTag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticle queries the article edge of a ArticleTag.
func (c *ArticleTagClient) QueryArticle(at *ArticleTag) *ArticleQuery {
	query := (&ArticleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := at.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(articletag.Table, articletag.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, articletag.ArticleTable, articletag.ArticleColumn),
		)
		fromV = sqlgraph.Neighbors(at.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArticleTagClient) Hooks() []Hook {
	return c.hooks.ArticleTag
}

// Interceptors returns the client interceptors.
func (c *ArticleTagClient) Interceptors() []Interceptor {
	return c.inters.ArticleTag
}

func (c *ArticleTagClient) mutate(ctx context.Context, m *ArticleTagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ArticleTagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ArticleTagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ArticleTagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ArticleTagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ArticleTag mutation op: %q", m.Op())
	}
}

// ReadArticleClient is a client for the ReadArticle schema.
type ReadArticleClient struct {
	config
}

// NewReadArticleClient returns a client for the ReadArticle from the given config.
func NewReadArticleClient(c config) *ReadArticleClient {
	return &ReadArticleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `readarticle.Hooks(f(g(h())))`.
func (c *ReadArticleClient) Use(hooks ...Hook) {
	c.hooks.ReadArticle = append(c.hooks.ReadArticle, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `readarticle.Intercept(f(g(h())))`.
func (c *ReadArticleClient) Intercept(interceptors ...Interceptor) {
	c.inters.ReadArticle = append(c.inters.ReadArticle, interceptors...)
}

// Create returns a builder for creating a ReadArticle entity.
func (c *ReadArticleClient) Create() *ReadArticleCreate {
	mutation := newReadArticleMutation(c.config, OpCreate)
	return &ReadArticleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ReadArticle entities.
func (c *ReadArticleClient) CreateBulk(builders ...*ReadArticleCreate) *ReadArticleCreateBulk {
	return &ReadArticleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ReadArticle.
func (c *ReadArticleClient) Update() *ReadArticleUpdate {
	mutation := newReadArticleMutation(c.config, OpUpdate)
	return &ReadArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReadArticleClient) UpdateOne(ra *ReadArticle) *ReadArticleUpdateOne {
	mutation := newReadArticleMutation(c.config, OpUpdateOne, withReadArticle(ra))
	return &ReadArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReadArticleClient) UpdateOneID(id uuid.UUID) *ReadArticleUpdateOne {
	mutation := newReadArticleMutation(c.config, OpUpdateOne, withReadArticleID(id))
	return &ReadArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ReadArticle.
func (c *ReadArticleClient) Delete() *ReadArticleDelete {
	mutation := newReadArticleMutation(c.config, OpDelete)
	return &ReadArticleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReadArticleClient) DeleteOne(ra *ReadArticle) *ReadArticleDeleteOne {
	return c.DeleteOneID(ra.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReadArticleClient) DeleteOneID(id uuid.UUID) *ReadArticleDeleteOne {
	builder := c.Delete().Where(readarticle.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReadArticleDeleteOne{builder}
}

// Query returns a query builder for ReadArticle.
func (c *ReadArticleClient) Query() *ReadArticleQuery {
	return &ReadArticleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeReadArticle},
		inters: c.Interceptors(),
	}
}

// Get returns a ReadArticle entity by its id.
func (c *ReadArticleClient) Get(ctx context.Context, id uuid.UUID) (*ReadArticle, error) {
	return c.Query().Where(readarticle.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReadArticleClient) GetX(ctx context.Context, id uuid.UUID) *ReadArticle {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArticle queries the article edge of a ReadArticle.
func (c *ReadArticleClient) QueryArticle(ra *ReadArticle) *ArticleQuery {
	query := (&ArticleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ra.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(readarticle.Table, readarticle.FieldID, id),
			sqlgraph.To(article.Table, article.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, readarticle.ArticleTable, readarticle.ArticleColumn),
		)
		fromV = sqlgraph.Neighbors(ra.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReadArticleClient) Hooks() []Hook {
	return c.hooks.ReadArticle
}

// Interceptors returns the client interceptors.
func (c *ReadArticleClient) Interceptors() []Interceptor {
	return c.inters.ReadArticle
}

func (c *ReadArticleClient) mutate(ctx context.Context, m *ReadArticleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ReadArticleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ReadArticleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ReadArticleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ReadArticleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ReadArticle mutation op: %q", m.Op())
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
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
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
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
		Article, ArticleTag, ReadArticle, User []ent.Hook
	}
	inters struct {
		Article, ArticleTag, ReadArticle, User []ent.Interceptor
	}
)
