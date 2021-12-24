package database

import (
	"todo-app-api/database/models"

	"gorm.io/gorm"
)

func RunMigrator(db *gorm.DB)  {
	db.AutoMigrate(
		&models.Activity{},
		&models.Todo{},
	)
}