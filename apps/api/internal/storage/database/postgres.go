package database

import (
	"fmt"

	"github.com/0xlebogang/gonvy/api/internal/storage/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Dsn string
}

type Database interface {
	Connect() (*gorm.DB, error)
	RunMigrations() error
	Close() error
}

type database struct {
	conf *Config
	conn *gorm.DB
}

func New(c *Config) Database {
	return &database{
		conf: c,
	}
}

func (d *database) Connect() (*gorm.DB, error) {
	if d.conn != nil {
		return nil, fmt.Errorf("Database connection already exist")
	}
	dialector := postgres.Open(d.conf.Dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	d.conn = db
	return db, nil
}

func (d *database) RunMigrations() error {
	return migrations.RunMigrations(d.conn)
}

func (d *database) Close() error {
	if d.conn == nil {
		return fmt.Errorf("No database connection to close")
	}
	sqlDB, err := d.conn.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
