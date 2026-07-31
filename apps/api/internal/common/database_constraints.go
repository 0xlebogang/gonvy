package common

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func IsDuplicateKey(err error) error {
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return gorm.ErrDuplicatedKey
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return gorm.ErrDuplicatedKey
		}
	}

	return err
}
