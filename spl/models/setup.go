package models

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	dbInitialized bool
	dbMutex       sync.Mutex
	DBConnections = map[string]*gorm.DB{}
)

func ConnectDatabase() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	if dbInitialized {
		return
	}

	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get database connection variables
	dbHost := os.Getenv("DB_HOST_SPL")
	dbPort := os.Getenv("DB_PORT_SPL")
	dbDatabase := os.Getenv("DB_DATABASE_SPL")
	dbUsername := os.Getenv("DB_USERNAME_SPL")
	dbPassword := os.Getenv("DB_PASSWORD_SPL")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsername, dbPassword, dbDatabase, dbPort)

	// Open connection to PostgreSQL database
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return
	}

	fmt.Printf("Connected to the %s database\n", dbDatabase)
	DB = conn
	DBConnections["spl"] = conn
	dbInitialized = true
}