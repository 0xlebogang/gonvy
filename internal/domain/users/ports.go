package users

import (
	"strings"
)

type RegisterUserCmd struct {
	Name     string
	Email    string
	Password string
}

func (cmd *RegisterUserCmd) normalize() {
	cmd.Name = strings.TrimSpace(cmd.Name)
	cmd.Email = strings.TrimSpace(strings.ToLower(cmd.Email))
}

type RegisteredUser struct {
	UserID    string
	AccountID string
	Name      string
	Email     string
}

type RegisterUserAccountInput struct {
	Body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
}
