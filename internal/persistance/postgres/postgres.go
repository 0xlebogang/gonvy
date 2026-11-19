package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"gonvy.internal/ent"
)

type postgres struct {
	dsn string
	db  DB
}

func New(dsn string) Postgres {
	if dsn == "" {
		panic("dsn is required")
	}
	return &postgres{
		db:  nil,
		dsn: dsn,
	}
}

func (p postgres) Connect() (DB, error) {
	sqlDB, err := sql.Open("pgx", p.dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("sqlDB.PingContext: %w", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, sqlDB)
	client := ent.NewClient(ent.Driver(drv))
	p.db = &db{
		client: client,
		sqlDB:  sqlDB,
	}
	return p.db, nil
}
