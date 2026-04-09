package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPg(connectionString string) (*gorm.DB, error) {
	dialector := postgres.Open(connectionString)
	return gorm.Open(dialector)
}

func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get open database connection: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("failed to close open database connection: %v", err)
	}

	return nil
}
