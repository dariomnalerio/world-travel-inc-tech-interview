package services_test

import (
	s "server/internal/api/services"
	e "server/internal/errors"
	"server/internal/models"
	testing_mocks "server/internal/testing"
	"testing"

	"github.com/stretchr/testify/assert"
)

const email = "test@example.com"
const validPass = "validPass123!"
const validPassHash = "$2a$10$Vlm2G.ULq2M9TbNTXCxlKu.mFv3g5CJw8/OEj02aTlfsF.zEsq9ly"

func TestRegister(t *testing.T) {
	user := &models.User{
		ID:           "1",
		Email:        email,
		PasswordHash: validPassHash,
	}

	t.Run("successful registration", func(t *testing.T) {
		builder := testing_mocks.NewMockBuilder()
		builder.WithSuccessfulUserNotFound(user.Email).WithSuccessfulCreate()
		service := s.NewUserService(builder.Build())

		response, err := service.Register(user.Email, validPass)

		assert.Equal(t, user.Email, response.Email)
		assert.Equal(t, user.ID, response.ID)
		assert.NoError(t, err)
		builder.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		builder := testing_mocks.NewMockBuilder()
		builder.WithSuccessfulUserNotFound(user.Email).WithDatabaseError()
		service := s.NewUserService(builder.Build())

		_, err := service.Register(user.Email, validPass)

		assert.Error(t, err)
		assert.IsType(t, &e.InternalError{}, err)
		userErr := err.(*e.InternalError)
		assert.Equal(t, e.DatabaseError, userErr.Code)
		assert.Equal(t, "failed to create user", userErr.Message)
		builder.AssertExpectations(t)
	})

	t.Run("duplicate email", func(t *testing.T) {
		builder := testing_mocks.NewMockBuilder()
		builder.WithDuplicateEmail("existing@example.com")
		service := s.NewUserService(builder.Build())

		_, err := service.Register("existing@example.com", validPass)

		assert.Error(t, err)
		assert.IsType(t, &e.UserError{}, err)
		userErr := err.(*e.UserError)
		assert.Equal(t, e.EmailAlreadyExists, userErr.Code)
		assert.Equal(t, "email already exists", userErr.Message)
		builder.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		builder := testing_mocks.NewMockBuilder()
		builder.WithUserFound(user.Email).WithInvalidPassword("wrongPass")
		service := s.NewUserService(builder.Build())

		token, err := service.Login(user.Email, "wrongPass")

		assert.Error(t, err)
		assert.IsType(t, &e.AuthError{}, err)
		authErr := err.(*e.AuthError)
		assert.Equal(t, e.InvalidCredentials, authErr.Code)
		assert.Equal(t, "invalid credentials", authErr.Message)
		assert.Empty(t, token)
		builder.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		builder := testing_mocks.NewMockBuilder()
		builder.WithUserNotFound("nonexistent@example.com")
		service := s.NewUserService(builder.Build())

		token, err := service.Login("nonexistent@example.com", "anyPass")

		assert.Error(t, err)
		assert.IsType(t, &e.UserError{}, err)
		userErr := err.(*e.UserError)
		assert.Equal(t, e.UserNotFound, userErr.Code)
		assert.Equal(t, "invalid credentials", userErr.Message)
		assert.Empty(t, token)
		builder.AssertExpectations(t)
	})
}
