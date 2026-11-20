package auth

import "errors"

var (
	ErrUnauthorized    = errors.New("unauthorized")
	ErrUserNotFound    = errors.New("user not found")
	ErrAccountNotFound = errors.New("account not found")
)
