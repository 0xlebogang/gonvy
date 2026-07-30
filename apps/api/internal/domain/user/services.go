package user

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/password"
)

type Service interface {
	Register(ctx context.Context, b *UserRequest) (*User, error)
	Authenticate(ctx context.Context, b *UserLoginRequest) error
	FetchAllUsers(ctx context.Context) (*[]User, error)
	FetchUserByID(ctx context.Context, id string) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	ModifyUser(ctx context.Context, id string, b *UserUpdateRequest) (*User, error)
	RemoveUser(ctx context.Context, id string) error
}

type service struct {
	repo     Repository
	password password.Password
}

func NewService(r Repository, p password.Password) Service {
	return &service{repo: r, password: p}
}

func (s *service) Register(ctx context.Context, b *UserRequest) (*User, error) {
	hashedPassword, err := s.password.Hash(b.Password)
	if err != nil {
		return nil, err
	}
	b.Password = hashedPassword
	return s.repo.CreateUser(ctx, b)
}

func (s *service) Authenticate(ctx context.Context, b *UserLoginRequest) error {
	user, err := s.repo.FindUserByEmail(ctx, b.Email)
	if err != nil {
		return err
	}
	return s.password.Check(user.Password, b.Password)
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
