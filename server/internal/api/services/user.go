package services

import (
	"server/internal/api/repositories"
	e "server/internal/errors"
	"server/internal/models"
	"server/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, password string) (models.CreateUserResponse, error)
	Login(email, password string) (string, error)
}

type userService struct {
	r repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{r}
}

func (s *userService) Register(email, password string) (models.CreateUserResponse, error) {
	existingUser, err := s.r.FindByEmail(email)
	if err != nil {
		return models.CreateUserResponse{}, e.NewError(e.InternalErr, e.DatabaseError, "internal server error", err)
	}
	if existingUser != nil {
		return models.CreateUserResponse{
			ID:    existingUser.ID,
			Email: existingUser.Email,
		}, e.NewError(e.UserErr, e.EmailAlreadyExists, "email already exists", nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.CreateUserResponse{}, e.NewError(e.InternalErr, e.FailedHash, "failed to hash password", err)
	}

	createdUser, err := s.r.Create(&models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return models.CreateUserResponse{}, e.NewError(e.InternalErr, e.DatabaseError, "failed to create user", err)
	}
	return createdUser, nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.r.FindByEmail(email)
	if err != nil {
		return "", e.NewError(e.InternalErr, e.DatabaseError, "internal server error", err)
	}
	if user == nil {
		return "", e.NewError(e.UserErr, e.UserNotFound, "invalid credentials", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", e.NewError(e.AuthorizationErr, e.InvalidCredentials, "invalid credentials", nil)
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return "", e.NewError(e.InternalErr, e.JWTError, "internal error authenticating user", err)
	}

	return token, nil
}
