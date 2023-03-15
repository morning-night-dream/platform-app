package main

import (
	"context"
	"fmt"
	"os"

	"github.com/morning-night-dream/platform-app/internal/driver/database"
	"github.com/morning-night-dream/platform-app/pkg/ent"
	"github.com/morning-night-dream/platform-app/pkg/log"
)

func main() {
	pDSN := os.Getenv("PRIMARY_DATABASE_URL")

	sDSN := os.Getenv("SECONDARY_DATABASE_URL")

	pClient := database.NewClient(pDSN)

	defer pClient.Close()

	sClient := database.NewClient(sDSN)

	defer sClient.Close()

	ctx := context.Background()

	// secondaryに対してマイグレーションを実行
	if err := sClient.Debug().Schema.Create(ctx); err != nil {
		log.Log().Panic(fmt.Sprintf("failed to create schema: %v", err))
	}

	// primaryからデータを取得
	tags, err := pClient.ArticleTag.Query().All(ctx)
	if err != nil {
		log.Log().Panic(fmt.Sprintf("failed to query tags: %v", err))
	}

	articles, err := pClient.Article.Query().All(ctx)
	if err != nil {
		log.Log().Panic(fmt.Sprintf("failed to query articles: %v", err))
	}

	// トランザクション開始

	log.Log().Info("start transaction")

	tx, err := sClient.BeginTx(ctx, nil)
	if err != nil {
		log.Log().Panic(fmt.Sprintf("failed to begin transaction: %v", err))
	}

	// めんどくさいのでsecondaryデータぶっぱ
	if _, err := tx.ArticleTag.Delete().Exec(ctx); err != nil {
		tx.Rollback()

		log.Log().Panic(fmt.Sprintf("failed to delete tags: %v", err))
	}

	if _, err := tx.Article.Delete().Exec(ctx); err != nil {
		tx.Rollback()

		log.Log().Panic(fmt.Sprintf("failed to delete articles: %v", err))
	}

	// secondaryにデータを投入

	articleBulk := make([]*ent.ArticleCreate, len(articles))
	for i, article := range articles {
		articleBulk[i] = tx.Article.Create().
			SetID(article.ID).
			SetTitle(article.Title).
			SetURL(article.URL).
			SetDescription(article.Description).
			SetThumbnail(article.Thumbnail).
			SetCreatedAt(article.CreatedAt).
			SetUpdatedAt(article.UpdatedAt)
	}

	if _, err := tx.Article.CreateBulk(articleBulk...).Save(ctx); err != nil {
		tx.Rollback()

		log.Log().Panic(fmt.Sprintf("failed to insert articles: %v", err))
	}

	tagBulk := make([]*ent.ArticleTagCreate, len(tags))
	for i, tag := range tags {
		tagBulk[i] = tx.ArticleTag.Create().
			SetArticleID(tag.ID).
			SetTag(tag.Tag).
			SetCreatedAt(tag.CreatedAt).
			SetUpdatedAt(tag.UpdatedAt).
			SetArticleID(tag.ArticleID)
	}

	if _, err := tx.ArticleTag.CreateBulk(tagBulk...).Save(ctx); err != nil {
		tx.Rollback()

		log.Log().Panic(fmt.Sprintf("failed to insert articles: %v", err))
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()

		log.Log().Panic(fmt.Sprintf("failed to commit transaction: %v", err))
	}

	log.Log().Info("commit transaction")
}
