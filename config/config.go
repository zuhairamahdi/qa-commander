// Description: This package contains the application configuration.
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config represents the application configuration.
type Config struct {
	DatabaseURL string
	AppPort     int
}

// NewConfig creates a new Config instance with default values.
func NewConfig() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/qa_commander?sslmode=disable"),
		AppPort:     getEnvAsInt("APP_PORT", 8080),
	}
}

// getEnv retrieves the value of an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt retrieves the value of an environment variable as an integer or returns a default value.
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// PrintConfig prints the configuration values to the console.
func PrintConfig(config *Config) {
	fmt.Printf("Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("App Port: %d\n", config.AppPort)
}
