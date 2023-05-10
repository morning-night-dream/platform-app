// Code generated by ent, DO NOT EDIT.

package readarticle

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldLTE(FieldID, id))
}

// ArticleID applies equality check predicate on the "article_id" field. It's identical to ArticleIDEQ.
func ArticleID(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldArticleID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldUserID, v))
}

// ReadAt applies equality check predicate on the "read_at" field. It's identical to ReadAtEQ.
func ReadAt(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldReadAt, v))
}

// ArticleIDEQ applies the EQ predicate on the "article_id" field.
func ArticleIDEQ(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldArticleID, v))
}

// ArticleIDNEQ applies the NEQ predicate on the "article_id" field.
func ArticleIDNEQ(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNEQ(FieldArticleID, v))
}

// ArticleIDIn applies the In predicate on the "article_id" field.
func ArticleIDIn(vs ...uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldIn(FieldArticleID, vs...))
}

// ArticleIDNotIn applies the NotIn predicate on the "article_id" field.
func ArticleIDNotIn(vs ...uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNotIn(FieldArticleID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldLTE(FieldUserID, v))
}

// ReadAtEQ applies the EQ predicate on the "read_at" field.
func ReadAtEQ(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldEQ(FieldReadAt, v))
}

// ReadAtNEQ applies the NEQ predicate on the "read_at" field.
func ReadAtNEQ(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNEQ(FieldReadAt, v))
}

// ReadAtIn applies the In predicate on the "read_at" field.
func ReadAtIn(vs ...time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldIn(FieldReadAt, vs...))
}

// ReadAtNotIn applies the NotIn predicate on the "read_at" field.
func ReadAtNotIn(vs ...time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldNotIn(FieldReadAt, vs...))
}

// ReadAtGT applies the GT predicate on the "read_at" field.
func ReadAtGT(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldGT(FieldReadAt, v))
}

// ReadAtGTE applies the GTE predicate on the "read_at" field.
func ReadAtGTE(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldGTE(FieldReadAt, v))
}

// ReadAtLT applies the LT predicate on the "read_at" field.
func ReadAtLT(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldLT(FieldReadAt, v))
}

// ReadAtLTE applies the LTE predicate on the "read_at" field.
func ReadAtLTE(v time.Time) predicate.ReadArticle {
	return predicate.ReadArticle(sql.FieldLTE(FieldReadAt, v))
}

// HasArticle applies the HasEdge predicate on the "article" edge.
func HasArticle() predicate.ReadArticle {
	return predicate.ReadArticle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArticleTable, ArticleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArticleWith applies the HasEdge predicate on the "article" edge with a given conditions (other predicates).
func HasArticleWith(preds ...predicate.Article) predicate.ReadArticle {
	return predicate.ReadArticle(func(s *sql.Selector) {
		step := newArticleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ReadArticle) predicate.ReadArticle {
	return predicate.ReadArticle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ReadArticle) predicate.ReadArticle {
	return predicate.ReadArticle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ReadArticle) predicate.ReadArticle {
	return predicate.ReadArticle(func(s *sql.Selector) {
		p(s.Not())
	})
}
