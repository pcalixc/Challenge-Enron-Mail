package profiling

import (
	"log"
	"os"
	"runtime/pprof"
)

func CPUProfiling() {
	cpuFile, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal("Could not create CPU profile: ", err)
	}
	err = pprof.StartCPUProfile(cpuFile)
	if err != nil {
		log.Fatal("Could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
}

func MeroryProfiling() {
	memoryFile, err := os.Create("memory.pprof")
	if err != nil {
		log.Fatal("Could not create Memory profile: ", err)
	}
	pprof.WriteHeapProfile(memoryFile)
}
