package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB            *gorm.DB
	DBConnections = map[string]*gorm.DB{}
)

func ConnectDatabase(app string) {
	err := godotenv.Load() // Memuat variabel lingkungan dari file .env
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	// Mendapatkan nilai variabel koneksi database
	dbHost := os.Getenv("DB_HOST_SPL")
	dbPort := os.Getenv("DB_PORT_SPL")
	dbDatabase := os.Getenv("DB_DATABASE_SPL")
	dbUsername := os.Getenv("DB_USERNAME_SPL")
	dbPassword := os.Getenv("DB_PASSWORD_SPL")

	if app == "user_man" {
		dbHost = os.Getenv("DB_HOST_USER")
		dbPort = os.Getenv("DB_PORT_USER")
		dbDatabase = os.Getenv("DB_DATABASE_USER")
		dbUsername = os.Getenv("DB_USERNAME_USER")
		dbPassword = os.Getenv("DB_PASSWORD_USER")
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUsername, dbPassword, dbDatabase, dbPort)

	// Membuka koneksi ke database PostgreSQL
	conn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("Gagal koneksi database %s: %v\n", app, err)
		return
	}

	fmt.Printf("terhubung dengan database %s\n", dbDatabase)
	DBConnections[app] = conn
}