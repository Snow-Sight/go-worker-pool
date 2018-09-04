package wp

import (
	"errors"
)

var ErrQueueClosed = errors.New("The queue has been closed, jobs can no longer be sent.")
