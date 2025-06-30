package main

import (
	"github.com/SEU_USUARIO/gotasker-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/config"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}
