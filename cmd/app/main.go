package main

import (
	"github.com/arkad0912/user-service/internal/database"

	"github.com/arkad0912/user-service/internal/userService"

	"github.com/arkad0912/user-service/internal/transport/grpc"
)

func main() {
	// Инициализация базы данных
	database.InitDB()

	// Репозитории и сервисы для пользователей
	userRepo := userService.NewUserRepository(database.DB)
	userService := userService.NewUserService(userRepo)
	userHandler := grpc.NewUserHandlers(userService)

	grpc.RunServer(userHandler, ":50051")
}
