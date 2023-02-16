// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/pkg/ent/article"
	"github.com/morning-night-dream/platform-app/pkg/ent/readarticle"
)

// ReadArticleCreate is the builder for creating a ReadArticle entity.
type ReadArticleCreate struct {
	config
	mutation *ReadArticleMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetArticleID sets the "article_id" field.
func (rac *ReadArticleCreate) SetArticleID(u uuid.UUID) *ReadArticleCreate {
	rac.mutation.SetArticleID(u)
	return rac
}

// SetUserID sets the "user_id" field.
func (rac *ReadArticleCreate) SetUserID(u uuid.UUID) *ReadArticleCreate {
	rac.mutation.SetUserID(u)
	return rac
}

// SetReadAt sets the "read_at" field.
func (rac *ReadArticleCreate) SetReadAt(t time.Time) *ReadArticleCreate {
	rac.mutation.SetReadAt(t)
	return rac
}

// SetNillableReadAt sets the "read_at" field if the given value is not nil.
func (rac *ReadArticleCreate) SetNillableReadAt(t *time.Time) *ReadArticleCreate {
	if t != nil {
		rac.SetReadAt(*t)
	}
	return rac
}

// SetID sets the "id" field.
func (rac *ReadArticleCreate) SetID(u uuid.UUID) *ReadArticleCreate {
	rac.mutation.SetID(u)
	return rac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rac *ReadArticleCreate) SetNillableID(u *uuid.UUID) *ReadArticleCreate {
	if u != nil {
		rac.SetID(*u)
	}
	return rac
}

// SetArticle sets the "article" edge to the Article entity.
func (rac *ReadArticleCreate) SetArticle(a *Article) *ReadArticleCreate {
	return rac.SetArticleID(a.ID)
}

// Mutation returns the ReadArticleMutation object of the builder.
func (rac *ReadArticleCreate) Mutation() *ReadArticleMutation {
	return rac.mutation
}

// Save creates the ReadArticle in the database.
func (rac *ReadArticleCreate) Save(ctx context.Context) (*ReadArticle, error) {
	rac.defaults()
	return withHooks[*ReadArticle, ReadArticleMutation](ctx, rac.sqlSave, rac.mutation, rac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rac *ReadArticleCreate) SaveX(ctx context.Context) *ReadArticle {
	v, err := rac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rac *ReadArticleCreate) Exec(ctx context.Context) error {
	_, err := rac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rac *ReadArticleCreate) ExecX(ctx context.Context) {
	if err := rac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rac *ReadArticleCreate) defaults() {
	if _, ok := rac.mutation.ReadAt(); !ok {
		v := readarticle.DefaultReadAt()
		rac.mutation.SetReadAt(v)
	}
	if _, ok := rac.mutation.ID(); !ok {
		v := readarticle.DefaultID()
		rac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rac *ReadArticleCreate) check() error {
	if _, ok := rac.mutation.ArticleID(); !ok {
		return &ValidationError{Name: "article_id", err: errors.New(`ent: missing required field "ReadArticle.article_id"`)}
	}
	if _, ok := rac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ReadArticle.user_id"`)}
	}
	if _, ok := rac.mutation.ReadAt(); !ok {
		return &ValidationError{Name: "read_at", err: errors.New(`ent: missing required field "ReadArticle.read_at"`)}
	}
	if _, ok := rac.mutation.ArticleID(); !ok {
		return &ValidationError{Name: "article", err: errors.New(`ent: missing required edge "ReadArticle.article"`)}
	}
	return nil
}

func (rac *ReadArticleCreate) sqlSave(ctx context.Context) (*ReadArticle, error) {
	if err := rac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rac.mutation.id = &_node.ID
	rac.mutation.done = true
	return _node, nil
}

func (rac *ReadArticleCreate) createSpec() (*ReadArticle, *sqlgraph.CreateSpec) {
	var (
		_node = &ReadArticle{config: rac.config}
		_spec = sqlgraph.NewCreateSpec(readarticle.Table, sqlgraph.NewFieldSpec(readarticle.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = rac.conflict
	if id, ok := rac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rac.mutation.UserID(); ok {
		_spec.SetField(readarticle.FieldUserID, field.TypeUUID, value)
		_node.UserID = value
	}
	if value, ok := rac.mutation.ReadAt(); ok {
		_spec.SetField(readarticle.FieldReadAt, field.TypeTime, value)
		_node.ReadAt = value
	}
	if nodes := rac.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   readarticle.ArticleTable,
			Columns: []string{readarticle.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ArticleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ReadArticle.Create().
//		SetArticleID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReadArticleUpsert) {
//			SetArticleID(v+v).
//		}).
//		Exec(ctx)
func (rac *ReadArticleCreate) OnConflict(opts ...sql.ConflictOption) *ReadArticleUpsertOne {
	rac.conflict = opts
	return &ReadArticleUpsertOne{
		create: rac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ReadArticle.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rac *ReadArticleCreate) OnConflictColumns(columns ...string) *ReadArticleUpsertOne {
	rac.conflict = append(rac.conflict, sql.ConflictColumns(columns...))
	return &ReadArticleUpsertOne{
		create: rac,
	}
}

type (
	// ReadArticleUpsertOne is the builder for "upsert"-ing
	//  one ReadArticle node.
	ReadArticleUpsertOne struct {
		create *ReadArticleCreate
	}

	// ReadArticleUpsert is the "OnConflict" setter.
	ReadArticleUpsert struct {
		*sql.UpdateSet
	}
)

// SetArticleID sets the "article_id" field.
func (u *ReadArticleUpsert) SetArticleID(v uuid.UUID) *ReadArticleUpsert {
	u.Set(readarticle.FieldArticleID, v)
	return u
}

// UpdateArticleID sets the "article_id" field to the value that was provided on create.
func (u *ReadArticleUpsert) UpdateArticleID() *ReadArticleUpsert {
	u.SetExcluded(readarticle.FieldArticleID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ReadArticleUpsert) SetUserID(v uuid.UUID) *ReadArticleUpsert {
	u.Set(readarticle.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadArticleUpsert) UpdateUserID() *ReadArticleUpsert {
	u.SetExcluded(readarticle.FieldUserID)
	return u
}

// SetReadAt sets the "read_at" field.
func (u *ReadArticleUpsert) SetReadAt(v time.Time) *ReadArticleUpsert {
	u.Set(readarticle.FieldReadAt, v)
	return u
}

// UpdateReadAt sets the "read_at" field to the value that was provided on create.
func (u *ReadArticleUpsert) UpdateReadAt() *ReadArticleUpsert {
	u.SetExcluded(readarticle.FieldReadAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ReadArticle.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(readarticle.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ReadArticleUpsertOne) UpdateNewValues() *ReadArticleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(readarticle.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ReadArticle.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ReadArticleUpsertOne) Ignore() *ReadArticleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReadArticleUpsertOne) DoNothing() *ReadArticleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReadArticleCreate.OnConflict
// documentation for more info.
func (u *ReadArticleUpsertOne) Update(set func(*ReadArticleUpsert)) *ReadArticleUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReadArticleUpsert{UpdateSet: update})
	}))
	return u
}

// SetArticleID sets the "article_id" field.
func (u *ReadArticleUpsertOne) SetArticleID(v uuid.UUID) *ReadArticleUpsertOne {
	return u.Update(func(s *ReadArticleUpsert) {
		s.SetArticleID(v)
	})
}

// UpdateArticleID sets the "article_id" field to the value that was provided on create.
func (u *ReadArticleUpsertOne) UpdateArticleID() *ReadArticleUpsertOne {
	return u.Update(func(s *ReadArticleUpsert) {
		s.UpdateArticleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ReadArticleUpsertOne) SetUserID(v uuid.UUID) *ReadArticleUpsertOne {
	return u.Update(func(s *ReadArticleUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadArticleUpsertOne) UpdateUserID() *ReadArticleUpsertOne {
	return u.Update(func(s *ReadArticleUpsert) {
		s.UpdateUserID()
	})
}

// SetReadAt sets the "read_at" field.
func (u *ReadArticleUpsertOne) SetReadAt(v time.Time) *ReadArticleUpsertOne {
	return u.Update(func(s *ReadArticleUpsert) {
		s.SetReadAt(v)
	})
}

// UpdateReadAt sets the "read_at" field to the value that was provided on create.
func (u *ReadArticleUpsertOne) UpdateReadAt() *ReadArticleUpsertOne {
	return u.Update(func(s *ReadArticleUpsert) {
		s.UpdateReadAt()
	})
}

// Exec executes the query.
func (u *ReadArticleUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReadArticleCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReadArticleUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReadArticleUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ReadArticleUpsertOne.ID is not supported by MySQL driver. Use ReadArticleUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReadArticleUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReadArticleCreateBulk is the builder for creating many ReadArticle entities in bulk.
type ReadArticleCreateBulk struct {
	config
	builders []*ReadArticleCreate
	conflict []sql.ConflictOption
}

// Save creates the ReadArticle entities in the database.
func (racb *ReadArticleCreateBulk) Save(ctx context.Context) ([]*ReadArticle, error) {
	specs := make([]*sqlgraph.CreateSpec, len(racb.builders))
	nodes := make([]*ReadArticle, len(racb.builders))
	mutators := make([]Mutator, len(racb.builders))
	for i := range racb.builders {
		func(i int, root context.Context) {
			builder := racb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReadArticleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, racb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = racb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, racb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, racb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (racb *ReadArticleCreateBulk) SaveX(ctx context.Context) []*ReadArticle {
	v, err := racb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (racb *ReadArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := racb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (racb *ReadArticleCreateBulk) ExecX(ctx context.Context) {
	if err := racb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ReadArticle.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReadArticleUpsert) {
//			SetArticleID(v+v).
//		}).
//		Exec(ctx)
func (racb *ReadArticleCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReadArticleUpsertBulk {
	racb.conflict = opts
	return &ReadArticleUpsertBulk{
		create: racb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ReadArticle.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (racb *ReadArticleCreateBulk) OnConflictColumns(columns ...string) *ReadArticleUpsertBulk {
	racb.conflict = append(racb.conflict, sql.ConflictColumns(columns...))
	return &ReadArticleUpsertBulk{
		create: racb,
	}
}

// ReadArticleUpsertBulk is the builder for "upsert"-ing
// a bulk of ReadArticle nodes.
type ReadArticleUpsertBulk struct {
	create *ReadArticleCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ReadArticle.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(readarticle.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ReadArticleUpsertBulk) UpdateNewValues() *ReadArticleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(readarticle.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ReadArticle.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ReadArticleUpsertBulk) Ignore() *ReadArticleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReadArticleUpsertBulk) DoNothing() *ReadArticleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReadArticleCreateBulk.OnConflict
// documentation for more info.
func (u *ReadArticleUpsertBulk) Update(set func(*ReadArticleUpsert)) *ReadArticleUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReadArticleUpsert{UpdateSet: update})
	}))
	return u
}

// SetArticleID sets the "article_id" field.
func (u *ReadArticleUpsertBulk) SetArticleID(v uuid.UUID) *ReadArticleUpsertBulk {
	return u.Update(func(s *ReadArticleUpsert) {
		s.SetArticleID(v)
	})
}

// UpdateArticleID sets the "article_id" field to the value that was provided on create.
func (u *ReadArticleUpsertBulk) UpdateArticleID() *ReadArticleUpsertBulk {
	return u.Update(func(s *ReadArticleUpsert) {
		s.UpdateArticleID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ReadArticleUpsertBulk) SetUserID(v uuid.UUID) *ReadArticleUpsertBulk {
	return u.Update(func(s *ReadArticleUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ReadArticleUpsertBulk) UpdateUserID() *ReadArticleUpsertBulk {
	return u.Update(func(s *ReadArticleUpsert) {
		s.UpdateUserID()
	})
}

// SetReadAt sets the "read_at" field.
func (u *ReadArticleUpsertBulk) SetReadAt(v time.Time) *ReadArticleUpsertBulk {
	return u.Update(func(s *ReadArticleUpsert) {
		s.SetReadAt(v)
	})
}

// UpdateReadAt sets the "read_at" field to the value that was provided on create.
func (u *ReadArticleUpsertBulk) UpdateReadAt() *ReadArticleUpsertBulk {
	return u.Update(func(s *ReadArticleUpsert) {
		s.UpdateReadAt()
	})
}

// Exec executes the query.
func (u *ReadArticleUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReadArticleCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReadArticleCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReadArticleUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
