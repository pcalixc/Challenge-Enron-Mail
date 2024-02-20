package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading env vars: %w", err)
	}

	fmt.Println("Environment vars successfully loaded")
	return nil
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
