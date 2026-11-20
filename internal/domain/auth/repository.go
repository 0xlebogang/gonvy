package auth

import (
	"context"
	"fmt"

	"gonvy.internal/ent"
)

type Repository interface {
	CreateSession(ctx context.Context, cmd *CreateSessionCmd) (*Session, error)
}

type repository struct {
	db *ent.Client
}

func newRepository(db *ent.Client) Repository {
	if db == nil {
		panic("db is nil")
	}
	return &repository{db: db}
}

func (r *repository) CreateSession(ctx context.Context, cmd *CreateSessionCmd) (*Session, error) {
	s, err := r.db.Session.Create().
		SetUserID(cmd.UserID).
		SetToken(cmd.Token).
		SetExpiresAt(cmd.ExpiresAt).
		SetNillableIPAddress(cmd.IPAddress).
		SetNillableUserAgent(cmd.UserAgent).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return &Session{
		ID:        s.ID,
		Token:     s.Token,
		ExpiresAt: s.ExpiresAt,
	}, nil
}
