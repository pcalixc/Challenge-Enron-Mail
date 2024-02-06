package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading indexer env vars")
	}

	fmt.Println("Environment vars succesfully loaded")
}
