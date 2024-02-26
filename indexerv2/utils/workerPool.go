package utils

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

// NewWorkerPool creates a new pool of workers with the specified number of workers and buffer capacity.
func NewWorkerPool(numWorkers, bufferCapacity int) *WorkerPool {
	return &WorkerPool{
		filePathQueue: make(chan string, bufferCapacity),
		resultChan:    make(chan Result, bufferCapacity),
		numWorkers:    numWorkers,
		emails:        make([]models.EnronMail, 0),
	}
}

// Start (workerPool method) starts each worker in the pool and starts listening in the results channel.
// Each result is processed accordingly: if there is an error, a message is printed; Otherwise, the emails are added to the email slice in the worker pool.
// If the number of emails in the slice reaches 1000, a batch of emails is sent by calling the sendBulk() method.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		worker := Worker{
			filePathsToProcess: wp.filePathQueue,
			resultChannel:      wp.resultChan}
		worker.Start()
	}

	go func() {
		for result := range wp.resultChan {
			if result.err != nil {
				fmt.Printf(" Path: %s, Error: %v\n", result.filePath, result.err)
			} else {
				wp.mu.Lock()
				wp.emails = append(wp.emails, result.email)
				if len(wp.emails) == batchSize {
					wp.sendBulk()
				}
				wp.mu.Unlock()
			}
		}
	}()
}

// SubmitFile (workerPool method) sends a file to the worker pool file queue channel for processing.
func (wp *WorkerPool) SubmitFile(filePath string) {
	wp.filePathQueue <- filePath
}

// Close (workerPool method) closes the worker pool and sends any pending batches of emails by calling the sendBulk() method.
func (wp *WorkerPool) Close() {
	if len(wp.emails) > 0 {
		wp.sendBulk()
	}
}

// sendBulk (workerPool method) method sends a batch of emails to the indexing system by calling SendDataToIndex() and then restarts the email slice.
func (wp *WorkerPool) sendBulk() {
	if err := SendDataToIndex(&wp.emails); err != nil {
		log.Printf("Error sending data: %v", err)
	}
	wp.emails = make([]models.EnronMail, 0)
}
