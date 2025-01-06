package main

import (
	"log"
	"server/config"
	"server/db"
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

	db, err := db.InitDB(cfg.DB.URL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	docs.SwaggerInfo.Schemes = []string{"http"}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	likedImagesRepo := repositories.NewLikedImagesRepository(db)
	likedImagesService := services.NewLikedImagesService(likedImagesRepo, userRepo)
	likedImagesHandler := handlers.NewLikedImagesHandler(likedImagesService)

	dogRepo := repositories.NewDogAPIRepository(cfg.DogApiBaseURL)
	dogService := services.NewDogService(dogRepo, likedImagesRepo)
	dogHandler := handlers.NewDogHandler(dogService)

	server := server.NewServer(*userHandler, *dogHandler, *likedImagesHandler)

	if err := server.Run(":8080", "api/v1"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
