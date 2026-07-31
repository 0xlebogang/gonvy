package auth

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/domain/user"
)

type Service interface {
	Authenticate(ctx context.Context, u *UserLogin) (*Tokens, error)
}

type svc struct {
	repo  user.Repository
	token Token
}

func NewSvc(r user.Repository, t Token) Service {
	return &svc{repo: r, token: t}
}

func (s *svc) Authenticate(ctx context.Context, u *UserLogin) (*Tokens, error) {
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
