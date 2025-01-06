package db

import (
	"database/sql"
	"embed"
	"fmt"
	"time"
)

// schemaFS adds the ability to embed files into the binary,
// bundling the schema files into the binary and reading them at runtime
//
//go:embed schema/init.sql
var schemaFS embed.FS

type RetryConfig struct {
	MaxAttempts       int
	InitialBackoff    time.Duration
	MaxBackoff        time.Duration
	BackoffMultiplier float64
}

var DefaultRetryConfig = RetryConfig{
	MaxAttempts:       5,
	InitialBackoff:    time.Second,
	MaxBackoff:        30 * time.Second,
	BackoffMultiplier: 2.0,
}

func InitDB(dbURL string) (*sql.DB, error) {
	return InitDBWithRetry(dbURL, DefaultRetryConfig)
}

func InitDBWithRetry(dbURL string, config RetryConfig) (*sql.DB, error) {
	db, err := connectToDBWithRetry(dbURL, config)

	if err != nil {
		return nil, err
	}

	if err := executeSchema(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("error executing schema: %w", err)
	}

	return db, nil
}

func connectToDBWithRetry(dbURL string, config RetryConfig) (*sql.DB, error) {
	var (
		db      *sql.DB
		err     error
		backoff = config.InitialBackoff
	)

	for attempt := 1; attempt <= config.MaxAttempts; attempt++ {
		db, err = connectToDB(dbURL)
		if err == nil {
			return db, nil
		}

		if attempt == config.MaxAttempts {
			return nil, fmt.Errorf("error connecting to the database: %w", err)
		}

		time.Sleep(backoff)

		backoff = time.Duration(float64(backoff) * config.BackoffMultiplier)
		if backoff > config.MaxBackoff {
			backoff = config.MaxBackoff
		}
	}
	return nil, fmt.Errorf("error connecting to the database: %w", err)
}

func connectToDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return db, nil
}

func executeSchema(db *sql.DB) error {
	schema, err := schemaFS.ReadFile("schema/init.sql")
	if err != nil {
		db.Close()
		return fmt.Errorf("error reading schema file: %w", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		db.Close()
		return fmt.Errorf("error executing schema: %w", err)
	}

	return nil
}
