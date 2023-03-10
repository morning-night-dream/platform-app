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
	"github.com/morning-night-dream/platform-app/pkg/ent/auth"
	"github.com/morning-night-dream/platform-app/pkg/ent/user"
)

// AuthCreate is the builder for creating a Auth entity.
type AuthCreate struct {
	config
	mutation *AuthMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (ac *AuthCreate) SetUserID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetUserID(u)
	return ac
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableUserID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetUserID(*u)
	}
	return ac
}

// SetLoginID sets the "login_id" field.
func (ac *AuthCreate) SetLoginID(s string) *AuthCreate {
	ac.mutation.SetLoginID(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AuthCreate) SetEmail(s string) *AuthCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AuthCreate) SetPassword(s string) *AuthCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AuthCreate) SetCreatedAt(t time.Time) *AuthCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AuthCreate) SetNillableCreatedAt(t *time.Time) *AuthCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AuthCreate) SetUpdatedAt(t time.Time) *AuthCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AuthCreate) SetNillableUpdatedAt(t *time.Time) *AuthCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AuthCreate) SetID(u uuid.UUID) *AuthCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AuthCreate) SetNillableID(u *uuid.UUID) *AuthCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AuthCreate) SetUser(u *User) *AuthCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AuthMutation object of the builder.
func (ac *AuthCreate) Mutation() *AuthMutation {
	return ac.mutation
}

// Save creates the Auth in the database.
func (ac *AuthCreate) Save(ctx context.Context) (*Auth, error) {
	ac.defaults()
	return withHooks[*Auth, AuthMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuthCreate) SaveX(ctx context.Context) *Auth {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuthCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuthCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuthCreate) defaults() {
	if _, ok := ac.mutation.UserID(); !ok {
		v := auth.DefaultUserID()
		ac.mutation.SetUserID(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := auth.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := auth.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := auth.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuthCreate) check() error {
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Auth.user_id"`)}
	}
	if _, ok := ac.mutation.LoginID(); !ok {
		return &ValidationError{Name: "login_id", err: errors.New(`ent: missing required field "Auth.login_id"`)}
	}
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Auth.email"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Auth.password"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Auth.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Auth.updated_at"`)}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Auth.user"`)}
	}
	return nil
}

func (ac *AuthCreate) sqlSave(ctx context.Context) (*Auth, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AuthCreate) createSpec() (*Auth, *sqlgraph.CreateSpec) {
	var (
		_node = &Auth{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(auth.Table, sqlgraph.NewFieldSpec(auth.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.LoginID(); ok {
		_spec.SetField(auth.FieldLoginID, field.TypeString, value)
		_node.LoginID = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(auth.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(auth.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(auth.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(auth.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   auth.UserTable,
			Columns: []string{auth.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Auth.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (ac *AuthCreate) OnConflict(opts ...sql.ConflictOption) *AuthUpsertOne {
	ac.conflict = opts
	return &AuthUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AuthCreate) OnConflictColumns(columns ...string) *AuthUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AuthUpsertOne{
		create: ac,
	}
}

type (
	// AuthUpsertOne is the builder for "upsert"-ing
	//  one Auth node.
	AuthUpsertOne struct {
		create *AuthCreate
	}

	// AuthUpsert is the "OnConflict" setter.
	AuthUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *AuthUpsert) SetUserID(v uuid.UUID) *AuthUpsert {
	u.Set(auth.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthUpsert) UpdateUserID() *AuthUpsert {
	u.SetExcluded(auth.FieldUserID)
	return u
}

// SetLoginID sets the "login_id" field.
func (u *AuthUpsert) SetLoginID(v string) *AuthUpsert {
	u.Set(auth.FieldLoginID, v)
	return u
}

// UpdateLoginID sets the "login_id" field to the value that was provided on create.
func (u *AuthUpsert) UpdateLoginID() *AuthUpsert {
	u.SetExcluded(auth.FieldLoginID)
	return u
}

// SetEmail sets the "email" field.
func (u *AuthUpsert) SetEmail(v string) *AuthUpsert {
	u.Set(auth.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AuthUpsert) UpdateEmail() *AuthUpsert {
	u.SetExcluded(auth.FieldEmail)
	return u
}

// SetPassword sets the "password" field.
func (u *AuthUpsert) SetPassword(v string) *AuthUpsert {
	u.Set(auth.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AuthUpsert) UpdatePassword() *AuthUpsert {
	u.SetExcluded(auth.FieldPassword)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthUpsert) SetCreatedAt(v time.Time) *AuthUpsert {
	u.Set(auth.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthUpsert) UpdateCreatedAt() *AuthUpsert {
	u.SetExcluded(auth.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthUpsert) SetUpdatedAt(v time.Time) *AuthUpsert {
	u.Set(auth.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthUpsert) UpdateUpdatedAt() *AuthUpsert {
	u.SetExcluded(auth.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auth.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthUpsertOne) UpdateNewValues() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(auth.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuthUpsertOne) Ignore() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthUpsertOne) DoNothing() *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthCreate.OnConflict
// documentation for more info.
func (u *AuthUpsertOne) Update(set func(*AuthUpsert)) *AuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *AuthUpsertOne) SetUserID(v uuid.UUID) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateUserID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUserID()
	})
}

// SetLoginID sets the "login_id" field.
func (u *AuthUpsertOne) SetLoginID(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetLoginID(v)
	})
}

// UpdateLoginID sets the "login_id" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateLoginID() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateLoginID()
	})
}

// SetEmail sets the "email" field.
func (u *AuthUpsertOne) SetEmail(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateEmail() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateEmail()
	})
}

// SetPassword sets the "password" field.
func (u *AuthUpsertOne) SetPassword(v string) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdatePassword() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdatePassword()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthUpsertOne) SetCreatedAt(v time.Time) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateCreatedAt() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthUpsertOne) SetUpdatedAt(v time.Time) *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthUpsertOne) UpdateUpdatedAt() *AuthUpsertOne {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *AuthUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AuthUpsertOne.ID is not supported by MySQL driver. Use AuthUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthCreateBulk is the builder for creating many Auth entities in bulk.
type AuthCreateBulk struct {
	config
	builders []*AuthCreate
	conflict []sql.ConflictOption
}

// Save creates the Auth entities in the database.
func (acb *AuthCreateBulk) Save(ctx context.Context) ([]*Auth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Auth, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuthCreateBulk) SaveX(ctx context.Context) []*Auth {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuthCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuthCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Auth.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (acb *AuthCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthUpsertBulk {
	acb.conflict = opts
	return &AuthUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AuthCreateBulk) OnConflictColumns(columns ...string) *AuthUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AuthUpsertBulk{
		create: acb,
	}
}

// AuthUpsertBulk is the builder for "upsert"-ing
// a bulk of Auth nodes.
type AuthUpsertBulk struct {
	create *AuthCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(auth.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuthUpsertBulk) UpdateNewValues() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(auth.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Auth.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuthUpsertBulk) Ignore() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthUpsertBulk) DoNothing() *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthCreateBulk.OnConflict
// documentation for more info.
func (u *AuthUpsertBulk) Update(set func(*AuthUpsert)) *AuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *AuthUpsertBulk) SetUserID(v uuid.UUID) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateUserID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUserID()
	})
}

// SetLoginID sets the "login_id" field.
func (u *AuthUpsertBulk) SetLoginID(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetLoginID(v)
	})
}

// UpdateLoginID sets the "login_id" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateLoginID() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateLoginID()
	})
}

// SetEmail sets the "email" field.
func (u *AuthUpsertBulk) SetEmail(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateEmail() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateEmail()
	})
}

// SetPassword sets the "password" field.
func (u *AuthUpsertBulk) SetPassword(v string) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdatePassword() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdatePassword()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AuthUpsertBulk) SetCreatedAt(v time.Time) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateCreatedAt() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AuthUpsertBulk) SetUpdatedAt(v time.Time) *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AuthUpsertBulk) UpdateUpdatedAt() *AuthUpsertBulk {
	return u.Update(func(s *AuthUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *AuthUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
