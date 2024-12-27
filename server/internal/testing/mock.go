package testing

import (
	"server/internal/models"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}
type MockUserRepository = Mock

// Create inserts a new user into the repository and returns a response containing
// the details of the created user or an error if the operation fails.
//
// Parameters:
//   - user: A pointer to the User model containing the details of the user to be created.
//
// Returns:
//   - models.CreateUserResponse: A response struct containing the details of the created user.
//   - error: An error object if the creation fails, otherwise nil.

func (m *MockUserRepository) Create(user *models.User) (models.CreateUserResponse, error) {
	args := m.Called(user)

	var resp models.CreateUserResponse
	if respInterface := args.Get(0); respInterface != nil {
		resp = respInterface.(models.CreateUserResponse)
	}

	return resp, args.Error(1)
}

// FindByEmail retrieves a user by their email address from the mock repository.
// It returns a pointer to the User model and an error if the user is not found or any other issue occurs.
//
// Parameters:
//   - email: The email address of the user to be retrieved.
//
// Returns:
//   - *models.User: A pointer to the User model if found, otherwise nil.
//   - error: An error if the user is not found or any other issue occurs.
func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}
