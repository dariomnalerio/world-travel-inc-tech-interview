package config

import (
	"log"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Logs          LogConfig
	DB            DBConfig
	JWTSecret     string
	Port          string
	DogApiBaseURL string
}

type LogConfig struct {
	Style string
	Level string
}

type DBConfig struct {
	Username string
	Password string
	Name     string
	Port     string
	URL      string
}

var (
	cfg  *Config
	once sync.Once
)

func LoadConfig() (*Config, error) {
	once.Do(func() {
		env := os.Getenv("SERVER_ENV")
		var dbURL string
		var postgresDB string

		switch env {
		case "production":
			dbURL = os.Getenv("DATABASE_URL_PROD")
			postgresDB = os.Getenv("POSTGRES_DB_PROD")
		case "development":
			dbURL = os.Getenv("DATABASE_URL_DEV")
			postgresDB = os.Getenv("POSTGRES_DB_DEV")
		case "testing":
			dbURL = os.Getenv("DATABASE_URL_TEST")
			postgresDB = os.Getenv("POSTGRES_DB_TEST")
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
				Port:     os.Getenv("POSTGRES_PORT"),
				Name:     postgresDB,
				URL:      dbURL,
			},
			DogApiBaseURL: "https://dog.ceo/api",
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
