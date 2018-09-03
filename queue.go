package wp

import "fmt"

// Queue is the top level struct
// It provides access to both the Job Queue and
// Stores the child dispatcher wich pulls jobs off the queue
type Queue struct {
	// JobQueue is the channel that jobs are sent over
	jobQueue   chan Job
	open       bool
	Dispatcher *Dispatcher
}

// NewQueue creates a new queue object
// With a max number of workers and a max queue buffer
func NewQueue(maxWorkers int, maxQueue int) *Queue {
	q := make(chan Job, maxQueue)
	return &Queue{
		q,
		true,
		&Dispatcher{
			jobQueue:   q,
			workerPool: make(chan chan Job, maxWorkers),
			MaxWorkers: maxWorkers,
		},
	}
}

// QueueJob adds a job onto the end of the job queue
// This is blocking and can take time
func (q *Queue) QueueJob(p Payload) error {
	if !q.open {
		return fmt.Errorf("Queue closed")
	}

	q.jobQueue <- Job{Payload: p}

	return nil
}

// Activate will start the queue dispatcher
func (q *Queue) Activate() {
	q.Dispatcher.Start()
}

// Stop stops the workers once the queue is empty, and
func (q *Queue) Stop() {
	if q.open == false {
		return
	}

	q.open = false

	q.Dispatcher.Stop(true)

	return
}
