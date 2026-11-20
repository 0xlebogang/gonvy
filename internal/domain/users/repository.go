package users

import (
	"context"
	"fmt"

	"gonvy.internal/ent"
)

type Repository interface {
	CreateUser(ctx context.Context, cmd *RegisterUserCmd) (*RegisterResult, error)
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

func (r *repository) CreateUser(ctx context.Context, cmd *RegisterUserCmd) (*RegisterResult, error) {
	var out *RegisterResult
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

func (r *repository) createUserTx(ctx context.Context, cmd *RegisterUserCmd) (*RegisterResult, error) {
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
		SetProviderID("credential").
		SetNillablePassword(&cmd.Password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("create account: %w", err)
	}
	return &RegisterResult{
		User:    *u,
		Account: *a,
	}, nil
}
