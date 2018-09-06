package wp

import (
	"testing"
)

type testPayload struct {
	num int
}

func (j *testPayload) Exec() error {
	return nil
}

func (j *testPayload) OnError(err error) {
	return
}

func TestNewQueue(t *testing.T) {
	q := NewQueue(1, 10)

	q.Activate()

	err := q.QueueJob(&testPayload{1})

	if err != nil {
		t.Errorf("Unable to queue job with err: %v", err)
	}

	q.Stop()
}

func TestStop(t *testing.T) {
	q := NewQueue(1, 10)

	q.Activate()

	if err := q.QueueJob(&testPayload{1}); err != nil {
		t.Errorf("Unable to queue job with err: %v", err)
	}

	q.Stop()

	if err := q.QueueJob(&testPayload{1}); err != ErrQueueClosed {
		t.Errorf("Recieved incorrect error when sending job over closed queue: %v", err)
	}
}
