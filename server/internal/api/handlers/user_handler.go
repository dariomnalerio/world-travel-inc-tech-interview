package handlers

import (
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

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   res.Token,
		"userID":  res.ID,
	})
}

// GetUsers godoc
//
//	@Summary		Retrieves a list of users.
//	@Description	Retrieves a list of users.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Security		Bearer
//	@Success		200		{object}	string
//
// @Security BearerAuth
//
//	@Router			/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {

	users, err := h.userService.FindAll()

	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
