package utils

import (
	"log"
	"net/http"
	"server/internal/errors"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error  string           `json:"error"`
	Code   errors.ErrorCode `json:"code"`
	Detail string           `json:"detail,omitempty"` // Optional field for detailed error messages
}

func HandleError(c *gin.Context, err error) {

	var (
		statusCode    int
		errorResponse ErrorResponse
	)

	switch e := err.(type) {
	case *errors.UserError:
		statusCode, errorResponse = handleUserError(e)
	case *errors.AuthError:
		statusCode, errorResponse = http.StatusUnauthorized, ErrorResponse{
			Error:  "Authentication error",
			Code:   e.Code,
			Detail: e.Error(),
		}
	case *errors.InternalError:
		// Log internal errors for debuggin purposes
		log.Printf("Internal error: Code=%s, Message=%s, Detail=%v", e.Code, e.Message, e.Err)

		statusCode, errorResponse = http.StatusInternalServerError, ErrorResponse{
			Error: "Internal Server Error",
			Code:  e.Code,
		}
	default:
		log.Printf("Unknown error: %v", err)
		statusCode, errorResponse = http.StatusInternalServerError, ErrorResponse{
			Error: "Unknown error ocurred",
		}
	}

	c.JSON(statusCode, errorResponse)
}

func handleUserError(e *errors.UserError) (int, ErrorResponse) {
	switch e.Code {
	case errors.InvalidEmail:
		return http.StatusBadRequest, ErrorResponse{
			Error:  "Invalid email format",
			Code:   e.Code,
			Detail: e.Error(),
		}
	case errors.EmailAlreadyExists:
		return http.StatusConflict, ErrorResponse{
			Error:  "Email already registered",
			Code:   e.Code,
			Detail: e.Error(),
		}
	default:
		return http.StatusBadRequest, ErrorResponse{
			Error:  "User error",
			Code:   e.Code,
			Detail: e.Error(),
		}
	}
}
