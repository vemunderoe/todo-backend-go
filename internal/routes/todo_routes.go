package routes

import (
	"todo-backend/internal/controllers"
	"todo-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.RouterGroup) {
	router.Use(middleware.JWTAuthMiddleware())
	router.GET("/", controllers.GetTodos)
	router.GET("/:id", controllers.GetTodoById)
	router.POST("/", controllers.CreateTodo)
}