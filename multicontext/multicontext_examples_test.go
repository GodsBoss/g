package multicontext_test

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GodsBoss/g/multicontext"
)

// The deadline of a multicontext is the earliest (if any) deadline of all parent contexts.
// A multicontext by itself does not have any deadline.
func Example_deadline() {
	earlyDeadline := time.Date(2067, 5, 6, 10, 0, 0, 0, time.UTC)
	lateDeadline := time.Date(2080, 11, 11, 16, 0, 0, 0, time.UTC)

	earlyDeadlineCtx, _ := context.WithDeadline(context.Background(), earlyDeadline)
	lateDeadlineCtx, _ := context.WithDeadline(context.Background(), lateDeadline)

	ctx, _ := multicontext.From(lateDeadlineCtx, earlyDeadlineCtx)

	deadline, _ := ctx.Deadline()
	fmt.Println(deadline.Format(time.RFC3339))

	// Output:
	// 2067-05-06T10:00:00Z
}

// Values are searched in the first parent first, then second parent etc., depth-first.
// This means that a parent may shadow values from parents after them.
func Example_value() {
	first := context.WithValue(context.Background(), "foo", "Hello, world!")
	second := context.WithValue(context.Background(), "foo", "Shadowed")

	ctx, _ := multicontext.From(first, second)

	fmt.Println(ctx.Value("foo"))

	// Output:
	// Hello, world!
}

// The context.CancelCauseFunc returned when creating a multicontext behaves as expected.
func Example_cause() {
	ctx, cancel := multicontext.From()

	cancel(errors.New("a very sad cause indeed"))

	fmt.Println(context.Cause(ctx))

	// Output:
	// a very sad cause indeed
}
