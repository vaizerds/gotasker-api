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

// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/vaizerds/gotasker-api/config"
// 	"github.com/vaizerds/gotasker-api/routes"
// )

//	func main() {
//		config.ConnectDB()
//		r := gin.Default()
//		routes.RegisterRoutes(r)
//		r.Run(":8080")
//	}
// package main

// import (
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/vaizerds/gotasker-api/config"
// 	"github.com/vaizerds/gotasker-api/routes"
// )

// func main() {
// 	// Conecta ao banco de dados
// 	config.ConnectDB()

// 	// Cria o router Gin
// 	r := gin.Default()

// 	// Registra as rotas do seu app
// 	routes.RegisterRoutes(r)

// 	// Inicia o servidor HTTP na porta 8080
// 	if err := r.Run(":8080"); err != nil {
// 		log.Fatalf("Failed to run server: %v", err)
// 	}
// }
