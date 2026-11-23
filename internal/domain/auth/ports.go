package auth

import (
	"strings"
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

type Session struct {
	ID        string    `json:"id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	IPAddress *string   `json:"ip_address"`
	UserAgent *string   `json:"user_agent"`
}

type RegisterResult struct {
	Token   string
	User    users.RegisteredUser
	Session Session
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

func (l *loginInput) nomalize() {
	l.Body.Email = strings.ToLower(strings.TrimSpace(l.Body.Email))
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginResult struct {
	Token   string
	User    User
	Session Session
}

type loginOutput struct {
	Body struct {
		Token   string  `json:"token"`
		User    User    `json:"user"`
		Session Session `json:"session"`
	}
}
