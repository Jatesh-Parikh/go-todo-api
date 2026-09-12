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