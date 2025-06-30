package routes

import (
	"github.com/vaizerds/gotasker-api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.GET("/tasks", controllers.GetTasks)
}
