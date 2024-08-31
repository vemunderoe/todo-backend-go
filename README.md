# Todo Backend

This is a simple Todo backend application written in Go, featuring user authentication and per-user todo management.

## Prerequisites

- Go 1.16 or higher
- PostgreSQL

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/vemunderoe/todo-backend-go.git
   cd todo-backend-go
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Set up the PostgreSQL database:

   - Create a new database for the project
   - Update the database connection string in the `.env` file (create one if it doesn't exist)

4. Run database migrations:

   Install golang-migrate:

   ```
   go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

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

- `POST /auth/register`: Register a new user
- `POST /auth/login`: Login and receive a JWT token
- `GET /todos`: Fetch all todos for the authenticated user
- `GET /todos/:id`: Fetch a specific todo by ID for the authenticated user
- `POST /todos`: Create a new todo for the authenticated user

All endpoints except `/auth/register` and `/auth/login` require authentication. Include the JWT token in the Authorization header:
