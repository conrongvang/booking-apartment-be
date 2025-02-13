package routers

import (
	"booking-apartment/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", controllers.CreateUser)      // Tạo user
		userGroup.GET("/", controllers.GetUsers)         // Lấy tất cả users
		userGroup.GET("/:id", controllers.GetUserByID)   // Lấy user theo ID
		userGroup.PUT("/:id", controllers.UpdateUser)    // Cập nhật user
		userGroup.DELETE("/:id", controllers.DeleteUser) // Xóa user
	}
}
