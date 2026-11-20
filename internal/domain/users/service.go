package users

import "context"

type Service interface {
	RegisterUserAccount(ctx context.Context, in *RegisterUserAccountInput) (*RegisteredUser, error)
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

func (s *service) RegisterUserAccount(ctx context.Context, in *RegisterUserAccountInput) (*RegisteredUser, error) {
	cmd := &RegisterUserCmd{
		Name:     in.Body.Name,
		Email:    in.Body.Email,
		Password: in.Body.Password,
	}
	cmd.normalize()
	return s.repo.CreateUser(ctx, cmd)
}
