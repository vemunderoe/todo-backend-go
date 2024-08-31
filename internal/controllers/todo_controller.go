package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"todo-backend/internal/models"
	"todo-backend/internal/repositories"
	"todo-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := services.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func GetTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")); 
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	todo, err := services.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := services.CreateTodo(input)
	if err != nil {
		if errors.Is(err, repositories.ErrTodoNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}