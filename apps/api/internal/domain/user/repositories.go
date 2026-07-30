package user

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	FindAllUsers(ctx context.Context) (*[]User, error)
	FindUserByID(ctx context.Context, id string) (*User, error)
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	UpdateUser(ctx context.Context, id string, u *UserUpdateRequest) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) CreateUser(ctx context.Context, u *User) (*User, error) {
	if err := r.db.WithContext(ctx).Create(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *repo) FindAllUsers(ctx context.Context) (*[]User, error) {
	var u []User
	if err := r.db.WithContext(ctx).Model(&u).Find(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repo) FindUserByID(ctx context.Context, id string) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).Model(&u).Where("ulid = ?", id).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repo) FindUserByEmail(ctx context.Context, email string) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).Model(&u).Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *repo) UpdateUser(ctx context.Context, id string, b *UserUpdateRequest) (*User, error) {
	u, err := r.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).Model(&u).Updates(b).Error; err != nil {
		return nil, err
	}

	if err := r.db.WithContext(ctx).First(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) DeleteUser(ctx context.Context, id string) error {
	u, err := r.FindUserByID(ctx, id)
	if err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).Delete(u).Error; err != nil {
		return err
	}

	return nil
}
