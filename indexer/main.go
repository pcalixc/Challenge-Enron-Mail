package main

import (
	"fmt"
	"indexer/config"
	"indexer/profiling"
	"indexer/utils"
	"log"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	config.LoadEnvVars()

	profiling.CPUProfiling()

	startTime := time.Now()

	if len(os.Args) < 2 {
		log.Println("DB Path is missing.")
		return
	}

	path := os.Args[1] + "/maildir"
	err := utils.IndexEmail(path)
	if err != nil {
		fmt.Println("Error indexing: ", path, err)
	}
	if len(utils.DataBatch) > 0 {
		err := utils.SendDataToIndex(&utils.DataBatch)
		if err != nil {
			return
		}
	}

	profiling.MeroryProfiling()

	endTime := time.Now()
	log.Printf("ELAPSED INDEXING TIME: %s", endTime.Sub(startTime))
}
