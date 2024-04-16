package multicontext_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/GodsBoss/g/multicontext"
)

func TestCancelContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := multicontext.From()
	cancel(nil)

	if err := ctx.Err(); err != context.Canceled {
		t.Errorf("expected context to return %s, got %+v", context.Canceled, err)
	}

	select {
	case <-ctx.Done():
	default:
		t.Error("expected done channel to be closed")
	}
}

func TestContextHasEarliestDeadlineOfParents(t *testing.T) {
	t.Parallel()

	deadline := time.Now().Add(time.Minute)

	ctxEarliestDeadline, _ := context.WithDeadline(context.Background(), deadline)
	ctxLaterDeadline, _ := context.WithDeadline(context.Background(), deadline.Add(time.Minute))

	ctx, _ := multicontext.From(context.Background(), ctxLaterDeadline, ctxEarliestDeadline, ctxLaterDeadline)

	ctxDeadline, ok := ctx.Deadline()
	if !ok {
		t.Errorf("expected context to have deadline")
	}
	if !deadline.Equal(ctxDeadline) {
		t.Errorf("expected context to have deadline %s, %s", deadline, ctxDeadline)
	}
}

func TestCancelingParentContext(t *testing.T) {
	parent, cancelParent := context.WithCancel(context.Background())

	ctx, _ := multicontext.From(context.Background(), parent, context.TODO())

	cancelParent()

	timer := time.NewTimer(time.Millisecond)

	select {
	case <-ctx.Done():
	case <-timer.C:
		t.Error("expected done channel to be closed")
	}

	if err := ctx.Err(); err != context.Canceled {
		t.Errorf("expected context to return %s, got %+v", context.Canceled, err)
	}
}

func TestCancelingChildContex(t *testing.T) {
	t.Parallel()

	ctx, cancel := multicontext.From()

	childCtx, _ := context.WithCancel(ctx)

	cancel(nil)

	timer := time.NewTimer(time.Millisecond)

	select {
	case <-childCtx.Done():
	case <-timer.C:
		t.Error("expected done channel to be closed")
	}

	if err := childCtx.Err(); err != context.Canceled {
		t.Errorf("expected context to return %s, got %+v", context.Canceled, err)
	}
}

func TestGettingValues(t *testing.T) {
	t.Parallel()

	fooCtx := context.WithValue(context.Background(), "foo", 1)
	foobazCtx := context.WithValue(context.WithValue(context.Background(), "foo", 3), "baz", 4)

	ctx, _ := multicontext.From(fooCtx, context.Background(), foobazCtx)

	rawFoo := ctx.Value("foo")
	if foo, ok := rawFoo.(int); !ok || foo != 1 {
		t.Errorf("expected foo to be %d, got %+v", 1, rawFoo)
	}

	rawBaz := ctx.Value("baz")
	if baz, ok := rawBaz.(int); !ok || baz != 4 {
		t.Errorf("expected baz to be %d, got %+v", 4, rawBaz)
	}
}

func TestCauseAfterCancel(t *testing.T) {
	t.Parallel()

	ctx, cancel := multicontext.From()
	err := errors.New("this is some error")

	cancel(err)

	cause := context.Cause(ctx)

	if cause != err {
		t.Errorf("expected cause %v, got %v", err, cause)
	}
}

func TestCauseAfterParentCancel(t *testing.T) {
	t.Parallel()

	parent, parentCancel := context.WithCancelCause(context.Background())
	ctx, cancel := multicontext.From(parent)

	parentCause := errors.New("parent error")
	ctxCause := errors.New("context error")

	parentCancel(parentCause)

	<-ctx.Done()

	cancel(ctxCause)

	if cause := context.Cause(ctx); cause != parentCause {
		t.Errorf("expected cause %v, got %v", parentCause, cause)
	}
}

func TestCausePassesToChildren(t *testing.T) {
	ctx, cancel := multicontext.From()

	child, _ := context.WithCancel(ctx)

	ctxCause := errors.New("child error")

	cancel(ctxCause)

	<-child.Done()

	if cause := context.Cause(child); cause != ctxCause {
		t.Errorf("expected cause %v, got %v", ctxCause, cause)
	}
}

func TestCancelWithNil(t *testing.T) {
	ctx, cancel := multicontext.From()

	cancel(nil)

	if err, cause := ctx.Err(), context.Cause(ctx); err != cause {
		t.Errorf("expected equal cause and error, got cause %v and error %v", cause, err)
	}
}

func TestStringifiedContextContainsParents(t *testing.T) {
	// We test whether stringification of the multicontext contains stringifications of the parent contexts.

	// deadlineCtx is a context with a deadline. Sadly, the string representation of a deadline contains the
	// difference to the current time with a nanosecond precision, so the resulting string differs between invokations.
	// We therefore just check date and time.
	deadlineCtx, _ := context.WithDeadline(context.Background(), time.Date(2025, 3, 5, 18, 30, 0, 0, time.UTC))
	deadlineCtxString := "2025-03-05 18:30:00"

	cancelCtx, _ := context.WithCancel(context.Background())
	cancelCtxString := fmt.Sprint(cancelCtx)

	ctx, _ := multicontext.From(deadlineCtx, cancelCtx)
	ctxString := fmt.Sprint(ctx)

	deadlinePos := strings.Index(ctxString, deadlineCtxString)
	cancelPos := strings.Index(ctxString, cancelCtxString)

	if deadlinePos == -1 {
		t.Errorf("expected context string to contain '%s', got '%s'", deadlineCtxString, ctxString)
	}

	if cancelPos == -1 {
		t.Errorf("expected context string to contain '%s', got '%s'", cancelCtxString, ctxString)
	}

	if deadlinePos == -1 || cancelPos == -1 {
		return
	}

	if deadlinePos > cancelPos {
		t.Errorf("expected '%s' before '%s', got '%s'", deadlineCtxString, cancelCtxString, ctxString)
	}
}
