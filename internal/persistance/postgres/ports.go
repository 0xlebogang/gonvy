package postgres

import (
	"database/sql"

	"gonvy.internal/ent"
)

type DB interface {
	Client() *ent.Client
	DB() *sql.DB
	Close()
}

type Postgres interface {
	Connect() (DB, error)
	Ping() error
}
