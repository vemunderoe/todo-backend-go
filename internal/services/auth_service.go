package services

import (
	"errors"
	"log"

	"todo-backend/internal/models"
	"todo-backend/internal/repositories"
	"todo-backend/internal/utils"
)

var (
	ErrUsernameTaken      = errors.New("username is already taken")
	ErrInvalidCredentials = errors.New("invalid username or password")
)

func RegisterAndLoginUser(user models.User) (string, error) {

	usernameAlreadyExists, err := repositories.CheckUsernameExists(user.Username)
	if err != nil {
		return "", err
	}
	if usernameAlreadyExists {
		return "", ErrUsernameTaken
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	newUser, err := repositories.CreateUser(user)
	if err != nil {
		log.Printf("Error registering user: %v\n", err)
		return "", err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(newUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func LoginUser(username, password string) (string, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
