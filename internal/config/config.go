package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigKey string

func (c ConfigKey) ToString() string {
	return string(c)
}

const (
	CFG_DB_HOST     = "DB_HOST"
	CFG_DB_PORT     = "DB_PORT"
	CFG_DB_USER     = "DB_USER"
	CFG_DB_PASS     = "DB_PASS"
	CFG_DB_NAME     = "DB_NAME"
	CFG_DB_SSL_MODE = "DB_SSL_MODE"
)

func GetenvString(key ConfigKey, fallback string) string {
	val := os.Getenv(key.ToString())
	if val == "" {
		return fallback
	}
	return val
}

func LoadEnvConfig(filenames ...string) error {
	return godotenv.Load(filenames...)
}
