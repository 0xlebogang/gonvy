package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Dsn string
}

type Database interface {
	Connect() (*gorm.DB, error)
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
	return gorm.Open(dialector, &gorm.Config{})
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
