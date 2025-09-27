package configs

import (
	"os"
	"strconv"
)

type Config struct {
	Env  string
	Port int
}

func LoadConfig() (*Config, error) {
	return &Config{
		Env:  getEnv("ENV", "development"),
		Port: getEnvInt("PORT", 8080),
	}, nil
}

func (c *Config) IsDevelopment() bool {
	return c.Env == "development"
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	port, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return port
}
