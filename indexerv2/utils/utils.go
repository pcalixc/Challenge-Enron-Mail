package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Main entrance to the indexing process. Start by creating a new pool of workers with the number of available CPUs and a buffer size of 1000.
// Then, start the pool of workers by calling its Start() method. After that, it invokes the readFilesInDir() function to process the files in the specified directory.
// Finally, it closes the worker pool and returns an error if there is one.
func ProcessEmailDirectory(pathName string) error {
	workerPool := NewWorkerPool(runtime.NumCPU(), bufferCapacity)
	workerPool.Start()

	if err := readFilesInDir(pathName, workerPool); err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	workerPool.Stop()
	return nil
}

// readFilesInDir recursively traverses the files in the given directory and their subdirectories.
// For each file found, it sends it to the worker pool for processing by calling the SubmitFile() method.
// Returns an error if there is one.
func readFilesInDir(dir string, workerPool *WorkerPool) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(dir, file.Name())
		if file.IsDir() {
			// If it's a directory, recursively call the function to read the files in the subdirectory
			if err := readFilesInDir(filePath, workerPool); err != nil {
				return err
			}
		} else {
			// If it's a file, send it to the worker pool for processing
			workerPool.filePathQueue <- filePath
		}
	}
	return nil
}

func ConvertDateFormat(date string) (time.Time, error) {
	time, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700 (MST)", date)

	return time, err
}
