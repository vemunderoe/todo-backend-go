package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func Init() {
	databaseURL := "postgres://postgres:12345678@localhost:5432/todo"
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	var err error

	// Set connection configuration options
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse DATABASE_URL: %v\n", err)
	}

	// Establish a connection to the database
	DB, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = DB.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping the database: %v\n")
	}

	log.Println("Successfully connedted to database")
}

func Close() {
	DB.Close()
}