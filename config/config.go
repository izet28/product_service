package config

import (
	"fmt"
	"log"
	"os"
	"product_service/models"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB menginisialisasi koneksi ke database dengan konfigurasi pool
func InitDB() *gorm.DB {
	// Muat variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Ambil nilai variabel lingkungan
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Buat DSN (Data Source Name) untuk MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Konfigurasi GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log query SQL
	})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Ambil konfigurasi pool koneksi dari variabel lingkungan
	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		maxIdleConns = 10 // Default value
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = 100 // Default value
	}

	connMaxLifetime, err := time.ParseDuration(os.Getenv("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		connMaxLifetime = 30 * time.Minute // Default value
	}

	// Mendapatkan objek sql.DB dari gorm.DB untuk mengatur pool koneksi
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Gagal mendapatkan objek sql.DB dari gorm.DB: %v", err)
	}

	// Atur konfigurasi pool koneksi
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	// Automigrate untuk membuat tabel jika belum ada
	db.AutoMigrate(&models.Product{})
	return db
}
