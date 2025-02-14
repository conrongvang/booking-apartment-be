package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"booking-apartment/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=127.0.0.1 user=developer password=123456 dbname=booking-apartment port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
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
