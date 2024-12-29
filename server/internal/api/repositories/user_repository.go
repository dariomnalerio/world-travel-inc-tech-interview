package repositories

import (
	"fmt"

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
type userRepository struct{}

// NewUserRepository creates a new instance of UserRepository.
//
// It returns a pointer to a userRepository struct that implements the UserRepository interface.
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// users is an in-memory map that stores user data with the user's ID as the key.
// This map acts as a temporary storage for user information during the application's runtime.
var users = map[string]*models.User{}

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
	if user, exists := users[user.Email]; exists {
		return models.CreateUserResponse{
			ID:    user.ID,
			Email: user.Email,
		}, e.NewError(e.UserErr, e.EmailAlreadyExists, "email already exists", nil)
	}

	users[user.Email] = user
	id := fmt.Sprint(len(users))
	return models.CreateUserResponse{
		ID:    id,
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
	user, exists := users[email]
	if !exists {
		return nil, nil
	}
	return user, nil
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
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, nil
}
