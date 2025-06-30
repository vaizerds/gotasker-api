package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vaizerds/gotasker-api/config"
	"github.com/vaizerds/gotasker-api/routes"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Run(":8080")
}
