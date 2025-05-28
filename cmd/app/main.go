package main

import (
	"log"
	"user-service/internal/database"
	"user-service/internal/handlers"
	"user-service/internal/userService"
	"user-service/internal/web/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Репозитории и сервисы для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := handlers.NewUserHandlers(userService)

	// Инициализация Echo
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Регистрация хендлеров для пользователей
	usersStrictHandler := users.NewStrictHandler(userHandler, nil)
	users.RegisterHandlers(e, usersStrictHandler)

	// Запуск сервера
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start with err: %v", err)
	}
}
