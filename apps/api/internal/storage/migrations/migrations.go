package migrations

import (
	"github.com/0xlebogang/gonvy/api/internal/domain/user"
	"gorm.io/gorm"
)

var tables = []interface{}{
	user.User{},
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(tables...)
}
