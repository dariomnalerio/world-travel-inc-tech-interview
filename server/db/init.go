package db

import (
	"database/sql"
	"embed"
	"fmt"
)

// schemaFS adds the ability to embed files into the binary,
// bundling the schema files into the binary and reading them at runtime
//
//go:embed schema/init.sql
var schemaFS embed.FS

func InitDB(dbURL string) (*sql.DB, error) {
	db, err := connectToDB(dbURL)

	if err != nil {
		return nil, err
	}

	if err := executeSchema(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("error executing schema: %w", err)
	}

	return db, nil
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
