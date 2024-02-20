package main

import (
	"challenge/api/config"
	"challenge/api/server"
	"fmt"
	"os"
)

func main() {

	if os.Getenv("GO_ENV") != "production" {
		config.LoadEnvVars()
	}

	router := server.New()

	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}
}
