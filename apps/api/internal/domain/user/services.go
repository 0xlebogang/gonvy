package user

import "context"

type Service interface {
	CreateNewUser(ctx context.Context, b *UserRequest) (*User, error)
	FetchAllUsers(ctx context.Context) (*[]User, error)
	FetchUserByID(ctx context.Context, id string) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	ModifyUser(ctx context.Context, id string, b *UserUpdateRequest) (*User, error)
	RemoveUser(ctx context.Context, id string) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) CreateNewUser(ctx context.Context, b *UserRequest) (*User, error) {
	return s.repo.CreateUser(ctx, b)
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
