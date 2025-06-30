package main

import (
    "github.com/SEU_USUARIO/gotasker-api/config"
    "github.com/SEU_USUARIO/gotasker-api/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()
    r := gin.Default()
    routes.RegisterRoutes(r)
    r.Run(":8080")
}
