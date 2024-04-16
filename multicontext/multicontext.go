// Package multicontext provides a context made from multiple parent contexts.
package multicontext

import (
	"context"
	"sync"
	"time"
)

// From creates a new context from the given parent contexts.
//
// If any of the parent contexts have deadlines, the resulting context has a deadline too,
// which is equal to the earliest deadline any parent has.
//
// Canceling any of the parent contexts cancels the multicontext.
//
// Values returned by the multicontext are taken from the parent contexts depth-first, left to right.
//
// Canceling this context releases resources associated with it, so code should call cancel as soon as
// the operations running in this Context complete.
func From(parents ...context.Context) (context.Context, context.CancelCauseFunc) {
	parents = filterNopContexts(parents)

	mc := &multicontext{
		parents:  parents,
		deadline: getDeadlineFromParents(parents),
		done:     make(chan struct{}),
	}

	for i := range parents {
		go func(parent context.Context, done <-chan struct{}) {
			select {
			case <-parent.Done():
				mc.cancel(parent.Err(), context.Cause(parent))
			case <-done:
			}
		}(parents[i], mc.done)
	}

	return mc, func(err error) {
		mc.cancel(context.Canceled, err)
	}
}

type multicontext struct {
	// parents stores the parent contexts. This is only used for fetching values from them.
	parents []context.Context

	deadline *time.Time

	done chan struct{}

	// cancelOnce assures this context is canceled exactly once.
	cancelOnce sync.Once

	err error

	lock sync.RWMutex
}

// cancel closes the done channel and sets this context's error to err. Only the first call takes effect.
func (mc *multicontext) cancel(err error, cause error) {
	mc.cancelOnce.Do(
		func() {
			mc.lock.Lock()
			mc.err = err
			close(mc.done)

			// We can't set a cause directly on this context as the cause is an unexported field of the unexported
			// context type for cancellations and the cancellation instance is retrieved via context.Value() using
			// an unexported variable. We therefore "cheat" by creating a new context that is immediately
			// canceled with a cause and prepend that to the list of parents. context.Cause(mc) then calls context.Value()
			// on our multicontext, that call is passed to the canceled context containing our cause and voilà,
			// the cause is extracted correctly.
			if cause != err {
				causeHolder, cancel := context.WithCancelCause(context.Background())
				cancel(cause)
				mc.parents = append([]context.Context{causeHolder}, mc.parents...)
			}

			mc.lock.Unlock()
		},
	)
}

// getDeadlineFromParents returns the earlierst deadline found in any of the parent contexts. Returns nil
// if none of the parents has a deadline.
func getDeadlineFromParents(parents []context.Context) *time.Time {
	var deadline *time.Time

	for i := range parents {
		ctxDeadline, ctxOK := parents[i].Deadline()

		if !ctxOK {
			continue
		}

		if deadline == nil {
			deadline = &ctxDeadline
			continue
		}

		if ctxDeadline.Before(*deadline) {
			deadline = &ctxDeadline
		}
	}

	return deadline
}

// filterNopContexts removes Background and TODO contexts from list of contexts, as they
// don't have deadlines and cannot be canceled.
func filterNopContexts(parents []context.Context) []context.Context {
	var result []context.Context

	for i := range parents {
		if parents[i] != context.Background() && parents[i] != context.TODO() {
			result = append(result, parents[i])
		}
	}

	return result
}

func (mc *multicontext) Deadline() (deadline time.Time, ok bool) {
	if mc.deadline != nil {
		deadline = *mc.deadline
		ok = true
	}

	return
}

func (mc *multicontext) Done() <-chan struct{} {
	return mc.done
}

func (mc *multicontext) Err() error {
	mc.lock.RLock()
	err := mc.err
	mc.lock.RUnlock()
	return err
}

func (mc *multicontext) Value(key any) any {
	mc.lock.RLock()
	defer mc.lock.RUnlock()

	for i := range mc.parents {
		if value := mc.parents[i].Value(key); value != nil {
			return value
		}
	}

	return nil
}
