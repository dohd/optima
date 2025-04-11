package config

import (
	"log"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDatabase() {
    dbPath := filepath.Join(GetProjectRoot(), "optima-app", "app.db")

    DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    log.Println("database connected successfully")
}
