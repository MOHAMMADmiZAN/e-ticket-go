package config

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase(models ...interface{}) *Database {
	dsn := BuildDSN()
	if !databaseExists(dsn) {
		createDatabase(dsn)
	}

	// Initialize the GORM logger
	//gormLogger := initializeLogger()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//Logger: gormLogger,
	})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not get database connection: %v", err)
	}
	setupConnectionPool(sqlDB)

	// AutoMigrate moved to a method on the Database struct
	err = db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Could not auto migrate models: %v", err)
	}

	return &Database{Conn: db}
}

func setupConnectionPool(sqlDB *sql.DB) {
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func databaseExists(dsn string) bool {
	tempDB, err := sql.Open("postgres", dsn)
	if err != nil {
		return false
	}
	defer func(tempDB *sql.DB) {
		err := tempDB.Close()
		if err != nil {
			log.Fatalf("Could not close temporary database connection: %v", err)
		}
	}(tempDB)
	return tempDB.Ping() == nil
}

func createDatabase(dsn string) {
	dbName := os.Getenv("DB_NAME")
	systemDSN := dsnForSystemDB(dsn)
	systemDB, err := sql.Open("postgres", systemDSN)
	if err != nil {
		log.Fatalf("Could not connect to system database: %v", err)
	}
	defer func(systemDB *sql.DB) {
		err := systemDB.Close()
		if err != nil {
			log.Fatalf("Could not close system database connection: %v", err)
		}
	}(systemDB)
	// Check if the database exists using a proper SQL query for PostgreSQL
	var exists bool
	err = systemDB.QueryRow("SELECT EXISTS(SELECT FROM pg_database WHERE datname = $1)", dbName).Scan(&exists)
	if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	}
	// Only create the database if it does not exist
	if !exists {
		// Proper SQL execution with error handling
		_, err = systemDB.Exec(fmt.Sprintf("CREATE DATABASE %s", pq.QuoteIdentifier(dbName)))
		if err != nil {
			log.Fatalf("Could not create database %s: %v", dbName, err)
		}
		log.Printf("Database %s created successfully", dbName)
	} else {
		log.Printf("Database %s already exists, no need to create.", dbName)
	}
}

func dsnForSystemDB(dsn string) string {
	// Remove the dbname and add the default one, 'postgres'
	return fmt.Sprintf("%s dbname=postgres", strings.TrimSuffix(dsn, os.Getenv("DB_NAME")))
}

// BuildDSN constructs the Data Source Name from environment variables.
func BuildDSN() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
}

// Close method for clean database disconnection.
func (database *Database) Close() {
	sqlDB, err := database.Conn.DB()
	if err != nil {
		log.Fatalf("Error getting DB from GORM: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}

func initializeLogger() logger.Interface {
	// Fetch log level from environment
	logLevelEnv := os.Getenv("GORM_LOG_LEVEL")
	var logLevel logger.LogLevel
	switch logLevelEnv {
	case "SILENT":
		logLevel = logger.Silent
	case "ERROR":
		logLevel = logger.Error
	case "WARN":
		logLevel = logger.Warn
	case "INFO":
		logLevel = logger.Info
	default:
		logLevel = logger.Warn // Default to Warn if not specified
	}

	// Setup the GORM logger
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Or use a more sophisticated logger
		logger.Config{
			SlowThreshold: time.Second, // Adjust based on your requirements
			LogLevel:      logLevel,
			Colorful:      true, // Set to false for production environments
		},
	)

	return gormLogger
}
