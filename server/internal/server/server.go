package server

import (
	"net/http"
	"server/config"
	h "server/internal/api/handlers"
	m "server/internal/api/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Server represents the HTTP server that handles incoming requests.
// It contains a router for routing the requests and handlers for processing requests.
type Server struct {
	router             *gin.Engine
	userHandler        *h.UserHandler
	dogHandler         *h.DogHandler
	likedImagesHandler *h.LikedImagesHandler
}

// NewServer creates a new instance of Server with the provided UserHandler.
// It initializes the router using the default settings from the gin framework.
//
// Parameters:
//   - userHandler: an instance of h.UserHandler to handle user-related routes.
//
// Returns:
//   - A pointer to a newly created Server instance.
func NewServer(userHandler h.UserHandler, dogHandler h.DogHandler, likedImagesHandler h.LikedImagesHandler) *Server {
	return &Server{
		router:             gin.Default(),
		userHandler:        &userHandler,
		dogHandler:         &dogHandler,
		likedImagesHandler: &likedImagesHandler,
	}
}

// setupRoutes initializes the API routes for the server.
func (s *Server) setupRoutes(baseRoute string) {
	v1 := s.router.Group(baseRoute)

	public := v1.Group("")
	{
		auth := public.Group("/auth")
		{
			auth.POST("register", s.userHandler.Register)
			auth.POST("login", s.userHandler.Login)
		}

		public.GET("/dog/random", s.dogHandler.GetRandomImage)
		public.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		public.GET("/health", s.healthCheck)
	}

	auth := m.NewAuthMiddleware(config.GetConfig().JWTSecret)
	protected := v1.Group("")
	protected.Use(auth.VerifyJWT())

	user := protected.Group("/user")
	{
		user.GET("/:id", s.userHandler.GetUser)
	}

	liked_images := protected.Group("/liked_images")
	liked_images.Use(auth.VerifyRequestOwnership())
	{
		liked_images.DELETE("/:id", s.likedImagesHandler.UnlikeImage)
		liked_images.GET("/:id", s.likedImagesHandler.GetLikedImages)
		liked_images.POST("/:id", s.likedImagesHandler.LikeImage)
	}
}

func (s *Server) Run(addr, baseRoute string) error {
	s.router.Use(cors.Default())
	s.setupRoutes(baseRoute)
	return s.router.Run(addr)
}

// HealthCheck godoc
//
//	@Summary		Checks the health of the server.
//	@Description	Verifies that the server is running and healthy.
//	@Tags			health
//	@Produces		json
//	@Success		200	{object}	string
//	@Router			/health [get]
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
