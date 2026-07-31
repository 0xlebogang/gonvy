package user

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/common"
)

type Service interface {
	CreateUser(ctx context.Context, b *User) (*UserResponse, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}

type service struct {
	repo   Repository
	hasher common.Hasher
}

func NewService(r Repository, h common.Hasher) Service {
	return &service{repo: r, hasher: h}
}

func (s *service) CreateUser(ctx context.Context, b *User) (*UserResponse, error) {
	hashedPassword, err := s.hasher.Hash(b.Password)
	if err != nil {
		return nil, err
	}

	b.Password = hashedPassword

	user, err := s.repo.CreateUser(ctx, b)
	if err != nil {
		return nil, err
	}

	return user.AsResponse(), nil
}

func (s *service) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.FindUserByEmail(ctx, email)
}
