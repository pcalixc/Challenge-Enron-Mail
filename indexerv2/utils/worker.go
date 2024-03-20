package utils

import (
	"fmt"
	"indexerv2/models"
)

type Worker struct {
	filePathsToProcess <-chan string
	processedEmails    chan<- *models.Email
}

// Start (Worker method) initiates a worker to process files. Each worker receives files from the file queue channel, processes them,
// and sends the results to the results channel.
func (w *Worker) Start() {
	go func() {
		for filePath := range w.filePathsToProcess {
			email, err := ConvertEmailFileToStruct(filePath)
			if err != nil {
				fmt.Printf("Error procesing file %s: %v", filePath, err)
			}
			w.processedEmails <- email
		}
	}()
}
