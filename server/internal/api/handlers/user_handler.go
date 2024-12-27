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

func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginUserRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		err := e.NewError(e.UserErr, e.InvalidCredentials, "correct email and password are required", err)
		utils.HandleError(c, err)
		return
	}

	token, err := h.userService.Login(req.Email, req.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.Header("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   token,
	})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	// quick return for demo
	c.JSON(http.StatusOK, gin.H{
		"users": []string{"user1", "user2"},
	})
}
