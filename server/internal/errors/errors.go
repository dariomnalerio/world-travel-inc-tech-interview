package errors

import (
	"fmt"
)

type ErrorType string

const (
	UserErr          ErrorType = "user_error"
	AuthorizationErr ErrorType = "authorization_error"
	InternalErr      ErrorType = "internal_error"
)

type ErrorCode string

const (
	InvalidEmail       ErrorCode = "invalid_email"
	FailedHash         ErrorCode = "failed_hash"
	EmailAlreadyExists ErrorCode = "email_already_exists"
	DatabaseError      ErrorCode = "database_error"
	UserNotFound       ErrorCode = "user_not_found"
	InvalidCredentials ErrorCode = "invalid_credentials"
	InvalidToken       ErrorCode = "invalid_token"
	JWTError           ErrorCode = "jwt_error"
)

type AppError interface {
	error
	Code() ErrorCode
	Unwrap() error
}

type UserError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *UserError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func Unwrap(e *UserError) error {
	return e.Err
}

type AuthError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *AuthError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AuthError) Unwrap() error {
	return e.Err
}

type InternalError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *InternalError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *InternalError) Unwrap() error {
	return e.Err
}

func NewError(errType ErrorType, code ErrorCode, message string, err error) error {
	switch errType {
	case "user_error":
		return &UserError{
			Code:    code,
			Message: message,
			Err:     err,
		}
	case "authorization_error":
		return &AuthError{
			Code:    code,
			Message: message,
			Err:     err,
		}
	default:
		return &InternalError{
			Code:    code,
			Message: message,
			Err:     err,
		}
	}
}
