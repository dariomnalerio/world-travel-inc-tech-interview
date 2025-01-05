package utils

import (
	"server/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// GenerateJWT generates a JSON Web Token (JWT) for a given user ID.
// The token is signed using the secret key from the configuration and includes standard claims:
// - "iss" (issuer): Identifies the principal that issued the JWT.
// - "sub" (subject): Identifies the subject of the JWT, in this case, the user ID.
// - "exp" (expiration time): Identifies the expiration time on or after which the JWT must not be accepted for processing.
// - "nbf" (not before): Identifies the time before which the JWT must not be accepted for processing.
// - "iat" (issued at): Identifies the time at which the JWT was issued.
// - "jti" (JWT ID): Provides a unique identifier for the JWT.
//
// Parameters:
// - userId: The ID of the user for whom the JWT is being generated.
//
// Returns:
// - A signed JWT as a string.
// - An error if there was a problem generating the token.
func GenerateJWT(userId string) (string, error) {
	cfg := config.GetConfig()
	jwtKey := []byte(cfg.JWTSecret)

	claims := jwt.MapClaims{
		"iss": "wti-tech-interview",
		"sub": userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"jti": uuid.New().String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// RefreshJWT generates a new JWT token with an extended expiration time.
// The new token is signed using the secret key from the configuration and includes the same claims as the original token.
//
// Parameters:
// - tokenString: The original JWT token that needs to be refreshed.
//
// Returns:
// - A signed JWT with an extended expiration time.
// - An error if there was a problem refreshing the token.
func RefreshJWT(tokenString string) (string, error) {
	cfg := config.GetConfig()
	jwtKey := []byte(cfg.JWTSecret)

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := newToken.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
