package main

import (
	"github.com/gin-gonic/gin"

	config "booking-apartment/configs"
	"booking-apartment/internal/routers"
)

func main() {
	config.ConnectDatabase()

	router := gin.Default()
	routers.SetupRoutes(router)

	router.Run(":8080")
}
