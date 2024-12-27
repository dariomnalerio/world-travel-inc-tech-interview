package utils

import (
	"server/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

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
