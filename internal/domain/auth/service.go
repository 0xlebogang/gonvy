package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/alexedwards/argon2id"
	"gonvy.internal/internal/domain/users"
)

const sessionTTL = 30 * 24 * time.Hour

type Service interface {
	RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*RegisterResult, error)
	Login(ctx context.Context, in *loginInput) (*LoginResult, error)
}

type service struct {
	repo     Repository
	userRepo users.Repository
}

func newService(repo Repository, userRepo users.Repository) Service {
	if repo == nil {
		panic("repo is nil")
	}
	if userRepo == nil {
		panic("user repo is nil")
	}
	return &service{repo: repo, userRepo: userRepo}
}

func (s *service) RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*RegisterResult, error) {
	hash, err := argon2id.CreateHash(in.Body.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, err
	}
	cmd := &users.RegisterUserCmd{
		Name:     in.Body.Name,
		Email:    in.Body.Email,
		Password: hash,
	}
	cmd.Normalize()
	user, err := s.userRepo.CreateUser(ctx, cmd)
	if err != nil {
		return nil, err
	}
	token, err := newSession()
	if err != nil {
		return nil, err
	}
	sess, err := s.repo.CreateSession(ctx, &CreateSessionCmd{
		UserID:    user.UserID,
		Token:     token,
		ExpiresAt: time.Now().UTC().Add(sessionTTL),
	})
	if err != nil {
		return nil, fmt.Errorf("create session after registration: %w", err)
	}
	return &RegisterResult{
		Token:   token,
		User:    *user,
		Session: *sess,
	}, nil
}

func (s *service) Login(ctx context.Context, in *loginInput) (*LoginResult, error) {
	//	Get user by Email
	// Compare user's password hash with provided password
	//	Create new session if password matches, otherwise return unauthorized domain error
	return nil, nil
}
