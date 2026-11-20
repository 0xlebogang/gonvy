package users

import (
	"context"
	"fmt"

	"github.com/alexedwards/argon2id"
)

type Service interface {
	RegisterUserAccount(ctx context.Context, in *RegisterInput) (*RegisteredUser, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	if r == nil {
		panic("repository is nil")
	}
	return &service{repo: r}
}

func (s *service) RegisterUserAccount(ctx context.Context, in *RegisterInput) (*RegisteredUser, error) {
	hash, err := argon2id.CreateHash(in.Body.Password, argon2id.DefaultParams)
	if err != nil {
		return nil, fmt.Errorf("create password hash: %w", err)
	}
	cmd := &RegisterUserCmd{
		Name:     in.Body.Name,
		Email:    in.Body.Email,
		Password: hash,
	}
	cmd.normalize()
	return s.repo.CreateUser(ctx, cmd)
}
