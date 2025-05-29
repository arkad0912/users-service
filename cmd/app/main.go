package main

import (
	"github.com/arkad0912/user-service/internal/database"

	"github.com/arkad0912/user-service/internal/userService"

	"github.com/arkad0912/user-service/internal/transport/grpc"
)

func main() {
	database.InitDB()                                      // 1. Подключение к БД
	userRepo := userService.NewUserRepository(database.DB) // 2. Репозиторий
	userService := userService.NewUserService(userRepo)    // 3. Сервис
	userHandler := grpc.NewUserHandlers(userService)       // 4. gRPC обработчики
	grpc.RunServer(userHandler, ":50051")                  // 5. Запуск сервера
}
