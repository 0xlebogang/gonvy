package postgres

import (
	"database/sql"

	"gonvy.internal/ent"
)

type DB interface {
	Client() *ent.Client
	DB() *sql.DB
	Ping() error
	Migrate() error
	Close()
}

type Postgres interface {
	Connect() (DB, error)
}
