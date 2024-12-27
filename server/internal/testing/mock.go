package testing

import (
	"server/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *models.User) (models.CreateUserResponse, error) {
	args := m.Called(user)

	var resp models.CreateUserResponse
	if respInterface := args.Get(0); respInterface != nil {
		resp = respInterface.(models.CreateUserResponse)
	}

	return resp, args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}
