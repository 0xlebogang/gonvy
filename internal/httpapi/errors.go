package httpapi

import (
	"errors"

	"github.com/danielgtaylor/huma/v2"
	"gonvy.internal/internal/httperr"
)

func setupErrors() {
	huma.NewError = func(status int, msg string, errs ...error) huma.StatusError {
		for _, err := range errs {
			if httpErr, ok := errors.AsType[httperr.Error](err); ok {
				return httpErr
			}
		}
		var cause error
		if msg != "" {
			cause = errors.New(msg)
		} else if len(errs) > 0 {
			cause = errors.Join(errs...)
		}
		return httperr.FromStatus(status, cause)
	}
}
