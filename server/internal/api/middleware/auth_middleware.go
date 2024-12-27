package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtSecret []byte
}

// NewAuthMiddleware creates a new instance of AuthMiddleware with the provided JWT secret.
// The JWT secret is used to sign and verify JWT tokens.
//
// Parameters:
//   - jwtSecret: A string representing the secret key used for JWT authentication.
//
// Returns:
//   - A pointer to an AuthMiddleware instance initialized with the provided JWT secret.
func NewAuthMiddleware(jwtSecret string) *AuthMiddleware {
	return &AuthMiddleware{
		jwtSecret: []byte(jwtSecret),
	}
}

// AuthMiddleware is a middleware function that handles
// authentication by validating the JWT token present in the "Authorization" header
// of the incoming request. If the token is missing, invalid, or expired, it aborts
// the request and responds with a 401 Unauthorized status. If the token is valid,
// it allows the request to proceed to the next handler.
//
// The JWT token is expected to be in the format "Bearer <token>" and is validated
// using the HMAC signing method with the secret key stored in the AuthMiddleware
// struct.
func (a *AuthMiddleware) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := jwt.Parse(tokenString[len("Bearer "):], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return a.jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

		c.Next()
	}
}
