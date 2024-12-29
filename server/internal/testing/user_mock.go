package testing

import (
	e "server/internal/errors"
	"server/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockBuilder struct {
	mock *MockUserRepository
}

// NewMockBuilder creates a new User Mock Builder
func NewMockBuilder() *MockBuilder {
	return &MockBuilder{
		mock: &MockUserRepository{},
	}
}

const email = "test@example.com"
const successHash = "$2a$10$Vlm2G.ULq2M9TbNTXCxlKu.mFv3g5CJw8/OEj02aTlfsF.zEsq9ly"

var user = &models.User{
	ID:           "1",
	Email:        email,
	PasswordHash: successHash,
}

// Successful not found - when user does not exist and should not exist on the database
func (b *MockBuilder) WithSuccessfulUserNotFound(email string) *MockBuilder {
	b.mock.On("FindByEmail", email).Return(nil, nil)
	return b
}

// Successful found - when user exists and should exist on the database
func (b *MockBuilder) WithUserFound(email string) *MockBuilder {

	user := &models.User{
		Email:        email,
		PasswordHash: successHash,
	}
	b.mock.On("FindByEmail", email).Return(user, nil)

	return b
}

// WithUserNotFound sets up the mock to return a UserNotFound error when FindByEmail is called with the given email.
// Used for login service mock, returns empty token and error
func (b *MockBuilder) WithUserNotFound(email string) *MockBuilder {
	b.mock.On("FindByEmail", email).Return(nil, nil)
	return b
}

// WithSuccessfulCreate sets up the mock to successfully create a User.
func (b *MockBuilder) WithSuccessfulCreate() *MockBuilder {
	b.mock.On("Create", mock.AnythingOfType("*models.User")).Return(models.CreateUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil)
	return b
}

func (b *MockBuilder) WithDuplicateEmail(email string) *MockBuilder {
	b.mock.On("FindByEmail", email).Return(&models.User{}, nil)
	return b
}

func (b *MockBuilder) WithDatabaseError() *MockBuilder {
	internalErr := &e.InternalError{
		Code:    e.DatabaseError,
		Message: "failed to create user",
	}

	b.mock.On("Create", mock.Anything).Return(models.CreateUserResponse{}, internalErr)
	return b
}

func (b *MockBuilder) WithInvalidPassword(password string) *MockBuilder {
	b.mock.On("FindByEmail", mock.Anything).Return(&models.User{
		PasswordHash: "",
	}, nil)
	return b
}

func (b *MockBuilder) WithFoundByID() *MockBuilder {
	b.mock.On("FindByID", user.ID).Return(user, nil)
	return b
}

func (b *MockBuilder) WithNotFoundByID() *MockBuilder {
	b.mock.On("FindByID", user.ID).Return(nil, nil)
	return b
}

func (b *MockBuilder) WithErrorFindByID() *MockBuilder {
	internalErr := &e.InternalError{
		Code:    e.DatabaseError,
		Message: "failed to find user",
	}
	b.mock.On("FindByID", user.ID).Return(nil, internalErr)
	return b
}

func (b *MockBuilder) Build() *MockUserRepository {
	return b.mock
}

func (b *MockBuilder) AssertExpectations(t mock.TestingT) {
	b.mock.AssertExpectations(t)
}
