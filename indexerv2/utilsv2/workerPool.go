package utilsv2

import (
	"fmt"
	"indexer/models"
	"log"
	"sync"
)

type WorkerPool struct {
	filePathQueue chan string
	resultChan    chan Result
	numWorkers    int
	emails        []models.EnronMail
	mu            sync.Mutex
}

func NewWorkerPool(numWorkers, bufferedSize int) *WorkerPool {
	return &WorkerPool{
		filePathQueue: make(chan string, bufferedSize),
		resultChan:    make(chan Result, bufferedSize),
		numWorkers:    numWorkers,
		emails:        make([]models.EnronMail, 0),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		worker := Worker{id: i, filePathsToProcess: wp.filePathQueue, resultChannel: wp.resultChan}
		worker.Start()
	}

	go func() {
		for result := range wp.resultChan {
			if result.err != nil {
				fmt.Printf(" Path: %s, Error: %v\n", result.filePath, result.err)
			} else {
				wp.mu.Lock()
				wp.emails = append(wp.emails, result.email)
				if len(wp.emails) == 1000 {
					wp.sendBulk()
				}
				wp.mu.Unlock()
			}
		}
	}()
}

func (wp *WorkerPool) SubmitFile(filePath string) {
	wp.filePathQueue <- filePath
}

func (wp *WorkerPool) Close() {
	if len(wp.emails) > 0 {
		wp.sendBulk()
	}
}

func (wp *WorkerPool) sendBulk() {
	if err := SendDataToIndex(&wp.emails); err != nil {
		log.Printf("Error sending data: %v", err)
	}
	wp.emails = make([]models.EnronMail, 0)
}
