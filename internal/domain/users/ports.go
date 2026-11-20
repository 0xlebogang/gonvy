package users

import "strings"

type CreateUserCmd struct {
	Name     string
	Email    string
	Password string
}

func (cmd *CreateUserCmd) normalize() {
	cmd.Name = strings.TrimSpace(cmd.Name)
	cmd.Email = strings.TrimSpace(strings.ToLower(cmd.Email))
	// Replace with password hashing logic
	cmd.Password = strings.TrimSpace(cmd.Password)
}

type CreatedUser struct {
	UserID    string
	AccountID string
	Name      string
	Email     string
}
