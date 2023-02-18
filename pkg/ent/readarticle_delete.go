// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/morning-night-dream/platform-app/pkg/ent/predicate"
	"github.com/morning-night-dream/platform-app/pkg/ent/readarticle"
)

// ReadArticleDelete is the builder for deleting a ReadArticle entity.
type ReadArticleDelete struct {
	config
	hooks    []Hook
	mutation *ReadArticleMutation
}

// Where appends a list predicates to the ReadArticleDelete builder.
func (rad *ReadArticleDelete) Where(ps ...predicate.ReadArticle) *ReadArticleDelete {
	rad.mutation.Where(ps...)
	return rad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rad *ReadArticleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ReadArticleMutation](ctx, rad.sqlExec, rad.mutation, rad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rad *ReadArticleDelete) ExecX(ctx context.Context) int {
	n, err := rad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rad *ReadArticleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(readarticle.Table, sqlgraph.NewFieldSpec(readarticle.FieldID, field.TypeUUID))
	if ps := rad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rad.mutation.done = true
	return affected, err
}

// ReadArticleDeleteOne is the builder for deleting a single ReadArticle entity.
type ReadArticleDeleteOne struct {
	rad *ReadArticleDelete
}

// Where appends a list predicates to the ReadArticleDelete builder.
func (rado *ReadArticleDeleteOne) Where(ps ...predicate.ReadArticle) *ReadArticleDeleteOne {
	rado.rad.mutation.Where(ps...)
	return rado
}

// Exec executes the deletion query.
func (rado *ReadArticleDeleteOne) Exec(ctx context.Context) error {
	n, err := rado.rad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{readarticle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rado *ReadArticleDeleteOne) ExecX(ctx context.Context) {
	if err := rado.Exec(ctx); err != nil {
		panic(err)
	}
}