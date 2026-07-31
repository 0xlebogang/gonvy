package user

import (
	"context"

	"github.com/0xlebogang/gonvy/api/internal/common"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, u *User) (*User, error) {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, common.IsDuplicateKey(err)
	}
	return u, nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
