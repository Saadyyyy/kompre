package config

import (
	"kompre/models"
	"log"

	"gorm.io/gorm"
)

func DBMigration(db *gorm.DB) {
	// Migrate Category
	err := db.AutoMigrate(&models.KinerjaCrud{})
	if err != nil {
		log.Fatalf("Failed to migrate Category: %v", err)
	}
}
