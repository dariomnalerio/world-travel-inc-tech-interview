package models

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserCredentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest = UserCredentials
type LoginUserRequest = UserCredentials
type CreateUserResponse = UserResponse

type LoginUserResponse struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}
