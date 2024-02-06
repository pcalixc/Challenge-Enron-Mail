package main

import (
	"challenge/api/config"
	"challenge/api/server"
	"fmt"
)

//var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	config.LoadEnvVars()

	router := server.New()
	err := router.Run()
	if err != nil {
		fmt.Println(err)
	}

}
