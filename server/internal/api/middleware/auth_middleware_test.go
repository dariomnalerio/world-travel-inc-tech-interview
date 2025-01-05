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
var authErrJSON = `{"error":"Authentication error","code":"invalid_token","detail":"unauthorized"}`
var forbiddenErrJSON = `{"error":"Forbidden", "code":"invalid_token", "detail":"forbidden"}`

func TestVerifyJWT(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("no authorization header", func(t *testing.T) {
		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)

		router.Use(a.VerifyJWT())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.JSONEq(t, authErrJSON, resp.Body.String())
	})

	t.Run("invalid token", func(t *testing.T) {
		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)
		router.Use(a.VerifyJWT())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer invalidtoken")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.JSONEq(t, authErrJSON, resp.Body.String())
	})

	t.Run("valid token", func(t *testing.T) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "1234567890",
			"iat": 1516239022,
		})
		tokenString, _ := token.SignedString([]byte(jwtSecret))

		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)

		router.Use(a.VerifyJWT())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
	})

	t.Run("userID set in gin context", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": "1234567890",
			"iat": 1516239022,
		})
		tokenString, _ := token.SignedString([]byte(jwtSecret))

		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)

		router.Use(a.VerifyJWT())
		router.GET("/test", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			c.JSON(http.StatusOK, gin.H{"userID": userID})
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"userID":"1234567890"}`, resp.Body.String())
	})

	t.Run("userID missing from token", func(t *testing.T) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iat": 1516239022,
		})
		tokenString, _ := token.SignedString([]byte(jwtSecret))

		router := gin.New()
		a := middleware.NewAuthMiddleware(jwtSecret)

		router.Use(a.VerifyJWT())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.JSONEq(t, authErrJSON, resp.Body.String())
	})

}

func TestVerifyRequestOwnership(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jwtSecret := "test_secret"
	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	validToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "123",
	})
	validTokenString, _ := validToken.SignedString([]byte(jwtSecret))

	invalidToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	invalidTokenString, _ := invalidToken.SignedString([]byte(jwtSecret))

	router := gin.New()
	router.Use(authMiddleware.VerifyJWT())
	router.Use(authMiddleware.VerifyRequestOwnership())
	router.GET("/test/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	t.Run("valid request ownership", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test/123", nil)
		req.Header.Set("Authorization", "Bearer "+validTokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.JSONEq(t, `{"message": "OK"}`, resp.Body.String())
	})

	t.Run("valid token but mistmatched user ID", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test/456", nil)
		req.Header.Set("Authorization", "Bearer "+validTokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusForbidden, resp.Code)
		assert.JSONEq(t, forbiddenErrJSON, resp.Body.String())
	})

	t.Run("missing userID in context", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/test/123", nil)
		req.Header.Set("Authorization", "Bearer "+invalidTokenString)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.JSONEq(t, authErrJSON, resp.Body.String())
	})

}
