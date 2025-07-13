# 🚀 GoTasker API

A RESTful API for task management, built in **Go** using **Gin** and **GORM**.

---

## ✨ Features

- 👤 User registration and login  
- 🔐 JWT-based authentication  
- 📝 Full CRUD operations for tasks (protected)  
- 🐘 Support for **PostgreSQL** (configurable)  
- 🐳 Ready to run with Docker  

---

## ⚙️ Tech Stack

- [Go](https://golang.org/)  
- [Gin](https://github.com/gin-gonic/gin) (Web Framework)  
- [GORM](https://gorm.io/) (ORM para banco de dados)  
- [PostgreSQL](https://www.postgresql.org/) (Banco de dados relacional)  
- [Docker](https://www.docker.com/) (Containerização)  

---

## 🛠️ How to Run Locally

```bash
git clone https://github.com/vaizerds/gotasker-api.git
cd gotasker-api



# Configure your environment variables for PostgreSQL:
# DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT

go run cmd/main.go
