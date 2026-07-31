package user

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/common"
)

type Service interface {
	CreateUser(ctx context.Context, b *User) (*UserResponse, error)
	FetchAllUsers(ctx context.Context) (*[]User, error)
	FetchUserByID(ctx context.Context, id string) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	ModifyUser(ctx context.Context, id string, b *UserUpdateRequest) (*User, error)
	RemoveUser(ctx context.Context, id string) error
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

func (s *service) FetchAllUsers(ctx context.Context) (*[]User, error) {
	return s.repo.FindAllUsers(ctx)
}

func (s *service) FetchUserByID(ctx context.Context, id string) (*User, error) {
	return s.repo.FindUserByID(ctx, id)
}

func (s *service) FetchUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.FindUserByEmail(ctx, email)
}

func (s *service) ModifyUser(ctx context.Context, id string, b *UserUpdateRequest) (*User, error) {
	return s.repo.UpdateUser(ctx, id, b)
}

func (s *service) RemoveUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}
