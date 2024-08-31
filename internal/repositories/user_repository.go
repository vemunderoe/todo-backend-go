package repositories

import (
	"context"
	"log"
	"todo-backend/internal/db"
	"todo-backend/internal/models"
)


func CreateUser(user models.User) (models.User, error) {
	var newUser models.User

	err := db.DB.QueryRow(
		context.Background(),
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username, password",
		user.Username, user.Password,
	).Scan(&newUser.ID, &newUser.Username, &newUser.Password)

	if err != nil {
		log.Printf("Failed to insert user: %v\n", err)
		return newUser, err
	}

	return newUser, nil
}

func CheckUsernameExists(username string) (bool, error) {
	var exists bool
	err := db.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", username).Scan(&exists)
	if err != nil {
		log.Printf("Failed to check if username exists: %v\n", err)
		return false, err
	}
	return exists, nil
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User

	row := db.DB.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", username)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Printf("Failed to fetch user by username %s: %v\n", username, err)
		return models.User{}, err
	}
	return user, nil
}