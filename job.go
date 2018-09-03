package wp

// Payload represents an interface executed by the worker.
// Exec is called on the interface to run the job.
// If Exec returns an error, OnError is called.
type Payload interface {
	Exec() error
	OnError(error)
}

// Job represents a job to be run by a worker.
// Job wraps a payload so that additional data can be added if need be.
type Job struct {
	Payload Payload
}
