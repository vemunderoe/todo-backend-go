package services

import (
	"log"

	"todo-backend/internal/models"
	"todo-backend/internal/repositories"
)

func GetTodos(user models.User) ([]models.Todo, error) {
	todos, err := repositories.GetAllTodosByUserId(user.ID)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodoById(id int, user models.User) (models.Todo, error) {
	todo, err := repositories.GetTodoByUserIdAndId(id, user.ID)
	if err != nil {
		return todo, err
	}

	return todo, nil
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	todo, err := repositories.CreateTodo(todo)
	if err != nil {
		log.Printf("Error creating todo: %v\n", err)
		return todo, err
	}
	return todo, nil
}