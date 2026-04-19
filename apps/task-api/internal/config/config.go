package config

import (
	"os"
	"strconv"

	"github.com/codingbart/todoapp/task-api/internal/logger"
)

type Config struct {
	Host  string
	Port  uint
	DBUrl string
}

type configLoader struct {
	log logger.Logger
}

func NewConfig(log logger.Logger) Config {
	cl := &configLoader{log: log}
	return Config{
		Host:  cl.getEnvAsString("API_HOST", "localhost"),
		Port:  cl.getEnvAsUint("API_PORT", 8080),
		DBUrl: cl.getEnvAsString("API_DB_URL", "postgres://todoapp:todoapp@localhost:5432/todoapp?sslmode=disable"),
	}
}

func (cl *configLoader) getEnvAsString(key, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	cl.log.Warn("env var not set, using default", "key", key, "default", defaultVal)
	return defaultVal
}

func (cl *configLoader) getEnvAsUint(key string, defaultVal uint) uint {
	if value, ok := os.LookupEnv(key); ok {
		parsed, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			cl.log.Warn("invalid uint env var, using default", "key", key, "value", value, "default", defaultVal)
			return defaultVal
		}
		return uint(parsed)
	}

	cl.log.Warn("env var not set, using default", "key", key, "default", defaultVal)
	return defaultVal
}
