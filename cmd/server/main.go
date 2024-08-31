package main

import (
	"fmt"
	"todo-backend/internal/db"
	"todo-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init database connection
	db.Init()
	defer db.Close()

	// Set the port
	const port int = 8080

	// Setup gin with default logging and recovery middleware
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Run the router and panic if error
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
