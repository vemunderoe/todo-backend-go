# Todo Backend

This is a simple Todo backend application written in Go.

## Prerequisites

- Go 1.16 or higher
- PostgreSQL

## Installation

1. Clone the repository:

   Using HTTPS:

   ```
   git clone https://github.com/vemunderoe/todo-backend-go.git
   cd todo-backend-go
   ```

   Or using SSH:

   ```
   git git@github.com:vemunderoe/todo-backend-go.git
   cd todo-backend-go
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Set up the PostgreSQL database:

   - Create a new database for the project
   - Update the database connection string in the configuration file (if applicable)

4. Run database migrations:

   Install golang-migrate:

   ```
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

   Create a new migration:

   ```
   migrate create -ext sql -dir db/migrations -seq create_todos_table
   ```

   Edit the newly created migration files in `db/migrations` to define your schema.

   Run migrations:

   ```
   migrate -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" -path db/migrations up
   ```

   Replace `username`, `password`, `localhost`, `5432`, and `database_name` with your actual PostgreSQL connection details.

## Running the Application

1. Start the server:

   ```
   go run cmd/server/main.go
   ```

2. The server should now be running on `http://localhost:8080` (or your configured port)

## API Endpoints

- `GET /todos`: Fetch all todos
- `GET /todos/:id`: Fetch a specific todo by ID
- `POST /todos`: Create a new todo

## Development

To run tests:
