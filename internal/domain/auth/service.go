package auth

import (
	"context"
	"fmt"
	"time"

	"gonvy.internal/internal/domain/users"
)

const sessionTTL = 30 * 24 * time.Hour

type Service interface {
	RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*RegisterResult, error)
	Login(ctx context.Context, in *loginInput) (*LoginResult, error)
}

type service struct {
	repo    Repository
	userSvc users.Service
}

func newService(repo Repository, userSvc users.Service) Service {
	if repo == nil {
		panic("repo is nil")
	}
	if userSvc == nil {
		panic("user service is nil")
	}
	return &service{repo: repo, userSvc: userSvc}
}

func (s service) RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*RegisterResult, error) {
	userRec, err := s.userSvc.RegisterUserAccount(ctx, in)
	if err != nil {
		return nil, err
	}
	token, err := newSession()
	if err != nil {
		return nil, err
	}
	sessionRec, err := s.repo.CreateSession(ctx, &CreateSessionCmd{
		UserID:    userRec.User.ID,
		Token:     token,
		ExpiresAt: time.Now().UTC().Add(sessionTTL),
	})
	if err != nil {
		return nil, fmt.Errorf("create session after registration: %w", err)
	}
	return &RegisterResult{
		User:    userRec.User,
		Session: *sessionRec,
	}, nil
}

func (s service) Login(ctx context.Context, in *loginInput) (*LoginResult, error) {
	//	Get user by Email
	// Compare user's password hash with provided password
	//	Create new session if password matches, otherwise return unauthorized domain error
	return nil, nil
}
