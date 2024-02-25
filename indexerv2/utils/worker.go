package utils

import (
	"fmt"
)

type Worker struct {
	filePathsToProcess <-chan string
	resultChannel      chan<- Result
}

// Start (Worker method) initiates a worker to process files. Each worker receives files from the file queue channel, processes them,
// and sends the results to the results channel.
func (w *Worker) Start() {
	go func() {
		for filePath := range w.filePathsToProcess {
			email, err := ConvertEmailFileToStruct(filePath)
			if err != nil {
				fmt.Printf("Error al procesar el archivo %s: %v", filePath, err)
			}
			w.resultChannel <- Result{filePath: filePath, email: *email}
		}
	}()
}
