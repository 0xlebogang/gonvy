//go:build ignore

package main

import (
	"context"
	"database/sql"
	"log/slog"
	"os"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/gin-gonic/gin"
	pgxstdlib "github.com/jackc/pgx/v5/stdlib"
	"gonvy.internal/ent/migrate"
	"gonvy.internal/internal/config"
)

func init() {
	sql.Register("postgres", pgxstdlib.GetDefaultDriver())
}

func main() {
	ctx := context.Background()
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: map[bool]slog.Level{true: slog.LevelInfo, false: slog.LevelDebug}[gin.Mode() == gin.ReleaseMode],
	}))
	dir, err := sqltool.NewGooseDir("ent/migrate/migrations")
	if err != nil {
		log.Error("atlas migration directory creation failed", "error", err)
		os.Exit(1)
	}
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.Postgres),
	}
	if len(os.Args) != 2 {
		log.Error("migration name is required. Use 'go run -mod=mod ent/migrate/main.go <migration-name>'")
		os.Exit(1)
	}
	dsn := config.Load().DatabaseURL
	err = migrate.NamedDiff(ctx, dsn, os.Args[1], opts...)
	if err != nil {
		log.Error("migration file generation failed", "error", err)
		os.Exit(1)
	}
}
