package wp

// Queue represents a worker pool.
// The queue contains the dispatcher, housing the workers and the jobQueue to send jobs on.
type Queue struct {
	// JobQueue is the channel that jobs are sent over.
	jobQueue chan Job
	// open marks whether the queue is able to be used.
	open bool
	// Dispatcher houses the workers and delegates jobs to them.
	Dispatcher *Dispatcher
}

// NewQueue creates a new queue object.
// With a max number of workers and a max queue buffer.
func NewQueue(maxWorkers int, maxQueue int) *Queue {
	// Make the queue
	q := make(chan Job, maxQueue)

	// Return a pointer to the queue
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

// QueueJob adds a job onto the job channel.
// This is blocking and can take time, if the channel is full.
// QueueJob will error if the queue is closed.
func (q *Queue) QueueJob(p Payload) error {
	if !q.open {
		return ErrQueueClosed
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
