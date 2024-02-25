package main

import (
	"fmt"
	"indexer/config"
	"indexer/profiling"
	"indexer/utilities"
	"log"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	config.LoadEnvVars()

	cpuProf := profiling.StartCPUProfile()
	memoryProf := profiling.StartMemoryProfile()

	startTime := time.Now()

	if len(os.Args) < 2 {
		log.Println("DB Path is missing.")
		return
	}

	path := os.Args[1] + "/maildir"
	err := utilities.IndexEmailDirectory(path)
	if err != nil {
		fmt.Println("Error indexing: ", path, err)
	}

	if len(utilities.DataBatch) > 0 {
		err := utilities.SendDataToIndex(&utilities.DataBatch)
		if err != nil {
			return
		}
	}

	profiling.StopCPUProfile(cpuProf)
	profiling.StopMemoryProfile(memoryProf)

	endTime := time.Now()
	log.Printf("ELAPSED INDEXING TIME: %s", endTime.Sub(startTime))
}
