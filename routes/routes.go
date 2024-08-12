package routes

import (
	"github.com/gin-gonic/gin"

	"email/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser)
	router.GET("user/:id", controllers.GetUserByID)
	router.PUT("user/:id", controllers.UpdateUser)
	router.DELETE("user/:id", controllers.DeleteUser)
	router.GET("user", controllers.GetAllUser)
}
