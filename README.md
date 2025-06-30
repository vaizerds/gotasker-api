# 🚀 GoTasker API

Uma API RESTful para gerenciamento de tarefas, construída em **Go** usando **Gin**, **GORM** e autenticação com **JWT**.

---

## ✨ Funcionalidades

- 👤 Registro e login de usuários  
- 🔐 Autenticação baseada em JWT  
- 📝 CRUD completo para tarefas (protegido)  
- 🐘 Suporte a **PostgreSQL** (configurável)  
- 🐳 Pronto para rodar com Docker  

---

## ⚙️ Tecnologias

- [Go](https://golang.org/)  
- [Gin](https://github.com/gin-gonic/gin) (Web Framework)  
- [GORM](https://gorm.io/) (ORM para banco de dados)  
- [PostgreSQL](https://www.postgresql.org/) (Banco de dados relacional)  
- [Docker](https://www.docker.com/) (Containerização)  

---

## 🛠️ Como rodar localmente

```bash
git clone https://github.com/vaizerds/gotasker-api.git
cd gotasker-api

# Configure as variáveis de ambiente para o PostgreSQL:
# DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT

go run cmd/main.go
