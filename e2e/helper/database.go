package helper

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/morning-night-dream/platform-app/internal/driver/database"
	"github.com/morning-night-dream/platform-app/pkg/ent"
)

type ArticleDB struct {
	T      *testing.T
	Client *ent.Client
}

func NewArticleDB(
	t *testing.T,
	dsn string,
) *ArticleDB {
	t.Helper()

	return &ArticleDB{
		T:      t,
		Client: database.NewClient(dsn),
	}
}

func (adb *ArticleDB) BulkInsert(ids []string) {
	adb.T.Helper()

	bulk := make([]*ent.ArticleCreate, len(ids))

	for i, id := range ids {
		bulk[i] = adb.Client.Article.Create().
			SetID(uuid.MustParse(id)).
			SetTitle("title-" + id).
			SetURL("https://example.com/" + id).
			SetDescription("description").
			SetThumbnail("https://example.com/" + id).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now())
	}

	if err := adb.Client.Article.CreateBulk(bulk...).OnConflict().UpdateNewValues().DoNothing().Exec(context.Background()); err != nil {
		adb.T.Fatal(err)
	}
}

func (adb ArticleDB) BulkDelete(ids []string) {
	adb.T.Helper()

	tx, err := adb.Client.Tx(context.Background())
	if err != nil {
		adb.T.Error(err)

		return
	}

	for _, id := range ids {
		if err := tx.Article.DeleteOneID(uuid.MustParse(id)).Exec(context.Background()); err != nil {
			adb.T.Error(err)

			_ = tx.Rollback()
		}
	}

	if err := tx.Commit(); err != nil {
		adb.T.Error(err)
	}
}

type UserDB struct {
	T      *testing.T
	Client *ent.Client
}

func NewUserDB(
	t *testing.T,
	dsn string,
) *UserDB {
	t.Helper()

	return &UserDB{
		T:      t,
		Client: database.NewClient(dsn),
	}
}

func (udb *UserDB) BulkDelete(ids []string) {
	tx, err := udb.Client.Tx(context.Background())
	if err != nil {
		udb.T.Error(err)

		return
	}

	for _, id := range ids {
		if err := tx.User.DeleteOneID(uuid.MustParse(id)).Exec(context.Background()); err != nil {
			udb.T.Error(err)

			_ = tx.Rollback()
		}
	}

	if err := tx.Commit(); err != nil {
		udb.T.Error(err)
	}
}
