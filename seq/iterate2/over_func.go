package iterate2

import "iter"

// OverFunc creates an iterator that calls f for values.
//
// Infinite iterator. Whether it is reusable or not depends on f.
func OverFunc[First any, Second any](f func() (First, Second)) iter.Seq2[First, Second] {
	return func(yield func(First, Second) bool) {
		for yield(f()) {
		}
	}
}
