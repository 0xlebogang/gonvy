package auth

import (
	"time"

	"gonvy.internal/ent"
)

type CreateSessionCmd struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
	IPAddress *string
	UserAgent *string
}

type RegisterResult struct {
	User    ent.User
	Session ent.Session
}

type registerOutput struct {
	Body struct {
		UserID    string    `json:"userId"`
		AccountID string    `json:"accountId"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
		ExpiresAt time.Time `json:"expiresAt"`
	}
}

type loginInput struct {
	Body struct {
		Email    string `json:"email" minLength:"3" maxLength:"254" format:"email"`
		Password string `json:"password" minLength:"8" maxLength:"128" format:"alphannum"`
	}
}

type User struct {
	ID    string
	Name  string
	Email string
}

type LoginResult struct {
	User    User
	Session Session
}
