package users

import "errors"

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)
