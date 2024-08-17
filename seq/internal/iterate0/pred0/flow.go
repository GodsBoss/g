package pred0

import (
	"context"
	"sync"
)

// ContextIsValid creates a predicate that holds true until a context is canceled.
func ContextIsValid(ctx context.Context) func() bool {
	return func() bool {
		select {
		case <-ctx.Done():
			return false
		default:
			return true
		}
	}
}

// UntilCanceled creates a predicate that holds true until cancel has been called.
// cancel can be called multiple times, even from different Go routines.
func UntilCanceled() (predicate func() bool, cancel func()) {
	canceled := make(chan struct{})

	predicate = func() bool {
		select {
		case <-canceled:
			return false
		default:
			return true
		}
	}

	once := &sync.Once{}

	cancel = func() {
		once.Do(func() {
			close(canceled)
		})
	}

	return
}
