package httperr

import (
	"errors"
	"net/http"
)

type HTTPError interface {
	error
	HTTPStatus() int
	ErrorCode() ErrorCode
}

type APIError struct {
	Status  int       `json:"status"`
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

func FromError(err error) *APIError {
	if err == nil {
		return nil
	}

	var httpErr HTTPError
	if errors.As(err, &httpErr) {
		return &APIError{
			Status:  httpErr.HTTPStatus(),
			Code:    httpErr.ErrorCode(),
			Message: httpErr.Error(),
			Err:     err,
		}
	}

	return &APIError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Message: "An unexpected error occured",
		Err:     err,
	}
}
