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
				mc.cancel(parent.Err())
			case <-done:
			}
		}(parents[i], mc.done)
	}

	return mc, func(err error) {
		mc.cancel(context.Canceled)
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

	errLock sync.RWMutex
}

// cancel closes the done channel and sets this context's error to err. Only the first call takes effect.
func (mc *multicontext) cancel(err error) {
	mc.cancelOnce.Do(
		func() {
			mc.errLock.Lock()
			mc.err = err
			close(mc.done)
			mc.errLock.Unlock()
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
	mc.errLock.RLock()
	err := mc.err
	mc.errLock.RUnlock()
	return err
}

func (mc *multicontext) Value(key any) any {
	for i := range mc.parents {
		if value := mc.parents[i].Value(key); value != nil {
			return value
		}
	}

	return nil
}
