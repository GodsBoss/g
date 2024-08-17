package pred2

import (
	"context"
	"sync"
)

// ContextIsValid creates a predicate that holds true until a context is canceled.
func ContextIsValid[First any, Second any](ctx context.Context) func(_ First, _ Second) bool {
	return func(_ First, _ Second) bool {
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
func UntilCanceled[First any, Second any]() (predicate func(_ First, _ Second) bool, cancel func()) {
	canceled := make(chan struct{})

	predicate = func(_ First, _ Second) bool {
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
