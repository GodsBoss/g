package throttle_test

import (
	"testing"
	"time"

	"github.com/GodsBoss/g/seq/throttle"
)

func TestThrottleIteration(t *testing.T) {
	t.Parallel()

	stopSeq := make(chan struct{})
	iterations := 0
	seq := func(yield func(struct{}) bool) {
		for {
			select {
			case <-stopSeq:
				return
			default:
				iterations++
			}
			if !yield(struct{}{}) {
				return
			}
		}
	}

	stopWait := make(chan struct{})
	count := 3
	s := strategy{
		waiter: func() {
			if count > 0 {
				count--
				return
			}

			<-stopWait
		},
	}

	go func() {
		for _ = range throttle.Iteration[struct{}](s)(seq) {
			select {
			case <-stopSeq:
				return
			default:
			}
		}
	}()

	time.Sleep(time.Millisecond)
	close(stopSeq)
	time.Sleep(time.Millisecond)
	close(stopWait)

	if iterations != 4 {
		t.Errorf("expected %d iterations, got %d", 4, iterations)
	}
}

func TestThrottleIteration2(t *testing.T) {
	t.Parallel()

	stopSeq := make(chan struct{})
	iterations := 0
	seq := func(yield func(struct{}, struct{}) bool) {
		for {
			select {
			case <-stopSeq:
				return
			default:
				iterations++
			}
			if !yield(struct{}{}, struct{}{}) {
				return
			}
		}
	}

	stopWait := make(chan struct{})
	count := 3
	s := strategy{
		waiter: func() {
			if count > 0 {
				count--
				return
			}

			<-stopWait
		},
	}

	go func() {
		for _ = range throttle.Iteration2[struct{}, struct{}](s)(seq) {
			select {
			case <-stopSeq:
				return
			default:
			}
		}
	}()

	time.Sleep(time.Millisecond)
	close(stopSeq)
	time.Sleep(time.Millisecond)
	close(stopWait)

	if iterations != 4 {
		t.Errorf("expected %d iterations, got %d", 4, iterations)
	}
}

type strategy struct {
	waiter throttle.WaiterFunc
}

func (s strategy) Waiter() throttle.Waiter {
	return s.waiter
}
