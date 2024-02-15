package main

import (
	"challenge/api/config"
	"challenge/api/server"
	"fmt"
)

func main() {
	config.LoadEnvVars()

	router := server.New()
	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}

}
