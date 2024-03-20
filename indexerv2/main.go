package main

import (
	"indexer/config"
	"indexer/profiling"
	"indexerv2/utils"

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

	// Check if the database path is provided as a command-line argument
	if len(os.Args) < 2 {
		log.Println("DB Path is missing.")
		return
	}

	// Construct the path to the email directory
	path := os.Args[1] + "/maildir"

	// Process emails in the specified directory
	if err := utils.ProcessEmailDirectory(path); err != nil {
		log.Printf("Error processing the email directory %v", err)
	}

	profiling.StopCPUProfile(cpuProf)
	profiling.StopMemoryProfile(memoryProf)

	endTime := time.Now()
	log.Printf("ELAPSED INDEXING TIME: %s", endTime.Sub(startTime))
}
