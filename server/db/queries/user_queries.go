package queries

import (
	"database/sql"
	"log"
	"server/internal/models"

	_ "github.com/lib/pq"
)

// GetUserByEmail retrieves a user from the database by their email address.
//
// If no user is found with the given email, it returns (nil, nil).
func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow("SELECT id, email, password_hash, created_at, updated_at FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user from the database by their ID.
//
// If no user is found with the given ID, it returns (nil, nil).
func GetUserByID(db *sql.DB, id string) (*models.User, error) {
	user := &models.User{}
	err := db.QueryRow("SELECT id, email, password_hash, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, user *models.User) (string, error) {
	log.Println("Creating user:", user.Email)
	var userID string
	err := db.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING id", user.Email, user.PasswordHash).
		Scan(&userID)

	if err != nil {
		return "", err
	}

	log.Println("User created with ID:", userID)
	return userID, nil
}
