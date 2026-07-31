package auth

import (
	"fmt"
	"time"

	"github.com/0xlebogang/gonvy/api/internal/config"
	jwt "github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string
	Email  string
	jwt.RegisteredClaims
}

type Token interface {
	GenerateAccessToken(userID, email string) (string, error)
	GenerateRefreshToken(userID, email string) (string, error)
	Verify(tokenString string) (*CustomClaims, error)
}

type token struct {
	conf *config.EnvConfig
}

func NewToken(c *config.EnvConfig) Token {
	return &token{conf: c}
}

func (tk *token) GenerateAccessToken(userID, email string) (string, error) {
	claims, err := tk.buildClaims(userID, email, tk.conf.AccessTokenLifespan)
	if err != nil {
		return "", err
	}

	return tk.generateToken(claims)
}

func (tk *token) GenerateRefreshToken(userID, email string) (string, error) {
	claims, err := tk.buildClaims(userID, email, tk.conf.RefreshTokenLifespan)
	if err != nil {
		return "", err
	}

	return tk.generateToken(claims)
}

func (tk *token) Verify(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(tk.conf.JWTSecret), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Token parsing failed: %v", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}

func (tk *token) buildClaims(userID, email, life string) (*CustomClaims, error) {
	tokenLife, err := time.ParseDuration(life)
	if err != nil {
		return nil, fmt.Errorf("Duration parsing failed: %v", err)
	}

	now := time.Now()

	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tk.conf.AppName,
			Subject:   userID,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(tokenLife)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	return &claims, nil
}

func (tk *token) generateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(tk.conf.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("Token signing failed: %v", err)
	}

	return signedToken, nil
}
