package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "github.com/SEU_USUARIO/gotasker-api/models"
)

var DB *gorm.DB

func ConnectDB() {
    database, err := gorm.Open(sqlite.Open("gotasker.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database")
    }
    DB = database
    DB.AutoMigrate(&models.User{}, &models.Task{})
}
