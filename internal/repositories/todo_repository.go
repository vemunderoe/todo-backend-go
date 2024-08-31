package repositories

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/db"
	"todo-backend/internal/models"

	"github.com/jackc/pgx/v4"
)


func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo

	rows, err := db.DB.Query(context.Background(), "SELECT id, title, completed FROM todos")
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			return nil, err
		}
		todos = append(todos, todo)
	}

	if todos == nil {
		return []models.Todo{}, nil
	}

	return todos, nil
}

var ErrTodoNotFound = errors.New("todo not found")

func GetTodoById(id int) (models.Todo, error) {
	var todo models.Todo

	row := db.DB.QueryRow(context.Background(), "SELECT id, title, completed FROM todos WHERE id = $1", id)
	
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("No todo found with ID %d\n", id)
			return todo, ErrTodoNotFound
		}
		log.Printf("Failed to fetch todo by id %d: %v\n", id, err)
		return todo, err
	}
	return todo, nil
}

func CreateTodo(todo models.Todo) (models.Todo, error) {
	var newTodo models.Todo

	err := db.DB.QueryRow(
		context.Background(),
		"INSERT INTO todos (title, completed) VALUES ($1, $2) RETURNING id, title, completed",
		todo.Title, todo.Completed,
	).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Completed)

	if err != nil {
		log.Printf("Failed to insert todo: %v\n", err)
		return newTodo, err
	}

	return newTodo, nil
}