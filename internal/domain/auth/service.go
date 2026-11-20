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
	RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*registerResult, error)
	Login(ctx context.Context, in *loginInput) (*loginResult, error)
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

func (s *service) RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*registerResult, error) {
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
	return &registerResult{
		User:    userRec.User,
		Session: *sessionRec,
	}, nil
}

func (s *service) Login(ctx context.Context, in *loginInput) (*loginResult, error) {
	u, err := s.repo.FindUserByEmail(ctx, in.Body.Email)
	if err != nil {
		return nil, err
	}
	hash := u.Edges.Accounts[0].Password
	if ok, err := argon2id.ComparePasswordAndHash(in.Body.Password, *hash); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrUnauthorized
	}
	token, err := newSession()
	if err != nil {
		return nil, err
	}
	cmd := &CreateSessionCmd{
		UserID:    u.ID,
		ExpiresAt: time.Now().UTC().Add(sessionTTL),
		Token:     hashToken(token),
	}
	_, err = s.repo.CreateSession(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return &loginResult{
		Token: token,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
