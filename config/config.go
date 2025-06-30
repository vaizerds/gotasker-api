// package config

// import (
// 	"log"

// 	"github.com/glebarez/sqlite v1.4.7"
// 	"github.com/vaizerds/gotasker-api/models"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {
// 	database, err := gorm.Open(sqlite.Open("gotasker.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	DB = database

// 	err = DB.AutoMigrate(&models.User{}, &models.Task{})
// 	if err != nil {
// 		log.Fatalf("AutoMigrate failed: %v", err)
// 	}
// }

// package config

// import (
// 	"log"

// 	"github.com/glebarez/sqlite"
// 	"github.com/vaizerds/gotasker-api/models"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {
// 	database, err := gorm.Open(sqlite.Open("gotasker.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	DB = database

// 	if err := DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
// 		log.Fatalf("AutoMigrate failed: %v", err)
// 	}
// }

package config

import (
	"fmt"
	"log"
	"os"

	"github.com/vaizerds/gotasker-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Exemplo de conexão (ajuste conforme seu ambiente)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "gotasker_db"),
		getEnv("DB_PORT", "5432"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database

	// AutoMigrate models
	err = DB.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

// Utilitário para ler variáveis de ambiente com fallback
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
