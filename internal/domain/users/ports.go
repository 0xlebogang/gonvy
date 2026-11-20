package users

import (
	"strings"

	"gonvy.internal/ent"
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

type RegisterResult struct {
	User    ent.User
	Account ent.Account
}

type RegisterInput struct {
	Body struct {
		Name     string `json:"name" minLength:"1" maxLength:"100"`
		Email    string `json:"email" format:"email"`
		Password string `json:"password" minLength:"8" maxLength:"128"`
	}
}

type RegisterOutput struct {
	Body struct {
		UserID    string `json:"user_id"`
		AccountID string `json:"account_id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
	}
}
