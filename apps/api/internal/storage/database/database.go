package database

import (
	"fmt"

	"github.com/0xlebogang/gonvy/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Connect() (*gorm.DB, error)
	Close() error
}

type database struct {
	instance *gorm.DB

	conf *config.Config
}

func New(c *config.Config) Database {
	return &database{
		instance: nil,
		conf:     c,
	}
}

func (d *database) Connect() (*gorm.DB, error) {
	if d.instance != nil {
		return nil, fmt.Errorf("A database connection already exists")
	}

	dialector := postgres.Open(d.conf.Dsn)
	return gorm.Open(dialector, &gorm.Config{})
}

func (d *database) Close() error {
	if d.instance == nil {
		return fmt.Errorf("No database connection to close")
	}

	sqlDB, err := d.instance.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
