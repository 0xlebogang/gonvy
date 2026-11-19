package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	goose "github.com/pressly/goose/v3"
	"gonvy.internal/ent"
	"gonvy.internal/ent/migrate/migrations"
)

type db struct {
	client *ent.Client
	sqlDB  *sql.DB
}

func (d *db) Client() *ent.Client { return d.client }
func (d *db) DB() *sql.DB         { return d.sqlDB }

func (d *db) Close() {
	_ = d.client.Close()
	_ = d.sqlDB.Close()
}

func (d *db) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := d.DB().PingContext(ctx)
	if err != nil {
		return fmt.Errorf("ping: %w", err)
	}
	return nil
}

func (d *db) Migrate() error {
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("migrate: set dialect: %w", err)
	}
	if err := goose.Up(d.DB(), "."); err != nil {
		return fmt.Errorf("migrate: up: %w", err)
	}
	return nil
}
