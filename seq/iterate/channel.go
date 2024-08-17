package iterate

import (
	"iter"
	"sync"
)

// FromChannel takes a channel and returns an iterator pulling values from the channel and yielding them.
func FromChannel[Value any](values <-chan Value) iter.Seq[Value] {
	return func(yield func(value Value) bool) {
		for valueFromChannel := range values {
			if !yield(valueFromChannel) {
				return
			}
		}
	}
}

// IntoChannel creates a channel from a sequence. Values are taken from the sequence and pushed to the channel.
// After no more values are available, the channel is closed.
//
// Call cancel() to free resources and/or to stop taking values from the sequence prematurely. Can be called
// multiple times, even from different Go routines.
func IntoChannel[Value any](sequence iter.Seq[Value]) (channel <-chan Value, cancel func()) {
	ch := make(chan Value)
	closed := make(chan struct{})

	go func() {
		sequence(
			func(value Value) bool {
				select {
				case ch <- value:
					return true
				case <-closed:
					return false
				}
			},
		)

		close(ch)
	}()

	once := &sync.Once{}

	return ch, func() {
		once.Do(
			func() {
				close(closed)
			},
		)
	}
}
