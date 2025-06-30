package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.GET("/tasks", controllers.GetTasks)
}
