package multicontext_test

import (
	"context"
	"fmt"
	"time"

	"github.com/GodsBoss/g/multicontext"
)

func Example_deadline() {
	earlyDeadline := time.Date(2067, 5, 6, 10, 0, 0, 0, time.UTC)
	lateDeadline := time.Date(2080, 11, 11, 16, 0, 0, 0, time.UTC)

	earlyDeadlineCtx, _ := context.WithDeadline(context.Background(), earlyDeadline)
	lateDeadlineCtx, _ := context.WithDeadline(context.Background(), lateDeadline)

	// ctx's deadline is the earliest deadline (if any) of all parent contexts.
	ctx, _ := multicontext.From(lateDeadlineCtx, earlyDeadlineCtx)

	deadline, _ := ctx.Deadline()
	fmt.Println(deadline.Format(time.RFC3339))

	// Output:
	// 2067-05-06T10:00:00Z
}

func Example_value() {
	first := context.WithValue(context.Background(), "foo", "Hello, world!")
	second := context.WithValue(context.Background(), "foo", "Shadowed")

	// ctx will search in parent contexts from first to last depth-first for values.
	// Subsequent values are shadowed.
	ctx, _ := multicontext.From(first, second)

	fmt.Println(ctx.Value("foo"))

	// Output:
	// Hello, world!
}
