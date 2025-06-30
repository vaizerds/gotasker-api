# GoTasker API

A RESTful API for task management, built in Go using Gin, GORM, and JWT authentication.

## Features
- User registration and login
- JWT-based authentication
- CRUD for tasks (protected)
- SQLite or PostgreSQL support
- Docker-ready

## Running locally

```bash
git clone https://github.com/SEU_USUARIO/gotasker-api.git
cd gotasker-api
go run cmd/main.go
```

## Docker

```bash
docker build -t gotasker-api .
docker run -p 8080:8080 gotasker-api
```
