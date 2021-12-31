package database

import (
	"log"
	"todo-app-api/database/models"

	"gorm.io/gorm"
)

func RunMigrator(db *gorm.DB)  {
	err := db.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)

	if err != nil {
		panic("failed to migrate tables")
	} else {
		log.Println("tables successfully migrated")
	}
}