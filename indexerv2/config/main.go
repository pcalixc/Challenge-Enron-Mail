package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

// LoadEnvVars loads environment variables from a .env file using the godotenv package.
// If an error occurs during the loading process, it logs a fatal error message and exits
// the program. Otherwise, it prints a success message indicating that environment variables
// have been successfully loaded.
func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading indexer env vars")
	}

	fmt.Println("Environment vars succesfully loaded")
}
