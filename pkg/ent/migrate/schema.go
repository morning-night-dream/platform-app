// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "article_title",
				Unique:  false,
				Columns: []*schema.Column{ArticlesColumns[1]},
			},
		},
	}
	// ArticleTagsColumns holds the columns for the "article_tags" table.
	ArticleTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "tag", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "article_id", Type: field.TypeUUID},
	}
	// ArticleTagsTable holds the schema information for the "article_tags" table.
	ArticleTagsTable = &schema.Table{
		Name:       "article_tags",
		Columns:    ArticleTagsColumns,
		PrimaryKey: []*schema.Column{ArticleTagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_tags_articles_tags",
				Columns:    []*schema.Column{ArticleTagsColumns[4]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "articletag_tag_article_id",
				Unique:  true,
				Columns: []*schema.Column{ArticleTagsColumns[1], ArticleTagsColumns[4]},
			},
			{
				Name:    "articletag_tag",
				Unique:  false,
				Columns: []*schema.Column{ArticleTagsColumns[1]},
			},
			{
				Name:    "articletag_article_id",
				Unique:  false,
				Columns: []*schema.Column{ArticleTagsColumns[4]},
			},
		},
	}
	// ReadArticlesColumns holds the columns for the "read_articles" table.
	ReadArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "read_at", Type: field.TypeTime},
		{Name: "article_id", Type: field.TypeUUID},
	}
	// ReadArticlesTable holds the schema information for the "read_articles" table.
	ReadArticlesTable = &schema.Table{
		Name:       "read_articles",
		Columns:    ReadArticlesColumns,
		PrimaryKey: []*schema.Column{ReadArticlesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "read_articles_articles_read_articles",
				Columns:    []*schema.Column{ReadArticlesColumns[3]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "readarticle_user_id_article_id",
				Unique:  true,
				Columns: []*schema.Column{ReadArticlesColumns[1], ReadArticlesColumns[3]},
			},
			{
				Name:    "readarticle_user_id",
				Unique:  false,
				Columns: []*schema.Column{ReadArticlesColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "last_logged_in_at", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		ArticleTagsTable,
		ReadArticlesTable,
		UsersTable,
	}
)

func init() {
	ArticleTagsTable.ForeignKeys[0].RefTable = ArticlesTable
	ReadArticlesTable.ForeignKeys[0].RefTable = ArticlesTable
}
