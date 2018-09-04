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
	q := NewQueue(1, 1)

	if err := q.QueueJob(&testPayload{1}); err != nil {
		t.Errorf("Unable to queue job with err: %v", err)
	}
}
