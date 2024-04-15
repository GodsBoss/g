package multicontext_test

import (
	"context"
	"testing"
	"time"

	"github.com/GodsBoss/g/multicontext"
)

func TestCancelContext(t *testing.T) {
	t.Parallel()

	ctx, cancel := multicontext.From()
	cancel()

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

	cancel()

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
