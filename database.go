package database

import (
	"context"
	"log"

	"github.com/lvdlee/fertilizer-store-data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrate(ctx context.Context, db *gorm.DB) {
	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Fertilizer{})
}

func Setup(ctx context.Context) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("fertilizer.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed opening db: %s", err)
	}

	migrate(ctx, db)

	return db
}
