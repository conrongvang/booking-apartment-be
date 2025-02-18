package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"booking-apartment/models"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host, user, password, dbname, port, timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌ Connect database failed: %v", err)
	}

	DB = db

	log.Println("✅ Connect database success!")

	var version string
	db.Raw("SELECT version()").Scan(&version)
	fmt.Println("📌 PostgreSQL Version:", version)

	err = db.AutoMigrate(&models.User{}, &models.Apartment{}, &models.Customer{}, &models.Booking{})

	if err != nil {
		log.Fatalf("❌ Migrate erorr: %v", err)
	}
}
