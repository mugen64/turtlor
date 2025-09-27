package configs

import (
	"os"
	"strconv"
)

type (
	Server struct {
		Protocol string
		Host     string
		Port     int
	}

	Config struct {
		Env      string
		LogLevel string
		Server   Server
	}
)

func LoadConfig() (*Config, error) {
	port := getEnvInt("PORT", 8080)
	host := getEnv("HOST", "0.0.0.0")
	protocol := getEnv("PROTOCOL", "http")

	return &Config{
		Env:      getEnv("ENV", "development"),
		LogLevel: getEnv("LOG_LEVEL", "INFO"),
		Server: Server{
			Protocol: protocol,
			Host:     host,
			Port:     port,
		},
	}, nil
}

func (s Server) Address() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

func (s Server) AddressWithProtocol() string {
	return s.Protocol + "://" + s.Address()
}

func (c Config) IsDevelopment() bool {
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
