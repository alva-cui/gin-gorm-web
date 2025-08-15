# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go web application built with the Gin web framework and GORM ORM for MySQL database operations. It implements a RESTful API for user management with full CRUD operations.

## Architecture

The project follows the Go standard project layout:

- `cmd/gin-gorm-web/main.go` - Application entry point
- `internal/app/` - Application setup and initialization
- `internal/config/` - Database configuration and connection
- `internal/dao/` - Data Access Objects for database operations
- `internal/service/` - Business logic and request handling
- `internal/controller/` - API endpoint definitions
- `internal/middlewares/` - Custom middleware functions for logging and error handling

## Key Components

1. **Database**: MySQL with GORM ORM
2. **Web Framework**: Gin
3. **Model**: User model with ID, Name, Email, Password, and Age fields
4. **API Endpoints**:
   - GET /users - Get all users
   - GET /users/:id - Get a specific user by ID
   - POST /users - Create a new user
   - PUT /users/:id - Update a user by ID
   - DELETE /users/:id - Delete a user by ID

## Common Development Tasks

### Building and Running
```bash
# Install dependencies
go mod tidy

# Run the application
go run cmd/gin-gorm-web/main.go
```

### Database Configuration
Update the database connection string in `internal/config/config.go` with your MySQL credentials:
```go
dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

## Dependencies
- github.com/gin-gonic/gin v1.10.1
- gorm.io/driver/mysql v1.6.0
- gorm.io/gorm v1.30.1