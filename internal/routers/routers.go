package routers

import (
	"booking-apartment/internal/controllers"
	"booking-apartment/internal/services"

	"github.com/gin-gonic/gin"
)

var userService services.UserService

func SetupRoutes(router *gin.Engine) {
	router.GET("/health", services.NewHealthHandler().HealthCheck)

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.NewUserController().CreateUser)
		userGroup.GET("/", controllers.NewUserController().GetUsers)
		userGroup.GET("/:id", controllers.NewUserController().GetUserByID)
		userGroup.PUT("/:id", controllers.NewUserController().UpdateUser)
		userGroup.DELETE("/:id", controllers.NewUserController().DeleteUser)
	}
}
