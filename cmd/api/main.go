package main

import (
	"log"
	"project/internal/repository"
	"project/internal/service"
	httpserver "project/internal/transport/http"
)

func main() {
	log.Println("Запуск http-сервера")

	// Инициализация
	userRepo := repository.NewMemoryUserRepository()
	userService := service.NewUserService(userRepo)
	server := httpserver.NewServer(":8008", userService)

	server.Start()
}
