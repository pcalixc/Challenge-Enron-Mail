package main

import (
	"fmt"
	"indexer/utils"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("DB Path is missing.")
		return
	}

	path := os.Args[1] + "/maildir/"

	user_list, err := utils.ListFiles(path)
	if err != nil {
		log.Printf("Error while indexing email: %v", err)
		return
	}

	fmt.Println("indexing...")

	for u := range user_list {
		folders, err := utils.ListFiles(path + user_list[u])
		if err != nil {
			log.Printf("Error while listing email: %v", err)
			return
		}

		for f := range folders {
			utils.IndexEmailFolder(path + user_list[u] + "/" + folders[f])
		}
	}
}
