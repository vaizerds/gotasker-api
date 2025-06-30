package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	// Rota raiz para teste se o servidor está ativo
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "API está rodando"})
	})

	r.POST("/register", controllers.Register)
	r.GET("/tasks", controllers.GetTasks)
}
