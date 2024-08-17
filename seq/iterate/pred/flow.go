package pred

import (
	"context"

	"github.com/GodsBoss/g/seq/internal/iterate0/pred0"
)

// ContextIsValid creates a predicate that holds true until a context is canceled.
func ContextIsValid[Value any](ctx context.Context) func(_ Value) bool {
	return pred0.IgnoreValue[Value](pred0.ContextIsValid(ctx))
}

// UntilCanceled creates a predicate that holds true until cancel has been called.
// cancel can be called multiple times, even from different Go routines.
func UntilCanceled[Value any]() (predicate func(_ Value) bool, cancel func()) {
	p, cancel := pred0.UntilCanceled()
	return pred0.IgnoreValue[Value](p), cancel
}
