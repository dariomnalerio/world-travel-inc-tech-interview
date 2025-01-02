package repositories

import (
	"database/sql"
	"server/db/queries"
	e "server/internal/errors"
	"server/internal/models"
)

// UserRepository defines the interface for user-related database operations.
type UserRepository interface {
	Create(user *models.User) (models.CreateUserResponse, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(id string) (*models.User, error)
}

// userRepository is a struct that provides methods to interact with the user data in the repository.
type userRepository struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository.
//
// It returns a pointer to a userRepository struct that implements the UserRepository interface.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Create adds a new user to the repository. If a user with the same email already exists,
// it returns an error indicating that the email is already in use. Otherwise, it stores
// the user and returns a response containing the user's ID and email.
//
// Parameters:
//
//	user (*models.User): A pointer to the User model to be created.
//
// Returns:
//
//	(models.CreateUserResponse, error): A response containing the user's ID and email,
//	and an error if the email already exists or if there was an issue creating the user.
func (r *userRepository) Create(user *models.User) (models.CreateUserResponse, error) {
	userID, err := queries.CreateUser(r.db, user)
	if err != nil {
		return models.CreateUserResponse{}, e.NewError(e.InternalErr, e.DatabaseError, "error creating user", err)
	}

	return models.CreateUserResponse{
		ID:    userID,
		Email: user.Email,
	}, nil
}

// FindByEmail retrieves a user by their email address.
// It returns a pointer to the User model if found, otherwise it returns nil.
// If the user does not exist, the error will also be nil.
//
// Parameters:
//   - email: The email address of the user to be retrieved.
//
// Returns:
//   - *models.User: A pointer to the User model if found, otherwise nil.
//   - error: An error if there was an issue retrieving the user
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	return queries.GetUserByEmail(r.db, email)
}

// FindByID retrieves a user by their ID.
// It returns a pointer to the User model if found, otherwise it returns nil.
// If the user does not exist, the error will also be nil.
//
// Parameters:
//   - id: The ID of the user to be retrieved.
//
// Returns:
//   - *models.User: A pointer to the User model if found, otherwise nil.
//   - error: An error if there was an issue retrieving the user
func (r *userRepository) FindByID(id string) (*models.User, error) {
	return queries.GetUserByID(r.db, id)
}
