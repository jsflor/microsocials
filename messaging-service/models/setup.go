package models

import (
	"fmt"
	"jsflor/messaging-service/lib"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var dbUser = lib.GetEnv("MESSAGING_DATABASE_USER", "postgres")
var dbPassword = lib.GetEnv("MESSAGING_DATABASE_PASSWORD", "postgres")
var dbName = lib.GetEnv("DATABASE_DB", "postgres")
var dbPort = lib.GetEnv("DATABASE_PORT", "5432")

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=messaging-database user=%s password=%s dbname=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %s", err.Error()))
	}

	// Migrate the schema
	db.AutoMigrate(&Message{})

	DB = db
}
