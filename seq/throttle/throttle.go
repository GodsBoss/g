// Package throttle allows for throttling iterators.
package throttle

import (
	"iter"
)

// Strategy describes how throttling is done.
type Strategy interface {
	// Waiter creates a Waiter. A Waiter should be independent from other Waiters created from the same Strategy.
	Waiter() Waiter
}

type Waiter interface {
	// Wait waits until the next iteration may be invoked. It is called before every yield.
	Wait()
}

// WaiterFunc wraps a function to implement Waiter.
type WaiterFunc func()

// Wait waits by calling f.
func (f WaiterFunc) Wait() {
	f()
}

// Iteration throttles the iteration of another sequence via the given strategy.
func Iteration[Value any](strategy Strategy) func(iter.Seq[Value]) iter.Seq[Value] {
	return func(sequence iter.Seq[Value]) iter.Seq[Value] {
		return func(yield func(Value) bool) {
			waiter := strategy.Waiter()
			sequence(
				func(value Value) bool {
					waiter.Wait()
					return yield(value)
				},
			)
		}
	}
}

// Iteration2 throttles the iteration of another sequence via the given strategy.
func Iteration2[First any, Second any](strategy Strategy) func(iter.Seq2[First, Second]) iter.Seq2[First, Second] {
	return func(sequence iter.Seq2[First, Second]) iter.Seq2[First, Second] {
		return func(yield func(First, Second) bool) {
			waiter := strategy.Waiter()
			sequence(
				func(first First, second Second) bool {
					waiter.Wait()
					return yield(first, second)
				},
			)
		}
	}
}
