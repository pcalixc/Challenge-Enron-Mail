package utils

import (
	"indexer/models"
	"log"
)

type WorkerPool struct {
	filePathQueue   chan string
	processedEmails chan *models.EnronMail
	numWorkers      int
	emails          []models.EnronMail
}

// NewWorkerPool creates a new pool of workers with the specified number of workers and buffer capacity.
func NewWorkerPool(numWorkers, bufferCapacity int) *WorkerPool {
	return &WorkerPool{
		filePathQueue:   make(chan string, bufferCapacity),
		processedEmails: make(chan *models.EnronMail, bufferCapacity),
		numWorkers:      numWorkers,
		emails:          make([]models.EnronMail, 0),
	}
}

// Start (workerPool method) starts each worker in the pool and starts listening in the results channel.
// Each result is processed accordingly: if there is an error, a message is printed; Otherwise, the emails are added to the email slice in the worker pool.
// If the number of emails in the slice reaches 1000, a batch of emails is sent by calling the sendBulk() method.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		worker := Worker{
			filePathsToProcess: wp.filePathQueue,
			processedEmails:    wp.processedEmails}
		worker.Start()
	}

	go func() {
		for email := range wp.processedEmails {
			wp.emails = append(wp.emails, *email)
			if len(wp.emails) == batchSize {
				wp.SendAndResetBatch()
			}
		}
	}()
}

// Stop (workerPool method) closes the worker pool and sends any pending batches of emails by calling the sendBulk() method.
func (wp *WorkerPool) Stop() {
	if len(wp.emails) > 0 {
		if err := SendDataToIndex(&wp.emails); err != nil {
			log.Printf("Error sending data: %v", err)
		}
		wp.emails = make([]models.EnronMail, 0)
	}
}

// SendAndResetBatch (workerPool method) method sends a batch of emails to the indexing system by calling SendDataToIndex() and then restarts the email slice.
func (wp *WorkerPool) SendAndResetBatch() {
	if err := SendDataToIndex(&wp.emails); err != nil {
		log.Printf("Error sending data: %v", err)
	}
	wp.emails = make([]models.EnronMail, 0)
}
