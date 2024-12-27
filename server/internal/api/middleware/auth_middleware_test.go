package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"server/internal/api/middleware"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var jwtSecret = "secret"

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("no authorization header", func(t *testing.T) {
		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)

		router.Use(a.AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.JSONEq(t, `{"error": "Unauthorized"}`, resp.Body.String())
	})

	t.Run("invalid token", func(t *testing.T) {
		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)
		router.Use(a.AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer invalidtoken")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.JSONEq(t, `{"error": "Unauthorized"}`, resp.Body.String())
	})

	t.Run("valid token", func(t *testing.T) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "1234567890",
			"iat": 1516239022,
		})
		tokenString, _ := token.SignedString([]byte(jwtSecret))

		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)

		router.Use(a.AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})
}
