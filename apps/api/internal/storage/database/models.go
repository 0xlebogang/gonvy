package database

import (
	"github.com/0xlebogang/gonvy/api/internal/domain/user"
)

var tables = []any{
	&user.User{},
}
