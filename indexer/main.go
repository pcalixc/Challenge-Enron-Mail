package main

import (
	"fmt"
	"indexer/utils"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	cpuFile, _ := os.Create("cpu.pprof")
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	startTime := time.Now()

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

	memoryFile, _ := os.Create("memory.pprof")
	pprof.WriteHeapProfile(memoryFile)
	defer pprof.StopCPUProfile()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	log.Printf("ElapsedTime: %s", elapsedTime)
}
