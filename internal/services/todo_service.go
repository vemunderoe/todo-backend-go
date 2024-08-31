package services

import (
	"log"

	"todo-backend/internal/models"
	"todo-backend/internal/repositories"
)

func GetTodos() ([]models.Todo, error) {
	todos, err := repositories.GetAllTodos()
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodoById(id int) (models.Todo, error) {
	todo, err := repositories.GetTodoById(id)
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