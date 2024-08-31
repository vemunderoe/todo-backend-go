package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Welcome to the TODO API")
	})
	todoRoutes := router.Group("/todos")
	RegisterTodoRoutes(todoRoutes)
}