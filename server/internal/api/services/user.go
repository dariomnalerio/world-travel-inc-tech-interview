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
	Login(email, password string) (models.LoginUserResponse, error)
	GetUserByID(id string) (*models.User, error)
}

type userService struct {
	r repositories.UserRepository
}

// NewUserService creates a new instance of UserService using the provided UserRepository.
// It returns a UserService interface which can be used to interact with user-related operations.
//
// Parameters:
//   - r: An implementation of the UserRepository interface.
//
// Returns:
//   - UserService: An instance of the UserService interface.
func NewUserService(r repositories.UserRepository) UserService {
	return &userService{r}
}

// Register registers a new user with the given email and password.
// It first checks if a user with the provided email already exists.
// If the user exists, it returns an error indicating that the email already exists.
// If the user does not exist, it hashes the password and creates a new user record in the database.
// Returns the created user response and an error, if any.
//
// Parameters:
//   - email: The email address of the user to be registered.
//   - password: The password of the user to be registered.
//
// Returns:
//   - models.CreateUserResponse: The response containing the created user's details.
//   - error: An error if the registration fails at any step.
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

// Login authenticates a user by their email and password.
// It returns a JWT token if the authentication is successful, or an error if it fails.
//
// Parameters:
//   - email: The email address of the user.
//   - password: The password of the user.
//
// Returns:
//   - string: A JWT token if authentication is successful.
//   - error: An error if authentication fails, which could be due to internal server errors,
//     database errors, user not found, or invalid credentials.
func (s *userService) Login(email, password string) (models.LoginUserResponse, error) {
	user, err := s.r.FindByEmail(email)
	if err != nil {
		return models.LoginUserResponse{}, e.NewError(e.InternalErr, e.DatabaseError, "internal server error", err)
	}
	if user == nil {
		return models.LoginUserResponse{}, e.NewError(e.UserErr, e.UserNotFound, "invalid credentials", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return models.LoginUserResponse{}, e.NewError(e.AuthorizationErr, e.InvalidCredentials, "invalid credentials", nil)
	}

	token, err := utils.GenerateJWT(user.ID)

	if err != nil {
		return models.LoginUserResponse{}, e.NewError(e.InternalErr, e.JWTError, "internal error authenticating user", err)
	}

	response := models.LoginUserResponse{
		Token: token,
		ID:    user.ID,
	}

	return response, nil
}

// GetUserByID retrieves a user from the database by their ID.
//
// If no user is found with the given ID, it returns (nil, nil).
func (s *userService) GetUserByID(id string) (*models.User, error) {
	user, err := s.r.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
