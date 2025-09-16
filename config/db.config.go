package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	
	// Construct DSN from environment variables
	dbHost := GetEnv("DB_HOST")
	dbPort := GetEnv("DB_PORT")
	dbUser := GetEnv("DB_USER")
	dbPassword := GetEnv("DB_PASSWORD")
	dbName := GetEnv("DB_NAME")
	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	
	// Custom gorm logger
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              60 * time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Silent | logger.Error | logger.Warn, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,           // Include params in the SQL log
			Colorful:                  true,        
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})


	sqlDB, err := DB.DB()
	if err != nil {
		log.Panic("failed to get database instance")
	}
	
	// Set connection pool settings
	sqlDB.SetMaxIdleConns(5)                  // Set the maximum number of idle connections
	sqlDB.SetMaxOpenConns(30)                 // Set the maximum number of open connections
	sqlDB.SetConnMaxLifetime(30 * time.Minute)        // Set the maximum lifetime of a connection
	sqlDB.SetConnMaxIdleTime(5 * time.Minute) // Set the maximum idle time of a connection

	if err != nil {
		log.Panic("failed to connect database")
	}

	log.Println("\033[32mSuccessfully connected to the database!\033[0m") // Green color ansi code.
}
