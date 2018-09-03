package wp

// Payload is the thing being executed
// It is part of the Job and is a subsection
type Payload interface {
	Exec() error
	OnError(error)
}

// Job represents the job to be run
type Job struct {
	// JobSpec *JobSpec
	// The payload can be anything
	// This allows for the job executor to have access to whatever data it needs
	Payload Payload
}

// NewJob prepares a new job with the provided payload
// func NewJob(p Payload) Job {
// 	return Job{p}
// }
