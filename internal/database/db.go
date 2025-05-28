package database

import (
	"log"

	"github.com/arkad0912/user-service/internal/userService" // Импортируйте ваш пакет userService
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=main port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Автомиграция создаст таблицу со всеми необходимыми полями
	if err := DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatal("Failed to auto-migrate models:", err)
	}

	log.Println("Database migration completed successfully")
}
