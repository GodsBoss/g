package pred2

import (
	"context"

	"github.com/GodsBoss/g/seq/internal/iterate0/pred0"
)

// ContextIsValid creates a predicate that holds true until a context is canceled.
func ContextIsValid[First any, Second any](ctx context.Context) func(_ First, _ Second) bool {
	return pred0.IgnoreValues[First, Second](pred0.ContextIsValid(ctx))
}

// UntilCanceled creates a predicate that holds true until cancel has been called.
// cancel can be called multiple times, even from different Go routines.
func UntilCanceled[First any, Second any]() (predicate func(_ First, _ Second) bool, cancel func()) {
	p, cancel := pred0.UntilCanceled()
	return pred0.IgnoreValues[First, Second](p), cancel
}
