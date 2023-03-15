package main

import (
	"context"
	"log"
	"os"

	"github.com/morning-night-dream/platform-app/internal/driver/database"
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
		log.Fatalf("Failed create schema: %v", err)
	}

	// primaryからデータを取得

	// secondaryにデータを投入
}
