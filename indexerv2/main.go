package main

import (
	"indexer/config"
	"indexer/profiling"
	utils "indexerv2/utilsv2"
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
	utils.ProcessEmailDirectory(path)

	profiling.StopCPUProfile(cpuProf)
	profiling.StopMemoryProfile(memoryProf)

	endTime := time.Now()
	log.Printf("ELAPSED INDEXING TIME: %s", endTime.Sub(startTime))
}
