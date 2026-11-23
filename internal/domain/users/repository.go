package users

import (
	"context"
	"fmt"
	"time"

	"gonvy.internal/ent"
	"gonvy.internal/ent/account"
	"gonvy.internal/ent/user"
)

type Repository interface {
	CreateUser(ctx context.Context, cmd *RegisterUserCmd) (*RegisteredUser, error)
	FindUserByEmail(ctx context.Context, email string) (*ent.User, error)
}

type repository struct {
	db   *ent.Client
	inTx bool
}

func NewRepository(db *ent.Client) Repository {
	if db == nil {
		panic("db is nil")
	}
	return &repository{
		db:   db,
		inTx: false,
	}
}

func newTxRepository(tx *ent.Client) *repository {
	return &repository{db: tx, inTx: true}
}

func (r *repository) runInTx(ctx context.Context, fn func(txRepo *repository) error) error {
	if r.inTx {
		return fn(r)
	}
	tx, err := r.db.Tx(ctx)
	if err != nil {
		return fmt.Errorf("users: begin tx: %w", err)
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	err = fn(r)
	if err != nil {
		return fmt.Errorf("users: run tx: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("users: commit tx: %w", err)
	}
	return nil
}

func (r *repository) CreateUser(ctx context.Context, cmd *RegisterUserCmd) (*RegisteredUser, error) {
	var out *RegisteredUser
	err := r.runInTx(ctx, func(txRepo *repository) error {
		created, err := txRepo.createUserTx(ctx, cmd)
		if err != nil {
			return err
		}
		out = created
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (r *repository) createUserTx(ctx context.Context, cmd *RegisterUserCmd) (*RegisteredUser, error) {
	u, err := r.db.User.Create().
		SetName(cmd.Name).
		SetEmail(cmd.Email).
		Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrDuplicateEmail
		}
		return nil, fmt.Errorf("create user: %w", err)
	}

	a, err := r.db.Account.Create().
		SetUser(u).
		SetAccountID(fmt.Sprintf("test-acc-id-%s", time.Now().String())).
		SetProviderID("credential").
		SetNillablePassword(&cmd.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}

	return &RegisteredUser{
		UserID:    u.ID,
		AccountID: a.AccountID,
		Name:      u.Name,
		Email:     u.Email,
	}, nil
}

func (r *repository) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	out, err := r.db.User.Query().
		Where(user.EmailEQ(email)).
		WithAccounts(func(query *ent.AccountQuery) {
			query.Where(account.HasUserWith(user.Email(email)))
		}).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("find user: %w", err)
	}
	return out, nil
}
