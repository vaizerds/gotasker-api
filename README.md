# GoTasker API

A RESTful API for task management, built in Go using Gin, GORM, and JWT authentication.

## Features
- User registration and login
- JWT-based authentication (em breve)
- CRUD for tasks
- PostgreSQL support
- Docker-ready

## Running locally

1. **Start PostgreSQL (Docker):**

```bash
docker run --name gotasker-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=gotasker_db -p 5432:5432 -d postgres
```

2. **Run the API:**

```bash
go run cmd/main.go
```

## Docker

```bash
docker build -t gotasker-api .
docker run --env-file .env -p 8080:8080 gotasker-api
```
