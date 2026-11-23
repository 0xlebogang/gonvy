package users

import (
	"strings"
)

type RegisterUserCmd struct {
	Name     string
	Email    string
	Password string
}

func (cmd *RegisterUserCmd) Normalize() {
	cmd.Name = strings.TrimSpace(cmd.Name)
	cmd.Email = strings.TrimSpace(strings.ToLower(cmd.Email))
}

type RegisteredUser struct {
	UserID    string
	AccountID string
	Name      string
	Email     string
}

type RegisterInput struct {
	Body struct {
		Name     string `json:"name" minLength:"1" maxLength:"100"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}

type RegisterOutput struct {
	Body RegisteredUser
}
