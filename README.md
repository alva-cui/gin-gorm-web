# Go Web Framework with Gin and GORM

This is a simple web framework built with Gin and GORM for MySQL database operations.

## Project Structure

```
gin-gorm-web/
├── cmd/
│   └── gin-gorm-web/
│       └── main.go
├── internal/
│   ├── app/
│   ├── config/
│   │   └── config.go
│   ├── controller/
│   │   └── routes.go
│   ├── dao/
│   │   └── user.go
│   ├── middlewares/
│   │   └── middlewares.go
│   ├── service/
│   │   └── user.go
│   └── go.mod
└── docs/
    ├── docs.go
    ├── swagger.json
    └── swagger.yaml
```

## Features

1. RESTful API design
2. MySQL database integration with GORM
3. CRUD operations for User model
4. Middleware for logging and error handling
5. Automatic database migration
6. API documentation with Swagger

## Setup Instructions

1. Make sure you have Go installed (version 1.16 or later)
2. Install MySQL database
3. Update the database connection string in `internal/config/config.go`
4. Run the following commands:

```bash
# Install dependencies
go mod tidy

# Generate Swagger documentation
swag init

# Run the application
go run cmd/gin-gorm-web/main.go
```

## API Endpoints

- `GET /users` - Get all users
- `GET /users/:id` - Get a specific user by ID
- `POST /users` - Create a new user
- `PUT /users/:id` - Update a user by ID
- `DELETE /users/:id` - Delete a user by ID

## API Documentation

After running the application, you can access the Swagger UI documentation at:
- `http://localhost:8080/swagger/index.html`

## Database Configuration

Update the database connection string in `internal/config/config.go`:

```go
dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

Replace `user`, `password`, `127.0.0.1`, `3306`, and `dbname` with your actual database credentials and details.