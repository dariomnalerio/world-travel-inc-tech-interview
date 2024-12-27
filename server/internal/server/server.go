package server

import (
	"net/http"
	h "server/internal/api/handlers"
	m "server/internal/api/middleware"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	router      *gin.Engine
	userHandler *h.UserHandler
}

func NewServer(userHandler h.UserHandler) *Server {
	return &Server{
		router:      gin.Default(),
		userHandler: &userHandler,
	}
}

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

	protected := v1.Group("/app")
	protected.Use(m.AuthMiddleware())
	{
		protected.GET("/users", s.userHandler.GetUsers)
	}
}

func (s *Server) Run(addr string) error {
	s.router.Use(cors.Default())
	s.setupRoutes()
	return s.router.Run(addr)
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
