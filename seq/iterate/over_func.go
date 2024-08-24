package iterate

import "iter"

// OverFunc creates an iterator that calls f for values.
//
// Infinite iterator. Whether it is reusable or not depends on f.
func OverFunc[Value any](f func() Value) iter.Seq[Value] {
	return func(yield func(Value) bool) {
		for yield(f()) {
		}
	}
}
