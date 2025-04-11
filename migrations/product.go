package main

import (
	"log"
	"optima-app/config"
	"optima-app/internal/models"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    // database connection
    dbPath := filepath.Join(config.GetProjectRoot(), "optima-app", "app.db")
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Auto-migrate Product model
    err2 := db.AutoMigrate(&models.Product{})
    if err2 != nil {
        log.Fatal("migration failed:", err)
    }
    log.Println("Products table migrated successfully.")
}
