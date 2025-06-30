package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/SEU_USUARIO/gotasker-api/config"
    "github.com/SEU_USUARIO/gotasker-api/models"
    "net/http"
)

func GetTasks(c *gin.Context) {
    var tasks []models.Task
    config.DB.Find(&tasks)
    c.JSON(http.StatusOK, tasks)
}
