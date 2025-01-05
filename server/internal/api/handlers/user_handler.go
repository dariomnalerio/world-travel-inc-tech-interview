package handlers

import (
	"log"
	"net/http"
	"server/internal/api/services"
	e "server/internal/errors"
	"server/internal/models"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService,
	}
}

// Register godoc
//
//	@Summary		Registers a new user.
//	@Description	Registers a new user with the provided email and password.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.CreateUserRequest	true	"User registration request"
//	@Success		201		{object}	models.CreateUserResponse
//	@Failure		400		{object}	utils.ErrorResponse
//	@Router			/auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		err := e.NewError(e.UserErr, e.InvalidCredentials, "correct email and password are required", err)
		utils.HandleError(c, err)
		return
	}

	if !utils.IsValidPassword(req.Password) {
		err := e.NewError(e.UserErr, e.InvalidCredentials, "password must contain at least 8 characters, at most 32 characters, at least one uppercase letter, at least one lowercase letter, at least one number, and at least one special character", nil)
		utils.HandleError(c, err)
		return
	}

	user, err := h.userService.Register(req.Email, req.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// Login godoc
//
//	@Summary		Logs in an existing user.
//	@Description	Logs in an existing user with the provided email and password.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.LoginUserRequest	true	"User login request"
//	@Success		200		{object}	models.LoginUserResponse
//	@Failure		400		{object}	utils.ErrorResponse
//	@Router			/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		err := e.NewError(e.UserErr, e.InvalidCredentials, "correct email and password are required", err)
		utils.HandleError(c, err)
		return
	}

	res, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Header("Authorization", "Bearer "+res.Token)
	// not http only, secure, localhost as domain
	// would update for prod app
	c.SetCookie("auth_token", "Bearer "+res.Token, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   res.Token,
		"userID":  res.ID,
	})
}

// GetUser godoc
//
//	@Summary		Retrieves a user by ID.
//	@Description	Retrieves a user by their ID.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		200		{object}	models.User
//	@Failure		400		{object}	utils.ErrorResponse
//
//	@Security		BearerAuth
//
//	@Router			/user/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	log.Println("Getting user by ID")
	var id = c.Param("id")
	if id == "" {
		err := e.NewError(e.UserErr, e.InvalidCredentials, "user ID is required", nil)
		utils.HandleError(c, err)
		return
	}

	user, err := h.userService.GetUserByID(id)

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": models.UserResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	})
}

// VerifyAuth godoc
//
//	@Summary		Verifies user authentication.
//	@Description	Verifies that the user is authenticated.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	string
//
//	@Security		BearerAuth
//
//	@Router			/auth/verify [get]
func (h *UserHandler) VerifyAuth(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.HandleError(c, e.NewError(e.UserErr, e.InvalidToken, "unauthorized", nil))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User authenticated",
		"userID":  userID,
	})
}
