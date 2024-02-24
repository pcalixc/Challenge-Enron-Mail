package utilsv2

import (
	"fmt"
)

// A goroutine that processes files and send the results through a channel.
type Worker struct {
	id                 int
	filePathsToProcess <-chan string
	resultChannel      chan<- Result
}

// Worker pool manages the workers, distributes file processing and collects the emails.

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
