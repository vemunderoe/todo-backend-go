package routes

import (
	"todo-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}