package main

import (
	"log"
	"server/config"
	"server/internal/api/handlers"
	"server/internal/api/repositories"
	"server/internal/api/services"
	"server/internal/server"
)

func main() {

	_, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	server := server.NewServer(*userHandler)

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
