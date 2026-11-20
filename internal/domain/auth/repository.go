package auth

import (
	"context"
	"fmt"

	"gonvy.internal/ent"
	"gonvy.internal/ent/account"
	"gonvy.internal/ent/user"
)

type Repository interface {
	CreateSession(ctx context.Context, cmd *CreateSessionCmd) (*ent.Session, error)
	FindAccountByEmail(ctx context.Context, email string) (*ent.Account, error)
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

func (r *repository) CreateSession(ctx context.Context, cmd *CreateSessionCmd) (*ent.Session, error) {
	rec, err := r.db.Session.Create().
		SetUserID(cmd.UserID).
		SetToken(cmd.Token).
		SetExpiresAt(cmd.ExpiresAt).
		SetNillableIPAddress(cmd.IPAddress).
		SetNillableUserAgent(cmd.UserAgent).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create session: %w", err)
	}
	return rec, nil
}

func (r *repository) FindAccountByEmail(ctx context.Context, email string) (*ent.Account, error) {
	acc, err := r.db.Account.Query().
		Where(
			account.HasUserWith(user.EmailEQ(email)),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrAccountNotFound
		}
		return nil, fmt.Errorf("find account: %w", err)
	}
	return acc, nil
}
