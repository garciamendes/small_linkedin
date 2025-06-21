package infra

import (
	"log"

	"github.com/garciamendes/small_linkedin/src/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.Auth{},
	)

	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
}
