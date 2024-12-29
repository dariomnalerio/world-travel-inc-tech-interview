package testing

import (
	"server/internal/models"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}
type MockUserRepository = Mock
type MockDogRepository = Mock
type MockLikedImagesRepository = Mock

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

// FindByID retrieves a user by their ID from the mock repository.
// It returns a pointer to the User model and an error if the user is not found or any other issue occurs.
//
// Parameters:
//   - id: The ID of the user to be retrieved.
//
// Returns:
//   - *models.User: A pointer to the User model if found, otherwise nil.
//   - error: An error if the user is not found or any other issue occurs.
func (m *MockUserRepository) FindByID(id string) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

// GetRandomPicture retrieves a random dog picture from the mock repository.
//
// Returns:
//   - string: A string containing the URL of the random dog picture.
//   - error: An error object if the operation fails, otherwise nil.
func (m *MockDogRepository) GetRandomPicture() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// AddLikedImage inserts a new liked image into the repository and returns an error if the operation fails.
//
// Parameters:
//   - userID: The ID of the user who liked the image.
//   - imageURL: The URL of the image that was liked.
//
// Returns:
//   - error: An error object if the operation fails, otherwise nil.
func (m *MockLikedImagesRepository) AddLikedImage(userID, imageURL string) error {
	args := m.Called(userID, imageURL)
	return args.Error(0)
}

// GetLikedImages retrieves a list of liked images for a specific user from the mock repository.
//
// Parameters:
//   - userID: The ID of the user whose liked images are to be retrieved.
//
// Returns:
//   - []string: A slice of strings containing the URLs of the liked images.
//   - error: An error object if the operation fails, otherwise nil.
func (m *MockLikedImagesRepository) GetLikedImages(userID string) ([]string, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}

// RemoveLikedImage removes a liked image from the repository and returns an error if the operation fails.
//
// Parameters:
//   - userID: The ID of the user who unliked the image.
//   - imageURL: The URL of the image that was unliked.
//
// Returns:
//   - error: An error object if the operation fails, otherwise nil.
func (m *MockLikedImagesRepository) RemoveLikedImage(userID, imageURL string) error {
	args := m.Called(userID, imageURL)
	return args.Error(0)
}
