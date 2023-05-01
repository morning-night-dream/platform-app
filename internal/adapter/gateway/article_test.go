package gateway_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/morning-night-dream/platform-app/internal/adapter/gateway"
	"github.com/morning-night-dream/platform-app/internal/domain/model"
	"github.com/morning-night-dream/platform-app/pkg/ent"
	"github.com/morning-night-dream/platform-app/pkg/ent/enttest"
	"github.com/morning-night-dream/platform-app/pkg/ent/migrate"
)

func TestArticleSave(t *testing.T) {
	t.Parallel()

	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		// trueにすると、no such table: sqlite_sequenceでこけるため、falseにしておく
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(false)),
	}

	t.Run("記事を保存できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := gateway.NewArticle(db)

		ctx := context.Background()

		item := &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title",
			Url:         "url",
			Description: "description",
			Thumbnail:   "thumbnail",
			Tags:        []string{"tag"},
		}

		if err := sa.Save(ctx, item); err != nil {
			t.Error(err)
		}

		got, err := sa.Find(ctx, item.ArticleId)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(item, got) {
			t.Errorf("Find() = %v, want %v", got, item)
		}

		if err := sa.Save(ctx, item); err != nil {
			t.Error(err)
		}

		item.Title = "updated"
		item.Description = "updated"
		item.Thumbnail = "updated"
		item.Tags = []string{"updated"}

		if err := sa.Save(ctx, item); err != nil {
			t.Error(err)
		}

		got, err = sa.Find(ctx, item.ArticleId)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(item, got) {
			t.Errorf("Find() = %v, want %v", got, item)
		}
	})

	t.Run("記事を取得できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := gateway.NewArticle(db)

		ctx := context.Background()

		if err := sa.Save(ctx, &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title1",
			Url:         "url1",
			Description: "description1",
			Thumbnail:   "thumbnail1",
		}); err != nil {
			t.Error(err)
		}

		if err := sa.Save(ctx, &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title1",
			Url:         "url1",
			Description: "description1",
			Thumbnail:   "thumbnail1",
		}); err != nil {
			t.Error(err)
		}

		got, err := sa.FindAll(ctx, 1, 0)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(1, len(got)) {
			t.Errorf("NewArticle() = %v, want %v", len(got), 1)
		}
	})

	t.Run("記事を論理削除できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := gateway.NewArticle(db)

		ctx := context.Background()

		if err := sa.Save(ctx, &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title1",
			Url:         "url1",
			Description: "description1",
			Thumbnail:   "thumbnail1",
		}); err != nil {
			t.Error(err)
		}

		if err := sa.Save(ctx, &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title2",
			Url:         "url2",
			Description: "description2",
			Thumbnail:   "thumbnail2",
		}); err != nil {
			t.Error(err)
		}

		articles, err := sa.FindAll(ctx, 10, 0)
		if err != nil {
			t.Error(err)
		}

		id := articles[0].ArticleId

		if err := sa.LogicalDelete(ctx, id); err != nil {
			t.Error(err)
		}

		got, err := sa.FindAll(ctx, 10, 0)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(1, len(got)) {
			t.Errorf("NewArticle() = %v, want %v", len(got), 1)
		}
	})

	t.Run("タグ一覧を取得できる", func(t *testing.T) {
		t.Parallel()

		dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared&_fk=1", uuid.NewString())

		db := enttest.Open(t, "sqlite3", dsn, opts...)

		sa := gateway.NewArticle(db)

		ctx := context.Background()

		if err := sa.Save(ctx, &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title1",
			Url:         "url1",
			Description: "description1",
			Thumbnail:   "thumbnail1",
			Tags:        []string{"tag1"},
		}); err != nil {
			t.Error(err)
		}

		if err := sa.Save(ctx, &model.Article{
			ArticleId:   uuid.NewString(),
			Title:       "title2",
			Url:         "url2",
			Description: "description2",
			Thumbnail:   "thumbnail2",
			Tags:        []string{"tag2", "tag3"},
		}); err != nil {
			t.Error(err)
		}

		got, err := sa.FindAllTag(ctx)
		if err != nil {
			t.Error(err)
		}

		if !reflect.DeepEqual(got, []string{"tag1", "tag2", "tag3"}) {
			t.Errorf("Find() = %v, want %v", got, []string{"tag1", "tag2", "tag3"})
		}
	})
}
