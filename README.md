# 🚀 GoTasker API

A RESTful API for task management, built with Go using Gin, GORM, and JWT authentication.

## ✨ Features
- 👤 User registration
- ✅ Task creation and listing
- 🗄️ SQLite or PostgreSQL support
- 🐳 Docker-ready

## 🧰 Tech Stack
- 💻 [Go](https://golang.org/)
- ⚙️ [Gin Gonic](https://github.com/gin-gonic/gin)
- 🧮 [GORM](https://gorm.io/)
- 🔐 [JWT (planned)](https://github.com/golang-jwt/jwt)
- 🐳 [Docker](https://www.docker.com/)

## 🚀 Getting Started

### 📋 Prerequisites
- ✅ Go >= 1.22
- 🗃️ SQLite (default) or PostgreSQL (optional)

### 📦 Installation
```bash
git clone https://github.com/seu-usuario/gotasker-api.git
cd gotasker-api
go mod download
go run cmd/main.go
```

### 🐳 Run with Docker
```bash
docker build -t gotasker-api .
docker run -p 8080:8080 gotasker-api
```

## 📡 API Endpoints

### 🔐 `POST /register`
Create a new user.

```json
{
  "username": "johndoe",
  "password": "securepassword"
}
```

### 📋 `GET /tasks`
List all tasks (no auth required yet).

## 📁 Folder Structure
```
cmd/          → Main entry point
config/       → Database connection setup
controllers/  → Business logic (user, task)
models/       → GORM models
routes/       → Route handlers
```

## 📄 License
MIT
