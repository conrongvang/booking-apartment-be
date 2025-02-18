package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	config "booking-apartment/configs"
	"booking-apartment/internal/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Error loading .env file: %v", err)
	}

	port := os.Getenv("GATEWAY_PORT")

	config.ConnectDatabase()

	router := gin.Default()
	routers.SetupRoutes(router)
	router.Run(":" + port)
	fmt.Print("port:::", port)
}
