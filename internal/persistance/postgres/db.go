package postgres

import (
	"database/sql"

	"gonvy.internal/ent"
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
