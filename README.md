# Golang Worker Pool Module

The golang-worker-pool package or wp aims to provide a simple, prepackaged Golang worker pool.
The initial pattern is inspired by this blog post from [marcio.io](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/).

# Usage

wp is easy to use.

To begin, create a queue using works.NewQueue(Number of Workers, Queue Buffer Length).
The queue contains a job queue channel and a dispatcher.
Start the dispatcher to launch the worker go routines and start working.
You can send jobs to the dispatcher using *Queue.QueueJob(Job{})

Jobs require a payload, with the methods `Exec()`, and `OnError()`.
This is how your code is able to be executed by the worker.
The payload can contain any data the exec function requires.
The exec function can return an error.

```go
type Payload struct {
	Num int
	wg *sync.WaitGroup
}

func (j *Payload) Exec() error {
	defer wg.Done()
	fmt.Printf("Job No. %d\n", j.Num)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func (j *Payload) OnError(err error) {
	fmt.Println(err)

	return 
}

func main() {
	q := wp.NewQueue(2, 100)

	q.Activate()

	var wg *sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		q.QueueJob(&Payload{i})
	}

	// The waitgroup is only used to ensure that the program doesn't exit
	// until all jobs are done
	wg.Wait()

}
```
