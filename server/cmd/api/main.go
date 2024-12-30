package main

import (
	"log"
	"server/config"
	"server/docs"
	"server/internal/api/handlers"
	"server/internal/api/repositories"
	"server/internal/api/services"
	"server/internal/server"
)

// @title						WTI-Tech-Interview API
// @version					1.0
// @description				API for WTI-Tech-Interview
// @BasePath					/api/v1
// @host						localhost:8080
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	docs.SwaggerInfo.Schemes = []string{"http"}

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	dogRepo := repositories.NewDogAPIRepository(cfg.DogApiBaseURL)
	dogService := services.NewDogService(dogRepo)
	dogHandler := handlers.NewDogHandler(dogService)

	likedImagesRepo := repositories.NewLikedImagesRepository()
	likedImagesService := services.NewLikedImagesService(likedImagesRepo, userRepo)
	likedImagesHandler := handlers.NewLikedImagesHandler(likedImagesService)

	server := server.NewServer(*userHandler, *dogHandler, *likedImagesHandler)

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
