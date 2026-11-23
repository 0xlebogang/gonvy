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
	sessionRec, err := s.repo.CreateSession(ctx, &CreateSessionCmd{
		UserID:    userRec.User.ID,
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
	in.nomalize()
	u, err := s.userRepo.FindUserByEmail(ctx, in.Body.Email)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
	// Compare user's password hash with provided password
	var hash string
	for _, acc := range u.Edges.Accounts {
		if acc.ProviderID == "credential" && acc.Password != nil {
			hash = *acc.Password
			break
		}
	}
	if hash == "" {
		return nil, ErrUnauthorized
	}
	match, err := argon2id.ComparePasswordAndHash(in.Body.Password, hash)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, ErrUnauthorized
	}
	//	Create new session if password matches, otherwise return unauthorized domain error
	token, err := newSession()
	if err != nil {
		return nil, err
	}
	cmd := &CreateSessionCmd{
		UserID:    u.ID,
		Token:     token,
		ExpiresAt: time.Now().UTC().Add(sessionTTL),
	}
	sess, err := s.repo.CreateSession(ctx, cmd)
	if err != nil {
		return nil, err
	}
	out := &LoginResult{
		Token:   token,
		Session: *sess,
		User:    *user,
	}
	return out, nil
}
