package wp

// Dispatcher is a struct that stores workers and their job channels
type Dispatcher struct {
	jobQueue chan Job
	// A pool of workers channel that are registered with the dispatcher
	workerPool chan chan Job
	// The max number of workers
	MaxWorkers int
	// Allows stopping of workers
	workers []worker
}

// Start will create workers for the dispatcher and start them up
func (d *Dispatcher) Start() {
	// starting n number of workers
	for i := 0; i < d.MaxWorkers; i++ {
		w := d.newWorker()
		w.start()
		d.workers = append(d.workers, w)
	}

	go d.dispatch()
}

// Stop cancels all workers and removes them from the job channel
// waitForEmptyQueue will hold stop until the queue is empty and all workers return
func (d *Dispatcher) Stop(waitForEmptyQueue bool) {
	if waitForEmptyQueue {
		// Wait until the job queue is empty
		for {
			if len(d.jobQueue) == 0 {
				// Wait until there is one free worker in the worker pool
				// This ensures that any job in the dispatcher waiting for a worker will recieve a worker
				for {
					if len(d.workerPool) >= 1 {
						for _, c := range d.workers {
							c.stop()
							d.workers = d.workers[1:]
						}

						return
					}
				}
			}
		}
	} else {
		for _, c := range d.workers {
			c.stop()
			d.workers = d.workers[1:]
		}
	}

}

// The private dispatch method
// Pulls jobs off of the job queue and attempts to enqueue
// them onto an available worker
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.jobQueue:
			/* NOTE: This section used to be wrapped in an anonymous go routine.
			This was removed to prevent the draining of the buffered job queue chan
			And subsequent excesive number of go routines waiting because
			they were blocked by a lack of ready workers */

			// a job request has been received
			// try to obtain a worker job channel that is available.
			// this will block until a worker is idle
			jobChannel := <-d.workerPool

			// dispatch the job to the worker job channel
			jobChannel <- job
		}
	}
}
