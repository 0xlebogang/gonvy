package config

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Port                 string
	AppName              string
	JWTSecret            string
	AccessTokenLifespan  string
	RefreshTokenLifespan string
	Dsn                  string
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
		Port:                 getEnv("GONVY__PORT", "8080"),
		AppName:              getEnv("GONVY__APP_NAME", "api.gonvy"),
		JWTSecret:            getEnv("GONVY__JWT_SECRET", "secret-key-for-dev"),
		AccessTokenLifespan:  getEnv("GONVY__ACCESS_TOKEN_LIFESPAN", "2h"),
		RefreshTokenLifespan: getEnv("GONVY__REFRESH_TOKEN_LIFESPAN", "45h"),
		Dsn:                  getEnv("GONVY__DSN", "postgresql://postgres:postgres@localhost:5432/postgres"),
	}
}
