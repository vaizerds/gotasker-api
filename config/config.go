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

package config

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/vaizerds/gotasker-api/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("gotasker.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	DB = database

	if err := DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
}
