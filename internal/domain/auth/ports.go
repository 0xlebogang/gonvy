package auth

import "time"

type CreateSessionCmd struct {
	UserID    string
	Token     string
	ExpiresAt time.Time
	IPAddress *string
	UserAgent *string
}

type Session struct {
	ID        string
	Token     string
	ExpiresAt time.Time
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
