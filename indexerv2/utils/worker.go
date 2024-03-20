package utils

import (
	"fmt"
	"indexer/models"
)

type Worker struct {
	filePathsToProcess <-chan string
	processedEmails    chan<- *models.EnronMail
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
			w.processedEmails <- email
		}
	}()
}
