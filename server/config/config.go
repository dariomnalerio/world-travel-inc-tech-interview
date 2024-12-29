package config

import (
	"log"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Logs      LogConfig
	DB        DBConfig
	JWTSecret string
	Port      string
	DogApiUrl string
}

type LogConfig struct {
	Style string
	Level string
}

type DBConfig struct {
	Username string
	Password string
	URL      string
	Port     string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() (*Config, error) {
	once.Do(func() {
		env := os.Getenv("SERVER_ENV")
		var dbURL string

		switch env {
		case "production":
			dbURL = os.Getenv("DATABASE_URL_PROD")
		case "development":
			dbURL = os.Getenv("DATABASE_URL_DEV")
		case "testing":
			dbURL = os.Getenv("DATABASE_URL_TEST")
		}

		cfg = &Config{
			Port: os.Getenv("PORT"),
			Logs: LogConfig{
				Style: os.Getenv("LOG_STYLE"),
				Level: os.Getenv("LOG_LEVEL"),
			},
			JWTSecret: os.Getenv("JWT_SECRET"),
			DB: DBConfig{
				Username: os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				URL:      dbURL,
				Port:     os.Getenv("POSTGRES_PORT"),
			},
			DogApiUrl: os.Getenv("DOG_API_URL"),
		}
	})

	return cfg, nil
}

func GetConfig() *Config {
	if cfg == nil {
		log.Fatal("Config not loaded. Call LoadConfig() first.")
	}
	return cfg
}
