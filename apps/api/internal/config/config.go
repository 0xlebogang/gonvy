package config

import "os"

type Config struct {
	Port string
	Dsn  string
}

func getEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}

func Load() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Dsn:  getEnv("DSN", "postgresql://postgres:postgres@localhost:5432/postgres"),
	}
}
