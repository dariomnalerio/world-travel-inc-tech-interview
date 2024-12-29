package errors

import (
	"fmt"
)

type ErrorType string

const (
	UserErr          ErrorType = "user_error"
	AuthorizationErr ErrorType = "authorization_error"
	InternalErr      ErrorType = "internal_error"
	ValidationErr    ErrorType = "validation_error"
)

type ErrorCode string

const (
	InvalidEmail          ErrorCode = "invalid_email"
	FailedHash            ErrorCode = "failed_hash"
	EmailAlreadyExists    ErrorCode = "email_already_exists"
	DatabaseError         ErrorCode = "database_error"
	UserNotFound          ErrorCode = "user_not_found"
	InvalidCredentials    ErrorCode = "invalid_credentials"
	InvalidToken          ErrorCode = "invalid_token"
	JWTError              ErrorCode = "jwt_error"
	ExternalAPIError      ErrorCode = "external_api_error"
	EmptyImageURL         ErrorCode = "empty_image_url"
	MalformedURL          ErrorCode = "malformed_url"
	InvalidImageExtension ErrorCode = "invalid_image_extension"
	InvalidProtocol       ErrorCode = "invalid_protocol"
	ImageAlreadyLiked     ErrorCode = "image_already_liked"
	ImageNotLiked         ErrorCode = "image_not_liked"
)

// AppError represents a custom error interface that extends the standard error interface.
// It includes methods to retrieve an error code and to unwrap the underlying error.
type AppError interface {
	error
	Code() ErrorCode
	Unwrap() error
}

// UserError represents an error with a specific code and message.
// It also wraps an underlying error.
//
// Fields:
// - Code: A specific error code of type ErrorCode.
// - Message: A human-readable message describing the error.
// - Err: The underlying error that triggered this UserError.
type UserError struct {
	Code    ErrorCode
	Message string
	Err     error
}

// Error returns the error message for the UserError type.
// If the underlying error (Err) is not nil, it includes the message and the error.
// Otherwise, it returns only the message.
func (e *UserError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap returns the underlying error of a UserError.
// It allows access to the original error that caused the UserError.
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

type ValidationError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *ValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *ValidationError) Unwrap() error {
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

// NewError creates a new error based on the provided error type, code, message, and underlying error.
// It returns an error of type UserError, AuthError, or InternalError depending on the errType parameter.
//
// Parameters:
//   - errType: a string representing the type of error ("user_error", "authorization_error", or other).
//   - code: an ErrorCode representing the specific error code.
//   - message: a string containing the error message.
//   - err: an error representing the underlying error.
//
// Returns:
//   - error: an error of type UserError, AuthError, or InternalError.
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
	case "validation_error":
		return &ValidationError{
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
