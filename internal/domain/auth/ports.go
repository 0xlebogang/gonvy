package auth

import (
	"context"

	"gonvy.internal/internal/domain/users"
)

type Service interface {
	RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*users.RegisteredUser, error)
}

type service struct {
	userSvc users.Service
}

func newService(userSvc users.Service) Service {
	if userSvc == nil {
		panic("user service is nil")
	}
	return &service{userSvc: userSvc}
}

func (s service) RegisterUserAccount(ctx context.Context, in *users.RegisterInput) (*users.RegisteredUser, error) {
	return s.userSvc.RegisterUserAccount(ctx, in)
}
