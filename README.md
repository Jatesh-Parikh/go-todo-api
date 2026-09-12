# Go Todo REST API

A secure, full-stack REST API built with Go, Gin, PostgreSQL, and JWT authentication. This project demonstrates the implementation of a complete todo management system with user authentication, password hashing, and protected routes.

> This project is based on a tutorial by [Muslim Halelee](https://www.youtube.com/watch?v=S069igHKUIw&list=PLHsjm_W8kcWZLPDxUplr8yredk95F-KIx&index=50).

## Features

- **User Authentication**: Secure registration and login with JWT tokens
- **Password Security**: bcrypt hashing for secure password storage
- **Protected Routes**: Middleware-based route protection
- **User-Specific Todos**: Each user has their own private todo collection
- **CRUD Operations**: Create, read, update, and delete todos
- **Database Migrations**: Version-controlled database schema changes
- **Hot Reloading**: Air integration for development

## Project Structure

```
Go-Gin-Postgres-Todo-REST-API/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Environment configuration
│   ├── database/
│   │   └── postgres.go          # Database connection
│   ├── handlers/
│   │   ├── todo_handler.go      # Todo route handlers
│   │   └── user_handler.go      # Auth route handlers
│   ├── middleware/
│   │   └── auth_middleware.go   # JWT authentication middleware
│   ├── models/
│   │   ├── todo.go              # Todo model
│   │   └── user.go              # User model
│   └── repository/
│       ├── todo_repository.go   # Todo database operations
│       └── user_repository.go   # User database operations
├── migrations/
│   ├── 000001_create_todos_api_table.up.sql
│   ├── 000001_create_todos_api_table.down.sql
│   ├── 000002_create_users_api_table.up.sql
│   ├── 000002_create_users_api_table.down.sql
│   ├── 000003_add_user_id_to_todos_table.up.sql
│   └── 000003_add_user_id_to_todos_table.down.sql
├── scripts/
│   └── migrate.ps1              # Migration helper script
├── .air.toml                    # Air configuration
├── .env                         # Environment variables (create this)
├── go.mod                       # Go module definition
└── go.sum                       # Go dependencies checksum
```

## Technologies Used

- **Go 1.21+**: Backend programming language
- **Gin**: HTTP web framework
- **PostgreSQL**: Relational database
- **pgx/v5**: PostgreSQL driver and connection pool
- **JWT**: JSON Web Tokens for authentication
- **bcrypt**: Password hashing
- **golang-migrate**: Database migrations
- **Air**: Hot reloading for development
- **godotenv**: Environment variable management