package routes

import (
	"todo-backend/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.RouterGroup) {
	router.GET("/", controllers.GetTodos)
	router.GET("/:id", controllers.GetTodoById)
	router.POST("/", controllers.CreateTodo)
}