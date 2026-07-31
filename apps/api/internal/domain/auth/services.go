package auth

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/domain/user"
)

type Service interface {
	Authenticate(ctx context.Context, u *UserLogin) (*Tokens, error)
}

type service struct {
	repo  user.Repository
	token Token
}

func NewService(r user.Repository, t Token) Service {
	return &service{repo: r, token: t}
}

func (s *service) Authenticate(ctx context.Context, u *UserLogin) (*Tokens, error) {
	user, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.token.GenerateAccessToken(user.PublicID, user.Email)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.token.GenerateRefreshToken(user.PublicID, user.Email)
	if err != nil {
		return nil, err
	}

	tokens := &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}
