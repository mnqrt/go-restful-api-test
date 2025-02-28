package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// NewDB initializes the database connection using GORM
func NewDB() *gorm.DB {
	// Correct DSN format for GORM
	dsn := "host=localhost user=postgres password=12345 dbname=belajargo port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	// Use the correct PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logging SQL queries
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set database connection pool settings
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	log.Println("Database connected successfully!")
	return db
}
