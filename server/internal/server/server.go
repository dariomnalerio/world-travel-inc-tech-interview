package server

import (
	"net/http"
	"server/config"
	h "server/internal/api/handlers"
	m "server/internal/api/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

// Server represents the HTTP server that handles incoming requests.
// It contains a router for routing the requests and handlers for processing requests.
type Server struct {
	router      *gin.Engine
	userHandler *h.UserHandler
}

// NewServer creates a new instance of Server with the provided UserHandler.
// It initializes the router using the default settings from the gin framework.
//
// Parameters:
//   - userHandler: an instance of h.UserHandler to handle user-related routes.
//
// Returns:
//   - A pointer to a newly created Server instance.
func NewServer(userHandler h.UserHandler) *Server {
	return &Server{
		router:      gin.Default(),
		userHandler: &userHandler,
	}
}

// setupRoutes initializes the API routes for the server.
// It sets up the following routes:
//
// - Public routes:
//   - POST /api/v1/auth/register: Registers a new user.
//   - POST /api/v1/auth/login: Logs in an existing user.
//   - GET /api/v1/health: Checks the health of the server.
//
// - Protected routes (requires authentication):
//   - GET /api/v1/users: Retrieves a list of users.
func (s *Server) setupRoutes() {
	v1 := s.router.Group("api/v1")

	public := v1.Group("")
	{
		auth := public.Group("/auth")
		{
			auth.POST("register", s.userHandler.Register)
			auth.POST("login", s.userHandler.Login)
		}

		public.GET("/health", s.healthCheck)
	}

	auth := m.NewAuthMiddleware(config.GetConfig().JWTSecret)
	protected := v1.Group("/")
	protected.Use(auth.AuthMiddleware())
	{
		protected.GET("/users", s.userHandler.GetUsers)
	}
}

// Run starts the server on the specified address.
// It sets up CORS middleware and initializes the routes.
//
// Parameters:
//
//	addr - the address to listen on, in the form "host:port".
//
// Returns:
//
//	error - if the server fails to start, an error is returned.
func (s *Server) Run(addr string) error {
	s.router.Use(cors.Default())
	s.setupRoutes()
	return s.router.Run(addr)
}

// healthCheck handles the health check endpoint.
// It responds with a JSON object indicating the server status.
//
// @param c *gin.Context - the context for the request
//
// @response 200 - JSON object with the server status
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
