package config

import (
	"os"

	"github.com/0xlebogang/gonvy/api/internal/domain/auth"
	"github.com/0xlebogang/gonvy/api/internal/server"
	"github.com/0xlebogang/gonvy/api/internal/storage/database"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerConf   server.Config
	AuthConf     auth.Config
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
			Port: getEnv("GONVY__PORT", "8080"),
		},
		AuthConf: auth.Config{
			AppName:              getEnv("GONVY__APP_NAME", "api.gonvy"),
			JWTSecret:            getEnv("GONVY__JWT_SECRET", "secret-key-for-dev"),
			AccessTokenLifespan:  getEnv("GONVY__ACCESS_TOKEN_LIFESPAN", "2h"),
			RefreshTokenLifespan: getEnv("GONVY__REFRESH_TOKEN_LIFESPAN", "45h"),
		},
		DatabaseConf: database.Config{
			Dsn: getEnv("GONVY__DSN", "postgresql://postgres:postgres@localhost:5432/postgres"),
		},
	}
}
