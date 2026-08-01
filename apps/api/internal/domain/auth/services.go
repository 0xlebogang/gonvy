package auth

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/common"
	"github.com/0xlebogang/gonvy/api/internal/domain/user"
)

type Service interface {
	Authenticate(ctx context.Context, u *UserLogin) (*Tokens, error)
}

type service struct {
	repo   user.Repository
	hasher common.Hasher
	token  Token
}

func NewService(r user.Repository, h common.Hasher, t Token) Service {
	return &service{repo: r, hasher: h, token: t}
}

func (s *service) Authenticate(ctx context.Context, u *UserLogin) (*Tokens, error) {
	user, err := s.repo.FindUserByEmail(ctx, u.Email)
	if err != nil {
		return nil, err
	}

	if err := s.hasher.Check(user.Password, u.Password); err != nil {
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
