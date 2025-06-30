

directory: models/
file: user.go

package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Password string
	Tasks    []Task
}

file: task.go

package models

type Task struct {
	ID       uint   `gorm:"primaryKey"`
	Title    string
	Done     bool
	UserID   uint
}

directory: controllers/
file: userController.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/config"
	"github.com/vaizerds/gotasker-api/models"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

directory: controllers/
file: taskController.go

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/config"
	"github.com/vaizerds/gotasker-api/models"
	"net/http"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

directory: routes/
file: routes.go

package routes

import (
	"github.com/vaizerds/gotasker-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.GET("/tasks", controllers.GetTasks)
}

directory: Dockerfile

FROM golang:1.22

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./cmd/main.go

CMD ["./main"]

EXPOSE 8080

directory: README.md

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
git clone https://github.com/vaizerds/gotasker-api.git
cd gotasker-api
go run cmd/main.go
```

## Docker

```bash
docker build -t gotasker-api .
docker run -p 8080:8080 gotasker-api
