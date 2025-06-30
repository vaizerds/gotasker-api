package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/config"
	"github.com/vaizerds/gotasker-api/models"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	config.DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}
