package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func SetupDB() {
	dbUsername := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DBNAME")

	if dbPort == "" {
		dbPort = "8090"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
						dbUsername, dbPassword, dbHost, dbPort, dbName)
	
	conn, e := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if e != nil {
		panic("failed to connect database")
	} else {
		log.Println("connected to database")
	}
	db = conn

	RunMigrator(db)
}

func GetDB() *gorm.DB {
	return db
}