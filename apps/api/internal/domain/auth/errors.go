package auth

import (
	"net/http"

	"github.com/0xlebogang/gonvy/api/internal/httperr"
)

type AuthError struct {
	Status  int
	Code    httperr.ErrorCode
	Message string
}

func (e *AuthError) Error() string {
	return e.Message
}

func (e *AuthError) HTTPStatus() int {
	return e.Status
}

func (e *AuthError) ErrorCode() httperr.ErrorCode {
	return e.Code
}

var (
	ErrInvalidInput = &AuthError{
		Status:  http.StatusBadRequest,
		Code:    httperr.CodeInvalidInput,
		Message: "Invalid input provided",
	}

	ErrEmailExists = &AuthError{
		Status:  http.StatusConflict,
		Code:    httperr.CodeEntityExists,
		Message: "An account with this email already exists",
	}

	ErrInvalidCredentials = &AuthError{
		Status:  http.StatusUnauthorized,
		Code:    httperr.ErrorCode("INVALID_CREDENTIALS"),
		Message: "Invalid email or password",
	}
)
