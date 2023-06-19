// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/pkg/ent/article"
	"github.com/morning-night-dream/platform-app/pkg/ent/predicate"
	"github.com/morning-night-dream/platform-app/pkg/ent/readarticle"
)

// ReadArticleUpdate is the builder for updating ReadArticle entities.
type ReadArticleUpdate struct {
	config
	hooks    []Hook
	mutation *ReadArticleMutation
}

// Where appends a list predicates to the ReadArticleUpdate builder.
func (rau *ReadArticleUpdate) Where(ps ...predicate.ReadArticle) *ReadArticleUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetArticleID sets the "article_id" field.
func (rau *ReadArticleUpdate) SetArticleID(u uuid.UUID) *ReadArticleUpdate {
	rau.mutation.SetArticleID(u)
	return rau
}

// SetUserID sets the "user_id" field.
func (rau *ReadArticleUpdate) SetUserID(u uuid.UUID) *ReadArticleUpdate {
	rau.mutation.SetUserID(u)
	return rau
}

// SetReadAt sets the "read_at" field.
func (rau *ReadArticleUpdate) SetReadAt(t time.Time) *ReadArticleUpdate {
	rau.mutation.SetReadAt(t)
	return rau
}

// SetNillableReadAt sets the "read_at" field if the given value is not nil.
func (rau *ReadArticleUpdate) SetNillableReadAt(t *time.Time) *ReadArticleUpdate {
	if t != nil {
		rau.SetReadAt(*t)
	}
	return rau
}

// SetArticle sets the "article" edge to the Article entity.
func (rau *ReadArticleUpdate) SetArticle(a *Article) *ReadArticleUpdate {
	return rau.SetArticleID(a.ID)
}

// Mutation returns the ReadArticleMutation object of the builder.
func (rau *ReadArticleUpdate) Mutation() *ReadArticleMutation {
	return rau.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (rau *ReadArticleUpdate) ClearArticle() *ReadArticleUpdate {
	rau.mutation.ClearArticle()
	return rau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *ReadArticleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rau.sqlSave, rau.mutation, rau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rau *ReadArticleUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *ReadArticleUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *ReadArticleUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rau *ReadArticleUpdate) check() error {
	if _, ok := rau.mutation.ArticleID(); rau.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ReadArticle.article"`)
	}
	return nil
}

func (rau *ReadArticleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(readarticle.Table, readarticle.Columns, sqlgraph.NewFieldSpec(readarticle.FieldID, field.TypeUUID))
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.UserID(); ok {
		_spec.SetField(readarticle.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := rau.mutation.ReadAt(); ok {
		_spec.SetField(readarticle.FieldReadAt, field.TypeTime, value)
	}
	if rau.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   readarticle.ArticleTable,
			Columns: []string{readarticle.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   readarticle.ArticleTable,
			Columns: []string{readarticle.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{readarticle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rau.mutation.done = true
	return n, nil
}

// ReadArticleUpdateOne is the builder for updating a single ReadArticle entity.
type ReadArticleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReadArticleMutation
}

// SetArticleID sets the "article_id" field.
func (rauo *ReadArticleUpdateOne) SetArticleID(u uuid.UUID) *ReadArticleUpdateOne {
	rauo.mutation.SetArticleID(u)
	return rauo
}

// SetUserID sets the "user_id" field.
func (rauo *ReadArticleUpdateOne) SetUserID(u uuid.UUID) *ReadArticleUpdateOne {
	rauo.mutation.SetUserID(u)
	return rauo
}

// SetReadAt sets the "read_at" field.
func (rauo *ReadArticleUpdateOne) SetReadAt(t time.Time) *ReadArticleUpdateOne {
	rauo.mutation.SetReadAt(t)
	return rauo
}

// SetNillableReadAt sets the "read_at" field if the given value is not nil.
func (rauo *ReadArticleUpdateOne) SetNillableReadAt(t *time.Time) *ReadArticleUpdateOne {
	if t != nil {
		rauo.SetReadAt(*t)
	}
	return rauo
}

// SetArticle sets the "article" edge to the Article entity.
func (rauo *ReadArticleUpdateOne) SetArticle(a *Article) *ReadArticleUpdateOne {
	return rauo.SetArticleID(a.ID)
}

// Mutation returns the ReadArticleMutation object of the builder.
func (rauo *ReadArticleUpdateOne) Mutation() *ReadArticleMutation {
	return rauo.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (rauo *ReadArticleUpdateOne) ClearArticle() *ReadArticleUpdateOne {
	rauo.mutation.ClearArticle()
	return rauo
}

// Where appends a list predicates to the ReadArticleUpdate builder.
func (rauo *ReadArticleUpdateOne) Where(ps ...predicate.ReadArticle) *ReadArticleUpdateOne {
	rauo.mutation.Where(ps...)
	return rauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *ReadArticleUpdateOne) Select(field string, fields ...string) *ReadArticleUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated ReadArticle entity.
func (rauo *ReadArticleUpdateOne) Save(ctx context.Context) (*ReadArticle, error) {
	return withHooks(ctx, rauo.sqlSave, rauo.mutation, rauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *ReadArticleUpdateOne) SaveX(ctx context.Context) *ReadArticle {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *ReadArticleUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *ReadArticleUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rauo *ReadArticleUpdateOne) check() error {
	if _, ok := rauo.mutation.ArticleID(); rauo.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ReadArticle.article"`)
	}
	return nil
}

func (rauo *ReadArticleUpdateOne) sqlSave(ctx context.Context) (_node *ReadArticle, err error) {
	if err := rauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(readarticle.Table, readarticle.Columns, sqlgraph.NewFieldSpec(readarticle.FieldID, field.TypeUUID))
	id, ok := rauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ReadArticle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, readarticle.FieldID)
		for _, f := range fields {
			if !readarticle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != readarticle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.UserID(); ok {
		_spec.SetField(readarticle.FieldUserID, field.TypeUUID, value)
	}
	if value, ok := rauo.mutation.ReadAt(); ok {
		_spec.SetField(readarticle.FieldReadAt, field.TypeTime, value)
	}
	if rauo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   readarticle.ArticleTable,
			Columns: []string{readarticle.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   readarticle.ArticleTable,
			Columns: []string{readarticle.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(article.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ReadArticle{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{readarticle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rauo.mutation.done = true
	return _node, nil
}
