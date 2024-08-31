package repositories

import (
	"context"
	"errors"
	"log"
	"todo-backend/internal/db"
	"todo-backend/internal/models"

	"github.com/jackc/pgx/v4"
)


func GetAllTodosByUserId(userId int) ([]models.Todo, error) {
	var todos []models.Todo

	rows, err := db.DB.Query(context.Background(), "SELECT id, title, completed, owner FROM todos WHERE owner = $1", userId)
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.Owner)
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

func GetTodoByUserIdAndId(id, userId int) (models.Todo, error) {
	var todo models.Todo

	row := db.DB.QueryRow(context.Background(), "SELECT id, title, completed, owner FROM todos WHERE id = $1 and owner = $2", id, userId)
	
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.Owner)
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
		"INSERT INTO todos (title, completed, owner) VALUES ($1, $2, $3) RETURNING id, title, completed, owner",
		todo.Title, todo.Completed, todo.Owner,
	).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Completed, &newTodo.Owner)

	if err != nil {
		log.Printf("Failed to insert todo: %v\n", err)
		return newTodo, err
	}

	return newTodo, nil
}