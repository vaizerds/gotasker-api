package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/SEU_USUARIO/gotasker-api/config"
    "github.com/SEU_USUARIO/gotasker-api/models"
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
