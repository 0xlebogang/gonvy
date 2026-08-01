package httperr

import (
	"errors"
	"log"
	"net/http"
)

type HTTPError interface {
	error
	HTTPStatus() int
	ErrorCode() ErrorCode
	Log()
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

func (e *APIError) Log() {
	log.Printf("[%s]: %v\n", e.Code, e.Err)
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
		}
	}

	return &APIError{
		Status:  http.StatusInternalServerError,
		Code:    "INTERNAL_ERROR",
		Message: "An unexpected error occured",
		Err:     err,
	}
}
