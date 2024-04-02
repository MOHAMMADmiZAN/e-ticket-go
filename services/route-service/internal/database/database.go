package db

import (
	_ "fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type Database struct {
	Conn *gorm.DB
}

// NewDatabase creates a new database connection.
func NewDatabase(models ...interface{}) *Database {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not get database connection: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Perform auto migration for given models
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	return &Database{Conn: db}
}

// Close method for clean database disconnection
func (database *Database) Close() {
	sqlDB, err := database.Conn.DB()
	if err != nil {
		log.Printf("Error getting DB from GORM: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}
}
