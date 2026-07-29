package config

import (
	"os"

	"github.com/0xlebogang/gonvy/api/internal/server"
	"github.com/0xlebogang/gonvy/api/internal/storage/database"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerConf   server.Config
	DatabaseConf database.Config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func Load() *EnvConfig {
	_ = godotenv.Load()
	return &EnvConfig{
		ServerConf: server.Config{
			Port: getEnv("PORT", "8080"),
		},
		DatabaseConf: database.Config{
			Dsn: getEnv("DSN", "postgresql://postgres:postgres@localhost:5432/postgres"),
		},
	}
}
