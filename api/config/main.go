package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading env vars: %w", err)
	}

	log.Print("Environment vars successfully loaded")
	return nil
}

// GetEnvVar returns the value of an environment variable
func GetEnvVar(key string) string {
	return os.Getenv(key)
}
