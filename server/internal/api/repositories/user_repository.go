package repositories

import (
	"fmt"

	e "server/internal/errors"
	"server/internal/models"
)

type UserRepository interface {
	Create(user *models.User) (models.CreateUserResponse, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

var users = map[string]*models.User{}

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

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	user, exists := users[email]
	if !exists {
		return nil, nil
	}
	return user, nil
}
